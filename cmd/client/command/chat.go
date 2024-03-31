package command

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/sparhokm/go-course-ms-chat-server/pkg/chat_v1"

	"github.com/fatih/color"
	"google.golang.org/grpc/metadata"
)

type chat struct {
	ctx        context.Context
	chatClient chat_v1.ChatV1Client
	chatId     int64
	user       *user
	userNames  *userNames
}

func NewChat(ctx context.Context, chatClient chat_v1.ChatV1Client, u *user, userNames *userNames, chatId int64) *chat {
	c := chat{ctx: ctx, chatClient: chatClient, user: u, userNames: userNames, chatId: chatId}

	md := metadata.New(map[string]string{"authorization": "Bearer " + u.GetAccessToken()})
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := chatClient.ConnectChat(ctx, &chat_v1.ConnectChatIn{
		ChatId: chatId,
	})
	if err != nil {
		log.Fatalf("Can't connect to chat: %s \n", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done chat")
				return
			default:
				message, errRecv := stream.Recv()
				if errRecv == io.EOF {
					return
				}
				if errRecv != nil {
					log.Println("failed to receive message from stream: ", errRecv)
					return
				}

				fmt.Printf("%v [from: %s]: %s\n",
					color.YellowString(message.GetCreatedAt().AsTime().Format(time.TimeOnly)),
					color.BlueString(c.userNames.GetName(c.ctx, message.GetFrom())),
					message.GetText(),
				)
			}
		}
	}()

	return &c
}

func (c chat) SendMessage(text string) {
	md := metadata.New(map[string]string{"authorization": "Bearer " + c.user.GetAccessToken()})
	ctx := metadata.NewOutgoingContext(c.ctx, md)
	_, err := c.chatClient.SendMessage(ctx, &chat_v1.SendMessageIn{Text: text, ChatId: c.chatId})
	if err != nil {
		log.Fatalf("Send message error: %s", err)
	}
}

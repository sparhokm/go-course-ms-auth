package command

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sparhokm/go-course-ms-auth/pkg/auth_v1"
	"github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
	"github.com/sparhokm/go-course-ms-chat-server/pkg/chat_v1"
)

const (
	authAddress = ":50051"
	chatAddress = ":50052"
)

func connect(cmd *cobra.Command, _ []string) {
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		log.Fatalf("failed to get email: %s\n", err.Error())
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		log.Fatalf("failed to get password: %s\n", err.Error())
	}
	chatId, err := cmd.Flags().GetString("chat")
	if err != nil {
		log.Fatalf("failed to get chatId: %s\n", err.Error())
	}
	cId, err := strconv.ParseInt(chatId, 10, 0)
	if err != nil {
		log.Fatalf("failed to get chatId: %s\n", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	conn, err := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v \n", err)
	}
	u := userLogin(ctx, conn, email, password)
	uNamesService := userNamesService(conn)
	c := chatConnect(ctx, u, uNamesService, cId)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line == "exit" {
			break
		}
		fmt.Printf("\033[1A\033[K")
		c.SendMessage(line)
	}

	cancel()
}

func userLogin(ctx context.Context, conn *grpc.ClientConn, email string, password string) *user {
	client := auth_v1.NewAuthV1Client(conn)

	auth, err := client.Login(ctx, &auth_v1.LoginIn{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
	}

	return NewUser(ctx, client, email, auth.RefreshToken)
}

func chatConnect(ctx context.Context, u *user, uNamesService *userNames, chatId int64) *chat {
	conn, err := grpc.Dial(chatAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v \n", err)
	}
	client := chat_v1.NewChatV1Client(conn)

	return NewChat(ctx, client, u, uNamesService, chatId)
}

func userNamesService(conn *grpc.ClientConn) *userNames {
	client := user_v1.NewUserV1Client(conn)
	return NewUserNames(client)
}

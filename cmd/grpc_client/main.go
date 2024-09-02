package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	desc "github.com/vterebey/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}()

	client := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = CreateChat(ctx, client)
	if err != nil {
		log.Fatalf("could not create chat: %v", err)
	}

	err = DeleteChat(ctx, client)
	if err != nil {
		log.Fatalf("could not delete chat: %v", err)
	}

	_, err = ListChats(ctx, client)
	if err != nil {
		log.Fatalf("could not list chats: %v", err)
	}

	err = Connect(ctx, client)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	err = SendMessage(ctx, client)
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}

	err = AddUser(ctx, client)
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}

	err = BanUser(ctx, client)
	if err != nil {
		log.Fatalf("could not ban user: %v", err)
	}

	_, err = ListUsers(ctx, client)
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
}

func CreateChat(ctx context.Context, client desc.ChatV1Client) (*desc.CreateChatResponse, error) {
	chat := &desc.ChatInfo{
		Name:     "TestChat",
		State:    desc.ChatState_CHAT_ACTIVE,
		CreateAt: timestamppb.New(gofakeit.Date()),
	}

	users := []*desc.UserInfo{
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE},
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE},
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_DELETE},
	}

	resp, err := client.CreateChat(ctx, &desc.CreateChatRequest{Chat: chat, Users: users})
	if err != nil {
		return nil, err
	}

	fmt.Printf("Create chat response: %v\n", resp)
	return resp, nil
}

func DeleteChat(ctx context.Context, client desc.ChatV1Client) error {
	id := gofakeit.Int64()

	_, err := client.DeleteChat(ctx, &desc.DeleteChatRequest{Id: id})
	if err != nil {
		return err
	}

	fmt.Printf("Chat deleted\n")
	return nil
}

func ListChats(ctx context.Context, client desc.ChatV1Client) (*desc.ListChatsResponse, error) {

	resp, err := client.ListChats(ctx, &desc.ListChatsRequest{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("List chat response: %v\n", resp)
	return resp, nil
}

func Connect(ctx context.Context, client desc.ChatV1Client) error {
	ID := gofakeit.Int64()
	resp, err := client.Connect(ctx, &desc.ConnectRequest{Id: ID})
	if err != nil {
		return err
	}
	fmt.Printf("Connect response: %v\n", resp)
	return nil
}

func SendMessage(ctx context.Context, client desc.ChatV1Client) error {
	message := &desc.Message{From: gofakeit.Int64(), Text: gofakeit.BeerName(), Timestamp: timestamppb.New(time.Now())}

	_, err := client.SendMessage(ctx, &desc.SendMessageRequest{Message: message})
	if err != nil {
		return err
	}

	fmt.Printf("Message sent\n")
	return nil
}

func AddUser(ctx context.Context, client desc.ChatV1Client) error {
	user := &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE}

	_, err := client.AddUser(ctx, &desc.AddUserRequest{User: user})
	if err != nil {
		return err
	}

	fmt.Printf("User added\n")
	return nil
}

func BanUser(ctx context.Context, client desc.ChatV1Client) error {
	ID := gofakeit.Int64()

	_, err := client.BanUser(ctx, &desc.BanUserRequest{Id: ID})
	if err != nil {
		return err
	}

	fmt.Printf("User banned\n")
	return nil
}

func ListUsers(ctx context.Context, client desc.ChatV1Client) (*desc.ListUsersResponse, error) {
	ID := gofakeit.Int64()

	resp, err := client.ListUsers(ctx, &desc.ListUsersRequest{Id: ID})
	if err != nil {
		return nil, err
	}

	fmt.Printf("List users response: %v\n", resp)
	return resp, nil
}

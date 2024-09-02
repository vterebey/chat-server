package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	desc "github.com/vterebey/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const (
	grpcPort = 50052
)

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("server chat listening at %v", lis.Addr())
}

func (s *server) CreateChat(_ context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	fmt.Printf(color.RedString("Create Chat: "))
	fmt.Print(color.GreenString("%+v\n", req.GetUsers()))

	return &desc.CreateChatResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) DeleteChat(_ context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	fmt.Printf(color.RedString("Delete Chat: "))
	fmt.Printf(color.GreenString("%d\n"), req.GetId())

	return &emptypb.Empty{}, nil
}

func (s *server) ListChats(_ context.Context, req *desc.ListChatsRequest) (*desc.ListChatsResponse, error) {
	fmt.Printf(color.RedString("List Chats: "))

	return &desc.ListChatsResponse{
		Chats: []*desc.ChatModel{
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), State: desc.ChatState_CHAT_ACTIVE, CreateAt: timestamppb.New(gofakeit.Date())}},
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), State: desc.ChatState_CHAT_ACTIVE, CreateAt: timestamppb.New(gofakeit.Date())}},
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), State: desc.ChatState_CHAT_ACTIVE, CreateAt: timestamppb.New(gofakeit.Date())}},
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), State: desc.ChatState_CHAT_DELETE, CreateAt: timestamppb.New(gofakeit.Date())}},
		},
	}, nil
}

func (s *server) Connect(_ context.Context, req *desc.ConnectRequest) (*emptypb.Empty, error) {
	fmt.Print(color.RedString("Ban user: "))
	fmt.Print(color.GreenString("info: %d\n", req.GetId()))

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf(color.RedString("Send Message: "))
	fmt.Print(color.GreenString("%+v\n", req.GetMessage()))

	return &emptypb.Empty{}, nil
}

func (s *server) AddUser(_ context.Context, req *desc.AddUserRequest) (*emptypb.Empty, error) {
	fmt.Printf(color.RedString("Add User: "))
	fmt.Printf(color.GreenString("%+v\n", req.GetUser()))

	return &emptypb.Empty{}, nil
}

func (s *server) BanUser(_ context.Context, req *desc.BanUserRequest) (*emptypb.Empty, error) {
	fmt.Printf(color.RedString("Ban User: "))
	fmt.Printf(color.GreenString("%+v\n", req.GetId()))

	return &emptypb.Empty{}, nil
}

func (s *server) ListUsers(_ context.Context, req *desc.ListUsersRequest) (*desc.ListUsersResponse, error) {
	fmt.Printf(color.RedString("List Users: "))
	fmt.Printf(color.GreenString("%+v", req.Id))

	return &desc.ListUsersResponse{
		Users: []*desc.UserModel{
			{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE}},
			{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE}},
			{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE}},
		},
	}, nil
}

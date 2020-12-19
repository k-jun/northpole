package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	pb "northpole/grpc"
	"os"
	"testing"

	"google.golang.org/grpc"
)

var (
	testServerPort                  = 8081
	conn           *grpc.ClientConn = nil
)

func TestMain(m *testing.M) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", testServerPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := NewServer()
	go grpcServer.Serve(listener)
	defer grpcServer.Stop()
	conn, err = grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	os.Exit(m.Run())
}

func TestJoinPublicMatch(t *testing.T) {

	client := pb.NewNorthPoleClient(conn)
	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
	stream, err := client.JoinPublicMatch(context.Background(), userInfo)
	if err != nil {
		t.Fatal(err)
	}
	for {
		matchInfo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		log.Println(matchInfo)
	}
}

func TestCreatePrivateMatch(t *testing.T) {

	client := pb.NewNorthPoleClient(conn)
	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
	stream, err := client.CreatePrivateMatch(context.Background(), userInfo)
	if err != nil {
		t.Fatal(err)
	}
	for {
		matchInfo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		log.Println(matchInfo)
	}
}

func TestJoinPrivateMatch(t *testing.T) {

	client := pb.NewNorthPoleClient(conn)
	midAndUid := &pb.MatchIDAndUserID{MatchId: "b4e21a50-9b55-3ecf-88e4-7342a8c4e8a5", UserId: "0d6a9e73-1d88-35bb-8d8d-440615dfee2d"}
	stream, err := client.JoinPrivateMatch(context.Background(), midAndUid)
	if err != nil {
		t.Fatal(err)
	}
	for {
		matchInfo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		log.Println(matchInfo)
	}
}

func TestLeaveMatch(t *testing.T) {

	client := pb.NewNorthPoleClient(conn)
	matchIdAndUserId := &pb.MatchIDAndUserID{
		MatchId: "8f96317e-7731-346b-85fa-24eb1ed5b6ec",
		UserId:  "6ffca27e-6e1b-30a4-9393-91ba7c59e1e6",
	}
	matchInfo, err := client.LeaveMatch(context.Background(), matchIdAndUserId)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(matchInfo)
}

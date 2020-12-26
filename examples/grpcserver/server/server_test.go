package server

import (
	"context"
	"fmt"
	pb "grpcserver/grpc"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestJoinPublicRoom(t *testing.T) {
	client := pb.NewNorthPoleClient(conn)
	matchInfo := &pb.MatchInfo{UserId: "83e38929-e746-3d31-9c21-49a180de2448"}
	stream, err := client.JoinPublicRoom(context.Background(), matchInfo)
	if err != nil {
		t.Fatal(err)
	}
	mi, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", mi.Id)
}

func TestCreatePrivateRoom(t *testing.T) {
	client := pb.NewNorthPoleClient(conn)
	roomCreateInfo := &pb.RoomCreateInfo{RoomId: "8dc04312-e067-3d14-b224-bb6fc6eb56a1", MaxNumberOfUsers: 4, UserId: "0562ae37-4e63-3b47-8d8a-e3140e4f5177"}
	stream, err := client.CreatePrivateRoom(context.Background(), roomCreateInfo)
	if err != nil {
		t.Fatal(err)
	}
	mi, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", mi.Id)
}

func TestJoinPrivateRoom(t *testing.T) {
	client := pb.NewNorthPoleClient(conn)
	// create room before joining
	createRoomInfo := &pb.RoomCreateInfo{RoomId: "cd18e2ff-0e5e-3b8f-ac3a-8113cfbae46d", UserId: "bd0688b7-b47d-3422-a389-aba512c644ab", MaxNumberOfUsers: 2}
	stream, err := client.CreatePrivateRoom(context.Background(), createRoomInfo)
	if err != nil {
		t.Fatal(err)
	}
	mi, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}

	// join test
	matchInfo := &pb.MatchInfo{RoomId: mi.Id, UserId: "0d6a9e73-1d88-35bb-8d8d-440615dfee2d"}
	stream, err = client.JoinPrivateRoom(context.Background(), matchInfo)
	if err != nil {
		t.Fatal(err)
	}
	mi2, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, mi.Id, mi2.Id)
}

func TestLeaveRoom(t *testing.T) {
	// create, and join the room before leaving
	client := pb.NewNorthPoleClient(conn)
	matchInfo := &pb.MatchInfo{UserId: "27396fe4-d7ca-3795-a6b2-7df2d5082157"}
	stream, err := client.JoinPublicRoom(context.Background(), matchInfo)
	if err != nil {
		t.Fatal(err)
	}
	mi, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "", mi.Id)
	matchInfo.RoomId = mi.Id
	// leave test
	_, err = client.LeaveRoom(context.Background(), matchInfo)
	if err != nil {
		t.Fatal(err)
	}
}

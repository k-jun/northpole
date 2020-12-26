package server

//
// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
// 	pb "northpole/grpc"
// 	"os"
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// 	"google.golang.org/grpc"
// )
//
// var (
// 	testServerPort                  = 8081
// 	conn           *grpc.ClientConn = nil
// )
//
// func TestMain(m *testing.M) {
// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", testServerPort))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	grpcServer := NewServer()
// 	go grpcServer.Serve(listener)
// 	defer grpcServer.Stop()
// 	conn, err = grpc.Dial("localhost:8081", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	os.Exit(m.Run())
// }
//
// func TestJoinPublicMatch(t *testing.T) {
//
// 	client := pb.NewNorthPoleClient(conn)
// 	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
// 	stream, err := client.JoinPublicMatch(context.Background(), userInfo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi, err := stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.NotEqual(t, "", mi.Id)
// }
//
// func TestCreatePrivateMatch(t *testing.T) {
//
// 	client := pb.NewNorthPoleClient(conn)
// 	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
// 	stream, err := client.CreatePrivateMatch(context.Background(), userInfo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi, err := stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.NotEqual(t, "", mi.Id)
// }
//
// func TestJoinPrivateMatch(t *testing.T) {
// 	client := pb.NewNorthPoleClient(conn)
// 	// create match before joining
// 	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
// 	stream, err := client.CreatePrivateMatch(context.Background(), userInfo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi, err := stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// join test
// 	midAndUid := &pb.MatchIDAndUserID{MatchId: mi.Id, UserId: "0d6a9e73-1d88-35bb-8d8d-440615dfee2d"}
// 	stream, err = client.JoinPrivateMatch(context.Background(), midAndUid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi2, err := stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, mi.Id, mi2.Id)
// }
//
// func TestLeavePublicMatch(t *testing.T) {
// 	// create, and join the match before leaving
// 	client := pb.NewNorthPoleClient(conn)
// 	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
// 	stream, err := client.JoinPublicMatch(context.Background(), userInfo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi, err := stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// leave test
// 	matchIdAndUserId := &pb.MatchIDAndUserID{MatchId: mi.Id, UserId: userInfo.Id}
// 	mi2, err := client.LeavePublicMatch(context.Background(), matchIdAndUserId)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, mi.Id, mi2.Id)
// }
//
// func TestLeavePrivateMatch(t *testing.T) {
// 	client := pb.NewNorthPoleClient(conn)
// 	// create match before joining
// 	userInfo := &pb.UserInfo{Id: "83e38929-e746-3d31-9c21-49a180de2448"}
// 	stream, err := client.CreatePrivateMatch(context.Background(), userInfo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi, err := stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// joining match before leaving
// 	midAndUid := &pb.MatchIDAndUserID{MatchId: mi.Id, UserId: "0d6a9e73-1d88-35bb-8d8d-440615dfee2d"}
// 	stream, err = client.JoinPrivateMatch(context.Background(), midAndUid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	mi, err = stream.Recv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// test leave
// 	matchIdAndUserId := &pb.MatchIDAndUserID{MatchId: mi.Id, UserId: midAndUid.UserId}
// 	mi2, err := client.LeavePrivateMatch(context.Background(), matchIdAndUserId)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, mi.Id, mi2.Id)
// }

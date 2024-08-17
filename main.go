package main

import (
	"context"
	"log"
	"net"

	app "github.com/moulid15/BetOnBet/app"
	pb "github.com/moulid15/BetOnBet/proto"
	"google.golang.org/grpc"
)

type theBetOnBetServer struct {
	pb.UnimplementedBetOnBetServiceServer
}

func (s theBetOnBetServer) CompletedScores(ctx context.Context, req *pb.CompletedScoresRequest) (*pb.CompletedScoresResponse, error) {
	Completed := pb.CompletedScoresResponse{}
	g := app.Game{}

	res, err := g.GetScores(req.League, req.Date)

	if err != nil {
		return nil, err
	}

	Completed.BoxScore = res
	return &Completed, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("error = ", err)
	}
	serverRegister := grpc.NewServer()
	server := &theBetOnBetServer{}
	pb.RegisterBetOnBetServiceServer(serverRegister, server)
	err = serverRegister.Serve(lis)

	if err != nil {
		log.Fatalf("error = ", err)
	}

}

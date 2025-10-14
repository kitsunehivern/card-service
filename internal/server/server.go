package server

import (
	"card-service/gen/proto"
	"card-service/internal/service"
	"net"

	"google.golang.org/grpc"
)

func NewRouter(cardSvc *service.CardService, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	proto.RegisterCardServiceServer(s, cardSvc)

	return s.Serve(lis)
}

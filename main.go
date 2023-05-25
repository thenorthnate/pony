package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/thenorthnate/pony/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterPubsubServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	api.UnimplementedPubsubServer
}

func (s *server) Publish(stream api.Pubsub_PublishServer) error {
	for {
		tmpNote, err := stream.Recv()
		if err == io.EOF {
			log.Info().Msgf("Received EOF")
			break
		}
		if err != nil {
			log.Error().Msgf("error occurred: %v", err)
			return err
		}
		log.Info().Msgf("received message with %v bytes", len(tmpNote.Data))
	}
	return nil
}

func (s *server) Subscribe(note *api.Note, stream api.Pubsub_SubscribeServer) error {
	log.Info().Msgf("received request to subscribe to %v", note.RoutingKey)
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for currentTime := range ticker.C {
		note := api.Note{
			Data: []byte(currentTime.Format(time.RFC3339Nano)),
		}
		err := stream.Send(&note)
		if err != nil {
			return err
		}
	}
	return nil
}

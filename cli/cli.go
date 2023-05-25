package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/thenorthnate/pony/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewPubsubClient(conn)
	note := api.Note{
		Data: []byte("hello, world!"),
	}
	publishNote(client, &note)
}

func publishNote(client api.PubsubClient, note *api.Note) {
	// log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Publish(ctx)
	if err != nil {
		log.Fatal().Msgf("client.Publish failed: %v", err)
	}
	for i := 0; i < 3; i++ {
		err = stream.Send(note)
		if err != nil {
			log.Fatal().Msgf("client.Send failed: %v", err)
		}
	}
	log.Info().Msg("published message")
}

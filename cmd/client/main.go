package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
	pb "voicekit-mock/pkg/api/recognizer/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


const (
	ChunkSize = 1024
	SampleRate = 16000
)
func main() {
	addr := "localhost:8080"
	sessionDurations := 7*time.Second

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewVoiceRecognizerClient(conn)

	runSession(client, sessionDurations)
	
}

func runSession(client pb.VoiceRecognizerClient, duration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	stream, err := client.Recognize(ctx)
	if err != nil {
		return err
	}

	waitReciever := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// Reciever
	go func() {
		defer wg.Done()
		defer close(waitReciever)
		
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println(err)
				return
			}

			if resp.IsFinal {
				if len(resp.Results) > 0 {
				fmt.Printf("%s ", resp.Results[0].Text)
				}
			}
			// else {}
		}

	}()

	// Sender
	go func() {
		defer wg.Done()
		cfg := &pb.RecognizeRequest{
			StreamingRequest: &pb.RecognizeRequest_Config{
				Config: &pb.StreamingConfig{
					Encoding: pb.AudioEncoding_LINEAR16,
					SampleRate: 16000,
					AudioChannelCount: 1,
					LanguageCode: "ru-RU",
				},
			},
		}
		err := stream.Send(cfg)
		if err != nil {
			fmt.Println(err)
			return
		}

		buff := make([]byte, ChunkSize)
		ticker := time.NewTicker(30*time.Millisecond)
		defer ticker.Stop()
		timeout := time.After(duration)

		for finished:=false; !finished; {
			select {
			case <-timeout:
				finished = true
			case <-ticker.C:
				rand.Read(buff)
				req := &pb.RecognizeRequest{
					StreamingRequest: &pb.RecognizeRequest_AudioContent{
						AudioContent: buff,
					},
				}
				stream.Send(req)

			}
		}
	}()
	
	wg.Wait()

	if err := stream.CloseSend(); err != nil {
		return err
	}

	<-waitReciever
	return nil
}


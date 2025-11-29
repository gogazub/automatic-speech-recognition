package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"voicekit-mock/config"
	handler "voicekit-mock/internal/service/recognizer"
	pb "voicekit-mock/pkg/api/recognizer/v1"
	"voicekit-mock/pkg/logger"

	"google.golang.org/grpc"
)

func main(){
	cfg := config.NewConfig()

	log, err := logger.New()
	if err != nil {
		fmt.Printf("create logger error: %v", err)
		os.Exit(1)
	}
	defer log.Sync()

	voiceService := handler.New()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HTTPCfg.Host, cfg.HTTPCfg.Port))
	if err != nil {
		log.Fatal("starting server", "err", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVoiceRecognizerServer(grpcServer, voiceService)

	log.Info("started server", "addr", lis.Addr().String())

	go func(){
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("grpc server failed to serve", "err", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	<-quit
	log.Info("shutting down server...")
	grpcServer.GracefulStop()
}
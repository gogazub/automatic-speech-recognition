package main

import (
	"fmt"
	"net"
	"voicekit-mock/config"
	recognizerservice "voicekit-mock/internal/service/recognizer"
	pb "voicekit-mock/pkg/api/recognizer/v1"

	"google.golang.org/grpc"
)

func main(){
	cfg := config.NewConfig()

	//logger := logger.NewLogger(cfg.LoggerCfg)

	voiceRecognizer := recognizerservice.New()

	lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HTTPCfg.Host, cfg.HTTPCfg.Port))

	grpcServer := grpc.NewServer()
	pb.RegisterVoiceRecognizerServer(grpcServer, voiceRecognizer)
	_ = grpcServer.Serve(lis)

}
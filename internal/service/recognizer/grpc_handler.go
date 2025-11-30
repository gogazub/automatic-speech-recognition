package handler

import (
	"fmt"
	"io"
	"voicekit-mock/internal/processor"
	pb "voicekit-mock/pkg/api/recognizer/v1"
	"voicekit-mock/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedVoiceRecognizerServer
	log *logger.Logger
}

func (s *Service) Recognize(stream pb.VoiceRecognizer_RecognizeServer) error {
	ctx := stream.Context()
	s.log.Info("Session started")

	var proc processor.Processor

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				s.log.Info("Session finished (EOF)")
				return nil
			}
			if status.Code(err) == codes.Canceled {
				s.log.Info("client canceled connection")
				return nil
			}
			s.log.Error("Stream error", "err", err)
			return err
		}

		switch v := req.StreamingRequest.(type) {
			
		case *pb.RecognizeRequest_Config:
			if proc != nil {
				return status.Error(codes.InvalidArgument, "config sent twice")
			}

			s.log.Info("Received config", "encoding", v.Config.Encoding, "sample rate", v.Config.SampleRate)

			procCfg := processor.Config{
				Encoding: int32(*v.Config.Encoding.Enum()),
				SampleRate: v.Config.SampleRate,
				AudioChannelCount: v.Config.AudioChannelCount,
			}

			proc = processor.NewProcessor(procCfg)

		case *pb.RecognizeRequest_AudioContent:

		if proc == nil {
			return status.Error(codes.FailedPrecondition, "stream error: no config Received")
		}

		result, err := proc.Process(ctx, v.AudioContent)
		if err != nil {
			s.log.Error("processing error", "err", err)
			return status.Error(codes.Internal, "processing error")
		}

		if result != nil {
			resp := &pb.RecognizeResponse{
				Results: []*pb.Transcript{
					{
						Text: result.Text,
						Confidence: result.Confidence,
					},
				},
				IsFinal: result.IsFinal,
			}
			if err := stream.Send(resp); err != nil {
				return fmt.Errorf("send response error: %w", err)
			}
		}
		case nil:
			s.log.Info("empty message")
		default:
			s.log.Warn("unexpected type")
		}
		
	}
}

func New(log *logger.Logger) *Service {
	return &Service{log: log}
}
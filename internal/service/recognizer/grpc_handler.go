package recognizerservice

import (
	pb "voicekit-mock/pkg/api/recognizer/v1"
	"voicekit-mock/pkg/logger"
)

type Service struct {
	pb.UnimplementedVoiceRecognizerServer
	log *logger.Logger
}

func (vr *Service) Recognize(stream pb.VoiceRecognizer_RecognizeServer) (error) {
	return nil
}

func New() *Service {
	return &Service{}
}
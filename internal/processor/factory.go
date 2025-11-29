package processor

import (
	"voicekit-mock/internal/processor/decoder"
	"voicekit-mock/internal/processor/recognizer"
)

type Config struct {
	Encoding int32
	SampleRate int32
	AudioChannelCount int32
}

func NewProcessor(cfg Config) Processor {
	var deco decoder.Decoder

	switch cfg.Encoding {
	case 1: // AudioEncoding LINEAR16. todo: mapper pb.AutoEncoding -> enum
		deco = decoder.NewLinear16Decoder()
	default:
		deco = decoder.NewNopDecoder()
	}

	reco := recognizer.NewRecognizer(cfg.SampleRate)
	return newProcessor(deco, reco)
}
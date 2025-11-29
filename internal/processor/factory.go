package processor

import (
	"voicekit-mock/internal/processor/decoder"
	"voicekit-mock/internal/processor/recognizer"
)

type Config struct {
	Encoding string
	SampleRate int32
	AudioChannelCount int32
}

func NewProcessor(cfg Config) Processor {
	var deco decoder.Decoder

	switch cfg.Encoding {
	case "Linear16", "1":
		deco = decoder.NewLinear16Decoder()
	default:
		deco = decoder.NewMockDecoder()
	}

	reco := recognizer.NewRecognizer(cfg.SampleRate)
	return newProcessor(deco, reco)
}
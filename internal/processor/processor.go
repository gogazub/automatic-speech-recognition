// contract
package processor

import (
	"context"
	"fmt"
	"voicekit-mock/internal/processor/decoder"
	"voicekit-mock/internal/processor/recognizer"
)

type Processor interface {
	Process(ctx context.Context, chunk []byte) (*recognizer.Result, error)
}

type processor struct {
	decoder    decoder.Decoder
	recognizer recognizer.Recognizer
}

func (p *processor) Process(ctx context.Context, chunk []byte) (*recognizer.Result, error) {
	pcm, err := p.decoder.Decode(chunk)
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return p.recognizer.Recognize(ctx, pcm)
}


func newProcessor(decoder decoder.Decoder, recognizer recognizer.Recognizer) *processor {
	return &processor{decoder: decoder, recognizer: recognizer}
}

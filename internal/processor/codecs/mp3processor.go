package codecs

import (
	"context"
	"voicekit-mock/internal/processor"
)

type MP3Processor struct {
	buf []byte
}

func (proc *MP3Processor) Process(ctx context.Context, chunk []byte) (*processor.Result, error) {
	return nil, nil
}
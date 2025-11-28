package codec

import (
	"context"
)

type MP3Processor struct {
	buf []byte
}

func (proc *MP3Processor) Process(ctx context.Context, chunk []byte) ([]byte, error) {
	return nil, nil
}


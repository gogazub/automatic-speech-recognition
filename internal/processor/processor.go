// contract
package processor

import (
	"context"
	"voicekit-mock/internal/processor/decoder"
	"voicekit-mock/internal/processor/recognizer"
	pb "voicekit-mock/pkg/api/recognizer/v1"
)

type Processor interface {
	Process(ctx context.Context, chunk []byte) (*recognizer.Result, error)
}

type processor struct {
	decoder decoder.Decoder
	recognizer Recognizer
}

type Recognizer interface {
}


func NewProcessor(cfg *pb.StreamingConfig) *Processor {
	return nil
}

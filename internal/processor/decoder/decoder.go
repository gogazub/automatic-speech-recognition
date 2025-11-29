package decoder

type DecoderConfig struct {
	//Encoding          int/enum
	SampleRate        int32
	AudioChannelCount int32
}

type Decoder interface {
	Decode(chunk []byte) ([]byte, error)
}

type Linear16Decoder struct{}

func NewLinear16Decoder() *Linear16Decoder {
	return &Linear16Decoder{}
}

// to implement
func (d *Linear16Decoder) Decode(chunk []byte) ([]byte, error) {
	return chunk, nil
}

type NopDecoder struct{}

func NewNopDecoder() *NopDecoder {
	return &NopDecoder{}
}

func (d *NopDecoder) Decode(chunk []byte) ([]byte, error) {
	return chunk, nil
}

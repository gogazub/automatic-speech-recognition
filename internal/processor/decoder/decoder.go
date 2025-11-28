package decoder

type DecoderConfig struct {
	//Encoding          int/enum
	SampleRate        int32
	AudioChannelCount int32
}

type Decoder interface {
	Decode(chunk []byte) []byte
}

func NewDecoder()
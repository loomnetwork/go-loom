package contractpb

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom/plugin"
)

var (
	errUnknownEncodingType = errors.New("unknown encoding type")
)

type PBMarshaler interface {
	Marshal(w io.Writer, pb proto.Message) error
}

type PBUnmarshaler interface {
	Unmarshal(r io.Reader, pb proto.Message) error
}

type BinaryPBMarshaler struct {
}

func (m *BinaryPBMarshaler) Marshal(w io.Writer, pb proto.Message) error {
	b, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

type BinaryPBUnmarshaler struct {
}

func (m *BinaryPBUnmarshaler) Unmarshal(r io.Reader, pb proto.Message) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, pb)
}

func marshalerFactory(encoding plugin.EncodingType) (PBMarshaler, error) {
	switch encoding {
	case plugin.EncodingType_JSON:
		return &jsonpb.Marshaler{}, nil
	case plugin.EncodingType_PROTOBUF3:
		return &BinaryPBMarshaler{}, nil
	}

	return nil, errUnknownEncodingType
}

func unmarshalerFactory(encoding plugin.EncodingType) (PBUnmarshaler, error) {
	switch encoding {
	case plugin.EncodingType_JSON:
		return &jsonpb.Unmarshaler{}, nil
	case plugin.EncodingType_PROTOBUF3:
		return &BinaryPBUnmarshaler{}, nil
	}

	return nil, errUnknownEncodingType
}

package gomicrosvc

import (
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/runtime/protoiface"
)

func Unmarshal(any *Data, m protoiface.MessageV1) error {
	err := ptypes.UnmarshalAny(any, m)
	if err != nil {
		return err
	}
	return nil
}

func Marshal(m protoiface.MessageV1) (*Data, error) {
	data, err := ptypes.MarshalAny(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}

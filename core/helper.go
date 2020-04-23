package core

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"strings"
)

func GetAnyType(a *any.Any) (interface{}, error) {
	switch strings.ToLower(a.TypeUrl) {
	case strings.ToLower("google.protobuf.StringValue"):
		var md wrappers.StringValue
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.BytesValue"):
		var md wrappers.BytesValue
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.BoolValue"):
		var md wrappers.BoolValue
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return &md.Value, nil
	case strings.ToLower("google.protobuf.UInt32Value"):
		var md wrappers.UInt32Value
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.Int32Value"):
		var md wrappers.Int32Value
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.UInt64Value"):
		var md wrappers.UInt64Value
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.Int64Value"):
		var md wrappers.Int64Value
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.FloatValue"):
		var md wrappers.FloatValue
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	case strings.ToLower("google.protobuf.DoubleValue"):
		var md wrappers.DoubleValue
		if err := ptypes.UnmarshalAny(a, &md); err != nil {
			return nil, err
		}
		return md.Value, nil
	}

	return nil, fmt.Errorf("use google.protobuf types list")
}

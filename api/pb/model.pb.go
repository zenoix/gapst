// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: model.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SepalLengthCm float32 `protobuf:"fixed32,1,opt,name=SepalLengthCm,proto3" json:"SepalLengthCm,omitempty"`
	SepalWidthCm  float32 `protobuf:"fixed32,2,opt,name=SepalWidthCm,proto3" json:"SepalWidthCm,omitempty"`
	PetalLengthCm float32 `protobuf:"fixed32,3,opt,name=PetalLengthCm,proto3" json:"PetalLengthCm,omitempty"`
	PetalWidthCm  float32 `protobuf:"fixed32,4,opt,name=PetalWidthCm,proto3" json:"PetalWidthCm,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetSepalLengthCm() float32 {
	if x != nil {
		return x.SepalLengthCm
	}
	return 0
}

func (x *PredictRequest) GetSepalWidthCm() float32 {
	if x != nil {
		return x.SepalWidthCm
	}
	return 0
}

func (x *PredictRequest) GetPetalLengthCm() float32 {
	if x != nil {
		return x.PetalLengthCm
	}
	return 0
}

func (x *PredictRequest) GetPetalWidthCm() float32 {
	if x != nil {
		return x.PetalWidthCm
	}
	return 0
}

type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Species string `protobuf:"bytes,1,opt,name=Species,proto3" json:"Species,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x70, 0x61, 0x6c,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x43, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x53, 0x65, 0x70, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x43, 0x6d, 0x12, 0x22, 0x0a,
	0x0c, 0x53, 0x65, 0x70, 0x61, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x43, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x53, 0x65, 0x70, 0x61, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x43,
	0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x65, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x43, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x50, 0x65, 0x74, 0x61, 0x6c, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x43, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x65, 0x74, 0x61, 0x6c,
	0x57, 0x69, 0x64, 0x74, 0x68, 0x43, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x50,
	0x65, 0x74, 0x61, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x43, 0x6d, 0x22, 0x2b, 0x0a, 0x0f, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x32, 0x41, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x38, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),  // 0: model.PredictRequest
	(*PredictResponse)(nil), // 1: model.PredictResponse
}
var file_model_proto_depIdxs = []int32{
	0, // 0: model.Model.Predict:input_type -> model.PredictRequest
	1, // 1: model.Model.Predict:output_type -> model.PredictResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}

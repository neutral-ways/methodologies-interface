// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: pb/methodology.proto

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

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methodology                      string   `protobuf:"bytes,1,opt,name=Methodology,proto3" json:"Methodology,omitempty"`
	Reference                        uint32   `protobuf:"varint,2,opt,name=Reference,proto3" json:"Reference,omitempty"`
	Version                          uint32   `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	AverageTemperature               *float64 `protobuf:"fixed64,10,opt,name=AverageTemperature,proto3,oneof" json:"AverageTemperature,omitempty"`                             // AVG_TEMPERATURE
	AccumulatedPrecipitation         *float64 `protobuf:"fixed64,11,opt,name=AccumulatedPrecipitation,proto3,oneof" json:"AccumulatedPrecipitation,omitempty"`                 //S UM_PRECIPITATION
	Evapotranspiration               *float64 `protobuf:"fixed64,12,opt,name=Evapotranspiration,proto3,oneof" json:"Evapotranspiration,omitempty"`                             // EVAPOTRANSPIRATION
	FractionMethaneRecovery          *float64 `protobuf:"fixed64,13,opt,name=FractionMethaneRecovery,proto3,oneof" json:"FractionMethaneRecovery,omitempty"`                   // BE_SWDS_CH4_RECOVERY
	MethaneCorrectionFactorDefault   *string  `protobuf:"bytes,14,opt,name=MethaneCorrectionFactorDefault,proto3,oneof" json:"MethaneCorrectionFactorDefault,omitempty"`       // BE_SWDS_MCF_DEFAULT
	OxidationFactor                  *float64 `protobuf:"fixed64,15,opt,name=OxidationFactor,proto3,oneof" json:"OxidationFactor,omitempty"`                                   // BE_SWDS_OXIDATION_FACTOR
	OverwriteMethaneCorrectionFactor *float64 `protobuf:"fixed64,16,opt,name=OverwriteMethaneCorrectionFactor,proto3,oneof" json:"OverwriteMethaneCorrectionFactor,omitempty"` // BE_SWDS_MCF_OVERWRITE
	WaterTable                       *string  `protobuf:"bytes,17,opt,name=WaterTable,proto3,oneof" json:"WaterTable,omitempty"`                                               // BE_SWDS_MCF_WATER_TABLE
	WaterTableHeight                 *float64 `protobuf:"fixed64,18,opt,name=WaterTableHeight,proto3,oneof" json:"WaterTableHeight,omitempty"`                                 // BE_SWDS_MCF_WATER_TABLE_HEIGHT
	WaterTableDepth                  *float64 `protobuf:"fixed64,19,opt,name=WaterTableDepth,proto3,oneof" json:"WaterTableDepth,omitempty"`                                   // BE_SWDS_MCF_DEPTH
	DecompositionYear                *int32   `protobuf:"varint,20,opt,name=DecompositionYear,proto3,oneof" json:"DecompositionYear,omitempty"`                                // WASTE_SWDS_DECOMPOSITION_YEAR
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_methodology_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_methodology_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_pb_methodology_proto_rawDescGZIP(), []int{0}
}

func (x *DataRequest) GetMethodology() string {
	if x != nil {
		return x.Methodology
	}
	return ""
}

func (x *DataRequest) GetReference() uint32 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *DataRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DataRequest) GetAverageTemperature() float64 {
	if x != nil && x.AverageTemperature != nil {
		return *x.AverageTemperature
	}
	return 0
}

func (x *DataRequest) GetAccumulatedPrecipitation() float64 {
	if x != nil && x.AccumulatedPrecipitation != nil {
		return *x.AccumulatedPrecipitation
	}
	return 0
}

func (x *DataRequest) GetEvapotranspiration() float64 {
	if x != nil && x.Evapotranspiration != nil {
		return *x.Evapotranspiration
	}
	return 0
}

func (x *DataRequest) GetFractionMethaneRecovery() float64 {
	if x != nil && x.FractionMethaneRecovery != nil {
		return *x.FractionMethaneRecovery
	}
	return 0
}

func (x *DataRequest) GetMethaneCorrectionFactorDefault() string {
	if x != nil && x.MethaneCorrectionFactorDefault != nil {
		return *x.MethaneCorrectionFactorDefault
	}
	return ""
}

func (x *DataRequest) GetOxidationFactor() float64 {
	if x != nil && x.OxidationFactor != nil {
		return *x.OxidationFactor
	}
	return 0
}

func (x *DataRequest) GetOverwriteMethaneCorrectionFactor() float64 {
	if x != nil && x.OverwriteMethaneCorrectionFactor != nil {
		return *x.OverwriteMethaneCorrectionFactor
	}
	return 0
}

func (x *DataRequest) GetWaterTable() string {
	if x != nil && x.WaterTable != nil {
		return *x.WaterTable
	}
	return ""
}

func (x *DataRequest) GetWaterTableHeight() float64 {
	if x != nil && x.WaterTableHeight != nil {
		return *x.WaterTableHeight
	}
	return 0
}

func (x *DataRequest) GetWaterTableDepth() float64 {
	if x != nil && x.WaterTableDepth != nil {
		return *x.WaterTableDepth
	}
	return 0
}

func (x *DataRequest) GetDecompositionYear() int32 {
	if x != nil && x.DecompositionYear != nil {
		return *x.DecompositionYear
	}
	return 0
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference uint32 `protobuf:"varint,2,opt,name=Reference,proto3" json:"Reference,omitempty"`
	Version   uint32 `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_methodology_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_methodology_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_pb_methodology_proto_rawDescGZIP(), []int{1}
}

func (x *DetailRequest) GetReference() uint32 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *DetailRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail        []byte  `protobuf:"bytes,1,opt,name=Detail,proto3" json:"Detail,omitempty"`
	ReductionUnit float64 `protobuf:"fixed64,2,opt,name=ReductionUnit,proto3" json:"ReductionUnit,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_methodology_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_methodology_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_pb_methodology_proto_rawDescGZIP(), []int{2}
}

func (x *DetailResponse) GetDetail() []byte {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *DetailResponse) GetReductionUnit() float64 {
	if x != nil {
		return x.ReductionUnit
	}
	return 0
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculate float64 `protobuf:"fixed64,1,opt,name=Calculate,proto3" json:"Calculate,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_methodology_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_methodology_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_pb_methodology_proto_rawDescGZIP(), []int{3}
}

func (x *CalculateResponse) GetCalculate() float64 {
	if x != nil {
		return x.Calculate
	}
	return 0
}

var File_pb_methodology_proto protoreflect.FileDescriptor

var file_pb_methodology_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x22, 0x8b, 0x08, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x12, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x12, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x3f, 0x0a, 0x18, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x50, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x18, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x12, 0x45, 0x76, 0x61, 0x70, 0x6f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x02, 0x52, 0x12, 0x45, 0x76, 0x61, 0x70, 0x6f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x17, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x17, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x1e, 0x4d, 0x65, 0x74, 0x68, 0x61,
	0x6e, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x1e, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x4f, 0x78, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x48, 0x05, 0x52,
	0x0f, 0x4f, 0x78, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x20, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x48, 0x06, 0x52,
	0x20, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e,
	0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x57, 0x61, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x57, 0x61, 0x74,
	0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x08, 0x52, 0x10, 0x57, 0x61, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x57, 0x61,
	0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x09, 0x52, 0x0f, 0x57, 0x61, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x44, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52, 0x11, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13,
	0x5f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x45, 0x76, 0x61, 0x70, 0x6f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x42, 0x21, 0x0a, 0x1f, 0x5f, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x43,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x4f, 0x78, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x23, 0x0a, 0x21, 0x5f, 0x4f,
	0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x61, 0x6e, 0x65, 0x43,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x57, 0x61, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x13,
	0x0a, 0x11, 0x5f, 0x57, 0x61, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x57, 0x61, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x44, 0x65, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x4a,
	0x04, 0x08, 0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a, 0x04, 0x08, 0x09, 0x10,
	0x0a, 0x22, 0x47, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x0e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x31, 0x0a, 0x11, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x32, 0xa1, 0x01,
	0x0a, 0x12, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x10, 0x5a, 0x0e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_methodology_proto_rawDescOnce sync.Once
	file_pb_methodology_proto_rawDescData = file_pb_methodology_proto_rawDesc
)

func file_pb_methodology_proto_rawDescGZIP() []byte {
	file_pb_methodology_proto_rawDescOnce.Do(func() {
		file_pb_methodology_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_methodology_proto_rawDescData)
	})
	return file_pb_methodology_proto_rawDescData
}

var file_pb_methodology_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_methodology_proto_goTypes = []interface{}{
	(*DataRequest)(nil),       // 0: methodology.DataRequest
	(*DetailRequest)(nil),     // 1: methodology.DetailRequest
	(*DetailResponse)(nil),    // 2: methodology.DetailResponse
	(*CalculateResponse)(nil), // 3: methodology.CalculateResponse
}
var file_pb_methodology_proto_depIdxs = []int32{
	0, // 0: methodology.MethodologyService.Calculate:input_type -> methodology.DataRequest
	1, // 1: methodology.MethodologyService.GetDetail:input_type -> methodology.DetailRequest
	3, // 2: methodology.MethodologyService.Calculate:output_type -> methodology.CalculateResponse
	2, // 3: methodology.MethodologyService.GetDetail:output_type -> methodology.DetailResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_methodology_proto_init() }
func file_pb_methodology_proto_init() {
	if File_pb_methodology_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_methodology_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_pb_methodology_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
		file_pb_methodology_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
		file_pb_methodology_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
	file_pb_methodology_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_methodology_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_methodology_proto_goTypes,
		DependencyIndexes: file_pb_methodology_proto_depIdxs,
		MessageInfos:      file_pb_methodology_proto_msgTypes,
	}.Build()
	File_pb_methodology_proto = out.File
	file_pb_methodology_proto_rawDesc = nil
	file_pb_methodology_proto_goTypes = nil
	file_pb_methodology_proto_depIdxs = nil
}

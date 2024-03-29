// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: outliers.proto

//import "google/protobuf/wrappers.proto";

package GoPythonGrpcTest

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reports int32

const (
	Reports_ClusterConfig       Reports = 0
	Reports_ClusterConfigMetric Reports = 1
)

// Enum value maps for Reports.
var (
	Reports_name = map[int32]string{
		0: "ClusterConfig",
		1: "ClusterConfigMetric",
	}
	Reports_value = map[string]int32{
		"ClusterConfig":       0,
		"ClusterConfigMetric": 1,
	}
)

func (x Reports) Enum() *Reports {
	p := new(Reports)
	*p = x
	return p
}

func (x Reports) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reports) Descriptor() protoreflect.EnumDescriptor {
	return file_outliers_proto_enumTypes[0].Descriptor()
}

func (Reports) Type() protoreflect.EnumType {
	return &file_outliers_proto_enumTypes[0]
}

func (x Reports) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reports.Descriptor instead.
func (Reports) EnumDescriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{0}
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Name  string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Metric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type OutliersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *OutliersRequest) Reset() {
	*x = OutliersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutliersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutliersRequest) ProtoMessage() {}

func (x *OutliersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutliersRequest.ProtoReflect.Descriptor instead.
func (*OutliersRequest) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{1}
}

func (x *OutliersRequest) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type OutliersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indices []int32 `protobuf:"varint,1,rep,packed,name=indices,proto3" json:"indices,omitempty"`
}

func (x *OutliersResponse) Reset() {
	*x = OutliersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutliersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutliersResponse) ProtoMessage() {}

func (x *OutliersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutliersResponse.ProtoReflect.Descriptor instead.
func (*OutliersResponse) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{2}
}

func (x *OutliersResponse) GetIndices() []int32 {
	if x != nil {
		return x.Indices
	}
	return nil
}

type ReportType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report Reports `protobuf:"varint,1,opt,name=report,proto3,enum=pb.Reports" json:"report,omitempty"`
	// google.protobuf.StringValue filter = 2;
	Param string `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	Start int64  `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	End   int64  `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ReportType) Reset() {
	*x = ReportType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportType) ProtoMessage() {}

func (x *ReportType) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportType.ProtoReflect.Descriptor instead.
func (*ReportType) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{3}
}

func (x *ReportType) GetReport() Reports {
	if x != nil {
		return x.Report
	}
	return Reports_ClusterConfig
}

func (x *ReportType) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *ReportType) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ReportType) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type ReportB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report []byte `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *ReportB) Reset() {
	*x = ReportB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outliers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportB) ProtoMessage() {}

func (x *ReportB) ProtoReflect() protoreflect.Message {
	mi := &file_outliers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportB.ProtoReflect.Descriptor instead.
func (*ReportB) Descriptor() ([]byte, []int) {
	return file_outliers_proto_rawDescGZIP(), []int{4}
}

func (x *ReportB) GetReport() []byte {
	if x != nil {
		return x.Report
	}
	return nil
}

var File_outliers_proto protoreflect.FileDescriptor

var file_outliers_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x4f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x22, 0x6f, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x06, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x22, 0x21, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2a, 0x35, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x11, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x10, 0x01, 0x32, 0x6a, 0x0a, 0x08, 0x4f,
	0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x6c,
	0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27,
	0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x6e, 0x6e, 0x2f, 0x47, 0x6f,
	0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x54, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outliers_proto_rawDescOnce sync.Once
	file_outliers_proto_rawDescData = file_outliers_proto_rawDesc
)

func file_outliers_proto_rawDescGZIP() []byte {
	file_outliers_proto_rawDescOnce.Do(func() {
		file_outliers_proto_rawDescData = protoimpl.X.CompressGZIP(file_outliers_proto_rawDescData)
	})
	return file_outliers_proto_rawDescData
}

var file_outliers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_outliers_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_outliers_proto_goTypes = []interface{}{
	(Reports)(0),                  // 0: pb.Reports
	(*Metric)(nil),                // 1: pb.Metric
	(*OutliersRequest)(nil),       // 2: pb.OutliersRequest
	(*OutliersResponse)(nil),      // 3: pb.OutliersResponse
	(*ReportType)(nil),            // 4: pb.ReportType
	(*ReportB)(nil),               // 5: pb.ReportB
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_outliers_proto_depIdxs = []int32{
	6, // 0: pb.Metric.time:type_name -> google.protobuf.Timestamp
	1, // 1: pb.OutliersRequest.metrics:type_name -> pb.Metric
	0, // 2: pb.ReportType.report:type_name -> pb.Reports
	2, // 3: pb.Outliers.Detect:input_type -> pb.OutliersRequest
	4, // 4: pb.Outliers.Report:input_type -> pb.ReportType
	3, // 5: pb.Outliers.Detect:output_type -> pb.OutliersResponse
	5, // 6: pb.Outliers.Report:output_type -> pb.ReportB
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_outliers_proto_init() }
func file_outliers_proto_init() {
	if File_outliers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outliers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_outliers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutliersRequest); i {
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
		file_outliers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutliersResponse); i {
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
		file_outliers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportType); i {
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
		file_outliers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportB); i {
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
			RawDescriptor: file_outliers_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_outliers_proto_goTypes,
		DependencyIndexes: file_outliers_proto_depIdxs,
		EnumInfos:         file_outliers_proto_enumTypes,
		MessageInfos:      file_outliers_proto_msgTypes,
	}.Build()
	File_outliers_proto = out.File
	file_outliers_proto_rawDesc = nil
	file_outliers_proto_goTypes = nil
	file_outliers_proto_depIdxs = nil
}

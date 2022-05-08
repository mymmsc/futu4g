// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: Qot_GetCapitalDistribution.proto

package Qot_GetCapitalDistribution

import (
	_ "github.com/mymmsc/futu4g/protocol/Common"
	Qot_Common "github.com/mymmsc/futu4g/protocol/Qot_Common"
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

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` //股票
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetCapitalDistribution_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetSecurity() *Qot_Common.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CapitalInBig    *float64 `protobuf:"fixed64,1,req,name=capitalInBig" json:"capitalInBig,omitempty"`       //流入资金额度，大单
	CapitalInMid    *float64 `protobuf:"fixed64,2,req,name=capitalInMid" json:"capitalInMid,omitempty"`       //流入资金额度，中单
	CapitalInSmall  *float64 `protobuf:"fixed64,3,req,name=capitalInSmall" json:"capitalInSmall,omitempty"`   //流入资金额度，小单
	CapitalOutBig   *float64 `protobuf:"fixed64,4,req,name=capitalOutBig" json:"capitalOutBig,omitempty"`     //流出资金额度，大单
	CapitalOutMid   *float64 `protobuf:"fixed64,5,req,name=capitalOutMid" json:"capitalOutMid,omitempty"`     //流出资金额度，中单
	CapitalOutSmall *float64 `protobuf:"fixed64,6,req,name=capitalOutSmall" json:"capitalOutSmall,omitempty"` //流出资金额度，小单
	UpdateTime      *string  `protobuf:"bytes,7,opt,name=updateTime" json:"updateTime,omitempty"`             //更新时间字符串
	UpdateTimestamp *float64 `protobuf:"fixed64,8,opt,name=updateTimestamp" json:"updateTimestamp,omitempty"` //更新时间戳
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetCapitalDistribution_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetCapitalInBig() float64 {
	if x != nil && x.CapitalInBig != nil {
		return *x.CapitalInBig
	}
	return 0
}

func (x *S2C) GetCapitalInMid() float64 {
	if x != nil && x.CapitalInMid != nil {
		return *x.CapitalInMid
	}
	return 0
}

func (x *S2C) GetCapitalInSmall() float64 {
	if x != nil && x.CapitalInSmall != nil {
		return *x.CapitalInSmall
	}
	return 0
}

func (x *S2C) GetCapitalOutBig() float64 {
	if x != nil && x.CapitalOutBig != nil {
		return *x.CapitalOutBig
	}
	return 0
}

func (x *S2C) GetCapitalOutMid() float64 {
	if x != nil && x.CapitalOutMid != nil {
		return *x.CapitalOutMid
	}
	return 0
}

func (x *S2C) GetCapitalOutSmall() float64 {
	if x != nil && x.CapitalOutSmall != nil {
		return *x.CapitalOutSmall
	}
	return 0
}

func (x *S2C) GetUpdateTime() string {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return ""
}

func (x *S2C) GetUpdateTimestamp() float64 {
	if x != nil && x.UpdateTimestamp != nil {
		return *x.UpdateTimestamp
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetCapitalDistribution_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCapitalDistribution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_GetCapitalDistribution_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetCapitalDistribution_proto protoreflect.FileDescriptor

var file_Qot_GetCapitalDistribution_proto_rawDesc = []byte{
	0x0a, 0x20, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f,
	0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0xb5, 0x02, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x42, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x6e,
	0x42, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x6e,
	0x4d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x49, 0x6e, 0x4d, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x49, 0x6e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x02, 0x28, 0x01, 0x52,
	0x0e, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x12,
	0x24, 0x0a, 0x0d, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x42, 0x69, 0x67,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x4f,
	0x75, 0x74, 0x42, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x4f, 0x75, 0x74, 0x4d, 0x69, 0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x61,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x4d, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x18, 0x06,
	0x20, 0x02, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x4f, 0x75, 0x74,
	0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x3c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x63, 0x32,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x8f, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30,
	0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d,
	0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x03,
	0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x51, 0x6f, 0x74, 0x5f,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42,
	0x53, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6d, 0x6d, 0x73, 0x63, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x34, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e,
}

var (
	file_Qot_GetCapitalDistribution_proto_rawDescOnce sync.Once
	file_Qot_GetCapitalDistribution_proto_rawDescData = file_Qot_GetCapitalDistribution_proto_rawDesc
)

func file_Qot_GetCapitalDistribution_proto_rawDescGZIP() []byte {
	file_Qot_GetCapitalDistribution_proto_rawDescOnce.Do(func() {
		file_Qot_GetCapitalDistribution_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetCapitalDistribution_proto_rawDescData)
	})
	return file_Qot_GetCapitalDistribution_proto_rawDescData
}

var file_Qot_GetCapitalDistribution_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Qot_GetCapitalDistribution_proto_goTypes = []interface{}{
	(*C2S)(nil),                 // 0: Qot_GetCapitalDistribution.C2S
	(*S2C)(nil),                 // 1: Qot_GetCapitalDistribution.S2C
	(*Request)(nil),             // 2: Qot_GetCapitalDistribution.Request
	(*Response)(nil),            // 3: Qot_GetCapitalDistribution.Response
	(*Qot_Common.Security)(nil), // 4: Qot_Common.Security
}
var file_Qot_GetCapitalDistribution_proto_depIdxs = []int32{
	4, // 0: Qot_GetCapitalDistribution.C2S.security:type_name -> Qot_Common.Security
	0, // 1: Qot_GetCapitalDistribution.Request.c2s:type_name -> Qot_GetCapitalDistribution.C2S
	1, // 2: Qot_GetCapitalDistribution.Response.s2c:type_name -> Qot_GetCapitalDistribution.S2C
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Qot_GetCapitalDistribution_proto_init() }
func file_Qot_GetCapitalDistribution_proto_init() {
	if File_Qot_GetCapitalDistribution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetCapitalDistribution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetCapitalDistribution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetCapitalDistribution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetCapitalDistribution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_GetCapitalDistribution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetCapitalDistribution_proto_goTypes,
		DependencyIndexes: file_Qot_GetCapitalDistribution_proto_depIdxs,
		MessageInfos:      file_Qot_GetCapitalDistribution_proto_msgTypes,
	}.Build()
	File_Qot_GetCapitalDistribution_proto = out.File
	file_Qot_GetCapitalDistribution_proto_rawDesc = nil
	file_Qot_GetCapitalDistribution_proto_goTypes = nil
	file_Qot_GetCapitalDistribution_proto_depIdxs = nil
}

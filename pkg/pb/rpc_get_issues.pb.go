// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: rpc_get_issues.proto

package issues4missions_backend

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

type GetIssueByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIssueByIDRequest) Reset() {
	*x = GetIssueByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_issues_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIssueByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIssueByIDRequest) ProtoMessage() {}

func (x *GetIssueByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_issues_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIssueByIDRequest.ProtoReflect.Descriptor instead.
func (*GetIssueByIDRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_issues_proto_rawDescGZIP(), []int{0}
}

func (x *GetIssueByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetIssueByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issue *Issue `protobuf:"bytes,1,opt,name=issue,proto3" json:"issue,omitempty"`
}

func (x *GetIssueByIDResponse) Reset() {
	*x = GetIssueByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_issues_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIssueByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIssueByIDResponse) ProtoMessage() {}

func (x *GetIssueByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_issues_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIssueByIDResponse.ProtoReflect.Descriptor instead.
func (*GetIssueByIDResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_issues_proto_rawDescGZIP(), []int{1}
}

func (x *GetIssueByIDResponse) GetIssue() *Issue {
	if x != nil {
		return x.Issue
	}
	return nil
}

type GetIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issue []*Issue `protobuf:"bytes,1,rep,name=issue,proto3" json:"issue,omitempty"`
}

func (x *GetIssuesResponse) Reset() {
	*x = GetIssuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_issues_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIssuesResponse) ProtoMessage() {}

func (x *GetIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_issues_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIssuesResponse.ProtoReflect.Descriptor instead.
func (*GetIssuesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_issues_proto_rawDescGZIP(), []int{2}
}

func (x *GetIssuesResponse) GetIssue() []*Issue {
	if x != nil {
		return x.Issue
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_issues_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_issues_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_rpc_get_issues_proto_rawDescGZIP(), []int{3}
}

var File_rpc_get_issues_proto protoreflect.FileDescriptor

var file_rpc_get_issues_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x70, 0x62,
	0x1a, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x05, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x70, 0x62, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x76, 0x69, 0x64, 0x4b, 0x6f, 0x72, 0x6f,
	0x63, 0x68, 0x69, 0x6b, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x34, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_issues_proto_rawDescOnce sync.Once
	file_rpc_get_issues_proto_rawDescData = file_rpc_get_issues_proto_rawDesc
)

func file_rpc_get_issues_proto_rawDescGZIP() []byte {
	file_rpc_get_issues_proto_rawDescOnce.Do(func() {
		file_rpc_get_issues_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_issues_proto_rawDescData)
	})
	return file_rpc_get_issues_proto_rawDescData
}

var file_rpc_get_issues_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_get_issues_proto_goTypes = []interface{}{
	(*GetIssueByIDRequest)(nil),  // 0: issue_pb.GetIssueByIDRequest
	(*GetIssueByIDResponse)(nil), // 1: issue_pb.GetIssueByIDResponse
	(*GetIssuesResponse)(nil),    // 2: issue_pb.GetIssuesResponse
	(*Empty)(nil),                // 3: issue_pb.Empty
	(*Issue)(nil),                // 4: issue_pb.Issue
}
var file_rpc_get_issues_proto_depIdxs = []int32{
	4, // 0: issue_pb.GetIssueByIDResponse.issue:type_name -> issue_pb.Issue
	4, // 1: issue_pb.GetIssuesResponse.issue:type_name -> issue_pb.Issue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_get_issues_proto_init() }
func file_rpc_get_issues_proto_init() {
	if File_rpc_get_issues_proto != nil {
		return
	}
	file_issue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_issues_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIssueByIDRequest); i {
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
		file_rpc_get_issues_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIssueByIDResponse); i {
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
		file_rpc_get_issues_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIssuesResponse); i {
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
		file_rpc_get_issues_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_rpc_get_issues_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_issues_proto_goTypes,
		DependencyIndexes: file_rpc_get_issues_proto_depIdxs,
		MessageInfos:      file_rpc_get_issues_proto_msgTypes,
	}.Build()
	File_rpc_get_issues_proto = out.File
	file_rpc_get_issues_proto_rawDesc = nil
	file_rpc_get_issues_proto_goTypes = nil
	file_rpc_get_issues_proto_depIdxs = nil
}

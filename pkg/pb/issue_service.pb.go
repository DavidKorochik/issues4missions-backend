// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: issue_service.proto

package issues4missions_backend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_issue_service_proto protoreflect.FileDescriptor

var file_issue_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x1a,
	0x16, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xea, 0x01,
	0x0a, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1c, 0x2e,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x76, 0x69, 0x64, 0x4b, 0x6f,
	0x72, 0x6f, 0x63, 0x68, 0x69, 0x6b, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x34, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_issue_service_proto_goTypes = []interface{}{
	(*CreateIssueRequest)(nil),   // 0: issue_pb.CreateIssueRequest
	(*GetIssueByIDRequest)(nil),  // 1: issue_pb.GetIssueByIDRequest
	(*Empty)(nil),                // 2: issue_pb.Empty
	(*CreateIssueResponse)(nil),  // 3: issue_pb.CreateIssueResponse
	(*GetIssueByIDResponse)(nil), // 4: issue_pb.GetIssueByIDResponse
	(*GetIssuesResponse)(nil),    // 5: issue_pb.GetIssuesResponse
}
var file_issue_service_proto_depIdxs = []int32{
	0, // 0: issue_pb.IssueService.CreateIssue:input_type -> issue_pb.CreateIssueRequest
	1, // 1: issue_pb.IssueService.GetIssueByID:input_type -> issue_pb.GetIssueByIDRequest
	2, // 2: issue_pb.IssueService.GetIssues:input_type -> issue_pb.Empty
	3, // 3: issue_pb.IssueService.CreateIssue:output_type -> issue_pb.CreateIssueResponse
	4, // 4: issue_pb.IssueService.GetIssueByID:output_type -> issue_pb.GetIssueByIDResponse
	5, // 5: issue_pb.IssueService.GetIssues:output_type -> issue_pb.GetIssuesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_issue_service_proto_init() }
func file_issue_service_proto_init() {
	if File_issue_service_proto != nil {
		return
	}
	file_rpc_create_issue_proto_init()
	file_rpc_get_issues_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_issue_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_issue_service_proto_goTypes,
		DependencyIndexes: file_issue_service_proto_depIdxs,
	}.Build()
	File_issue_service_proto = out.File
	file_issue_service_proto_rawDesc = nil
	file_issue_service_proto_goTypes = nil
	file_issue_service_proto_depIdxs = nil
}

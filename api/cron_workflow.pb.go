// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.3
// source: cron_workflow.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CronWorkflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manifest          string             `protobuf:"bytes,2,opt,name=manifest,proto3" json:"manifest,omitempty"`
	WorkflowExecution *WorkflowExecution `protobuf:"bytes,3,opt,name=workflowExecution,proto3" json:"workflowExecution,omitempty"`
	Labels            []*KeyValue        `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *CronWorkflow) Reset() {
	*x = CronWorkflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronWorkflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronWorkflow) ProtoMessage() {}

func (x *CronWorkflow) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronWorkflow.ProtoReflect.Descriptor instead.
func (*CronWorkflow) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *CronWorkflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronWorkflow) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

func (x *CronWorkflow) GetWorkflowExecution() *WorkflowExecution {
	if x != nil {
		return x.WorkflowExecution
	}
	return nil
}

func (x *CronWorkflow) GetLabels() []*KeyValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

type CreateCronWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string        `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CronWorkflow *CronWorkflow `protobuf:"bytes,2,opt,name=cronWorkflow,proto3" json:"cronWorkflow,omitempty"`
}

func (x *CreateCronWorkflowRequest) Reset() {
	*x = CreateCronWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCronWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCronWorkflowRequest) ProtoMessage() {}

func (x *CreateCronWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCronWorkflowRequest.ProtoReflect.Descriptor instead.
func (*CreateCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCronWorkflowRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateCronWorkflowRequest) GetCronWorkflow() *CronWorkflow {
	if x != nil {
		return x.CronWorkflow
	}
	return nil
}

type GetCronWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetCronWorkflowRequest) Reset() {
	*x = GetCronWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCronWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCronWorkflowRequest) ProtoMessage() {}

func (x *GetCronWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCronWorkflowRequest.ProtoReflect.Descriptor instead.
func (*GetCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *GetCronWorkflowRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetCronWorkflowRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type UpdateCronWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string        `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid          string        `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	CronWorkflow *CronWorkflow `protobuf:"bytes,3,opt,name=cronWorkflow,proto3" json:"cronWorkflow,omitempty"`
}

func (x *UpdateCronWorkflowRequest) Reset() {
	*x = UpdateCronWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCronWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCronWorkflowRequest) ProtoMessage() {}

func (x *UpdateCronWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCronWorkflowRequest.ProtoReflect.Descriptor instead.
func (*UpdateCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCronWorkflowRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *UpdateCronWorkflowRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UpdateCronWorkflowRequest) GetCronWorkflow() *CronWorkflow {
	if x != nil {
		return x.CronWorkflow
	}
	return nil
}

type TerminateCronWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid       string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *TerminateCronWorkflowRequest) Reset() {
	*x = TerminateCronWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateCronWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateCronWorkflowRequest) ProtoMessage() {}

func (x *TerminateCronWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateCronWorkflowRequest.ProtoReflect.Descriptor instead.
func (*TerminateCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *TerminateCronWorkflowRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *TerminateCronWorkflowRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ListCronWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace            string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplateName string `protobuf:"bytes,2,opt,name=workflow_template_name,json=workflowTemplateName,proto3" json:"workflow_template_name,omitempty"`
	PageSize             int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Page                 int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListCronWorkflowRequest) Reset() {
	*x = ListCronWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCronWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCronWorkflowRequest) ProtoMessage() {}

func (x *ListCronWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCronWorkflowRequest.ProtoReflect.Descriptor instead.
func (*ListCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{5}
}

func (x *ListCronWorkflowRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListCronWorkflowRequest) GetWorkflowTemplateName() string {
	if x != nil {
		return x.WorkflowTemplateName
	}
	return ""
}

func (x *ListCronWorkflowRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCronWorkflowRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListCronWorkflowsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count         int32           `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	CronWorkflows []*CronWorkflow `protobuf:"bytes,2,rep,name=cronWorkflows,proto3" json:"cronWorkflows,omitempty"`
	Page          int32           `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Pages         int32           `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	TotalCount    int32           `protobuf:"varint,5,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListCronWorkflowsResponse) Reset() {
	*x = ListCronWorkflowsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCronWorkflowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCronWorkflowsResponse) ProtoMessage() {}

func (x *ListCronWorkflowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cron_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCronWorkflowsResponse.ProtoReflect.Descriptor instead.
func (*ListCronWorkflowsResponse) Descriptor() ([]byte, []int) {
	return file_cron_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *ListCronWorkflowsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListCronWorkflowsResponse) GetCronWorkflows() []*CronWorkflow {
	if x != nil {
		return x.CronWorkflows
	}
	return nil
}

func (x *ListCronWorkflowsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCronWorkflowsResponse) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *ListCronWorkflowsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_cron_workflow_proto protoreflect.FileDescriptor

var file_cron_workflow_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x22, 0x70, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x63,
	0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x0c, 0x63, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x22, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x0c, 0x63, 0x72,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x0c, 0x63, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x22, 0x4e, 0x0a, 0x1c, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x9d, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0xb4, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x63, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x0d, 0x63, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x89, 0x06, 0x0a, 0x13, 0x43, 0x72, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x86, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x37, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x63, 0x72,
	0x6f, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x0c, 0x63, 0x72, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x8c, 0x01, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x1a, 0x2d, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x3a, 0x0c, 0x63, 0x72, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x78, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43,
	0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x63,
	0x72, 0x6f, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x7b, 0x75, 0x69,
	0x64, 0x7d, 0x12, 0xc8, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x6f, 0x12, 0x28,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5a, 0x43, 0x12, 0x41, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x94, 0x01,
	0x0a, 0x15, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x2a, 0x38, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cron_workflow_proto_rawDescOnce sync.Once
	file_cron_workflow_proto_rawDescData = file_cron_workflow_proto_rawDesc
)

func file_cron_workflow_proto_rawDescGZIP() []byte {
	file_cron_workflow_proto_rawDescOnce.Do(func() {
		file_cron_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_cron_workflow_proto_rawDescData)
	})
	return file_cron_workflow_proto_rawDescData
}

var file_cron_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cron_workflow_proto_goTypes = []interface{}{
	(*CronWorkflow)(nil),                 // 0: api.CronWorkflow
	(*CreateCronWorkflowRequest)(nil),    // 1: api.CreateCronWorkflowRequest
	(*GetCronWorkflowRequest)(nil),       // 2: api.GetCronWorkflowRequest
	(*UpdateCronWorkflowRequest)(nil),    // 3: api.UpdateCronWorkflowRequest
	(*TerminateCronWorkflowRequest)(nil), // 4: api.TerminateCronWorkflowRequest
	(*ListCronWorkflowRequest)(nil),      // 5: api.ListCronWorkflowRequest
	(*ListCronWorkflowsResponse)(nil),    // 6: api.ListCronWorkflowsResponse
	(*WorkflowExecution)(nil),            // 7: api.WorkflowExecution
	(*KeyValue)(nil),                     // 8: api.KeyValue
	(*empty.Empty)(nil),                  // 9: google.protobuf.Empty
}
var file_cron_workflow_proto_depIdxs = []int32{
	7,  // 0: api.CronWorkflow.workflowExecution:type_name -> api.WorkflowExecution
	8,  // 1: api.CronWorkflow.labels:type_name -> api.KeyValue
	0,  // 2: api.CreateCronWorkflowRequest.cronWorkflow:type_name -> api.CronWorkflow
	0,  // 3: api.UpdateCronWorkflowRequest.cronWorkflow:type_name -> api.CronWorkflow
	0,  // 4: api.ListCronWorkflowsResponse.cronWorkflows:type_name -> api.CronWorkflow
	1,  // 5: api.CronWorkflowService.CreateCronWorkflow:input_type -> api.CreateCronWorkflowRequest
	3,  // 6: api.CronWorkflowService.UpdateCronWorkflow:input_type -> api.UpdateCronWorkflowRequest
	2,  // 7: api.CronWorkflowService.GetCronWorkflow:input_type -> api.GetCronWorkflowRequest
	5,  // 8: api.CronWorkflowService.ListCronWorkflows:input_type -> api.ListCronWorkflowRequest
	4,  // 9: api.CronWorkflowService.TerminateCronWorkflow:input_type -> api.TerminateCronWorkflowRequest
	0,  // 10: api.CronWorkflowService.CreateCronWorkflow:output_type -> api.CronWorkflow
	0,  // 11: api.CronWorkflowService.UpdateCronWorkflow:output_type -> api.CronWorkflow
	0,  // 12: api.CronWorkflowService.GetCronWorkflow:output_type -> api.CronWorkflow
	6,  // 13: api.CronWorkflowService.ListCronWorkflows:output_type -> api.ListCronWorkflowsResponse
	9,  // 14: api.CronWorkflowService.TerminateCronWorkflow:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_cron_workflow_proto_init() }
func file_cron_workflow_proto_init() {
	if File_cron_workflow_proto != nil {
		return
	}
	file_workflow_proto_init()
	file_label_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cron_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronWorkflow); i {
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
		file_cron_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCronWorkflowRequest); i {
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
		file_cron_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCronWorkflowRequest); i {
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
		file_cron_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCronWorkflowRequest); i {
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
		file_cron_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateCronWorkflowRequest); i {
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
		file_cron_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCronWorkflowRequest); i {
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
		file_cron_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCronWorkflowsResponse); i {
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
			RawDescriptor: file_cron_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cron_workflow_proto_goTypes,
		DependencyIndexes: file_cron_workflow_proto_depIdxs,
		MessageInfos:      file_cron_workflow_proto_msgTypes,
	}.Build()
	File_cron_workflow_proto = out.File
	file_cron_workflow_proto_rawDesc = nil
	file_cron_workflow_proto_goTypes = nil
	file_cron_workflow_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CronWorkflowServiceClient is the client API for CronWorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CronWorkflowServiceClient interface {
	CreateCronWorkflow(ctx context.Context, in *CreateCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error)
	UpdateCronWorkflow(ctx context.Context, in *UpdateCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error)
	GetCronWorkflow(ctx context.Context, in *GetCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error)
	ListCronWorkflows(ctx context.Context, in *ListCronWorkflowRequest, opts ...grpc.CallOption) (*ListCronWorkflowsResponse, error)
	TerminateCronWorkflow(ctx context.Context, in *TerminateCronWorkflowRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type cronWorkflowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCronWorkflowServiceClient(cc grpc.ClientConnInterface) CronWorkflowServiceClient {
	return &cronWorkflowServiceClient{cc}
}

func (c *cronWorkflowServiceClient) CreateCronWorkflow(ctx context.Context, in *CreateCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error) {
	out := new(CronWorkflow)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/CreateCronWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronWorkflowServiceClient) UpdateCronWorkflow(ctx context.Context, in *UpdateCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error) {
	out := new(CronWorkflow)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/UpdateCronWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronWorkflowServiceClient) GetCronWorkflow(ctx context.Context, in *GetCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error) {
	out := new(CronWorkflow)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/GetCronWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronWorkflowServiceClient) ListCronWorkflows(ctx context.Context, in *ListCronWorkflowRequest, opts ...grpc.CallOption) (*ListCronWorkflowsResponse, error) {
	out := new(ListCronWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/ListCronWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronWorkflowServiceClient) TerminateCronWorkflow(ctx context.Context, in *TerminateCronWorkflowRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/TerminateCronWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronWorkflowServiceServer is the server API for CronWorkflowService service.
type CronWorkflowServiceServer interface {
	CreateCronWorkflow(context.Context, *CreateCronWorkflowRequest) (*CronWorkflow, error)
	UpdateCronWorkflow(context.Context, *UpdateCronWorkflowRequest) (*CronWorkflow, error)
	GetCronWorkflow(context.Context, *GetCronWorkflowRequest) (*CronWorkflow, error)
	ListCronWorkflows(context.Context, *ListCronWorkflowRequest) (*ListCronWorkflowsResponse, error)
	TerminateCronWorkflow(context.Context, *TerminateCronWorkflowRequest) (*empty.Empty, error)
}

// UnimplementedCronWorkflowServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCronWorkflowServiceServer struct {
}

func (*UnimplementedCronWorkflowServiceServer) CreateCronWorkflow(context.Context, *CreateCronWorkflowRequest) (*CronWorkflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCronWorkflow not implemented")
}
func (*UnimplementedCronWorkflowServiceServer) UpdateCronWorkflow(context.Context, *UpdateCronWorkflowRequest) (*CronWorkflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCronWorkflow not implemented")
}
func (*UnimplementedCronWorkflowServiceServer) GetCronWorkflow(context.Context, *GetCronWorkflowRequest) (*CronWorkflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCronWorkflow not implemented")
}
func (*UnimplementedCronWorkflowServiceServer) ListCronWorkflows(context.Context, *ListCronWorkflowRequest) (*ListCronWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCronWorkflows not implemented")
}
func (*UnimplementedCronWorkflowServiceServer) TerminateCronWorkflow(context.Context, *TerminateCronWorkflowRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateCronWorkflow not implemented")
}

func RegisterCronWorkflowServiceServer(s *grpc.Server, srv CronWorkflowServiceServer) {
	s.RegisterService(&_CronWorkflowService_serviceDesc, srv)
}

func _CronWorkflowService_CreateCronWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).CreateCronWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/CreateCronWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).CreateCronWorkflow(ctx, req.(*CreateCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronWorkflowService_UpdateCronWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).UpdateCronWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/UpdateCronWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).UpdateCronWorkflow(ctx, req.(*UpdateCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronWorkflowService_GetCronWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).GetCronWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/GetCronWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).GetCronWorkflow(ctx, req.(*GetCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronWorkflowService_ListCronWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).ListCronWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/ListCronWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).ListCronWorkflows(ctx, req.(*ListCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronWorkflowService_TerminateCronWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).TerminateCronWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/TerminateCronWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).TerminateCronWorkflow(ctx, req.(*TerminateCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CronWorkflowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CronWorkflowService",
	HandlerType: (*CronWorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCronWorkflow",
			Handler:    _CronWorkflowService_CreateCronWorkflow_Handler,
		},
		{
			MethodName: "UpdateCronWorkflow",
			Handler:    _CronWorkflowService_UpdateCronWorkflow_Handler,
		},
		{
			MethodName: "GetCronWorkflow",
			Handler:    _CronWorkflowService_GetCronWorkflow_Handler,
		},
		{
			MethodName: "ListCronWorkflows",
			Handler:    _CronWorkflowService_ListCronWorkflows_Handler,
		},
		{
			MethodName: "TerminateCronWorkflow",
			Handler:    _CronWorkflowService_TerminateCronWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cron_workflow.proto",
}

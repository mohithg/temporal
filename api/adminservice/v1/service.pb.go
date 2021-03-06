// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/adminservice/v1/service.proto

package adminservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("temporal/server/api/adminservice/v1/service.proto", fileDescriptor_cf5ca5e0c737570d)
}

var fileDescriptor_cf5ca5e0c737570d = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x3f, 0x6f, 0x13, 0x3f,
	0x18, 0xc7, 0xcf, 0xcb, 0x6f, 0xb0, 0x7e, 0xfc, 0x91, 0x8b, 0x90, 0xe8, 0x60, 0x10, 0xec, 0x17,
	0xb5, 0x48, 0x45, 0xb4, 0x40, 0x9b, 0xb6, 0x21, 0x95, 0x68, 0x11, 0x5c, 0x11, 0x48, 0x2c, 0xc8,
	0xbd, 0x3c, 0x4d, 0x4e, 0xbd, 0xc4, 0x87, 0xed, 0xa4, 0x74, 0x82, 0x11, 0x09, 0x09, 0xc1, 0x84,
	0x84, 0xc4, 0xc4, 0xc2, 0xc0, 0x6b, 0x40, 0x62, 0x63, 0xcc, 0xd8, 0x91, 0x5c, 0x96, 0x8e, 0x7d,
	0x09, 0x28, 0x5c, 0xec, 0x5c, 0x93, 0x34, 0xf8, 0x2e, 0xdd, 0x72, 0x92, 0x3f, 0xdf, 0xe7, 0xf3,
	0xe4, 0x7c, 0x8f, 0x8d, 0xe7, 0x14, 0xd4, 0x23, 0x2e, 0x58, 0x58, 0x90, 0x20, 0x5a, 0x20, 0x0a,
	0x2c, 0x0a, 0x0a, 0xac, 0x52, 0x0f, 0x1a, 0xbd, 0xe7, 0xc0, 0x87, 0x42, 0x6b, 0xae, 0xd0, 0xff,
	0xe9, 0x46, 0x82, 0x2b, 0x4e, 0x6e, 0x68, 0xc4, 0x4d, 0x10, 0x97, 0x45, 0x81, 0x9b, 0x46, 0xdc,
	0xd6, 0xdc, 0xec, 0xa2, 0x4d, 0xae, 0x80, 0x97, 0x4d, 0x90, 0xea, 0x85, 0x00, 0x19, 0xf1, 0x86,
	0xec, 0x17, 0x98, 0x3f, 0x9a, 0xc1, 0xff, 0x17, 0x7b, 0x4b, 0xb7, 0x93, 0xa5, 0xe4, 0x3b, 0xc2,
	0x57, 0xd6, 0x41, 0xfa, 0x22, 0xd8, 0x81, 0x67, 0x5c, 0xec, 0xed, 0x86, 0x7c, 0xbf, 0xf4, 0x0a,
	0xfc, 0xa6, 0x0a, 0x78, 0x83, 0x94, 0x5c, 0x0b, 0x21, 0xf7, 0x54, 0xde, 0x4b, 0x24, 0x66, 0xef,
	0x4f, 0x1b, 0x93, 0xf4, 0x70, 0xdd, 0x21, 0x9f, 0x11, 0x9e, 0xd1, 0xeb, 0x36, 0x02, 0xa9, 0xb8,
	0x38, 0xd8, 0xe0, 0x52, 0x91, 0xe5, 0x4c, 0x15, 0x52, 0xa4, 0x56, 0x5c, 0xc9, 0x1f, 0x60, 0xe4,
	0x5e, 0x63, 0xbc, 0x16, 0x72, 0x09, 0xdb, 0x35, 0x26, 0x2a, 0x64, 0xc1, 0x2a, 0x71, 0x00, 0x68,
	0x93, 0x5b, 0x99, 0xb9, 0xb4, 0x80, 0x07, 0x75, 0xde, 0x82, 0x27, 0x4c, 0xee, 0x59, 0x0a, 0x0c,
	0x80, 0x6c, 0x02, 0x69, 0xce, 0x08, 0xfc, 0x44, 0xf8, 0x5a, 0x19, 0xd4, 0xe8, 0x1b, 0x64, 0xfb,
	0xfd, 0xbf, 0xec, 0xe9, 0x3c, 0xd9, 0xb4, 0xca, 0xff, 0x57, 0x8c, 0xb6, 0xdd, 0x3a, 0xa3, 0x34,
	0xd3, 0xc3, 0x57, 0x84, 0x2f, 0x97, 0x41, 0x79, 0x10, 0x85, 0x81, 0xcf, 0x7a, 0x0b, 0xb7, 0x40,
	0x4a, 0x56, 0x05, 0x49, 0x56, 0x6d, 0x6b, 0x8d, 0x81, 0xb5, 0xef, 0xda, 0x54, 0x19, 0xc6, 0xf2,
	0x07, 0xc2, 0x57, 0xcb, 0xa0, 0x1e, 0xb2, 0x3a, 0xc8, 0x88, 0xf9, 0x30, 0x4e, 0xf7, 0x81, 0x6d,
	0xa9, 0x49, 0x29, 0xda, 0x7b, 0xf3, 0x6c, 0xc2, 0x4c, 0x03, 0xbd, 0xc1, 0x53, 0x06, 0xb5, 0xbe,
	0xf9, 0x78, 0x9c, 0x7a, 0xc9, 0xb6, 0xda, 0x78, 0x3e, 0xdb, 0xe0, 0x99, 0x10, 0x63, 0x74, 0xdf,
	0x22, 0x7c, 0xce, 0x03, 0x16, 0x45, 0xe1, 0x41, 0xa9, 0x05, 0x0d, 0x25, 0xc9, 0x6d, 0xcb, 0xcf,
	0x24, 0xc5, 0x68, 0xad, 0xc5, 0x3c, 0xa8, 0x51, 0xf9, 0x84, 0x30, 0x29, 0x56, 0x2a, 0xdb, 0xc0,
	0x84, 0x5f, 0x2b, 0x2a, 0x25, 0x82, 0x9d, 0xa6, 0x02, 0x72, 0xcf, 0x2a, 0x74, 0x14, 0xd4, 0x52,
	0xcb, 0xb9, 0x79, 0x63, 0xf6, 0x1e, 0xe1, 0x0b, 0x7a, 0x44, 0xae, 0x85, 0x4d, 0xa9, 0x40, 0x90,
	0xa5, 0x4c, 0x83, 0xb5, 0x4f, 0x69, 0xa7, 0x3b, 0xf9, 0x60, 0x23, 0xf4, 0x0e, 0xe1, 0xf3, 0xc9,
	0xdb, 0x35, 0x3b, 0x6b, 0x31, 0xc3, 0x96, 0x18, 0xde, 0x4e, 0x4b, 0xb9, 0x58, 0x63, 0xf3, 0x11,
	0xe1, 0x8b, 0x8f, 0x9a, 0xa2, 0x0a, 0x69, 0x1f, 0xbb, 0x16, 0x87, 0x31, 0x6d, 0x74, 0x37, 0x27,
	0x7d, 0xc2, 0x69, 0x0b, 0x72, 0x39, 0x0d, 0x63, 0xd9, 0x9c, 0x46, 0x69, 0xe3, 0xf4, 0x05, 0xe1,
	0x4b, 0x1e, 0xec, 0x0a, 0x90, 0x35, 0x3d, 0xb4, 0x7b, 0xe7, 0x8c, 0x24, 0x2b, 0x96, 0xdf, 0xcd,
	0x28, 0xaa, 0xdd, 0x8a, 0x53, 0x24, 0x9c, 0x38, 0x21, 0x3c, 0x90, 0xd0, 0xa8, 0xa4, 0x66, 0x46,
	0x62, 0xb8, 0x6a, 0x99, 0x3f, 0x0e, 0xce, 0x76, 0x42, 0x9c, 0x96, 0xa1, 0x2d, 0x57, 0xc3, 0x76,
	0x87, 0x3a, 0x87, 0x1d, 0xea, 0x1c, 0x77, 0x28, 0x7a, 0x13, 0x53, 0xf4, 0x2d, 0xa6, 0xe8, 0x57,
	0x4c, 0x51, 0x3b, 0xa6, 0xe8, 0x77, 0x4c, 0xd1, 0x51, 0x4c, 0x9d, 0xe3, 0x98, 0xa2, 0x0f, 0x5d,
	0xea, 0xb4, 0xbb, 0xd4, 0x39, 0xec, 0x52, 0xe7, 0xf9, 0x42, 0x95, 0x0f, 0xca, 0x07, 0x7c, 0xc2,
	0x15, 0x73, 0x29, 0xfd, 0xbc, 0xf3, 0xdf, 0xdf, 0xfb, 0xe5, 0xcd, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0xd9, 0x39, 0x56, 0xf5, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	// DescribeWorkflowExecution returns information about the internal states of workflow execution.
	DescribeWorkflowExecution(ctx context.Context, in *DescribeWorkflowExecutionRequest, opts ...grpc.CallOption) (*DescribeWorkflowExecutionResponse, error)
	// DescribeHistoryHost returns information about the internal states of a history host
	DescribeHistoryHost(ctx context.Context, in *DescribeHistoryHostRequest, opts ...grpc.CallOption) (*DescribeHistoryHostResponse, error)
	CloseShard(ctx context.Context, in *CloseShardRequest, opts ...grpc.CallOption) (*CloseShardResponse, error)
	RemoveTask(ctx context.Context, in *RemoveTaskRequest, opts ...grpc.CallOption) (*RemoveTaskResponse, error)
	// Returns the raw history of specified workflow execution.  It fails with 'NotFound' if specified workflow
	// execution in unknown to the service.
	// StartEventId defines the beginning of the event to fetch. The first event is inclusive.
	// EndEventId and EndEventVersion defines the end of the event to fetch. The end event is exclusive.
	GetWorkflowExecutionRawHistoryV2(ctx context.Context, in *GetWorkflowExecutionRawHistoryV2Request, opts ...grpc.CallOption) (*GetWorkflowExecutionRawHistoryV2Response, error)
	// GetReplicationMessages returns new replication tasks since the read level provided in the token.
	GetReplicationMessages(ctx context.Context, in *GetReplicationMessagesRequest, opts ...grpc.CallOption) (*GetReplicationMessagesResponse, error)
	// GetNamespaceReplicationMessages returns new namespace replication tasks since last retrieved task Id.
	GetNamespaceReplicationMessages(ctx context.Context, in *GetNamespaceReplicationMessagesRequest, opts ...grpc.CallOption) (*GetNamespaceReplicationMessagesResponse, error)
	// GetDLQReplicationMessages return replication messages based on DLQ info.
	GetDLQReplicationMessages(ctx context.Context, in *GetDLQReplicationMessagesRequest, opts ...grpc.CallOption) (*GetDLQReplicationMessagesResponse, error)
	// ReapplyEvents applies stale events to the current workflow and current run.
	ReapplyEvents(ctx context.Context, in *ReapplyEventsRequest, opts ...grpc.CallOption) (*ReapplyEventsResponse, error)
	// AddSearchAttribute whitelist search attribute in request.
	AddSearchAttribute(ctx context.Context, in *AddSearchAttributeRequest, opts ...grpc.CallOption) (*AddSearchAttributeResponse, error)
	// DescribeCluster returns information about Temporal cluster.
	DescribeCluster(ctx context.Context, in *DescribeClusterRequest, opts ...grpc.CallOption) (*DescribeClusterResponse, error)
	// GetDLQMessages returns messages from DLQ.
	GetDLQMessages(ctx context.Context, in *GetDLQMessagesRequest, opts ...grpc.CallOption) (*GetDLQMessagesResponse, error)
	// PurgeDLQMessages purges messages from DLQ.
	PurgeDLQMessages(ctx context.Context, in *PurgeDLQMessagesRequest, opts ...grpc.CallOption) (*PurgeDLQMessagesResponse, error)
	// MergeDLQMessages merges messages from DLQ.
	MergeDLQMessages(ctx context.Context, in *MergeDLQMessagesRequest, opts ...grpc.CallOption) (*MergeDLQMessagesResponse, error)
	// RefreshWorkflowTasks refreshes all tasks of a workflow.
	RefreshWorkflowTasks(ctx context.Context, in *RefreshWorkflowTasksRequest, opts ...grpc.CallOption) (*RefreshWorkflowTasksResponse, error)
	// ResendReplicationTasks requests replication tasks from remote cluster and apply tasks to current cluster.
	ResendReplicationTasks(ctx context.Context, in *ResendReplicationTasksRequest, opts ...grpc.CallOption) (*ResendReplicationTasksResponse, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) DescribeWorkflowExecution(ctx context.Context, in *DescribeWorkflowExecutionRequest, opts ...grpc.CallOption) (*DescribeWorkflowExecutionResponse, error) {
	out := new(DescribeWorkflowExecutionResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/DescribeWorkflowExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DescribeHistoryHost(ctx context.Context, in *DescribeHistoryHostRequest, opts ...grpc.CallOption) (*DescribeHistoryHostResponse, error) {
	out := new(DescribeHistoryHostResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/DescribeHistoryHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CloseShard(ctx context.Context, in *CloseShardRequest, opts ...grpc.CallOption) (*CloseShardResponse, error) {
	out := new(CloseShardResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/CloseShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RemoveTask(ctx context.Context, in *RemoveTaskRequest, opts ...grpc.CallOption) (*RemoveTaskResponse, error) {
	out := new(RemoveTaskResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/RemoveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetWorkflowExecutionRawHistoryV2(ctx context.Context, in *GetWorkflowExecutionRawHistoryV2Request, opts ...grpc.CallOption) (*GetWorkflowExecutionRawHistoryV2Response, error) {
	out := new(GetWorkflowExecutionRawHistoryV2Response)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/GetWorkflowExecutionRawHistoryV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetReplicationMessages(ctx context.Context, in *GetReplicationMessagesRequest, opts ...grpc.CallOption) (*GetReplicationMessagesResponse, error) {
	out := new(GetReplicationMessagesResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/GetReplicationMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetNamespaceReplicationMessages(ctx context.Context, in *GetNamespaceReplicationMessagesRequest, opts ...grpc.CallOption) (*GetNamespaceReplicationMessagesResponse, error) {
	out := new(GetNamespaceReplicationMessagesResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/GetNamespaceReplicationMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetDLQReplicationMessages(ctx context.Context, in *GetDLQReplicationMessagesRequest, opts ...grpc.CallOption) (*GetDLQReplicationMessagesResponse, error) {
	out := new(GetDLQReplicationMessagesResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/GetDLQReplicationMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ReapplyEvents(ctx context.Context, in *ReapplyEventsRequest, opts ...grpc.CallOption) (*ReapplyEventsResponse, error) {
	out := new(ReapplyEventsResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/ReapplyEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddSearchAttribute(ctx context.Context, in *AddSearchAttributeRequest, opts ...grpc.CallOption) (*AddSearchAttributeResponse, error) {
	out := new(AddSearchAttributeResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/AddSearchAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DescribeCluster(ctx context.Context, in *DescribeClusterRequest, opts ...grpc.CallOption) (*DescribeClusterResponse, error) {
	out := new(DescribeClusterResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/DescribeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetDLQMessages(ctx context.Context, in *GetDLQMessagesRequest, opts ...grpc.CallOption) (*GetDLQMessagesResponse, error) {
	out := new(GetDLQMessagesResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/GetDLQMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) PurgeDLQMessages(ctx context.Context, in *PurgeDLQMessagesRequest, opts ...grpc.CallOption) (*PurgeDLQMessagesResponse, error) {
	out := new(PurgeDLQMessagesResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/PurgeDLQMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) MergeDLQMessages(ctx context.Context, in *MergeDLQMessagesRequest, opts ...grpc.CallOption) (*MergeDLQMessagesResponse, error) {
	out := new(MergeDLQMessagesResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/MergeDLQMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RefreshWorkflowTasks(ctx context.Context, in *RefreshWorkflowTasksRequest, opts ...grpc.CallOption) (*RefreshWorkflowTasksResponse, error) {
	out := new(RefreshWorkflowTasksResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/RefreshWorkflowTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ResendReplicationTasks(ctx context.Context, in *ResendReplicationTasksRequest, opts ...grpc.CallOption) (*ResendReplicationTasksResponse, error) {
	out := new(ResendReplicationTasksResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.adminservice.v1.AdminService/ResendReplicationTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	// DescribeWorkflowExecution returns information about the internal states of workflow execution.
	DescribeWorkflowExecution(context.Context, *DescribeWorkflowExecutionRequest) (*DescribeWorkflowExecutionResponse, error)
	// DescribeHistoryHost returns information about the internal states of a history host
	DescribeHistoryHost(context.Context, *DescribeHistoryHostRequest) (*DescribeHistoryHostResponse, error)
	CloseShard(context.Context, *CloseShardRequest) (*CloseShardResponse, error)
	RemoveTask(context.Context, *RemoveTaskRequest) (*RemoveTaskResponse, error)
	// Returns the raw history of specified workflow execution.  It fails with 'NotFound' if specified workflow
	// execution in unknown to the service.
	// StartEventId defines the beginning of the event to fetch. The first event is inclusive.
	// EndEventId and EndEventVersion defines the end of the event to fetch. The end event is exclusive.
	GetWorkflowExecutionRawHistoryV2(context.Context, *GetWorkflowExecutionRawHistoryV2Request) (*GetWorkflowExecutionRawHistoryV2Response, error)
	// GetReplicationMessages returns new replication tasks since the read level provided in the token.
	GetReplicationMessages(context.Context, *GetReplicationMessagesRequest) (*GetReplicationMessagesResponse, error)
	// GetNamespaceReplicationMessages returns new namespace replication tasks since last retrieved task Id.
	GetNamespaceReplicationMessages(context.Context, *GetNamespaceReplicationMessagesRequest) (*GetNamespaceReplicationMessagesResponse, error)
	// GetDLQReplicationMessages return replication messages based on DLQ info.
	GetDLQReplicationMessages(context.Context, *GetDLQReplicationMessagesRequest) (*GetDLQReplicationMessagesResponse, error)
	// ReapplyEvents applies stale events to the current workflow and current run.
	ReapplyEvents(context.Context, *ReapplyEventsRequest) (*ReapplyEventsResponse, error)
	// AddSearchAttribute whitelist search attribute in request.
	AddSearchAttribute(context.Context, *AddSearchAttributeRequest) (*AddSearchAttributeResponse, error)
	// DescribeCluster returns information about Temporal cluster.
	DescribeCluster(context.Context, *DescribeClusterRequest) (*DescribeClusterResponse, error)
	// GetDLQMessages returns messages from DLQ.
	GetDLQMessages(context.Context, *GetDLQMessagesRequest) (*GetDLQMessagesResponse, error)
	// PurgeDLQMessages purges messages from DLQ.
	PurgeDLQMessages(context.Context, *PurgeDLQMessagesRequest) (*PurgeDLQMessagesResponse, error)
	// MergeDLQMessages merges messages from DLQ.
	MergeDLQMessages(context.Context, *MergeDLQMessagesRequest) (*MergeDLQMessagesResponse, error)
	// RefreshWorkflowTasks refreshes all tasks of a workflow.
	RefreshWorkflowTasks(context.Context, *RefreshWorkflowTasksRequest) (*RefreshWorkflowTasksResponse, error)
	// ResendReplicationTasks requests replication tasks from remote cluster and apply tasks to current cluster.
	ResendReplicationTasks(context.Context, *ResendReplicationTasksRequest) (*ResendReplicationTasksResponse, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) DescribeWorkflowExecution(ctx context.Context, req *DescribeWorkflowExecutionRequest) (*DescribeWorkflowExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeWorkflowExecution not implemented")
}
func (*UnimplementedAdminServiceServer) DescribeHistoryHost(ctx context.Context, req *DescribeHistoryHostRequest) (*DescribeHistoryHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeHistoryHost not implemented")
}
func (*UnimplementedAdminServiceServer) CloseShard(ctx context.Context, req *CloseShardRequest) (*CloseShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseShard not implemented")
}
func (*UnimplementedAdminServiceServer) RemoveTask(ctx context.Context, req *RemoveTaskRequest) (*RemoveTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTask not implemented")
}
func (*UnimplementedAdminServiceServer) GetWorkflowExecutionRawHistoryV2(ctx context.Context, req *GetWorkflowExecutionRawHistoryV2Request) (*GetWorkflowExecutionRawHistoryV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowExecutionRawHistoryV2 not implemented")
}
func (*UnimplementedAdminServiceServer) GetReplicationMessages(ctx context.Context, req *GetReplicationMessagesRequest) (*GetReplicationMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplicationMessages not implemented")
}
func (*UnimplementedAdminServiceServer) GetNamespaceReplicationMessages(ctx context.Context, req *GetNamespaceReplicationMessagesRequest) (*GetNamespaceReplicationMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespaceReplicationMessages not implemented")
}
func (*UnimplementedAdminServiceServer) GetDLQReplicationMessages(ctx context.Context, req *GetDLQReplicationMessagesRequest) (*GetDLQReplicationMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDLQReplicationMessages not implemented")
}
func (*UnimplementedAdminServiceServer) ReapplyEvents(ctx context.Context, req *ReapplyEventsRequest) (*ReapplyEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReapplyEvents not implemented")
}
func (*UnimplementedAdminServiceServer) AddSearchAttribute(ctx context.Context, req *AddSearchAttributeRequest) (*AddSearchAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSearchAttribute not implemented")
}
func (*UnimplementedAdminServiceServer) DescribeCluster(ctx context.Context, req *DescribeClusterRequest) (*DescribeClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCluster not implemented")
}
func (*UnimplementedAdminServiceServer) GetDLQMessages(ctx context.Context, req *GetDLQMessagesRequest) (*GetDLQMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDLQMessages not implemented")
}
func (*UnimplementedAdminServiceServer) PurgeDLQMessages(ctx context.Context, req *PurgeDLQMessagesRequest) (*PurgeDLQMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeDLQMessages not implemented")
}
func (*UnimplementedAdminServiceServer) MergeDLQMessages(ctx context.Context, req *MergeDLQMessagesRequest) (*MergeDLQMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeDLQMessages not implemented")
}
func (*UnimplementedAdminServiceServer) RefreshWorkflowTasks(ctx context.Context, req *RefreshWorkflowTasksRequest) (*RefreshWorkflowTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshWorkflowTasks not implemented")
}
func (*UnimplementedAdminServiceServer) ResendReplicationTasks(ctx context.Context, req *ResendReplicationTasksRequest) (*ResendReplicationTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendReplicationTasks not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_DescribeWorkflowExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeWorkflowExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DescribeWorkflowExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/DescribeWorkflowExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DescribeWorkflowExecution(ctx, req.(*DescribeWorkflowExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DescribeHistoryHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeHistoryHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DescribeHistoryHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/DescribeHistoryHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DescribeHistoryHost(ctx, req.(*DescribeHistoryHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CloseShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CloseShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/CloseShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CloseShard(ctx, req.(*CloseShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RemoveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RemoveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/RemoveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RemoveTask(ctx, req.(*RemoveTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetWorkflowExecutionRawHistoryV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowExecutionRawHistoryV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetWorkflowExecutionRawHistoryV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/GetWorkflowExecutionRawHistoryV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetWorkflowExecutionRawHistoryV2(ctx, req.(*GetWorkflowExecutionRawHistoryV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetReplicationMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplicationMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetReplicationMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/GetReplicationMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetReplicationMessages(ctx, req.(*GetReplicationMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetNamespaceReplicationMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceReplicationMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetNamespaceReplicationMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/GetNamespaceReplicationMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetNamespaceReplicationMessages(ctx, req.(*GetNamespaceReplicationMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetDLQReplicationMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDLQReplicationMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetDLQReplicationMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/GetDLQReplicationMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetDLQReplicationMessages(ctx, req.(*GetDLQReplicationMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ReapplyEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReapplyEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ReapplyEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/ReapplyEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ReapplyEvents(ctx, req.(*ReapplyEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddSearchAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSearchAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddSearchAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/AddSearchAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddSearchAttribute(ctx, req.(*AddSearchAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DescribeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DescribeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/DescribeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DescribeCluster(ctx, req.(*DescribeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetDLQMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDLQMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetDLQMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/GetDLQMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetDLQMessages(ctx, req.(*GetDLQMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_PurgeDLQMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeDLQMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PurgeDLQMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/PurgeDLQMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PurgeDLQMessages(ctx, req.(*PurgeDLQMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_MergeDLQMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeDLQMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).MergeDLQMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/MergeDLQMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).MergeDLQMessages(ctx, req.(*MergeDLQMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RefreshWorkflowTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshWorkflowTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RefreshWorkflowTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/RefreshWorkflowTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RefreshWorkflowTasks(ctx, req.(*RefreshWorkflowTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ResendReplicationTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendReplicationTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ResendReplicationTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.adminservice.v1.AdminService/ResendReplicationTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ResendReplicationTasks(ctx, req.(*ResendReplicationTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.server.api.adminservice.v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeWorkflowExecution",
			Handler:    _AdminService_DescribeWorkflowExecution_Handler,
		},
		{
			MethodName: "DescribeHistoryHost",
			Handler:    _AdminService_DescribeHistoryHost_Handler,
		},
		{
			MethodName: "CloseShard",
			Handler:    _AdminService_CloseShard_Handler,
		},
		{
			MethodName: "RemoveTask",
			Handler:    _AdminService_RemoveTask_Handler,
		},
		{
			MethodName: "GetWorkflowExecutionRawHistoryV2",
			Handler:    _AdminService_GetWorkflowExecutionRawHistoryV2_Handler,
		},
		{
			MethodName: "GetReplicationMessages",
			Handler:    _AdminService_GetReplicationMessages_Handler,
		},
		{
			MethodName: "GetNamespaceReplicationMessages",
			Handler:    _AdminService_GetNamespaceReplicationMessages_Handler,
		},
		{
			MethodName: "GetDLQReplicationMessages",
			Handler:    _AdminService_GetDLQReplicationMessages_Handler,
		},
		{
			MethodName: "ReapplyEvents",
			Handler:    _AdminService_ReapplyEvents_Handler,
		},
		{
			MethodName: "AddSearchAttribute",
			Handler:    _AdminService_AddSearchAttribute_Handler,
		},
		{
			MethodName: "DescribeCluster",
			Handler:    _AdminService_DescribeCluster_Handler,
		},
		{
			MethodName: "GetDLQMessages",
			Handler:    _AdminService_GetDLQMessages_Handler,
		},
		{
			MethodName: "PurgeDLQMessages",
			Handler:    _AdminService_PurgeDLQMessages_Handler,
		},
		{
			MethodName: "MergeDLQMessages",
			Handler:    _AdminService_MergeDLQMessages_Handler,
		},
		{
			MethodName: "RefreshWorkflowTasks",
			Handler:    _AdminService_RefreshWorkflowTasks_Handler,
		},
		{
			MethodName: "ResendReplicationTasks",
			Handler:    _AdminService_ResendReplicationTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/server/api/adminservice/v1/service.proto",
}

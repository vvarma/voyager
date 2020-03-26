// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/voyager.proto

package voyager

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Meta struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{0}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProbeRequest struct {
	ResponseSize         int64    `protobuf:"varint,1,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	ResponseCode         int32    `protobuf:"varint,2,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	ShouldKillServer     bool     `protobuf:"varint,3,opt,name=should_kill_server,json=shouldKillServer,proto3" json:"should_kill_server,omitempty"`
	Meta                 *Meta    `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeRequest) Reset()         { *m = ProbeRequest{} }
func (m *ProbeRequest) String() string { return proto.CompactTextString(m) }
func (*ProbeRequest) ProtoMessage()    {}
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{1}
}

func (m *ProbeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeRequest.Unmarshal(m, b)
}
func (m *ProbeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeRequest.Marshal(b, m, deterministic)
}
func (m *ProbeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeRequest.Merge(m, src)
}
func (m *ProbeRequest) XXX_Size() int {
	return xxx_messageInfo_ProbeRequest.Size(m)
}
func (m *ProbeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeRequest proto.InternalMessageInfo

func (m *ProbeRequest) GetResponseSize() int64 {
	if m != nil {
		return m.ResponseSize
	}
	return 0
}

func (m *ProbeRequest) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return 0
}

func (m *ProbeRequest) GetShouldKillServer() bool {
	if m != nil {
		return m.ShouldKillServer
	}
	return false
}

func (m *ProbeRequest) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ProbeResponse struct {
	NodeName             string               `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PodName              string               `protobuf:"bytes,2,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	ServerTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	ProcessDuration      *duration.Duration   `protobuf:"bytes,4,opt,name=process_duration,json=processDuration,proto3" json:"process_duration,omitempty"`
	Meta                 *Meta                `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProbeResponse) Reset()         { *m = ProbeResponse{} }
func (m *ProbeResponse) String() string { return proto.CompactTextString(m) }
func (*ProbeResponse) ProtoMessage()    {}
func (*ProbeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{2}
}

func (m *ProbeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeResponse.Unmarshal(m, b)
}
func (m *ProbeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeResponse.Marshal(b, m, deterministic)
}
func (m *ProbeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeResponse.Merge(m, src)
}
func (m *ProbeResponse) XXX_Size() int {
	return xxx_messageInfo_ProbeResponse.Size(m)
}
func (m *ProbeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeResponse proto.InternalMessageInfo

func (m *ProbeResponse) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ProbeResponse) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *ProbeResponse) GetServerTime() *timestamp.Timestamp {
	if m != nil {
		return m.ServerTime
	}
	return nil
}

func (m *ProbeResponse) GetProcessDuration() *duration.Duration {
	if m != nil {
		return m.ProcessDuration
	}
	return nil
}

func (m *ProbeResponse) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type InitiateNetworkProbe struct {
	RequestSize          int64    `protobuf:"varint,1,opt,name=request_size,json=requestSize,proto3" json:"request_size,omitempty"`
	ResponseSize         int64    `protobuf:"varint,2,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	HostPorts            []string `protobuf:"bytes,3,rep,name=host_ports,json=hostPorts,proto3" json:"host_ports,omitempty"`
	NumberRequests       int64    `protobuf:"varint,4,opt,name=number_requests,json=numberRequests,proto3" json:"number_requests,omitempty"`
	NumThreads           int32    `protobuf:"varint,5,opt,name=num_threads,json=numThreads,proto3" json:"num_threads,omitempty"`
	ErrorRate            float32  `protobuf:"fixed32,6,opt,name=error_rate,json=errorRate,proto3" json:"error_rate,omitempty"`
	PanicRate            float32  `protobuf:"fixed32,7,opt,name=panic_rate,json=panicRate,proto3" json:"panic_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitiateNetworkProbe) Reset()         { *m = InitiateNetworkProbe{} }
func (m *InitiateNetworkProbe) String() string { return proto.CompactTextString(m) }
func (*InitiateNetworkProbe) ProtoMessage()    {}
func (*InitiateNetworkProbe) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{3}
}

func (m *InitiateNetworkProbe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitiateNetworkProbe.Unmarshal(m, b)
}
func (m *InitiateNetworkProbe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitiateNetworkProbe.Marshal(b, m, deterministic)
}
func (m *InitiateNetworkProbe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitiateNetworkProbe.Merge(m, src)
}
func (m *InitiateNetworkProbe) XXX_Size() int {
	return xxx_messageInfo_InitiateNetworkProbe.Size(m)
}
func (m *InitiateNetworkProbe) XXX_DiscardUnknown() {
	xxx_messageInfo_InitiateNetworkProbe.DiscardUnknown(m)
}

var xxx_messageInfo_InitiateNetworkProbe proto.InternalMessageInfo

func (m *InitiateNetworkProbe) GetRequestSize() int64 {
	if m != nil {
		return m.RequestSize
	}
	return 0
}

func (m *InitiateNetworkProbe) GetResponseSize() int64 {
	if m != nil {
		return m.ResponseSize
	}
	return 0
}

func (m *InitiateNetworkProbe) GetHostPorts() []string {
	if m != nil {
		return m.HostPorts
	}
	return nil
}

func (m *InitiateNetworkProbe) GetNumberRequests() int64 {
	if m != nil {
		return m.NumberRequests
	}
	return 0
}

func (m *InitiateNetworkProbe) GetNumThreads() int32 {
	if m != nil {
		return m.NumThreads
	}
	return 0
}

func (m *InitiateNetworkProbe) GetErrorRate() float32 {
	if m != nil {
		return m.ErrorRate
	}
	return 0
}

func (m *InitiateNetworkProbe) GetPanicRate() float32 {
	if m != nil {
		return m.PanicRate
	}
	return 0
}

type NetworkProbeStats struct {
	Stats                *NetworkProbeStats_Stats       `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
	HostSplits           []*NetworkProbeStats_HostSplit `protobuf:"bytes,5,rep,name=host_splits,json=hostSplits,proto3" json:"host_splits,omitempty"`
	CreatedAt            *timestamp.Timestamp           `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *NetworkProbeStats) Reset()         { *m = NetworkProbeStats{} }
func (m *NetworkProbeStats) String() string { return proto.CompactTextString(m) }
func (*NetworkProbeStats) ProtoMessage()    {}
func (*NetworkProbeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{4}
}

func (m *NetworkProbeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProbeStats.Unmarshal(m, b)
}
func (m *NetworkProbeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProbeStats.Marshal(b, m, deterministic)
}
func (m *NetworkProbeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProbeStats.Merge(m, src)
}
func (m *NetworkProbeStats) XXX_Size() int {
	return xxx_messageInfo_NetworkProbeStats.Size(m)
}
func (m *NetworkProbeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProbeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProbeStats proto.InternalMessageInfo

func (m *NetworkProbeStats) GetStats() *NetworkProbeStats_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *NetworkProbeStats) GetHostSplits() []*NetworkProbeStats_HostSplit {
	if m != nil {
		return m.HostSplits
	}
	return nil
}

func (m *NetworkProbeStats) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type NetworkProbeStats_Stats struct {
	NumRequestsCompleted int64                              `protobuf:"varint,1,opt,name=num_requests_completed,json=numRequestsCompleted,proto3" json:"num_requests_completed,omitempty"`
	NumSuccess           int64                              `protobuf:"varint,2,opt,name=num_success,json=numSuccess,proto3" json:"num_success,omitempty"`
	NumFailure           int64                              `protobuf:"varint,3,opt,name=num_failure,json=numFailure,proto3" json:"num_failure,omitempty"`
	Latencies            *NetworkProbeStats_Stats_Latencies `protobuf:"bytes,5,opt,name=latencies,proto3" json:"latencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *NetworkProbeStats_Stats) Reset()         { *m = NetworkProbeStats_Stats{} }
func (m *NetworkProbeStats_Stats) String() string { return proto.CompactTextString(m) }
func (*NetworkProbeStats_Stats) ProtoMessage()    {}
func (*NetworkProbeStats_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{4, 0}
}

func (m *NetworkProbeStats_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProbeStats_Stats.Unmarshal(m, b)
}
func (m *NetworkProbeStats_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProbeStats_Stats.Marshal(b, m, deterministic)
}
func (m *NetworkProbeStats_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProbeStats_Stats.Merge(m, src)
}
func (m *NetworkProbeStats_Stats) XXX_Size() int {
	return xxx_messageInfo_NetworkProbeStats_Stats.Size(m)
}
func (m *NetworkProbeStats_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProbeStats_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProbeStats_Stats proto.InternalMessageInfo

func (m *NetworkProbeStats_Stats) GetNumRequestsCompleted() int64 {
	if m != nil {
		return m.NumRequestsCompleted
	}
	return 0
}

func (m *NetworkProbeStats_Stats) GetNumSuccess() int64 {
	if m != nil {
		return m.NumSuccess
	}
	return 0
}

func (m *NetworkProbeStats_Stats) GetNumFailure() int64 {
	if m != nil {
		return m.NumFailure
	}
	return 0
}

func (m *NetworkProbeStats_Stats) GetLatencies() *NetworkProbeStats_Stats_Latencies {
	if m != nil {
		return m.Latencies
	}
	return nil
}

type NetworkProbeStats_Stats_Latencies struct {
	P70                  *duration.Duration `protobuf:"bytes,1,opt,name=p70,proto3" json:"p70,omitempty"`
	P90                  *duration.Duration `protobuf:"bytes,2,opt,name=p90,proto3" json:"p90,omitempty"`
	P99                  *duration.Duration `protobuf:"bytes,3,opt,name=p99,proto3" json:"p99,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NetworkProbeStats_Stats_Latencies) Reset()         { *m = NetworkProbeStats_Stats_Latencies{} }
func (m *NetworkProbeStats_Stats_Latencies) String() string { return proto.CompactTextString(m) }
func (*NetworkProbeStats_Stats_Latencies) ProtoMessage()    {}
func (*NetworkProbeStats_Stats_Latencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{4, 0, 0}
}

func (m *NetworkProbeStats_Stats_Latencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProbeStats_Stats_Latencies.Unmarshal(m, b)
}
func (m *NetworkProbeStats_Stats_Latencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProbeStats_Stats_Latencies.Marshal(b, m, deterministic)
}
func (m *NetworkProbeStats_Stats_Latencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProbeStats_Stats_Latencies.Merge(m, src)
}
func (m *NetworkProbeStats_Stats_Latencies) XXX_Size() int {
	return xxx_messageInfo_NetworkProbeStats_Stats_Latencies.Size(m)
}
func (m *NetworkProbeStats_Stats_Latencies) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProbeStats_Stats_Latencies.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProbeStats_Stats_Latencies proto.InternalMessageInfo

func (m *NetworkProbeStats_Stats_Latencies) GetP70() *duration.Duration {
	if m != nil {
		return m.P70
	}
	return nil
}

func (m *NetworkProbeStats_Stats_Latencies) GetP90() *duration.Duration {
	if m != nil {
		return m.P90
	}
	return nil
}

func (m *NetworkProbeStats_Stats_Latencies) GetP99() *duration.Duration {
	if m != nil {
		return m.P99
	}
	return nil
}

type NetworkProbeStats_HostSplit struct {
	NodeName             string                   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PodName              string                   `protobuf:"bytes,2,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Stats                *NetworkProbeStats_Stats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NetworkProbeStats_HostSplit) Reset()         { *m = NetworkProbeStats_HostSplit{} }
func (m *NetworkProbeStats_HostSplit) String() string { return proto.CompactTextString(m) }
func (*NetworkProbeStats_HostSplit) ProtoMessage()    {}
func (*NetworkProbeStats_HostSplit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3994cf10bd07ee6, []int{4, 1}
}

func (m *NetworkProbeStats_HostSplit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProbeStats_HostSplit.Unmarshal(m, b)
}
func (m *NetworkProbeStats_HostSplit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProbeStats_HostSplit.Marshal(b, m, deterministic)
}
func (m *NetworkProbeStats_HostSplit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProbeStats_HostSplit.Merge(m, src)
}
func (m *NetworkProbeStats_HostSplit) XXX_Size() int {
	return xxx_messageInfo_NetworkProbeStats_HostSplit.Size(m)
}
func (m *NetworkProbeStats_HostSplit) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProbeStats_HostSplit.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProbeStats_HostSplit proto.InternalMessageInfo

func (m *NetworkProbeStats_HostSplit) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NetworkProbeStats_HostSplit) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *NetworkProbeStats_HostSplit) GetStats() *NetworkProbeStats_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*Meta)(nil), "voyager.Meta")
	proto.RegisterType((*ProbeRequest)(nil), "voyager.ProbeRequest")
	proto.RegisterType((*ProbeResponse)(nil), "voyager.ProbeResponse")
	proto.RegisterType((*InitiateNetworkProbe)(nil), "voyager.InitiateNetworkProbe")
	proto.RegisterType((*NetworkProbeStats)(nil), "voyager.NetworkProbeStats")
	proto.RegisterType((*NetworkProbeStats_Stats)(nil), "voyager.NetworkProbeStats.Stats")
	proto.RegisterType((*NetworkProbeStats_Stats_Latencies)(nil), "voyager.NetworkProbeStats.Stats.Latencies")
	proto.RegisterType((*NetworkProbeStats_HostSplit)(nil), "voyager.NetworkProbeStats.HostSplit")
}

func init() { proto.RegisterFile("pkg/voyager.proto", fileDescriptor_c3994cf10bd07ee6) }

var fileDescriptor_c3994cf10bd07ee6 = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xed, 0x4e, 0xdb, 0x48,
	0x14, 0x5d, 0xc7, 0x31, 0xe0, 0x9b, 0xb0, 0xc0, 0x88, 0x45, 0xc1, 0x2b, 0x96, 0x90, 0x5d, 0x69,
	0xa3, 0xdd, 0x55, 0x40, 0xd9, 0xd5, 0xb6, 0x51, 0x7f, 0x21, 0x68, 0x05, 0x6a, 0x8b, 0xd0, 0x84,
	0xff, 0xd6, 0x24, 0xbe, 0x04, 0x0b, 0xdb, 0xe3, 0xce, 0x8c, 0xa9, 0x4a, 0x7f, 0xf7, 0x05, 0xfa,
	0x16, 0x55, 0xdf, 0xa9, 0x6f, 0xd0, 0x77, 0xa8, 0x3c, 0x33, 0x0e, 0xdf, 0x41, 0xed, 0x9f, 0x68,
	0x72, 0xce, 0x99, 0xf1, 0xb9, 0xe7, 0xde, 0x19, 0x58, 0xc9, 0xcf, 0x27, 0xdb, 0x17, 0xfc, 0x1d,
	0x9b, 0xa0, 0xe8, 0xe5, 0x82, 0x2b, 0x4e, 0xe6, 0xed, 0xdf, 0x60, 0x73, 0xc2, 0xf9, 0x24, 0xc1,
	0x6d, 0x0d, 0x8f, 0x8a, 0xd3, 0x6d, 0x15, 0xa7, 0x28, 0x15, 0x4b, 0x73, 0xa3, 0x0c, 0x7e, 0xbb,
	0x2d, 0x88, 0x0a, 0xc1, 0x54, 0xcc, 0x33, 0xc3, 0x77, 0x02, 0xa8, 0xbf, 0x46, 0xc5, 0x08, 0x81,
	0x7a, 0xc4, 0x14, 0x6b, 0x39, 0x6d, 0xa7, 0xdb, 0xa4, 0x7a, 0xdd, 0xf9, 0xe4, 0x40, 0xf3, 0x58,
	0xf0, 0x11, 0x52, 0x7c, 0x53, 0xa0, 0x54, 0xe4, 0x77, 0x58, 0x14, 0x28, 0x73, 0x9e, 0x49, 0x0c,
	0x65, 0x7c, 0x89, 0x5a, 0xed, 0xd2, 0x66, 0x05, 0x0e, 0xe3, 0x4b, 0xbc, 0x21, 0x1a, 0xf3, 0x08,
	0x5b, 0xb5, 0xb6, 0xd3, 0xf5, 0xae, 0x44, 0x7b, 0x3c, 0x42, 0xf2, 0x0f, 0x10, 0x79, 0xc6, 0x8b,
	0x24, 0x0a, 0xcf, 0xe3, 0x24, 0x09, 0x25, 0x8a, 0x0b, 0x14, 0x2d, 0xb7, 0xed, 0x74, 0x17, 0xe8,
	0xb2, 0x61, 0x5e, 0xc6, 0x49, 0x32, 0xd4, 0x38, 0xd9, 0x82, 0x7a, 0x8a, 0x8a, 0xb5, 0xea, 0x6d,
	0xa7, 0xdb, 0xe8, 0x2f, 0xf6, 0xaa, 0x30, 0x4a, 0xe7, 0x54, 0x53, 0x9d, 0xaf, 0x0e, 0x2c, 0x5a,
	0xaf, 0xe6, 0x33, 0xe4, 0x57, 0xf0, 0x33, 0x1e, 0x61, 0x98, 0xb1, 0xd4, 0x18, 0xf5, 0xe9, 0x42,
	0x09, 0x1c, 0xb1, 0x14, 0xc9, 0x3a, 0x2c, 0xe4, 0x3c, 0x32, 0x5c, 0x4d, 0x73, 0xf3, 0x39, 0x8f,
	0x34, 0xf5, 0x0c, 0x1a, 0xc6, 0x4e, 0x58, 0x66, 0xa9, 0x3d, 0x35, 0xfa, 0x41, 0xcf, 0xe4, 0xd8,
	0xab, 0x72, 0xec, 0x9d, 0x54, 0x41, 0x53, 0x30, 0xf2, 0x12, 0x20, 0xfb, 0xb0, 0x9c, 0x0b, 0x3e,
	0x46, 0x29, 0xc3, 0x2a, 0x68, 0xeb, 0x7a, 0xfd, 0xce, 0x09, 0xfb, 0x56, 0x40, 0x97, 0xec, 0x96,
	0x0a, 0x98, 0xd6, 0xeb, 0x3d, 0x5c, 0xef, 0x87, 0x1a, 0xac, 0x1e, 0x66, 0xb1, 0x8a, 0x99, 0xc2,
	0x23, 0x54, 0x6f, 0xb9, 0x38, 0xd7, 0xe5, 0x93, 0x2d, 0x68, 0x0a, 0xd3, 0xae, 0xeb, 0x2d, 0x6a,
	0x58, 0xec, 0x4e, 0x87, 0xb4, 0xa6, 0x76, 0x4f, 0x1b, 0x37, 0x00, 0xce, 0xb8, 0x54, 0x61, 0xce,
	0x85, 0x92, 0x2d, 0xb7, 0xed, 0x76, 0x7d, 0xea, 0x97, 0xc8, 0x71, 0x09, 0x90, 0x3f, 0x61, 0x29,
	0x2b, 0xd2, 0x11, 0x8a, 0xd0, 0x9e, 0x2c, 0x75, 0x9d, 0x2e, 0xfd, 0xd9, 0xc0, 0x76, 0x64, 0x24,
	0xd9, 0x84, 0x46, 0x56, 0xa4, 0xa1, 0x3a, 0x13, 0xc8, 0x22, 0xa9, 0x4b, 0xf2, 0x28, 0x64, 0x45,
	0x7a, 0x62, 0x90, 0xf2, 0x43, 0x28, 0x04, 0x17, 0xa1, 0x60, 0x0a, 0x5b, 0x73, 0x6d, 0xa7, 0x5b,
	0xa3, 0xbe, 0x46, 0x28, 0x53, 0xda, 0x47, 0xce, 0xb2, 0x78, 0x6c, 0xe8, 0x79, 0x43, 0x6b, 0xa4,
	0xa4, 0x3b, 0x9f, 0x3d, 0x58, 0xb9, 0x5e, 0xff, 0x50, 0x31, 0x25, 0xc9, 0xff, 0xe0, 0xc9, 0x72,
	0x61, 0xb3, 0x6f, 0x4f, 0x13, 0xbc, 0x23, 0xed, 0xe9, 0x5f, 0x6a, 0xe4, 0xe4, 0x39, 0x34, 0x74,
	0xd1, 0x32, 0x4f, 0x62, 0x55, 0x9a, 0x75, 0xbb, 0x8d, 0xfe, 0x1f, 0x33, 0x76, 0x1f, 0x70, 0xa9,
	0x86, 0xa5, 0x98, 0xea, 0xb4, 0xf4, 0x52, 0x92, 0x01, 0xc0, 0x58, 0x20, 0x53, 0x18, 0x85, 0x4c,
	0xe9, 0x92, 0x66, 0x4f, 0x90, 0x6f, 0xd5, 0xbb, 0x2a, 0xf8, 0x52, 0x03, 0xcf, 0xd4, 0xf0, 0x1f,
	0xac, 0x95, 0xc1, 0x55, 0xf1, 0x86, 0x63, 0x9e, 0xe6, 0x09, 0x2a, 0x8c, 0x6c, 0x4b, 0x57, 0xb3,
	0x22, 0xad, 0x52, 0xde, 0xab, 0xb8, 0x2a, 0x6e, 0x59, 0x8c, 0xcb, 0x89, 0xb2, 0x9d, 0x2d, 0xe3,
	0x1e, 0x1a, 0xa4, 0x12, 0x9c, 0xb2, 0x38, 0x29, 0x84, 0x19, 0x6f, 0x23, 0x78, 0x61, 0x10, 0x72,
	0x00, 0x7e, 0xc2, 0x14, 0x66, 0xe3, 0x18, 0xa5, 0x9d, 0xc0, 0xbf, 0x1e, 0xcb, 0xaf, 0xf7, 0xaa,
	0xda, 0x41, 0xaf, 0x36, 0x07, 0x1f, 0x1d, 0xf0, 0xa7, 0x04, 0xf9, 0x1b, 0xdc, 0xfc, 0xc9, 0x8e,
	0x36, 0x3f, 0xf3, 0x36, 0x94, 0x2a, 0x2d, 0x1e, 0xec, 0x68, 0xfb, 0x8f, 0x88, 0x07, 0x56, 0x3c,
	0xb0, 0x37, 0x75, 0xb6, 0x78, 0x10, 0xbc, 0x07, 0x7f, 0xda, 0xb4, 0x1f, 0x7e, 0x23, 0xa6, 0xf3,
	0xe5, 0x7e, 0xd7, 0x7c, 0xf5, 0x0f, 0xa0, 0x79, 0xe3, 0xb2, 0x3e, 0x05, 0xcf, 0x2c, 0x7e, 0x99,
	0x9e, 0x70, 0xfd, 0xc1, 0x0d, 0xd6, 0x6e, 0xc3, 0xe6, 0x82, 0x76, 0x7e, 0xea, 0x53, 0xf0, 0x76,
	0xa3, 0x34, 0xce, 0xc8, 0x21, 0xc0, 0x50, 0x31, 0xa1, 0xcc, 0x39, 0x1b, 0xd3, 0x0d, 0xf7, 0x3d,
	0x0e, 0x41, 0xf0, 0xb0, 0xd1, 0x1d, 0x67, 0x34, 0xa7, 0x23, 0xfb, 0xf7, 0x5b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0xd7, 0x8d, 0xaf, 0x71, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkProbeClient is the client API for NetworkProbe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkProbeClient interface {
	Probe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeResponse, error)
}

type networkProbeClient struct {
	cc *grpc.ClientConn
}

func NewNetworkProbeClient(cc *grpc.ClientConn) NetworkProbeClient {
	return &networkProbeClient{cc}
}

func (c *networkProbeClient) Probe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeResponse, error) {
	out := new(ProbeResponse)
	err := c.cc.Invoke(ctx, "/voyager.NetworkProbe/Probe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkProbeServer is the server API for NetworkProbe service.
type NetworkProbeServer interface {
	Probe(context.Context, *ProbeRequest) (*ProbeResponse, error)
}

// UnimplementedNetworkProbeServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkProbeServer struct {
}

func (*UnimplementedNetworkProbeServer) Probe(ctx context.Context, req *ProbeRequest) (*ProbeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Probe not implemented")
}

func RegisterNetworkProbeServer(s *grpc.Server, srv NetworkProbeServer) {
	s.RegisterService(&_NetworkProbe_serviceDesc, srv)
}

func _NetworkProbe_Probe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkProbeServer).Probe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voyager.NetworkProbe/Probe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkProbeServer).Probe(ctx, req.(*ProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkProbe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "voyager.NetworkProbe",
	HandlerType: (*NetworkProbeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Probe",
			Handler:    _NetworkProbe_Probe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/voyager.proto",
}

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	StartProbe(ctx context.Context, in *InitiateNetworkProbe, opts ...grpc.CallOption) (Admin_StartProbeClient, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) StartProbe(ctx context.Context, in *InitiateNetworkProbe, opts ...grpc.CallOption) (Admin_StartProbeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Admin_serviceDesc.Streams[0], "/voyager.Admin/StartProbe", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminStartProbeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_StartProbeClient interface {
	Recv() (*NetworkProbeStats, error)
	grpc.ClientStream
}

type adminStartProbeClient struct {
	grpc.ClientStream
}

func (x *adminStartProbeClient) Recv() (*NetworkProbeStats, error) {
	m := new(NetworkProbeStats)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	StartProbe(*InitiateNetworkProbe, Admin_StartProbeServer) error
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) StartProbe(req *InitiateNetworkProbe, srv Admin_StartProbeServer) error {
	return status.Errorf(codes.Unimplemented, "method StartProbe not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_StartProbe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitiateNetworkProbe)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).StartProbe(m, &adminStartProbeServer{stream})
}

type Admin_StartProbeServer interface {
	Send(*NetworkProbeStats) error
	grpc.ServerStream
}

type adminStartProbeServer struct {
	grpc.ServerStream
}

func (x *adminStartProbeServer) Send(m *NetworkProbeStats) error {
	return x.ServerStream.SendMsg(m)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "voyager.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartProbe",
			Handler:       _Admin_StartProbe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/voyager.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/clusters.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_type "github.com/cilium/proxy/go/envoy/type"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Admin endpoint uses this wrapper for `/clusters` to display cluster status information.
// See :ref:`/clusters <operations_admin_interface_clusters>` for more information.
type Clusters struct {
	// Mapping from cluster name to each cluster's status.
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{0}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

// Details an individual cluster's current status.
type ClusterStatus struct {
	// Name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Denotes whether this cluster was added via API or configured statically.
	AddedViaApi bool `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	// The success rate threshold used in the last interval. The threshold is used to eject hosts
	// based on their success rate. See
	// :ref:`Cluster outlier detection <arch_overview_outlier_detection>` statistics
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	SuccessRateEjectionThreshold *_type.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	// Mapping from host address to the host's current status.
	HostStatuses         []*HostStatus `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{1}
}

func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

// Current state of a particular host.
type HostStatus struct {
	// Address of this host.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// List of stats specific to this host.
	Stats []*SimpleMetric `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	// The host's current health status.
	HealthStatus *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	// Request success rate for this host over the last calculated interval.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	SuccessRate          *_type.Percent `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{2}
}

func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() []*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *_type.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

// Health status for a host.
type HostHealthStatus struct {
	// The host is currently failing active health checks.
	FailedActiveHealthCheck bool `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	// The host is currently considered an outlier and has been ejected.
	FailedOutlierCheck bool `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	// Health status as reported by EDS. Note: only HEALTHY and UNHEALTHY are currently supported
	// here.
	// TODO(mrice32): pipe through remaining EDS health status possibilities.
	EdsHealthStatus      core.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{3}
}

func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostHealthStatus.Unmarshal(m, b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
}
func (m *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(m, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return xxx_messageInfo_HostHealthStatus.Size(m)
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() core.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v2alpha.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v2alpha.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v2alpha.HostStatus")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v2alpha.HostHealthStatus")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/clusters.proto", fileDescriptor_c6251a3a957f478b) }

var fileDescriptor_c6251a3a957f478b = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0x95, 0xad, 0x40, 0xb9, 0x5d, 0xd9, 0xf0, 0x90, 0x88, 0x2a, 0xa4, 0x76, 0x11, 0x48,
	0x7d, 0x4a, 0x50, 0x40, 0xf0, 0xc0, 0x53, 0x35, 0x90, 0x26, 0xd0, 0x04, 0xca, 0x10, 0x12, 0xbc,
	0x58, 0xc6, 0xbe, 0x28, 0x86, 0x34, 0xb6, 0x6c, 0x37, 0xd2, 0xfe, 0x24, 0xff, 0x83, 0x5f, 0x01,
	0xaa, 0xed, 0xae, 0x19, 0xeb, 0xde, 0x92, 0xdc, 0xef, 0x58, 0xe7, 0x9c, 0xeb, 0x40, 0x86, 0x6d,
	0xa7, 0x2e, 0x0b, 0x26, 0x96, 0xb2, 0x2d, 0xba, 0x92, 0x35, 0xba, 0x66, 0x05, 0x6f, 0x56, 0xd6,
	0xa1, 0xb1, 0xb9, 0x36, 0xca, 0x29, 0x72, 0xec, 0x99, 0xdc, 0x33, 0x79, 0x64, 0x26, 0x27, 0xbb,
	0x84, 0x4b, 0x74, 0x46, 0xf2, 0xa8, 0x9b, 0x4c, 0x23, 0xa2, 0x65, 0xd1, 0x95, 0x05, 0x57, 0x06,
	0x0b, 0x26, 0x84, 0x41, 0xbb, 0x01, 0x9e, 0xde, 0x04, 0x6a, 0x64, 0x8d, 0xab, 0x29, 0xaf, 0x91,
	0xff, 0x8a, 0x54, 0x1a, 0x28, 0x77, 0xa9, 0xb1, 0xd0, 0x68, 0x38, 0xb6, 0x2e, 0x4c, 0xb2, 0xaf,
	0x30, 0x3c, 0x8d, 0x56, 0xc9, 0x39, 0x1c, 0x45, 0xdb, 0xd4, 0x3a, 0xe6, 0x56, 0x16, 0x6d, 0x9a,
	0xcc, 0xf6, 0xe7, 0xa3, 0x32, 0xcb, 0x77, 0xf8, 0xcf, 0xa3, 0xf0, 0xc2, 0xb3, 0xd5, 0x21, 0xef,
	0xbf, 0xa2, 0xcd, 0xfe, 0x24, 0x30, 0xbe, 0x86, 0x10, 0x02, 0x83, 0x96, 0x2d, 0x31, 0x4d, 0x66,
	0xc9, 0xfc, 0x7e, 0xe5, 0x9f, 0x49, 0x06, 0x63, 0x26, 0x04, 0x0a, 0xda, 0x49, 0x46, 0x99, 0x96,
	0xe9, 0xde, 0x2c, 0x99, 0x0f, 0xab, 0x91, 0xff, 0xf8, 0x45, 0xb2, 0x85, 0x96, 0xe4, 0x1b, 0x4c,
	0xed, 0x8a, 0x73, 0xb4, 0x96, 0x1a, 0xe6, 0x90, 0xe2, 0x4f, 0xe4, 0x4e, 0xaa, 0x96, 0xba, 0xda,
	0xa0, 0xad, 0x55, 0x23, 0xd2, 0xfd, 0x59, 0x32, 0x1f, 0x95, 0xc7, 0xd1, 0xe7, 0x3a, 0x68, 0xfe,
	0x29, 0x04, 0xad, 0x9e, 0x44, 0x6d, 0xc5, 0x1c, 0xbe, 0x8b, 0xca, 0xcf, 0x1b, 0x21, 0x79, 0x0b,
	0xe3, 0x5a, 0x59, 0xb7, 0x4d, 0x3c, 0xf0, 0x89, 0xa7, 0x3b, 0x13, 0x9f, 0x29, 0xeb, 0x62, 0xdc,
	0x83, 0xfa, 0xea, 0x19, 0x6d, 0xf6, 0x37, 0x01, 0xd8, 0x0e, 0xc9, 0x4b, 0xb8, 0x17, 0xd7, 0xe4,
	0xb3, 0x8e, 0xca, 0xc9, 0xe6, 0x38, 0x2d, 0xf3, 0xae, 0xcc, 0xd7, 0x7b, 0xca, 0x17, 0x81, 0xa8,
	0x36, 0x28, 0x79, 0x0d, 0x77, 0xd6, 0x2e, 0x6c, 0xba, 0xe7, 0x2d, 0x9c, 0xec, 0xb4, 0x70, 0x21,
	0x97, 0xba, 0xc1, 0x73, 0x7f, 0x4b, 0xaa, 0xc0, 0x93, 0xf7, 0x30, 0x8e, 0x4b, 0x0f, 0x29, 0x62,
	0x1b, 0xcf, 0x6e, 0xcd, 0x70, 0xe6, 0xe9, 0xab, 0x24, 0xbd, 0x37, 0xf2, 0x0a, 0x0e, 0xfa, 0x5d,
	0xa7, 0x83, 0xdb, 0x8b, 0x1d, 0xf5, 0x8a, 0xcd, 0x7e, 0x27, 0x70, 0xf4, 0xff, 0xd1, 0xe4, 0x0d,
	0x4c, 0x7e, 0x30, 0xd9, 0xa0, 0xa0, 0x8c, 0x3b, 0xd9, 0x21, 0xed, 0xdf, 0x4d, 0x5f, 0xcd, 0xb0,
	0x7a, 0x1c, 0x88, 0x85, 0x07, 0x82, 0xfa, 0x74, 0x3d, 0x26, 0xcf, 0xe1, 0x51, 0x14, 0xab, 0x95,
	0x6b, 0x24, 0x9a, 0x28, 0x0b, 0x17, 0x84, 0x84, 0xd9, 0xc7, 0x30, 0x0a, 0x8a, 0x0f, 0xf0, 0x10,
	0x85, 0xa5, 0x37, 0xbb, 0x78, 0xb0, 0xdd, 0x67, 0x6f, 0x01, 0xd7, 0x5a, 0x38, 0x44, 0x61, 0xfb,
	0x1f, 0xbe, 0xdf, 0xf5, 0x3f, 0xc8, 0x8b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xc9, 0x9e,
	0x73, 0xdf, 0x03, 0x00, 0x00,
}

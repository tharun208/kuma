// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/mesh.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Mesh defines configuration of a single mesh.
type Mesh struct {
	// mTLS settings.
	// +optional
	Mtls *Mesh_Mtls `protobuf:"bytes,1,opt,name=mtls,proto3" json:"mtls,omitempty"`
	// Tracing settings.
	// +optional
	Tracing *Tracing `protobuf:"bytes,2,opt,name=tracing,proto3" json:"tracing,omitempty"`
	// Logging settings.
	// +optional
	Logging *Logging `protobuf:"bytes,3,opt,name=logging,proto3" json:"logging,omitempty"`
	// Configuration for metrics collected and exposed by dataplanes.
	//
	// Settings defined here become defaults for every dataplane in a given Mesh.
	// Additionally, it is also possible to further customize this configuration
	// for each dataplane individually using Dataplane resource.
	// +optional
	Metrics              *Metrics `protobuf:"bytes,4,opt,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{0}
}

func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (m *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(m, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

func (m *Mesh) GetMtls() *Mesh_Mtls {
	if m != nil {
		return m.Mtls
	}
	return nil
}

func (m *Mesh) GetTracing() *Tracing {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Mesh) GetLogging() *Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *Mesh) GetMetrics() *Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// mTLS settings of a Mesh.
type Mesh_Mtls struct {
	// Name of the enabled backend
	EnabledBackend string `protobuf:"bytes,1,opt,name=enabledBackend,proto3" json:"enabledBackend,omitempty"`
	// List of available Certificate Authority backends
	Backends             []*CertificateAuthorityBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Mesh_Mtls) Reset()         { *m = Mesh_Mtls{} }
func (m *Mesh_Mtls) String() string { return proto.CompactTextString(m) }
func (*Mesh_Mtls) ProtoMessage()    {}
func (*Mesh_Mtls) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{0, 0}
}

func (m *Mesh_Mtls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh_Mtls.Unmarshal(m, b)
}
func (m *Mesh_Mtls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh_Mtls.Marshal(b, m, deterministic)
}
func (m *Mesh_Mtls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh_Mtls.Merge(m, src)
}
func (m *Mesh_Mtls) XXX_Size() int {
	return xxx_messageInfo_Mesh_Mtls.Size(m)
}
func (m *Mesh_Mtls) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh_Mtls.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh_Mtls proto.InternalMessageInfo

func (m *Mesh_Mtls) GetEnabledBackend() string {
	if m != nil {
		return m.EnabledBackend
	}
	return ""
}

func (m *Mesh_Mtls) GetBackends() []*CertificateAuthorityBackend {
	if m != nil {
		return m.Backends
	}
	return nil
}

// CertificateAuthorityBackend defines Certificate Authority backend
type CertificateAuthorityBackend struct {
	// Name of the backend
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the backend. Has to be one of the loaded plugins (Kuma ships with
	// builtin and provided)
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Configuration of the backend
	Config               *_struct.Struct `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CertificateAuthorityBackend) Reset()         { *m = CertificateAuthorityBackend{} }
func (m *CertificateAuthorityBackend) String() string { return proto.CompactTextString(m) }
func (*CertificateAuthorityBackend) ProtoMessage()    {}
func (*CertificateAuthorityBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{1}
}

func (m *CertificateAuthorityBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateAuthorityBackend.Unmarshal(m, b)
}
func (m *CertificateAuthorityBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateAuthorityBackend.Marshal(b, m, deterministic)
}
func (m *CertificateAuthorityBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateAuthorityBackend.Merge(m, src)
}
func (m *CertificateAuthorityBackend) XXX_Size() int {
	return xxx_messageInfo_CertificateAuthorityBackend.Size(m)
}
func (m *CertificateAuthorityBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateAuthorityBackend.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateAuthorityBackend proto.InternalMessageInfo

func (m *CertificateAuthorityBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateAuthorityBackend) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CertificateAuthorityBackend) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Tracing defines tracing configuration of the mesh.
type Tracing struct {
	// Name of the default backend
	DefaultBackend string `protobuf:"bytes,1,opt,name=defaultBackend,proto3" json:"defaultBackend,omitempty"`
	// List of available tracing backends
	Backends             []*TracingBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{2}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetDefaultBackend() string {
	if m != nil {
		return m.DefaultBackend
	}
	return ""
}

func (m *Tracing) GetBackends() []*TracingBackend {
	if m != nil {
		return m.Backends
	}
	return nil
}

// TracingBackend defines tracing backend available to mesh. Backends can be
// used in TrafficTrace rules.
type TracingBackend struct {
	// Name of the backend, can be then used in Mesh.tracing.defaultBackend or in
	// TrafficTrace
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Percentage of traces that will be sent to the backend (range 0.0 - 100.0).
	// Empty value defaults to 100.0%
	Sampling *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=sampling,proto3" json:"sampling,omitempty"`
	// Type of the backend (Kuma ships with 'zipkin')
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Configuration of the backend
	Config               *_struct.Struct `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TracingBackend) Reset()         { *m = TracingBackend{} }
func (m *TracingBackend) String() string { return proto.CompactTextString(m) }
func (*TracingBackend) ProtoMessage()    {}
func (*TracingBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{3}
}

func (m *TracingBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracingBackend.Unmarshal(m, b)
}
func (m *TracingBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracingBackend.Marshal(b, m, deterministic)
}
func (m *TracingBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracingBackend.Merge(m, src)
}
func (m *TracingBackend) XXX_Size() int {
	return xxx_messageInfo_TracingBackend.Size(m)
}
func (m *TracingBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_TracingBackend.DiscardUnknown(m)
}

var xxx_messageInfo_TracingBackend proto.InternalMessageInfo

func (m *TracingBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TracingBackend) GetSampling() *wrappers.DoubleValue {
	if m != nil {
		return m.Sampling
	}
	return nil
}

func (m *TracingBackend) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TracingBackend) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type ZipkinTracingBackendConfig struct {
	// Address of Zipkin collector.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Generate 128bit traces. Default: false
	TraceId128Bit bool `protobuf:"varint,2,opt,name=traceId128bit,proto3" json:"traceId128bit,omitempty"`
	// Version of the API. values: httpJson, httpJsonV1, httpProto. Default:
	// httpJson see
	// https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/trace/v3/trace.proto#envoy-v3-api-enum-config-trace-v3-zipkinconfig-collectorendpointversion
	ApiVersion           string   `protobuf:"bytes,3,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZipkinTracingBackendConfig) Reset()         { *m = ZipkinTracingBackendConfig{} }
func (m *ZipkinTracingBackendConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinTracingBackendConfig) ProtoMessage()    {}
func (*ZipkinTracingBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{4}
}

func (m *ZipkinTracingBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinTracingBackendConfig.Unmarshal(m, b)
}
func (m *ZipkinTracingBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinTracingBackendConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinTracingBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinTracingBackendConfig.Merge(m, src)
}
func (m *ZipkinTracingBackendConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinTracingBackendConfig.Size(m)
}
func (m *ZipkinTracingBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinTracingBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinTracingBackendConfig proto.InternalMessageInfo

func (m *ZipkinTracingBackendConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ZipkinTracingBackendConfig) GetTraceId128Bit() bool {
	if m != nil {
		return m.TraceId128Bit
	}
	return false
}

func (m *ZipkinTracingBackendConfig) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

type Logging struct {
	// Name of the default backend
	DefaultBackend string `protobuf:"bytes,1,opt,name=defaultBackend,proto3" json:"defaultBackend,omitempty"`
	// List of available logging backends
	Backends             []*LoggingBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Logging) Reset()         { *m = Logging{} }
func (m *Logging) String() string { return proto.CompactTextString(m) }
func (*Logging) ProtoMessage()    {}
func (*Logging) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{5}
}

func (m *Logging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logging.Unmarshal(m, b)
}
func (m *Logging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logging.Marshal(b, m, deterministic)
}
func (m *Logging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logging.Merge(m, src)
}
func (m *Logging) XXX_Size() int {
	return xxx_messageInfo_Logging.Size(m)
}
func (m *Logging) XXX_DiscardUnknown() {
	xxx_messageInfo_Logging.DiscardUnknown(m)
}

var xxx_messageInfo_Logging proto.InternalMessageInfo

func (m *Logging) GetDefaultBackend() string {
	if m != nil {
		return m.DefaultBackend
	}
	return ""
}

func (m *Logging) GetBackends() []*LoggingBackend {
	if m != nil {
		return m.Backends
	}
	return nil
}

// LoggingBackend defines logging backend available to mesh. Backends can be
// used in TrafficLog rules.
type LoggingBackend struct {
	// Name of the backend, can be then used in Mesh.logging.defaultBackend or in
	// TrafficLogging
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Format of access logs. Placehodlers available on
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log
	Format string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	// Type of the backend (Kuma ships with 'tcp' and 'file')
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Configuration of the backend
	Config               *_struct.Struct `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoggingBackend) Reset()         { *m = LoggingBackend{} }
func (m *LoggingBackend) String() string { return proto.CompactTextString(m) }
func (*LoggingBackend) ProtoMessage()    {}
func (*LoggingBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{6}
}

func (m *LoggingBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoggingBackend.Unmarshal(m, b)
}
func (m *LoggingBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoggingBackend.Marshal(b, m, deterministic)
}
func (m *LoggingBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggingBackend.Merge(m, src)
}
func (m *LoggingBackend) XXX_Size() int {
	return xxx_messageInfo_LoggingBackend.Size(m)
}
func (m *LoggingBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggingBackend.DiscardUnknown(m)
}

var xxx_messageInfo_LoggingBackend proto.InternalMessageInfo

func (m *LoggingBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoggingBackend) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *LoggingBackend) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LoggingBackend) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// FileLoggingBackendConfig defines configuration for file based access logs
type FileLoggingBackendConfig struct {
	// Path to a file that logs will be written to
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileLoggingBackendConfig) Reset()         { *m = FileLoggingBackendConfig{} }
func (m *FileLoggingBackendConfig) String() string { return proto.CompactTextString(m) }
func (*FileLoggingBackendConfig) ProtoMessage()    {}
func (*FileLoggingBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{7}
}

func (m *FileLoggingBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileLoggingBackendConfig.Unmarshal(m, b)
}
func (m *FileLoggingBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileLoggingBackendConfig.Marshal(b, m, deterministic)
}
func (m *FileLoggingBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileLoggingBackendConfig.Merge(m, src)
}
func (m *FileLoggingBackendConfig) XXX_Size() int {
	return xxx_messageInfo_FileLoggingBackendConfig.Size(m)
}
func (m *FileLoggingBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileLoggingBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileLoggingBackendConfig proto.InternalMessageInfo

func (m *FileLoggingBackendConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// TcpLoggingBackendConfig defines configuration for TCP based access logs
type TcpLoggingBackendConfig struct {
	// Address to TCP service that will receive logs
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpLoggingBackendConfig) Reset()         { *m = TcpLoggingBackendConfig{} }
func (m *TcpLoggingBackendConfig) String() string { return proto.CompactTextString(m) }
func (*TcpLoggingBackendConfig) ProtoMessage()    {}
func (*TcpLoggingBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{8}
}

func (m *TcpLoggingBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpLoggingBackendConfig.Unmarshal(m, b)
}
func (m *TcpLoggingBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpLoggingBackendConfig.Marshal(b, m, deterministic)
}
func (m *TcpLoggingBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpLoggingBackendConfig.Merge(m, src)
}
func (m *TcpLoggingBackendConfig) XXX_Size() int {
	return xxx_messageInfo_TcpLoggingBackendConfig.Size(m)
}
func (m *TcpLoggingBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpLoggingBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpLoggingBackendConfig proto.InternalMessageInfo

func (m *TcpLoggingBackendConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Mesh)(nil), "kuma.mesh.v1alpha1.Mesh")
	proto.RegisterType((*Mesh_Mtls)(nil), "kuma.mesh.v1alpha1.Mesh.Mtls")
	proto.RegisterType((*CertificateAuthorityBackend)(nil), "kuma.mesh.v1alpha1.CertificateAuthorityBackend")
	proto.RegisterType((*Tracing)(nil), "kuma.mesh.v1alpha1.Tracing")
	proto.RegisterType((*TracingBackend)(nil), "kuma.mesh.v1alpha1.TracingBackend")
	proto.RegisterType((*ZipkinTracingBackendConfig)(nil), "kuma.mesh.v1alpha1.ZipkinTracingBackendConfig")
	proto.RegisterType((*Logging)(nil), "kuma.mesh.v1alpha1.Logging")
	proto.RegisterType((*LoggingBackend)(nil), "kuma.mesh.v1alpha1.LoggingBackend")
	proto.RegisterType((*FileLoggingBackendConfig)(nil), "kuma.mesh.v1alpha1.FileLoggingBackendConfig")
	proto.RegisterType((*TcpLoggingBackendConfig)(nil), "kuma.mesh.v1alpha1.TcpLoggingBackendConfig")
}

func init() { proto.RegisterFile("mesh/v1alpha1/mesh.proto", fileDescriptor_ae9b3cd8c92bbf6a) }

var fileDescriptor_ae9b3cd8c92bbf6a = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0x26, 0x89, 0x49, 0xd2, 0x57, 0x16, 0x86, 0x0e, 0xab, 0x49, 0xba, 0x52, 0xc2, 0x18, 0x3d,
	0xd9, 0x24, 0x65, 0xd0, 0xd3, 0x60, 0xed, 0x18, 0x8c, 0x2d, 0x17, 0xaf, 0xf4, 0xd0, 0x9b, 0x6c,
	0x2b, 0x89, 0x88, 0x6c, 0x69, 0x92, 0xdc, 0x51, 0x76, 0xda, 0x6f, 0xd9, 0xcf, 0xda, 0x9f, 0x19,
	0x92, 0xa5, 0x6c, 0x4e, 0xdc, 0xd2, 0x43, 0x6f, 0x4f, 0xef, 0x7d, 0x9f, 0xfd, 0xbd, 0xef, 0x7b,
	0x10, 0x16, 0x44, 0xad, 0xe3, 0xbb, 0x19, 0x66, 0x62, 0x8d, 0x67, 0xb1, 0x79, 0x45, 0x42, 0x72,
	0xcd, 0x11, 0xda, 0x54, 0x05, 0x8e, 0x6c, 0xc3, 0x8f, 0xc7, 0x93, 0x5d, 0xb4, 0x96, 0x34, 0x53,
	0x35, 0x61, 0x7c, 0xb2, 0xe2, 0x7c, 0xc5, 0x48, 0x6c, 0x5f, 0x69, 0xb5, 0x8c, 0x7f, 0x48, 0x2c,
	0x04, 0x91, 0x7e, 0x7e, 0xbc, 0x3b, 0x57, 0x5a, 0x56, 0x99, 0xae, 0xa7, 0xd3, 0x3f, 0x5d, 0x08,
	0x16, 0x44, 0xad, 0xd1, 0x0c, 0x82, 0x42, 0x33, 0x15, 0x76, 0x4e, 0x3b, 0x67, 0x87, 0xf3, 0xd7,
	0xd1, 0xbe, 0x8c, 0xc8, 0xe0, 0xa2, 0x85, 0x66, 0x2a, 0xb1, 0x50, 0xf4, 0x0e, 0x06, 0x5a, 0xe2,
	0x8c, 0x96, 0xab, 0xb0, 0x6b, 0x59, 0x93, 0x36, 0xd6, 0x75, 0x0d, 0x49, 0x3c, 0xd6, 0xd0, 0x18,
	0x5f, 0xad, 0x0c, 0xad, 0xf7, 0x30, 0xed, 0x6b, 0x0d, 0x49, 0x3c, 0xd6, 0xd0, 0xdc, 0xe2, 0x61,
	0xf0, 0x30, 0x6d, 0x51, 0x43, 0x12, 0x8f, 0x1d, 0xff, 0x84, 0xc0, 0x48, 0x46, 0x6f, 0x61, 0x44,
	0x4a, 0x9c, 0x32, 0x92, 0x5f, 0xe2, 0x6c, 0x43, 0xca, 0xdc, 0x6e, 0x7a, 0x90, 0xec, 0x74, 0xd1,
	0x17, 0x18, 0xa6, 0x75, 0xa9, 0xc2, 0xee, 0x69, 0xef, 0xec, 0x70, 0x1e, 0xb7, 0xfd, 0xe7, 0x8a,
	0x48, 0x4d, 0x97, 0x34, 0xc3, 0x9a, 0x7c, 0xa8, 0xf4, 0x9a, 0x4b, 0xaa, 0xef, 0xdd, 0x27, 0x92,
	0xed, 0x07, 0xa6, 0x77, 0x30, 0x79, 0x04, 0x88, 0x10, 0x04, 0x25, 0x2e, 0x88, 0x53, 0x62, 0x6b,
	0xd3, 0xd3, 0xf7, 0x82, 0x58, 0x47, 0x0f, 0x12, 0x5b, 0xa3, 0x18, 0xfa, 0x19, 0x2f, 0x97, 0xd4,
	0x1b, 0x76, 0x14, 0xd5, 0x99, 0x46, 0x3e, 0xd3, 0xe8, 0x9b, 0xcd, 0x34, 0x71, 0xb0, 0xe9, 0x77,
	0x18, 0x38, 0xdb, 0xcd, 0xde, 0x39, 0x59, 0xe2, 0x8a, 0xe9, 0x9d, 0xbd, 0x9b, 0x5d, 0xf4, 0x7e,
	0x6f, 0xef, 0xe9, 0x23, 0x69, 0xee, 0xaf, 0xfa, 0xbb, 0x03, 0xa3, 0xe6, 0xb0, 0x75, 0xbd, 0x0b,
	0x18, 0x2a, 0x5c, 0x08, 0xf6, 0xef, 0x68, 0x8e, 0xf7, 0x96, 0xf9, 0xc8, 0xab, 0x94, 0x91, 0x1b,
	0xcc, 0x2a, 0x92, 0x6c, 0xd1, 0x5b, 0x63, 0x7a, 0xad, 0xc6, 0x04, 0x4f, 0x33, 0x46, 0xc3, 0xf8,
	0x96, 0x8a, 0x0d, 0x2d, 0x9b, 0x52, 0xaf, 0xec, 0x14, 0xbd, 0x84, 0x5e, 0x25, 0x99, 0xd3, 0x6b,
	0x4a, 0xf4, 0x06, 0x5e, 0x98, 0xb3, 0x25, 0x9f, 0xf3, 0xd9, 0xfc, 0x22, 0xa5, 0xda, 0x6a, 0x1e,
	0x26, 0xcd, 0x26, 0x3a, 0x01, 0xc0, 0x82, 0xde, 0x10, 0xa9, 0x28, 0x2f, 0x9d, 0xc0, 0xff, 0x3a,
	0x26, 0x0e, 0x77, 0xce, 0xcf, 0x1d, 0x87, 0xfb, 0xec, 0x7e, 0x1c, 0xbf, 0x3a, 0x30, 0x6a, 0x0e,
	0x5b, 0xe3, 0x78, 0x05, 0xfd, 0x25, 0x97, 0x05, 0xd6, 0xee, 0xde, 0xdc, 0xeb, 0x79, 0xcc, 0x8e,
	0x20, 0xfc, 0x44, 0x19, 0x69, 0xca, 0x70, 0x56, 0x23, 0x08, 0x04, 0xd6, 0x6b, 0x2f, 0xc6, 0xd4,
	0xd3, 0x73, 0x38, 0xba, 0xce, 0x44, 0x2b, 0x3c, 0x84, 0x01, 0xce, 0x73, 0x49, 0x94, 0x72, 0x0c,
	0xff, 0xbc, 0x84, 0xdb, 0xa1, 0x77, 0x23, 0xed, 0x5b, 0x25, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x35, 0x11, 0x57, 0x5e, 0x05, 0x00, 0x00,
}

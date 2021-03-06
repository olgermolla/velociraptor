// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifact_collector.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
	proto2 "www.velocidex.com/golang/velociraptor/crypto/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type ArtifactCollectorContext_State int32

const (
	ArtifactCollectorContext_UNSET      ArtifactCollectorContext_State = 0
	ArtifactCollectorContext_RUNNING    ArtifactCollectorContext_State = 1
	ArtifactCollectorContext_TERMINATED ArtifactCollectorContext_State = 2
	ArtifactCollectorContext_ERROR      ArtifactCollectorContext_State = 3
	ArtifactCollectorContext_ARCHIVED   ArtifactCollectorContext_State = 4
)

var ArtifactCollectorContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
	4: "ARCHIVED",
}

var ArtifactCollectorContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
	"ARCHIVED":   4,
}

func (x ArtifactCollectorContext_State) String() string {
	return proto.EnumName(ArtifactCollectorContext_State_name, int32(x))
}

func (ArtifactCollectorContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{4, 0}
}

type FlowContext_State int32

const (
	FlowContext_UNSET      FlowContext_State = 0
	FlowContext_RUNNING    FlowContext_State = 1
	FlowContext_TERMINATED FlowContext_State = 2
	FlowContext_ERROR      FlowContext_State = 3
	FlowContext_ARCHIVED   FlowContext_State = 4
)

var FlowContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
	4: "ARCHIVED",
}

var FlowContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
	"ARCHIVED":   4,
}

func (x FlowContext_State) String() string {
	return proto.EnumName(FlowContext_State_name, int32(x))
}

func (FlowContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{7, 0}
}

type ArtifactParameters struct {
	Env                  []*proto1.VQLEnv `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactParameters) Reset()         { *m = ArtifactParameters{} }
func (m *ArtifactParameters) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameters) ProtoMessage()    {}
func (*ArtifactParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{0}
}

func (m *ArtifactParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameters.Unmarshal(m, b)
}
func (m *ArtifactParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameters.Marshal(b, m, deterministic)
}
func (m *ArtifactParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameters.Merge(m, src)
}
func (m *ArtifactParameters) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameters.Size(m)
}
func (m *ArtifactParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameters proto.InternalMessageInfo

func (m *ArtifactParameters) GetEnv() []*proto1.VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

type ArtifactCollectorArgs struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// If set we send an urgent request to the client.
	Urgent               bool                `protobuf:"varint,21,opt,name=urgent,proto3" json:"urgent,omitempty"`
	Artifacts            []string            `protobuf:"bytes,2,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	Parameters           *ArtifactParameters `protobuf:"bytes,5,opt,name=parameters,proto3" json:"parameters,omitempty"`
	OpsPerSecond         float32             `protobuf:"fixed32,6,opt,name=ops_per_second,json=opsPerSecond,proto3" json:"ops_per_second,omitempty"`
	Timeout              uint64              `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	AllowCustomOverrides bool                `protobuf:"varint,8,opt,name=allow_custom_overrides,json=allowCustomOverrides,proto3" json:"allow_custom_overrides,omitempty"`
	// A place to cache the compiled request. If this is provided we
	// do not compile the artifacts at all, we just use it as is.
	CompiledCollectorArgs *proto1.VQLCollectorArgs `protobuf:"bytes,20,opt,name=compiled_collector_args,json=compiledCollectorArgs,proto3" json:"compiled_collector_args,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *ArtifactCollectorArgs) Reset()         { *m = ArtifactCollectorArgs{} }
func (m *ArtifactCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorArgs) ProtoMessage()    {}
func (*ArtifactCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{1}
}

func (m *ArtifactCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorArgs.Unmarshal(m, b)
}
func (m *ArtifactCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorArgs.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorArgs.Merge(m, src)
}
func (m *ArtifactCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorArgs.Size(m)
}
func (m *ArtifactCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorArgs proto.InternalMessageInfo

func (m *ArtifactCollectorArgs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ArtifactCollectorArgs) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ArtifactCollectorArgs) GetUrgent() bool {
	if m != nil {
		return m.Urgent
	}
	return false
}

func (m *ArtifactCollectorArgs) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetParameters() *ArtifactParameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetOpsPerSecond() float32 {
	if m != nil {
		return m.OpsPerSecond
	}
	return 0
}

func (m *ArtifactCollectorArgs) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ArtifactCollectorArgs) GetAllowCustomOverrides() bool {
	if m != nil {
		return m.AllowCustomOverrides
	}
	return false
}

func (m *ArtifactCollectorArgs) GetCompiledCollectorArgs() *proto1.VQLCollectorArgs {
	if m != nil {
		return m.CompiledCollectorArgs
	}
	return nil
}

type ArtifactCollectorResponse struct {
	FlowId               string                 `protobuf:"bytes,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Request              *ArtifactCollectorArgs `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ArtifactCollectorResponse) Reset()         { *m = ArtifactCollectorResponse{} }
func (m *ArtifactCollectorResponse) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorResponse) ProtoMessage()    {}
func (*ArtifactCollectorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{2}
}

func (m *ArtifactCollectorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorResponse.Unmarshal(m, b)
}
func (m *ArtifactCollectorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorResponse.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorResponse.Merge(m, src)
}
func (m *ArtifactCollectorResponse) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorResponse.Size(m)
}
func (m *ArtifactCollectorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorResponse proto.InternalMessageInfo

func (m *ArtifactCollectorResponse) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *ArtifactCollectorResponse) GetRequest() *ArtifactCollectorArgs {
	if m != nil {
		return m.Request
	}
	return nil
}

type ArtifactUploadedFileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VfsPath              string   `protobuf:"bytes,2,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactUploadedFileInfo) Reset()         { *m = ArtifactUploadedFileInfo{} }
func (m *ArtifactUploadedFileInfo) String() string { return proto.CompactTextString(m) }
func (*ArtifactUploadedFileInfo) ProtoMessage()    {}
func (*ArtifactUploadedFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{3}
}

func (m *ArtifactUploadedFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactUploadedFileInfo.Unmarshal(m, b)
}
func (m *ArtifactUploadedFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactUploadedFileInfo.Marshal(b, m, deterministic)
}
func (m *ArtifactUploadedFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactUploadedFileInfo.Merge(m, src)
}
func (m *ArtifactUploadedFileInfo) XXX_Size() int {
	return xxx_messageInfo_ArtifactUploadedFileInfo.Size(m)
}
func (m *ArtifactUploadedFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactUploadedFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactUploadedFileInfo proto.InternalMessageInfo

func (m *ArtifactUploadedFileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactUploadedFileInfo) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *ArtifactUploadedFileInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// This context is serialized into the data store.
type ArtifactCollectorContext struct {
	ClientId  string                 `protobuf:"bytes,27,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SessionId string                 `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Request   *ArtifactCollectorArgs `protobuf:"bytes,11,opt,name=request,proto3" json:"request,omitempty"`
	// If an error occurs this is the backtrace.
	Backtrace     string `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	CreateTime    uint64 `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ActiveTime    uint64 `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	KillTimestamp uint64 `protobuf:"varint,4,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	// Uploads are now permanently stored in a csv file. This field is
	// never serialized - it is just a place holder during processing.
	UploadedFiles []*ArtifactUploadedFileInfo `protobuf:"bytes,24,rep,name=uploaded_files,json=uploadedFiles,proto3" json:"uploaded_files,omitempty"`
	// A total count of files uploaded by this query.
	TotalUploadedFiles         uint64 `protobuf:"varint,23,opt,name=total_uploaded_files,json=totalUploadedFiles,proto3" json:"total_uploaded_files,omitempty"`
	TotalExpectedUploadedBytes uint64 `protobuf:"varint,25,opt,name=total_expected_uploaded_bytes,json=totalExpectedUploadedBytes,proto3" json:"total_expected_uploaded_bytes,omitempty"`
	TotalUploadedBytes         uint64 `protobuf:"varint,26,opt,name=total_uploaded_bytes,json=totalUploadedBytes,proto3" json:"total_uploaded_bytes,omitempty"`
	// Logs are stored in their own CSV file. This is just a
	// placeholder during processing.
	Logs         []*proto2.LogMessage           `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs,omitempty"`
	State        ArtifactCollectorContext_State `protobuf:"varint,14,opt,name=state,proto3,enum=proto.ArtifactCollectorContext_State" json:"state,omitempty"`
	Status       string                         `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified bool                           `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	// Some of the collected artifacts may not have results.
	ArtifactsWithResults []string `protobuf:"bytes,22,rep,name=artifacts_with_results,json=artifactsWithResults,proto3" json:"artifacts_with_results,omitempty"`
	Dirty                bool     `protobuf:"varint,2,opt,name=dirty,proto3" json:"dirty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactCollectorContext) Reset()         { *m = ArtifactCollectorContext{} }
func (m *ArtifactCollectorContext) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorContext) ProtoMessage()    {}
func (*ArtifactCollectorContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{4}
}

func (m *ArtifactCollectorContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorContext.Unmarshal(m, b)
}
func (m *ArtifactCollectorContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorContext.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorContext.Merge(m, src)
}
func (m *ArtifactCollectorContext) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorContext.Size(m)
}
func (m *ArtifactCollectorContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorContext.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorContext proto.InternalMessageInfo

func (m *ArtifactCollectorContext) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ArtifactCollectorContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *ArtifactCollectorContext) GetRequest() *ArtifactCollectorArgs {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ArtifactCollectorContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *ArtifactCollectorContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *ArtifactCollectorContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

func (m *ArtifactCollectorContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *ArtifactCollectorContext) GetUploadedFiles() []*ArtifactUploadedFileInfo {
	if m != nil {
		return m.UploadedFiles
	}
	return nil
}

func (m *ArtifactCollectorContext) GetTotalUploadedFiles() uint64 {
	if m != nil {
		return m.TotalUploadedFiles
	}
	return 0
}

func (m *ArtifactCollectorContext) GetTotalExpectedUploadedBytes() uint64 {
	if m != nil {
		return m.TotalExpectedUploadedBytes
	}
	return 0
}

func (m *ArtifactCollectorContext) GetTotalUploadedBytes() uint64 {
	if m != nil {
		return m.TotalUploadedBytes
	}
	return 0
}

func (m *ArtifactCollectorContext) GetLogs() []*proto2.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *ArtifactCollectorContext) GetState() ArtifactCollectorContext_State {
	if m != nil {
		return m.State
	}
	return ArtifactCollectorContext_UNSET
}

func (m *ArtifactCollectorContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ArtifactCollectorContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *ArtifactCollectorContext) GetArtifactsWithResults() []string {
	if m != nil {
		return m.ArtifactsWithResults
	}
	return nil
}

func (m *ArtifactCollectorContext) GetDirty() bool {
	if m != nil {
		return m.Dirty
	}
	return false
}

// This is stored in the ArtifactCollector state.
type ClientEventTable struct {
	Version              uint64                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Artifacts            *ArtifactCollectorArgs `protobuf:"bytes,2,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ClientEventTable) Reset()         { *m = ClientEventTable{} }
func (m *ClientEventTable) String() string { return proto.CompactTextString(m) }
func (*ClientEventTable) ProtoMessage()    {}
func (*ClientEventTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{5}
}

func (m *ClientEventTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventTable.Unmarshal(m, b)
}
func (m *ClientEventTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventTable.Marshal(b, m, deterministic)
}
func (m *ClientEventTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventTable.Merge(m, src)
}
func (m *ClientEventTable) XXX_Size() int {
	return xxx_messageInfo_ClientEventTable.Size(m)
}
func (m *ClientEventTable) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventTable.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventTable proto.InternalMessageInfo

func (m *ClientEventTable) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ClientEventTable) GetArtifacts() *ArtifactCollectorArgs {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type UploadedFileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VfsPath              string   `protobuf:"bytes,2,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadedFileInfo) Reset()         { *m = UploadedFileInfo{} }
func (m *UploadedFileInfo) String() string { return proto.CompactTextString(m) }
func (*UploadedFileInfo) ProtoMessage()    {}
func (*UploadedFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{6}
}

func (m *UploadedFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadedFileInfo.Unmarshal(m, b)
}
func (m *UploadedFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadedFileInfo.Marshal(b, m, deterministic)
}
func (m *UploadedFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadedFileInfo.Merge(m, src)
}
func (m *UploadedFileInfo) XXX_Size() int {
	return xxx_messageInfo_UploadedFileInfo.Size(m)
}
func (m *UploadedFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadedFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UploadedFileInfo proto.InternalMessageInfo

func (m *UploadedFileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadedFileInfo) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *UploadedFileInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// The flow context.
// Next field: 19
type FlowContext struct {
	Backtrace        string `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	CreateTime       uint64 `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	KillTimestamp    uint64 `protobuf:"varint,6,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	NetworkBytesSent uint64 `protobuf:"varint,7,opt,name=network_bytes_sent,json=networkBytesSent,proto3" json:"network_bytes_sent,omitempty"`
	// Uploads are now permanently stored in a csv file. This field is
	// never serialized - it is just a place holder during processing.
	UploadedFiles              []*UploadedFileInfo `protobuf:"bytes,24,rep,name=uploaded_files,json=uploadedFiles,proto3" json:"uploaded_files,omitempty"`
	TotalUploadedFiles         uint64              `protobuf:"varint,23,opt,name=total_uploaded_files,json=totalUploadedFiles,proto3" json:"total_uploaded_files,omitempty"`
	TotalExpectedUploadedBytes uint64              `protobuf:"varint,25,opt,name=total_expected_uploaded_bytes,json=totalExpectedUploadedBytes,proto3" json:"total_expected_uploaded_bytes,omitempty"`
	TotalUploadedBytes         uint64              `protobuf:"varint,26,opt,name=total_uploaded_bytes,json=totalUploadedBytes,proto3" json:"total_uploaded_bytes,omitempty"`
	// Logs are stored in their own CSV file. This is just a
	// placeholder during processing.
	Logs                 []*proto2.LogMessage `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs,omitempty"`
	SessionId            string               `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	State                FlowContext_State    `protobuf:"varint,14,opt,name=state,proto3,enum=proto.FlowContext_State" json:"state,omitempty"`
	Status               string               `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified         bool                 `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	ActiveTime           uint64               `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	Artifacts            []string             `protobuf:"bytes,21,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	ArtifactsWithResults []string             `protobuf:"bytes,22,rep,name=artifacts_with_results,json=artifactsWithResults,proto3" json:"artifacts_with_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FlowContext) Reset()         { *m = FlowContext{} }
func (m *FlowContext) String() string { return proto.CompactTextString(m) }
func (*FlowContext) ProtoMessage()    {}
func (*FlowContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{7}
}

func (m *FlowContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowContext.Unmarshal(m, b)
}
func (m *FlowContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowContext.Marshal(b, m, deterministic)
}
func (m *FlowContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowContext.Merge(m, src)
}
func (m *FlowContext) XXX_Size() int {
	return xxx_messageInfo_FlowContext.Size(m)
}
func (m *FlowContext) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowContext.DiscardUnknown(m)
}

var xxx_messageInfo_FlowContext proto.InternalMessageInfo

func (m *FlowContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *FlowContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FlowContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *FlowContext) GetNetworkBytesSent() uint64 {
	if m != nil {
		return m.NetworkBytesSent
	}
	return 0
}

func (m *FlowContext) GetUploadedFiles() []*UploadedFileInfo {
	if m != nil {
		return m.UploadedFiles
	}
	return nil
}

func (m *FlowContext) GetTotalUploadedFiles() uint64 {
	if m != nil {
		return m.TotalUploadedFiles
	}
	return 0
}

func (m *FlowContext) GetTotalExpectedUploadedBytes() uint64 {
	if m != nil {
		return m.TotalExpectedUploadedBytes
	}
	return 0
}

func (m *FlowContext) GetTotalUploadedBytes() uint64 {
	if m != nil {
		return m.TotalUploadedBytes
	}
	return 0
}

func (m *FlowContext) GetLogs() []*proto2.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *FlowContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *FlowContext) GetState() FlowContext_State {
	if m != nil {
		return m.State
	}
	return FlowContext_UNSET
}

func (m *FlowContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FlowContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *FlowContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

func (m *FlowContext) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *FlowContext) GetArtifactsWithResults() []string {
	if m != nil {
		return m.ArtifactsWithResults
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ArtifactCollectorContext_State", ArtifactCollectorContext_State_name, ArtifactCollectorContext_State_value)
	proto.RegisterEnum("proto.FlowContext_State", FlowContext_State_name, FlowContext_State_value)
	proto.RegisterType((*ArtifactParameters)(nil), "proto.ArtifactParameters")
	proto.RegisterType((*ArtifactCollectorArgs)(nil), "proto.ArtifactCollectorArgs")
	proto.RegisterType((*ArtifactCollectorResponse)(nil), "proto.ArtifactCollectorResponse")
	proto.RegisterType((*ArtifactUploadedFileInfo)(nil), "proto.ArtifactUploadedFileInfo")
	proto.RegisterType((*ArtifactCollectorContext)(nil), "proto.ArtifactCollectorContext")
	proto.RegisterType((*ClientEventTable)(nil), "proto.ClientEventTable")
	proto.RegisterType((*UploadedFileInfo)(nil), "proto.UploadedFileInfo")
	proto.RegisterType((*FlowContext)(nil), "proto.FlowContext")
}

func init() {
	proto.RegisterFile("artifact_collector.proto", fileDescriptor_ecbffc830d4bd8f6)
}

var fileDescriptor_ecbffc830d4bd8f6 = []byte{
	// 1517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xeb, 0x6e, 0xdb, 0xc6,
	0x12, 0x3e, 0xf2, 0x45, 0xb2, 0xd6, 0xb1, 0x8f, 0xb2, 0xc7, 0x8e, 0x69, 0x27, 0x41, 0xf6, 0x28,
	0xc9, 0x39, 0x6a, 0x90, 0x52, 0x8e, 0x73, 0x29, 0x90, 0xa2, 0x2d, 0x2c, 0x5b, 0x6e, 0x84, 0x3a,
	0xb6, 0x43, 0xcb, 0x09, 0x5a, 0xb4, 0x10, 0x56, 0xe4, 0x50, 0xda, 0x98, 0xe2, 0x32, 0xbb, 0x4b,
	0x29, 0xce, 0x3b, 0xf4, 0x19, 0xfa, 0x32, 0xfd, 0xdf, 0x57, 0x28, 0xda, 0xd7, 0x68, 0x81, 0x62,
	0x97, 0xa4, 0x75, 0xf1, 0x25, 0x2d, 0xd0, 0x02, 0xf1, 0x1f, 0x91, 0x73, 0xfb, 0x38, 0xb3, 0x33,
	0xdf, 0xac, 0x91, 0x45, 0x85, 0x62, 0x3e, 0x75, 0x55, 0xcb, 0xe5, 0x41, 0x00, 0xae, 0xe2, 0xc2,
	0x8e, 0x04, 0x57, 0x1c, 0xcf, 0x9a, 0x9f, 0xb5, 0x25, 0xf3, 0x53, 0x95, 0xd0, 0xa3, 0xa1, 0x62,
	0x6e, 0xa2, 0x5c, 0x5b, 0xa1, 0xae, 0x62, 0x3c, 0x94, 0xd5, 0x44, 0xdb, 0x7f, 0x13, 0x64, 0x0a,
	0x57, 0x9c, 0x44, 0x8a, 0xa7, 0xf2, 0xd7, 0xbc, 0x2d, 0x13, 0x45, 0xf9, 0x1d, 0xc2, 0x9b, 0x29,
	0xd4, 0x01, 0x15, 0xb4, 0x07, 0x0a, 0x84, 0xc4, 0x1e, 0x9a, 0x86, 0xb0, 0x6f, 0x4d, 0x93, 0xe9,
	0xca, 0xfc, 0xc6, 0x42, 0x62, 0x6a, 0xbf, 0x7c, 0xb1, 0x5b, 0x0f, 0xfb, 0xb5, 0xad, 0x5f, 0x7e,
	0xfb, 0xf5, 0xc7, 0xdc, 0x67, 0xf8, 0x61, 0x3d, 0xec, 0x33, 0xc1, 0xc3, 0x1e, 0x84, 0x8a, 0xf4,
	0xa9, 0x60, 0xb4, 0x1d, 0x80, 0x24, 0x8a, 0x93, 0x36, 0x90, 0x48, 0xf0, 0x3e, 0xf3, 0xc0, 0x23,
	0x3e, 0x17, 0x44, 0x75, 0x81, 0xbc, 0x89, 0x41, 0x9c, 0xd8, 0xe5, 0xbc, 0x01, 0x91, 0x8e, 0x0e,
	0x5f, 0xfe, 0x69, 0x16, 0x2d, 0x67, 0xe0, 0x5b, 0x59, 0x9a, 0x9b, 0xa2, 0x23, 0xb1, 0x85, 0x0a,
	0xae, 0x00, 0xaa, 0xb8, 0xb0, 0x72, 0x24, 0x57, 0x29, 0x3a, 0xd9, 0x2b, 0xbe, 0x8e, 0x8a, 0x6e,
	0xc0, 0x20, 0x54, 0x2d, 0xe6, 0x59, 0xd3, 0x46, 0x37, 0x97, 0x08, 0x1a, 0x1e, 0xbe, 0x86, 0xf2,
	0xb1, 0xe8, 0x40, 0xa8, 0xac, 0x65, 0x92, 0xab, 0xcc, 0x39, 0xe9, 0x1b, 0xde, 0x41, 0xc5, 0xac,
	0x9e, 0xd2, 0x9a, 0x22, 0xd3, 0x95, 0x62, 0xad, 0x62, 0xb2, 0x28, 0x63, 0xab, 0xd9, 0x05, 0x72,
	0xaa, 0xd4, 0x5f, 0x1f, 0xd0, 0x38, 0x74, 0xbb, 0x76, 0x39, 0xbf, 0x6b, 0x1e, 0x9c, 0xa1, 0x2b,
	0x3e, 0x46, 0x28, 0x3a, 0x2d, 0x92, 0x35, 0x4b, 0x72, 0x95, 0xf9, 0x8d, 0xd5, 0xb4, 0x3a, 0x67,
	0xab, 0x58, 0x5b, 0x37, 0x18, 0xf7, 0xf0, 0x8d, 0xa1, 0x4c, 0x03, 0xa8, 0x51, 0x44, 0xbb, 0x8c,
	0x86, 0x5a, 0x67, 0x24, 0x3c, 0x7e, 0x8d, 0x16, 0x79, 0x24, 0x5b, 0x11, 0x88, 0x96, 0x04, 0x97,
	0x87, 0x9e, 0x95, 0x27, 0xb9, 0xca, 0x54, 0x6d, 0xdb, 0x44, 0xfd, 0x1c, 0xdf, 0xde, 0x8f, 0x40,
	0x50, 0x73, 0xdc, 0x24, 0x02, 0x41, 0x12, 0x23, 0x52, 0x69, 0x76, 0x05, 0x57, 0x2a, 0x60, 0x61,
	0xe7, 0x23, 0xbb, 0xbc, 0xb8, 0x1f, 0x49, 0x72, 0x00, 0x82, 0x1c, 0x1a, 0xed, 0x46, 0xe1, 0xc1,
	0xba, 0xf9, 0x73, 0xae, 0xf0, 0x48, 0x1e, 0x80, 0x48, 0xc4, 0x18, 0x50, 0x41, 0xb1, 0x1e, 0xf0,
	0x58, 0x59, 0x05, 0x92, 0xab, 0xcc, 0xd4, 0xbe, 0x32, 0x20, 0x75, 0xfc, 0x78, 0x2f, 0xee, 0xb5,
	0x41, 0x10, 0xee, 0xa7, 0xf1, 0x4d, 0x06, 0x22, 0x0e, 0x49, 0x1b, 0x7c, 0x2e, 0x80, 0xb8, 0x34,
	0x74, 0x21, 0xd0, 0x68, 0xa3, 0xc7, 0x5c, 0x68, 0x26, 0xd1, 0x36, 0xa6, 0x9f, 0xac, 0xaf, 0x3b,
	0x59, 0x6c, 0xfc, 0x7d, 0x0e, 0x5d, 0xa3, 0x41, 0xc0, 0x07, 0x2d, 0x37, 0x96, 0x8a, 0xf7, 0x5a,
	0xbc, 0x0f, 0x42, 0x30, 0x0f, 0xa4, 0x35, 0xa7, 0x0f, 0xac, 0xf6, 0xca, 0xc0, 0xbe, 0xc0, 0xfb,
	0x0d, 0x9f, 0x28, 0x11, 0x03, 0x19, 0x00, 0x19, 0xb0, 0x20, 0x20, 0xb1, 0x04, 0x42, 0x49, 0xe2,
	0x75, 0x5a, 0x3c, 0xc2, 0x7c, 0x12, 0x09, 0x90, 0xba, 0x01, 0x59, 0x28, 0x15, 0x50, 0x4f, 0x7f,
	0xa8, 0xfe, 0x8e, 0x90, 0xf6, 0xc0, 0x3b, 0x35, 0xb4, 0x9d, 0x25, 0x03, 0xbb, 0x65, 0xfc, 0xf7,
	0x33, 0x50, 0xbc, 0x8f, 0x56, 0x5c, 0xde, 0x8b, 0x58, 0x00, 0xde, 0x70, 0xce, 0x5a, 0x54, 0x74,
	0xa4, 0xb5, 0x64, 0x0e, 0x77, 0x65, 0xd8, 0xfa, 0x63, 0x0d, 0xea, 0x2c, 0x67, 0x7e, 0x63, 0xe2,
	0x72, 0x80, 0x56, 0xcf, 0x34, 0xb4, 0x03, 0x32, 0xe2, 0xa1, 0x04, 0xbc, 0x82, 0x0a, 0xbe, 0xce,
	0x9d, 0x79, 0x69, 0x53, 0xe7, 0xf5, 0x6b, 0xc3, 0xc3, 0x4f, 0x50, 0x41, 0xc0, 0x9b, 0x18, 0xa4,
	0xb2, 0xa6, 0x0c, 0xec, 0x8d, 0x89, 0x9e, 0x1a, 0xc7, 0xce, 0x8c, 0xcb, 0xdf, 0x21, 0x2b, 0xb3,
	0x38, 0x8a, 0x02, 0x4e, 0x3d, 0xf0, 0x76, 0x58, 0x00, 0x8d, 0xd0, 0xe7, 0x18, 0xa3, 0x19, 0x5d,
	0x83, 0x14, 0xc9, 0x3c, 0xe3, 0x55, 0x34, 0xd7, 0xf7, 0x65, 0x2b, 0xa2, 0xaa, 0x6b, 0x80, 0x8a,
	0x4e, 0xa1, 0xef, 0xcb, 0x03, 0xaa, 0xba, 0xda, 0x5c, 0xb2, 0x77, 0x60, 0x26, 0x6a, 0xc6, 0x31,
	0xcf, 0xe5, 0x9f, 0x0b, 0xc3, 0xf8, 0xa7, 0x5f, 0xb0, 0xc5, 0x43, 0x05, 0x6f, 0xd5, 0xf8, 0x1c,
	0x5e, 0x9f, 0x98, 0xc3, 0x9b, 0x08, 0x49, 0x90, 0x92, 0xf1, 0x50, 0x6b, 0x17, 0x8c, 0xb6, 0x98,
	0x4a, 0xc6, 0xf3, 0x9d, 0xff, 0x0b, 0xf9, 0xe2, 0x1b, 0xa8, 0xd8, 0xa6, 0xee, 0xb1, 0x12, 0xd4,
	0xcd, 0x12, 0x1b, 0x0a, 0xf0, 0x2d, 0x34, 0x6f, 0x48, 0x02, 0x5a, 0xba, 0xdd, 0xd2, 0x4c, 0x50,
	0x22, 0xd2, 0xed, 0xa8, 0x0d, 0x34, 0x3d, 0xf6, 0x53, 0x83, 0xab, 0x89, 0x41, 0x22, 0x32, 0x06,
	0x77, 0xd1, 0xe2, 0x31, 0x0b, 0x02, 0xa3, 0x96, 0x8a, 0xf6, 0x22, 0x6b, 0xc6, 0xd8, 0x2c, 0x68,
	0x69, 0x33, 0x13, 0xe2, 0x1d, 0xb4, 0x18, 0xa7, 0xe5, 0x6e, 0xf9, 0x2c, 0x00, 0x69, 0x59, 0x86,
	0x27, 0x6f, 0x4d, 0x64, 0x31, 0x79, 0x26, 0xce, 0x42, 0x3c, 0x22, 0x91, 0x78, 0x1d, 0x2d, 0x29,
	0xae, 0x68, 0xd0, 0x9a, 0x88, 0xb6, 0x62, 0x40, 0xb1, 0xd1, 0x1d, 0x8d, 0x79, 0x6c, 0xa2, 0x9b,
	0x89, 0x07, 0xbc, 0x8d, 0xc0, 0x55, 0xe0, 0x0d, 0x5d, 0xdb, 0x27, 0x0a, 0xa4, 0xb5, 0x6a, 0x5c,
	0xd7, 0x8c, 0x51, 0x3d, 0xb5, 0xc9, 0x42, 0xd4, 0xb4, 0xc5, 0x39, 0xa0, 0x89, 0xe7, 0xda, 0x39,
	0xa0, 0x89, 0xc7, 0x5d, 0x34, 0x13, 0x70, 0x33, 0x11, 0x3a, 0xc9, 0xab, 0x69, 0x92, 0xbb, 0xbc,
	0xf3, 0x1c, 0xa4, 0xa4, 0x1d, 0x70, 0x8c, 0x1a, 0x7f, 0x8a, 0x66, 0xa5, 0xa2, 0x0a, 0xac, 0x45,
	0x92, 0xab, 0x2c, 0x6e, 0xdc, 0xbd, 0xe8, 0x48, 0xd3, 0x06, 0xb2, 0x0f, 0xb5, 0xb1, 0x93, 0xf8,
	0xe0, 0x7d, 0x94, 0xd7, 0x0f, 0xb1, 0xb4, 0xfe, 0xad, 0x8f, 0xb5, 0xf6, 0x89, 0xe1, 0x81, 0x07,
	0xb8, 0x6a, 0xac, 0x43, 0x25, 0xf5, 0x5c, 0xd3, 0x90, 0x80, 0x10, 0x5c, 0x90, 0xc4, 0x94, 0x98,
	0xb9, 0x6f, 0x9f, 0x98, 0x79, 0x4f, 0x5a, 0xcf, 0x76, 0xd2, 0x30, 0xf8, 0x36, 0x5a, 0x88, 0x25,
	0x88, 0x56, 0xc8, 0x15, 0xf3, 0x19, 0x78, 0x56, 0xc9, 0x2c, 0x84, 0x2b, 0x5a, 0xb8, 0x97, 0xca,
	0xf0, 0x0f, 0x9a, 0x8e, 0x32, 0x22, 0x6e, 0x0d, 0x98, 0xea, 0xb6, 0x04, 0xc8, 0x38, 0x50, 0xd2,
	0xba, 0x66, 0x96, 0x04, 0x33, 0x9f, 0xe1, 0x62, 0xaa, 0x97, 0x84, 0x1f, 0x07, 0x01, 0xd1, 0xb3,
	0x73, 0x86, 0xc3, 0x89, 0x76, 0xd5, 0x22, 0x26, 0x48, 0x1a, 0xc0, 0x26, 0xcd, 0x2e, 0x93, 0x44,
	0xd1, 0x63, 0xbd, 0x12, 0xb5, 0x37, 0x17, 0x23, 0xc4, 0xa5, 0xc7, 0xb1, 0x2a, 0x79, 0x2c, 0xdc,
	0x84, 0xaa, 0x34, 0x41, 0x65, 0xd1, 0x5e, 0x31, 0xd5, 0x75, 0x92, 0x28, 0x78, 0x09, 0xcd, 0x7a,
	0x4c, 0xa8, 0x13, 0x33, 0xae, 0x73, 0x4e, 0xf2, 0x52, 0x7e, 0x86, 0x66, 0x4d, 0xf5, 0x70, 0x11,
	0xcd, 0x1e, 0xed, 0x1d, 0xd6, 0x9b, 0xa5, 0x7f, 0xe1, 0x79, 0x54, 0x70, 0x8e, 0xf6, 0xf6, 0x1a,
	0x7b, 0x5f, 0x96, 0x72, 0x78, 0x11, 0xa1, 0x66, 0xdd, 0x79, 0xde, 0xd8, 0xdb, 0x6c, 0xd6, 0xb7,
	0x4b, 0x53, 0xda, 0xae, 0xee, 0x38, 0xfb, 0x4e, 0x69, 0x1a, 0x5f, 0x41, 0x73, 0x9b, 0xce, 0xd6,
	0xb3, 0xc6, 0xcb, 0xfa, 0x76, 0x69, 0xa6, 0xdc, 0x45, 0xa5, 0x2d, 0x53, 0xb9, 0x7a, 0x1f, 0x42,
	0xd5, 0xd4, 0x1b, 0x5c, 0xef, 0xde, 0x3e, 0x08, 0x3d, 0xaa, 0x66, 0xc6, 0x66, 0x9c, 0xec, 0x15,
	0x3f, 0x1d, 0x5f, 0xa3, 0xef, 0x9f, 0xdc, 0xa1, 0x79, 0xf9, 0x08, 0x95, 0xfe, 0x09, 0x8e, 0xfa,
	0xbd, 0x88, 0xe6, 0x77, 0x34, 0xb3, 0xa7, 0xb4, 0x74, 0x39, 0x45, 0x3c, 0x3a, 0x87, 0x22, 0x6a,
	0xff, 0x31, 0x87, 0xbc, 0x80, 0xe6, 0x9d, 0xed, 0x9d, 0x6d, 0xaa, 0x40, 0xab, 0xc6, 0x78, 0xe3,
	0xe9, 0x19, 0x5a, 0xc8, 0x5f, 0xec, 0x38, 0xc1, 0x15, 0xf7, 0x11, 0x0e, 0x41, 0x0d, 0xb8, 0x38,
	0x4e, 0xe6, 0xac, 0xa5, 0x9b, 0x36, 0xd9, 0xb1, 0x4e, 0x29, 0xd5, 0x98, 0x31, 0x3b, 0xd4, 0xf7,
	0x94, 0xf8, 0x02, 0x66, 0xc9, 0xd6, 0xd0, 0x64, 0x05, 0x6b, 0x8f, 0xcd, 0x27, 0x54, 0xf1, 0xc7,
	0x9b, 0x24, 0x60, 0x52, 0xe9, 0x29, 0x31, 0x7e, 0x24, 0x0b, 0x43, 0xa8, 0x24, 0x11, 0x15, 0x2a,
	0xdb, 0x8b, 0x7a, 0xe3, 0xd8, 0x93, 0x44, 0xf4, 0xf5, 0x65, 0x44, 0x54, 0xfb, 0xbf, 0xc1, 0xf8,
	0x2f, 0xbe, 0xd5, 0xd4, 0x36, 0x24, 0x3c, 0xbd, 0x10, 0x9c, 0x62, 0x18, 0x6b, 0xfb, 0x5c, 0xc6,
	0xe2, 0x7f, 0x8a, 0xb1, 0x6a, 0xf7, 0x0d, 0xc6, 0xff, 0xf0, 0x9d, 0x67, 0x7c, 0x40, 0x7a, 0x34,
	0x3c, 0x21, 0x46, 0xab, 0xd7, 0x7f, 0xe2, 0x68, 0x6e, 0x1d, 0xe0, 0x02, 0xeb, 0x83, 0x7d, 0x29,
	0xbf, 0x7d, 0x7b, 0x19, 0xbf, 0xd5, 0xee, 0x19, 0x9c, 0x3b, 0xb8, 0x7c, 0x16, 0x27, 0x8d, 0xee,
	0x11, 0xc9, 0x89, 0x4f, 0x85, 0x7d, 0x2e, 0x17, 0x7e, 0xf1, 0x1e, 0x2e, 0xac, 0x59, 0x06, 0x00,
	0xe3, 0x85, 0xdd, 0xf4, 0x38, 0xb4, 0xb9, 0xbd, 0x96, 0x9b, 0x4a, 0x59, 0xf2, 0x3d, 0x9b, 0xd1,
	0x1e, 0x27, 0x51, 0x2b, 0x05, 0x18, 0xe9, 0xf0, 0x0f, 0x81, 0x37, 0x1f, 0x9d, 0xb3, 0x48, 0x2f,
	0x18, 0xa3, 0x91, 0xed, 0xda, 0x19, 0x65, 0x8f, 0x65, 0xc3, 0xaf, 0x0d, 0xe3, 0xb3, 0x85, 0x37,
	0x1b, 0xc3, 0xee, 0x24, 0x02, 0x54, 0x2c, 0x42, 0x49, 0x28, 0x91, 0x60, 0x4a, 0x38, 0x72, 0x3b,
	0xd7, 0x9c, 0xea, 0x33, 0x08, 0xbc, 0xe4, 0x56, 0x68, 0x7a, 0x5e, 0x75, 0xa1, 0x67, 0x8f, 0xde,
	0xd2, 0x3f, 0x78, 0x5a, 0xff, 0xfb, 0x08, 0xbc, 0xf6, 0xf0, 0x9b, 0x07, 0x83, 0xc1, 0xc0, 0xee,
	0x43, 0xc0, 0x5d, 0xe6, 0xc1, 0x5b, 0xdb, 0xe5, 0xbd, 0x6a, 0x87, 0x07, 0x34, 0xec, 0x54, 0x13,
	0xa1, 0xa0, 0x91, 0xe2, 0xa2, 0xaa, 0x4b, 0x9b, 0xfe, 0x5b, 0xd8, 0xce, 0x9b, 0x9f, 0x87, 0x7f,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0x69, 0xfe, 0xc7, 0x65, 0x0e, 0x00, 0x00,
}

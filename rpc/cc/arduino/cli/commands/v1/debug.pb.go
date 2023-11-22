// This file is part of arduino-cli.
//
// Copyright 2023 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: cc/arduino/cli/commands/v1/debug.proto

package commands

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The top-level message sent by the client for the `Debug` method.
// Multiple `DebugRequest` messages can be sent but the first message
// must contain a `GetDebugConfigRequest` message to initialize the debug
// session. All subsequent messages must contain bytes to be sent to the debug
// session and must not contain a `GetDebugConfigRequest` message.
type DebugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provides information to the debug that specifies which is the target.
	// The first `DebugRequest` message must contain a `GetDebugConfigRequest`
	// message.
	DebugRequest *GetDebugConfigRequest `protobuf:"bytes,1,opt,name=debug_request,json=debugRequest,proto3" json:"debug_request,omitempty"`
	// The data to be sent to the target being monitored.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Set this to true to send and Interrupt signal to the debugger process
	SendInterrupt bool `protobuf:"varint,3,opt,name=send_interrupt,json=sendInterrupt,proto3" json:"send_interrupt,omitempty"`
}

func (x *DebugRequest) Reset() {
	*x = DebugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugRequest) ProtoMessage() {}

func (x *DebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugRequest.ProtoReflect.Descriptor instead.
func (*DebugRequest) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP(), []int{0}
}

func (x *DebugRequest) GetDebugRequest() *GetDebugConfigRequest {
	if x != nil {
		return x.DebugRequest
	}
	return nil
}

func (x *DebugRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DebugRequest) GetSendInterrupt() bool {
	if x != nil {
		return x.SendInterrupt
	}
	return false
}

type GetDebugConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Arduino Core Service instance from the `Init` response.
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// Fully qualified board name of the board in use
	// (e.g., `arduino:samd:mkr1000`). If this is omitted, the FQBN attached to
	// the sketch will be used.
	Fqbn string `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	// Path to the sketch that is running on the board. The compiled executable
	// is expected to be located under this path.
	SketchPath string `protobuf:"bytes,3,opt,name=sketch_path,json=sketchPath,proto3" json:"sketch_path,omitempty"`
	// Port of the debugger (optional).
	Port *Port `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	// Which GDB command interpreter to use.
	Interpreter string `protobuf:"bytes,5,opt,name=interpreter,proto3" json:"interpreter,omitempty"`
	// Directory containing the compiled executable. If `import_dir` is not
	// specified, the executable is assumed to be in
	// `{sketch_path}/build/{fqbn}/`.
	ImportDir string `protobuf:"bytes,8,opt,name=import_dir,json=importDir,proto3" json:"import_dir,omitempty"`
	// The programmer to use for debugging.
	Programmer string `protobuf:"bytes,9,opt,name=programmer,proto3" json:"programmer,omitempty"`
}

func (x *GetDebugConfigRequest) Reset() {
	*x = GetDebugConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDebugConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDebugConfigRequest) ProtoMessage() {}

func (x *GetDebugConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDebugConfigRequest.ProtoReflect.Descriptor instead.
func (*GetDebugConfigRequest) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP(), []int{1}
}

func (x *GetDebugConfigRequest) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *GetDebugConfigRequest) GetFqbn() string {
	if x != nil {
		return x.Fqbn
	}
	return ""
}

func (x *GetDebugConfigRequest) GetSketchPath() string {
	if x != nil {
		return x.SketchPath
	}
	return ""
}

func (x *GetDebugConfigRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *GetDebugConfigRequest) GetInterpreter() string {
	if x != nil {
		return x.Interpreter
	}
	return ""
}

func (x *GetDebugConfigRequest) GetImportDir() string {
	if x != nil {
		return x.ImportDir
	}
	return ""
}

func (x *GetDebugConfigRequest) GetProgrammer() string {
	if x != nil {
		return x.Programmer
	}
	return ""
}

type DebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Incoming data from the debugger tool.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Incoming error output from the debugger tool.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DebugResponse) Reset() {
	*x = DebugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugResponse) ProtoMessage() {}

func (x *DebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugResponse.ProtoReflect.Descriptor instead.
func (*DebugResponse) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP(), []int{2}
}

func (x *DebugResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DebugResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetDebugConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The executable binary to debug
	Executable string `protobuf:"bytes,1,opt,name=executable,proto3" json:"executable,omitempty"`
	// The toolchain type used for the build (for example "gcc")
	Toolchain string `protobuf:"bytes,2,opt,name=toolchain,proto3" json:"toolchain,omitempty"`
	// The toolchain directory
	ToolchainPath string `protobuf:"bytes,3,opt,name=toolchain_path,json=toolchainPath,proto3" json:"toolchain_path,omitempty"`
	// The toolchain architecture prefix (for example "arm-none-eabi")
	ToolchainPrefix string `protobuf:"bytes,4,opt,name=toolchain_prefix,json=toolchainPrefix,proto3" json:"toolchain_prefix,omitempty"`
	// The GDB server type used to connect to the programmer/board (for example
	// "openocd")
	Server string `protobuf:"bytes,5,opt,name=server,proto3" json:"server,omitempty"`
	// The GDB server directory
	ServerPath string `protobuf:"bytes,6,opt,name=server_path,json=serverPath,proto3" json:"server_path,omitempty"`
	// Extra configuration parameters wrt toolchain
	ToolchainConfiguration *anypb.Any `protobuf:"bytes,7,opt,name=toolchain_configuration,json=toolchainConfiguration,proto3" json:"toolchain_configuration,omitempty"`
	// Extra configuration parameters wrt GDB server
	ServerConfiguration *anypb.Any `protobuf:"bytes,8,opt,name=server_configuration,json=serverConfiguration,proto3" json:"server_configuration,omitempty"`
	// Custom debugger configurations (not handled directly by Arduino CLI but
	// provided for 3rd party plugins/debuggers). The map keys identifies which
	// 3rd party plugin/debugger is referred, the map string values contains a
	// JSON configuration for it.
	CustomConfigs map[string]string `protobuf:"bytes,9,rep,name=custom_configs,json=customConfigs,proto3" json:"custom_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the SVD file to use
	SvdFile string `protobuf:"bytes,10,opt,name=svd_file,json=svdFile,proto3" json:"svd_file,omitempty"`
	// The programmer specified in the request
	Programmer string `protobuf:"bytes,11,opt,name=programmer,proto3" json:"programmer,omitempty"`
}

func (x *GetDebugConfigResponse) Reset() {
	*x = GetDebugConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDebugConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDebugConfigResponse) ProtoMessage() {}

func (x *GetDebugConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDebugConfigResponse.ProtoReflect.Descriptor instead.
func (*GetDebugConfigResponse) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP(), []int{3}
}

func (x *GetDebugConfigResponse) GetExecutable() string {
	if x != nil {
		return x.Executable
	}
	return ""
}

func (x *GetDebugConfigResponse) GetToolchain() string {
	if x != nil {
		return x.Toolchain
	}
	return ""
}

func (x *GetDebugConfigResponse) GetToolchainPath() string {
	if x != nil {
		return x.ToolchainPath
	}
	return ""
}

func (x *GetDebugConfigResponse) GetToolchainPrefix() string {
	if x != nil {
		return x.ToolchainPrefix
	}
	return ""
}

func (x *GetDebugConfigResponse) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *GetDebugConfigResponse) GetServerPath() string {
	if x != nil {
		return x.ServerPath
	}
	return ""
}

func (x *GetDebugConfigResponse) GetToolchainConfiguration() *anypb.Any {
	if x != nil {
		return x.ToolchainConfiguration
	}
	return nil
}

func (x *GetDebugConfigResponse) GetServerConfiguration() *anypb.Any {
	if x != nil {
		return x.ServerConfiguration
	}
	return nil
}

func (x *GetDebugConfigResponse) GetCustomConfigs() map[string]string {
	if x != nil {
		return x.CustomConfigs
	}
	return nil
}

func (x *GetDebugConfigResponse) GetSvdFile() string {
	if x != nil {
		return x.SvdFile
	}
	return ""
}

func (x *GetDebugConfigResponse) GetProgrammer() string {
	if x != nil {
		return x.Programmer
	}
	return ""
}

// Configurations specific for the 'gcc' toolchain
type DebugGCCToolchainConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DebugGCCToolchainConfiguration) Reset() {
	*x = DebugGCCToolchainConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugGCCToolchainConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugGCCToolchainConfiguration) ProtoMessage() {}

func (x *DebugGCCToolchainConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugGCCToolchainConfiguration.ProtoReflect.Descriptor instead.
func (*DebugGCCToolchainConfiguration) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP(), []int{4}
}

// Configuration specific for the 'openocd` server
type DebugOpenOCDServerConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path to openocd
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// path to openocd scripts
	ScriptsDir string `protobuf:"bytes,2,opt,name=scripts_dir,json=scriptsDir,proto3" json:"scripts_dir,omitempty"`
	// list of scripts to execute by openocd
	Scripts []string `protobuf:"bytes,3,rep,name=scripts,proto3" json:"scripts,omitempty"`
}

func (x *DebugOpenOCDServerConfiguration) Reset() {
	*x = DebugOpenOCDServerConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugOpenOCDServerConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugOpenOCDServerConfiguration) ProtoMessage() {}

func (x *DebugOpenOCDServerConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugOpenOCDServerConfiguration.ProtoReflect.Descriptor instead.
func (*DebugOpenOCDServerConfiguration) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP(), []int{5}
}

func (x *DebugOpenOCDServerConfiguration) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DebugOpenOCDServerConfiguration) GetScriptsDir() string {
	if x != nil {
		return x.ScriptsDir
	}
	return ""
}

func (x *DebugOpenOCDServerConfiguration) GetScripts() []string {
	if x != nil {
		return x.Scripts
	}
	return nil
}

var File_cc_arduino_cli_commands_v1_debug_proto protoreflect.FileDescriptor

var file_cc_arduino_cli_commands_v1_debug_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64,
	0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f,
	0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63,
	0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa1, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x56, 0x0a, 0x0d, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64,
	0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72,
	0x75, 0x70, 0x74, 0x22, 0xa5, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x71, 0x62, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x71, 0x62, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x34, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e,
	0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0d, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xe4, 0x04, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x4d, 0x0a, 0x17, 0x74, 0x6f,
	0x6f, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x16, 0x74, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x13, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x63, 0x2e,
	0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x76, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x76, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x72, 0x1a, 0x40, 0x0a, 0x12, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20, 0x0a,
	0x1e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x47, 0x43, 0x43, 0x54, 0x6f, 0x6f, 0x6c, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x70, 0x0a, 0x1f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4f, 0x70, 0x65, 0x6e, 0x4f, 0x43, 0x44, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x73, 0x44, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d,
	0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69,
	0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cc_arduino_cli_commands_v1_debug_proto_rawDescOnce sync.Once
	file_cc_arduino_cli_commands_v1_debug_proto_rawDescData = file_cc_arduino_cli_commands_v1_debug_proto_rawDesc
)

func file_cc_arduino_cli_commands_v1_debug_proto_rawDescGZIP() []byte {
	file_cc_arduino_cli_commands_v1_debug_proto_rawDescOnce.Do(func() {
		file_cc_arduino_cli_commands_v1_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_arduino_cli_commands_v1_debug_proto_rawDescData)
	})
	return file_cc_arduino_cli_commands_v1_debug_proto_rawDescData
}

var file_cc_arduino_cli_commands_v1_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cc_arduino_cli_commands_v1_debug_proto_goTypes = []interface{}{
	(*DebugRequest)(nil),                    // 0: cc.arduino.cli.commands.v1.DebugRequest
	(*GetDebugConfigRequest)(nil),           // 1: cc.arduino.cli.commands.v1.GetDebugConfigRequest
	(*DebugResponse)(nil),                   // 2: cc.arduino.cli.commands.v1.DebugResponse
	(*GetDebugConfigResponse)(nil),          // 3: cc.arduino.cli.commands.v1.GetDebugConfigResponse
	(*DebugGCCToolchainConfiguration)(nil),  // 4: cc.arduino.cli.commands.v1.DebugGCCToolchainConfiguration
	(*DebugOpenOCDServerConfiguration)(nil), // 5: cc.arduino.cli.commands.v1.DebugOpenOCDServerConfiguration
	nil,                                     // 6: cc.arduino.cli.commands.v1.GetDebugConfigResponse.CustomConfigsEntry
	(*Instance)(nil),                        // 7: cc.arduino.cli.commands.v1.Instance
	(*Port)(nil),                            // 8: cc.arduino.cli.commands.v1.Port
	(*anypb.Any)(nil),                       // 9: google.protobuf.Any
}
var file_cc_arduino_cli_commands_v1_debug_proto_depIdxs = []int32{
	1, // 0: cc.arduino.cli.commands.v1.DebugRequest.debug_request:type_name -> cc.arduino.cli.commands.v1.GetDebugConfigRequest
	7, // 1: cc.arduino.cli.commands.v1.GetDebugConfigRequest.instance:type_name -> cc.arduino.cli.commands.v1.Instance
	8, // 2: cc.arduino.cli.commands.v1.GetDebugConfigRequest.port:type_name -> cc.arduino.cli.commands.v1.Port
	9, // 3: cc.arduino.cli.commands.v1.GetDebugConfigResponse.toolchain_configuration:type_name -> google.protobuf.Any
	9, // 4: cc.arduino.cli.commands.v1.GetDebugConfigResponse.server_configuration:type_name -> google.protobuf.Any
	6, // 5: cc.arduino.cli.commands.v1.GetDebugConfigResponse.custom_configs:type_name -> cc.arduino.cli.commands.v1.GetDebugConfigResponse.CustomConfigsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cc_arduino_cli_commands_v1_debug_proto_init() }
func file_cc_arduino_cli_commands_v1_debug_proto_init() {
	if File_cc_arduino_cli_commands_v1_debug_proto != nil {
		return
	}
	file_cc_arduino_cli_commands_v1_common_proto_init()
	file_cc_arduino_cli_commands_v1_port_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugRequest); i {
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
		file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDebugConfigRequest); i {
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
		file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugResponse); i {
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
		file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDebugConfigResponse); i {
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
		file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugGCCToolchainConfiguration); i {
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
		file_cc_arduino_cli_commands_v1_debug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugOpenOCDServerConfiguration); i {
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
			RawDescriptor: file_cc_arduino_cli_commands_v1_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cc_arduino_cli_commands_v1_debug_proto_goTypes,
		DependencyIndexes: file_cc_arduino_cli_commands_v1_debug_proto_depIdxs,
		MessageInfos:      file_cc_arduino_cli_commands_v1_debug_proto_msgTypes,
	}.Build()
	File_cc_arduino_cli_commands_v1_debug_proto = out.File
	file_cc_arduino_cli_commands_v1_debug_proto_rawDesc = nil
	file_cc_arduino_cli_commands_v1_debug_proto_goTypes = nil
	file_cc_arduino_cli_commands_v1_debug_proto_depIdxs = nil
}
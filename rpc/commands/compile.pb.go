// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: commands/compile.proto

package commands

import (
	proto "github.com/golang/protobuf/proto"
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

type CompileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance                      *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`                                                                                      // Arduino Core Service instance from the `Init` response.
	Fqbn                          string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`                                                                                              // Fully Qualified Board Name, e.g.: `arduino:avr:uno`. If this field is not defined, the FQBN of the board attached to the sketch via the `BoardAttach` method is used.
	SketchPath                    string    `protobuf:"bytes,3,opt,name=sketchPath,proto3" json:"sketchPath,omitempty"`                                                                                  // The path where the sketch is stored.
	ShowProperties                bool      `protobuf:"varint,4,opt,name=showProperties,proto3" json:"showProperties,omitempty"`                                                                         // Show all build preferences used instead of compiling.
	Preprocess                    bool      `protobuf:"varint,5,opt,name=preprocess,proto3" json:"preprocess,omitempty"`                                                                                 // Print preprocessed code to stdout instead of compiling.
	BuildCachePath                string    `protobuf:"bytes,6,opt,name=buildCachePath,proto3" json:"buildCachePath,omitempty"`                                                                          // Builds of 'core.a' are saved into this path to be cached and reused.
	BuildPath                     string    `protobuf:"bytes,7,opt,name=buildPath,proto3" json:"buildPath,omitempty"`                                                                                    // Path to use to store the files used for the compilation. If omitted, a directory will be created in the operating system's default temporary path.
	BuildProperties               []string  `protobuf:"bytes,8,rep,name=buildProperties,proto3" json:"buildProperties,omitempty"`                                                                        // List of custom build properties separated by commas.
	Warnings                      string    `protobuf:"bytes,9,opt,name=warnings,proto3" json:"warnings,omitempty"`                                                                                      // Used to tell gcc which warning level to use. The level names are: "none", "default", "more" and "all".
	Verbose                       bool      `protobuf:"varint,10,opt,name=verbose,proto3" json:"verbose,omitempty"`                                                                                      // Turns on verbose mode.
	Quiet                         bool      `protobuf:"varint,11,opt,name=quiet,proto3" json:"quiet,omitempty"`                                                                                          // Suppresses almost every output.
	VidPid                        string    `protobuf:"bytes,12,opt,name=vidPid,proto3" json:"vidPid,omitempty"`                                                                                         // VID/PID specific build properties.
	Jobs                          int32     `protobuf:"varint,14,opt,name=jobs,proto3" json:"jobs,omitempty"`                                                                                            // The max number of concurrent compiler instances to run (as `make -jx`). If jobs is set to 0, it will use the number of available CPUs as the maximum.
	Libraries                     []string  `protobuf:"bytes,15,rep,name=libraries,proto3" json:"libraries,omitempty"`                                                                                   // List of custom libraries paths separated by commas.
	OptimizeForDebug              bool      `protobuf:"varint,16,opt,name=optimizeForDebug,proto3" json:"optimizeForDebug,omitempty"`                                                                    // Optimize compile output for debug, not for release.
	ExportDir                     string    `protobuf:"bytes,18,opt,name=export_dir,json=exportDir,proto3" json:"export_dir,omitempty"`                                                                  // Optional: save the build artifacts in this directory, the directory must exist.
	Clean                         bool      `protobuf:"varint,19,opt,name=clean,proto3" json:"clean,omitempty"`                                                                                          // Optional: cleanup the build folder and do not use any previously cached build
	ExportBinaries                bool      `protobuf:"varint,20,opt,name=export_binaries,json=exportBinaries,proto3" json:"export_binaries,omitempty"`                                                  // When set to `true` the compiled binary will be copied to the export directory.
	CreateCompilationDatabaseOnly bool      `protobuf:"varint,21,opt,name=create_compilation_database_only,json=createCompilationDatabaseOnly,proto3" json:"create_compilation_database_only,omitempty"` // When set to `true` only the compilation database will be produced and no actual build will be performed.
}

func (x *CompileReq) Reset() {
	*x = CompileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_compile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileReq) ProtoMessage() {}

func (x *CompileReq) ProtoReflect() protoreflect.Message {
	mi := &file_commands_compile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileReq.ProtoReflect.Descriptor instead.
func (*CompileReq) Descriptor() ([]byte, []int) {
	return file_commands_compile_proto_rawDescGZIP(), []int{0}
}

func (x *CompileReq) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *CompileReq) GetFqbn() string {
	if x != nil {
		return x.Fqbn
	}
	return ""
}

func (x *CompileReq) GetSketchPath() string {
	if x != nil {
		return x.SketchPath
	}
	return ""
}

func (x *CompileReq) GetShowProperties() bool {
	if x != nil {
		return x.ShowProperties
	}
	return false
}

func (x *CompileReq) GetPreprocess() bool {
	if x != nil {
		return x.Preprocess
	}
	return false
}

func (x *CompileReq) GetBuildCachePath() string {
	if x != nil {
		return x.BuildCachePath
	}
	return ""
}

func (x *CompileReq) GetBuildPath() string {
	if x != nil {
		return x.BuildPath
	}
	return ""
}

func (x *CompileReq) GetBuildProperties() []string {
	if x != nil {
		return x.BuildProperties
	}
	return nil
}

func (x *CompileReq) GetWarnings() string {
	if x != nil {
		return x.Warnings
	}
	return ""
}

func (x *CompileReq) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

func (x *CompileReq) GetQuiet() bool {
	if x != nil {
		return x.Quiet
	}
	return false
}

func (x *CompileReq) GetVidPid() string {
	if x != nil {
		return x.VidPid
	}
	return ""
}

func (x *CompileReq) GetJobs() int32 {
	if x != nil {
		return x.Jobs
	}
	return 0
}

func (x *CompileReq) GetLibraries() []string {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *CompileReq) GetOptimizeForDebug() bool {
	if x != nil {
		return x.OptimizeForDebug
	}
	return false
}

func (x *CompileReq) GetExportDir() string {
	if x != nil {
		return x.ExportDir
	}
	return ""
}

func (x *CompileReq) GetClean() bool {
	if x != nil {
		return x.Clean
	}
	return false
}

func (x *CompileReq) GetExportBinaries() bool {
	if x != nil {
		return x.ExportBinaries
	}
	return false
}

func (x *CompileReq) GetCreateCompilationDatabaseOnly() bool {
	if x != nil {
		return x.CreateCompilationDatabaseOnly
	}
	return false
}

type CompileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutStream              []byte                   `protobuf:"bytes,1,opt,name=out_stream,json=outStream,proto3" json:"out_stream,omitempty"`                                          // The output of the compilation process.
	ErrStream              []byte                   `protobuf:"bytes,2,opt,name=err_stream,json=errStream,proto3" json:"err_stream,omitempty"`                                          // The error output of the compilation process.
	BuildPath              string                   `protobuf:"bytes,3,opt,name=build_path,json=buildPath,proto3" json:"build_path,omitempty"`                                          // The compiler build path
	UsedLibraries          []*Library               `protobuf:"bytes,4,rep,name=used_libraries,json=usedLibraries,proto3" json:"used_libraries,omitempty"`                              // The libraries used in the build
	ExecutableSectionsSize []*ExecutableSectionSize `protobuf:"bytes,5,rep,name=executable_sections_size,json=executableSectionsSize,proto3" json:"executable_sections_size,omitempty"` // The size of the executable split by sections
}

func (x *CompileResp) Reset() {
	*x = CompileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_compile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileResp) ProtoMessage() {}

func (x *CompileResp) ProtoReflect() protoreflect.Message {
	mi := &file_commands_compile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileResp.ProtoReflect.Descriptor instead.
func (*CompileResp) Descriptor() ([]byte, []int) {
	return file_commands_compile_proto_rawDescGZIP(), []int{1}
}

func (x *CompileResp) GetOutStream() []byte {
	if x != nil {
		return x.OutStream
	}
	return nil
}

func (x *CompileResp) GetErrStream() []byte {
	if x != nil {
		return x.ErrStream
	}
	return nil
}

func (x *CompileResp) GetBuildPath() string {
	if x != nil {
		return x.BuildPath
	}
	return ""
}

func (x *CompileResp) GetUsedLibraries() []*Library {
	if x != nil {
		return x.UsedLibraries
	}
	return nil
}

func (x *CompileResp) GetExecutableSectionsSize() []*ExecutableSectionSize {
	if x != nil {
		return x.ExecutableSectionsSize
	}
	return nil
}

type ExecutableSectionSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size    int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	MaxSize int64  `protobuf:"varint,3,opt,name=maxSize,proto3" json:"maxSize,omitempty"`
}

func (x *ExecutableSectionSize) Reset() {
	*x = ExecutableSectionSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_compile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutableSectionSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutableSectionSize) ProtoMessage() {}

func (x *ExecutableSectionSize) ProtoReflect() protoreflect.Message {
	mi := &file_commands_compile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutableSectionSize.ProtoReflect.Descriptor instead.
func (*ExecutableSectionSize) Descriptor() ([]byte, []int) {
	return file_commands_compile_proto_rawDescGZIP(), []int{2}
}

func (x *ExecutableSectionSize) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecutableSectionSize) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ExecutableSectionSize) GetMaxSize() int64 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

var File_commands_compile_proto protoreflect.FileDescriptor

var file_commands_compile_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64,
	0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x05, 0x0a,
	0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x3d, 0x0a, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71,
	0x62, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x62, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x50, 0x61, 0x74, 0x68, 0x12, 0x26,
	0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x0f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x69, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x71, 0x75, 0x69,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x50, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x69, 0x64, 0x50, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x6f,
	0x62, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10,
	0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65,
	0x46, 0x6f, 0x72, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x6e,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x1d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22,
	0x9d, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x65, 0x72, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x47, 0x0a, 0x0e,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e,
	0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x64, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x18, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64,
	0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x16, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x59, 0x0a, 0x15, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f,
	0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commands_compile_proto_rawDescOnce sync.Once
	file_commands_compile_proto_rawDescData = file_commands_compile_proto_rawDesc
)

func file_commands_compile_proto_rawDescGZIP() []byte {
	file_commands_compile_proto_rawDescOnce.Do(func() {
		file_commands_compile_proto_rawDescData = protoimpl.X.CompressGZIP(file_commands_compile_proto_rawDescData)
	})
	return file_commands_compile_proto_rawDescData
}

var file_commands_compile_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commands_compile_proto_goTypes = []interface{}{
	(*CompileReq)(nil),            // 0: cc.arduino.cli.commands.CompileReq
	(*CompileResp)(nil),           // 1: cc.arduino.cli.commands.CompileResp
	(*ExecutableSectionSize)(nil), // 2: cc.arduino.cli.commands.ExecutableSectionSize
	(*Instance)(nil),              // 3: cc.arduino.cli.commands.Instance
	(*Library)(nil),               // 4: cc.arduino.cli.commands.Library
}
var file_commands_compile_proto_depIdxs = []int32{
	3, // 0: cc.arduino.cli.commands.CompileReq.instance:type_name -> cc.arduino.cli.commands.Instance
	4, // 1: cc.arduino.cli.commands.CompileResp.used_libraries:type_name -> cc.arduino.cli.commands.Library
	2, // 2: cc.arduino.cli.commands.CompileResp.executable_sections_size:type_name -> cc.arduino.cli.commands.ExecutableSectionSize
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_commands_compile_proto_init() }
func file_commands_compile_proto_init() {
	if File_commands_compile_proto != nil {
		return
	}
	file_commands_common_proto_init()
	file_commands_lib_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commands_compile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileReq); i {
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
		file_commands_compile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileResp); i {
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
		file_commands_compile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutableSectionSize); i {
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
			RawDescriptor: file_commands_compile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commands_compile_proto_goTypes,
		DependencyIndexes: file_commands_compile_proto_depIdxs,
		MessageInfos:      file_commands_compile_proto_msgTypes,
	}.Build()
	File_commands_compile_proto = out.File
	file_commands_compile_proto_rawDesc = nil
	file_commands_compile_proto_goTypes = nil
	file_commands_compile_proto_depIdxs = nil
}

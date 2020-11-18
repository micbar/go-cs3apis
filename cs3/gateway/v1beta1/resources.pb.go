// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/gateway/v1beta1/resources.proto

package gatewayv1beta1

import (
	fmt "fmt"
	_ "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A file upload protocol object stores information about
// uploading resources using a specific protocol.
type FileUploadProtocol struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The protocol to be followed.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// REQUIRED.
	// The endpoint where to upload the data.
	// The value MUST be a Uniform Resource Identifier (URI)
	// as specified in RFC 3986.
	UploadEndpoint string `protobuf:"bytes,3,opt,name=upload_endpoint,json=uploadEndpoint,proto3" json:"upload_endpoint,omitempty"`
	// REQUIRED.
	// List of available checksums
	// the client can use when sending
	// the file.
	AvailableChecksums []*v1beta11.ResourceChecksumPriority `protobuf:"bytes,4,rep,name=available_checksums,json=availableChecksums,proto3" json:"available_checksums,omitempty"`
	// OPTIONAL.
	// A token that MUST be validated by the data gateway for the upload.
	// Only makes sense for uploads passing through the data gateway.
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileUploadProtocol) Reset()         { *m = FileUploadProtocol{} }
func (m *FileUploadProtocol) String() string { return proto.CompactTextString(m) }
func (*FileUploadProtocol) ProtoMessage()    {}
func (*FileUploadProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dff5533c2cc40ba, []int{0}
}

func (m *FileUploadProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileUploadProtocol.Unmarshal(m, b)
}
func (m *FileUploadProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileUploadProtocol.Marshal(b, m, deterministic)
}
func (m *FileUploadProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileUploadProtocol.Merge(m, src)
}
func (m *FileUploadProtocol) XXX_Size() int {
	return xxx_messageInfo_FileUploadProtocol.Size(m)
}
func (m *FileUploadProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_FileUploadProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_FileUploadProtocol proto.InternalMessageInfo

func (m *FileUploadProtocol) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *FileUploadProtocol) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *FileUploadProtocol) GetUploadEndpoint() string {
	if m != nil {
		return m.UploadEndpoint
	}
	return ""
}

func (m *FileUploadProtocol) GetAvailableChecksums() []*v1beta11.ResourceChecksumPriority {
	if m != nil {
		return m.AvailableChecksums
	}
	return nil
}

func (m *FileUploadProtocol) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// A file download protocol object stores information about
// downloading resources using a specific protocol.
type FileDownloadProtocol struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The protocol to be followed.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// REQUIRED.
	// The endpoint where to download the data.
	// The value MUST be a Uniform Resource Identifier (URI)
	// as specified in RFC 3986.
	DownloadEndpoint string `protobuf:"bytes,3,opt,name=download_endpoint,json=downloadEndpoint,proto3" json:"download_endpoint,omitempty"`
	// OPTIONAL.
	// A token that MUST be validated by the data gateway for the download.
	// Only makes sense for downloads passing through the data gateway.
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDownloadProtocol) Reset()         { *m = FileDownloadProtocol{} }
func (m *FileDownloadProtocol) String() string { return proto.CompactTextString(m) }
func (*FileDownloadProtocol) ProtoMessage()    {}
func (*FileDownloadProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dff5533c2cc40ba, []int{1}
}

func (m *FileDownloadProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDownloadProtocol.Unmarshal(m, b)
}
func (m *FileDownloadProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDownloadProtocol.Marshal(b, m, deterministic)
}
func (m *FileDownloadProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDownloadProtocol.Merge(m, src)
}
func (m *FileDownloadProtocol) XXX_Size() int {
	return xxx_messageInfo_FileDownloadProtocol.Size(m)
}
func (m *FileDownloadProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDownloadProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_FileDownloadProtocol proto.InternalMessageInfo

func (m *FileDownloadProtocol) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *FileDownloadProtocol) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *FileDownloadProtocol) GetDownloadEndpoint() string {
	if m != nil {
		return m.DownloadEndpoint
	}
	return ""
}

func (m *FileDownloadProtocol) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*FileUploadProtocol)(nil), "cs3.gateway.v1beta1.FileUploadProtocol")
	proto.RegisterType((*FileDownloadProtocol)(nil), "cs3.gateway.v1beta1.FileDownloadProtocol")
}

func init() {
	proto.RegisterFile("cs3/gateway/v1beta1/resources.proto", fileDescriptor_6dff5533c2cc40ba)
}

var fileDescriptor_6dff5533c2cc40ba = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xce, 0x0f, 0x9b, 0x02, 0xd9, 0xa6, 0x04, 0xe6, 0x99, 0x0d, 0x42, 0x76, 0xb1, 0xc0,
	0x86, 0x8c, 0x63, 0xd8, 0x03, 0x24, 0xdb, 0x72, 0xb9, 0x60, 0x68, 0x29, 0x25, 0x10, 0x14, 0x59,
	0xa4, 0x26, 0x8e, 0xa5, 0x4a, 0x72, 0x42, 0x9e, 0xa1, 0x6f, 0xd1, 0xab, 0xd2, 0x47, 0xe9, 0x53,
	0x15, 0xcb, 0xc7, 0xa6, 0x94, 0xf6, 0xb2, 0x97, 0xe7, 0xfb, 0xd1, 0xf9, 0xce, 0x87, 0xd0, 0x77,
	0xa6, 0xa3, 0x60, 0x4b, 0x0d, 0x3f, 0xd2, 0x53, 0x70, 0x08, 0x37, 0xdc, 0xd0, 0x30, 0x50, 0x5c,
	0x8b, 0x42, 0x31, 0xae, 0x89, 0x54, 0xc2, 0x08, 0x3c, 0x60, 0x3a, 0x22, 0x20, 0x22, 0x20, 0xf2,
	0xbf, 0x96, 0x4e, 0x25, 0x59, 0xe3, 0xd2, 0x86, 0x9a, 0x02, 0x2c, 0xfe, 0xaf, 0x92, 0xd5, 0x46,
	0x28, 0xba, 0xe5, 0x81, 0x54, 0xe2, 0x90, 0x26, 0x5c, 0xbd, 0xb6, 0xc0, 0xff, 0x56, 0xaa, 0xcd,
	0x49, 0x72, 0xdd, 0x48, 0xec, 0x54, 0xd1, 0xe3, 0x1b, 0x17, 0xe1, 0x7f, 0x69, 0xc6, 0xcf, 0x64,
	0x26, 0x68, 0xb2, 0x2c, 0x31, 0x26, 0x32, 0x1c, 0xa2, 0xae, 0x90, 0xf4, 0xba, 0xe0, 0x9e, 0x33,
	0x72, 0x26, 0xbd, 0xe9, 0x17, 0x52, 0xe6, 0xac, 0x8c, 0xf0, 0x0c, 0xf9, 0x6f, 0x05, 0x31, 0x08,
	0xb1, 0x8f, 0xde, 0x49, 0xb0, 0x7b, 0xee, 0xc8, 0x99, 0xbc, 0x8f, 0x9b, 0x19, 0xff, 0x40, 0x1f,
	0x0a, 0xbb, 0x60, 0xcd, 0xf3, 0x44, 0x8a, 0x34, 0x37, 0x5e, 0xcb, 0x4a, 0xfa, 0x15, 0xfc, 0x17,
	0x50, 0xbc, 0x45, 0x03, 0x7a, 0xa0, 0x69, 0x46, 0x37, 0x19, 0x5f, 0xb3, 0x2b, 0xce, 0x76, 0xba,
	0xd8, 0x6b, 0xaf, 0x3d, 0x6a, 0x4d, 0x7a, 0xd3, 0xdf, 0x36, 0x04, 0x5c, 0x4e, 0xea, 0xcb, 0x9b,
	0x3c, 0x31, 0x5c, 0x3e, 0x07, 0xdb, 0x52, 0xa5, 0x42, 0xa5, 0xe6, 0x14, 0xe3, 0xe6, 0xc9, 0x9a,
	0xd2, 0x78, 0x88, 0x3a, 0x46, 0xec, 0x78, 0xee, 0x75, 0x6c, 0x8e, 0x6a, 0x18, 0xdf, 0x39, 0x68,
	0x58, 0xb6, 0xf1, 0x47, 0x1c, 0xf3, 0xb7, 0xec, 0xe3, 0x27, 0xfa, 0x94, 0xc0, 0x8a, 0xe7, 0x8d,
	0x7c, 0xac, 0x89, 0xa6, 0x93, 0x26, 0x6a, 0xfb, 0x49, 0xd4, 0xd9, 0x1e, 0x7d, 0x66, 0x62, 0x4f,
	0x5e, 0xf8, 0x3e, 0xb3, 0x7e, 0xdd, 0x84, 0xb6, 0xf9, 0x97, 0xce, 0x65, 0x1f, 0x24, 0xa0, 0xb8,
	0x75, 0x5b, 0xf3, 0xc5, 0xc5, 0xbd, 0x3b, 0x98, 0xeb, 0x88, 0x2c, 0xc0, 0x7d, 0x1e, 0xce, 0x4a,
	0xee, 0xc1, 0xa2, 0x2b, 0x40, 0x57, 0x80, 0x6e, 0xba, 0x36, 0x7b, 0xf4, 0x18, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x4b, 0x76, 0x9b, 0xd5, 0x02, 0x00, 0x00,
}

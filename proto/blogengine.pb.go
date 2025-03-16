// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/blogengine.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlogRequest) Reset() {
	*x = ListBlogRequest{}
	mi := &file_proto_blogengine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogRequest) ProtoMessage() {}

func (x *ListBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blogengine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogRequest.ProtoReflect.Descriptor instead.
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_blogengine_proto_rawDescGZIP(), []int{0}
}

type ListBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Blogs         []*Blog                `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlogResponse) Reset() {
	*x = ListBlogResponse{}
	mi := &file_proto_blogengine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogResponse) ProtoMessage() {}

func (x *ListBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blogengine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogResponse.ProtoReflect.Descriptor instead.
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_blogengine_proto_rawDescGZIP(), []int{1}
}

func (x *ListBlogResponse) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

type Blog struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Author        string                 `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Date          string                 `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blog) Reset() {
	*x = Blog{}
	mi := &file_proto_blogengine_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blogengine_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_proto_blogengine_proto_rawDescGZIP(), []int{2}
}

func (x *Blog) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Blog) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Blog) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBlogRequest) Reset() {
	*x = CreateBlogRequest{}
	mi := &file_proto_blogengine_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRequest) ProtoMessage() {}

func (x *CreateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blogengine_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_blogengine_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBlogRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateBlogRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type CreateBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBlogResponse) Reset() {
	*x = CreateBlogResponse{}
	mi := &file_proto_blogengine_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogResponse) ProtoMessage() {}

func (x *CreateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blogengine_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_blogengine_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBlogResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_blogengine_proto protoreflect.FileDescriptor

var file_proto_blogengine_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x6c, 0x0a, 0x04,
	0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x76, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x73, 0x12, 0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x61, 0x76, 0x6c, 0x69, 0x64, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_blogengine_proto_rawDescOnce sync.Once
	file_proto_blogengine_proto_rawDescData []byte
)

func file_proto_blogengine_proto_rawDescGZIP() []byte {
	file_proto_blogengine_proto_rawDescOnce.Do(func() {
		file_proto_blogengine_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_blogengine_proto_rawDesc), len(file_proto_blogengine_proto_rawDesc)))
	})
	return file_proto_blogengine_proto_rawDescData
}

var file_proto_blogengine_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_blogengine_proto_goTypes = []any{
	(*ListBlogRequest)(nil),    // 0: ListBlogRequest
	(*ListBlogResponse)(nil),   // 1: ListBlogResponse
	(*Blog)(nil),               // 2: Blog
	(*CreateBlogRequest)(nil),  // 3: CreateBlogRequest
	(*CreateBlogResponse)(nil), // 4: CreateBlogResponse
}
var file_proto_blogengine_proto_depIdxs = []int32{
	2, // 0: ListBlogResponse.blogs:type_name -> Blog
	0, // 1: BlogService.ListBlogs:input_type -> ListBlogRequest
	3, // 2: BlogService.CreateBlog:input_type -> CreateBlogRequest
	1, // 3: BlogService.ListBlogs:output_type -> ListBlogResponse
	4, // 4: BlogService.CreateBlog:output_type -> CreateBlogResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_blogengine_proto_init() }
func file_proto_blogengine_proto_init() {
	if File_proto_blogengine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_blogengine_proto_rawDesc), len(file_proto_blogengine_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blogengine_proto_goTypes,
		DependencyIndexes: file_proto_blogengine_proto_depIdxs,
		MessageInfos:      file_proto_blogengine_proto_msgTypes,
	}.Build()
	File_proto_blogengine_proto = out.File
	file_proto_blogengine_proto_goTypes = nil
	file_proto_blogengine_proto_depIdxs = nil
}

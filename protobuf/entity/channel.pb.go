// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: entity/channel.proto

package entity

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

type Channel struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ChannelId       string                 `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsFamilySafe    bool                   `protobuf:"varint,4,opt,name=is_family_safe,json=isFamilySafe,proto3" json:"is_family_safe,omitempty"`
	Keywords        string                 `protobuf:"bytes,5,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Thumbnails      []string               `protobuf:"bytes,6,rep,name=thumbnails,proto3" json:"thumbnails,omitempty"`
	Links           []string               `protobuf:"bytes,7,rep,name=links,proto3" json:"links,omitempty"`
	ViewCount       int32                  `protobuf:"varint,8,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	TotalSubscriber int32                  `protobuf:"varint,9,opt,name=total_subscriber,json=totalSubscriber,proto3" json:"total_subscriber,omitempty"`
	TotalVideo      int32                  `protobuf:"varint,10,opt,name=total_video,json=totalVideo,proto3" json:"total_video,omitempty"`
	Jointed         *Channel_Joined        `protobuf:"bytes,11,opt,name=jointed,proto3" json:"jointed,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Channel) Reset() {
	*x = Channel{}
	mi := &file_entity_channel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_entity_channel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_entity_channel_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Channel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Channel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Channel) GetIsFamilySafe() bool {
	if x != nil {
		return x.IsFamilySafe
	}
	return false
}

func (x *Channel) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *Channel) GetThumbnails() []string {
	if x != nil {
		return x.Thumbnails
	}
	return nil
}

func (x *Channel) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Channel) GetViewCount() int32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *Channel) GetTotalSubscriber() int32 {
	if x != nil {
		return x.TotalSubscriber
	}
	return 0
}

func (x *Channel) GetTotalVideo() int32 {
	if x != nil {
		return x.TotalVideo
	}
	return 0
}

func (x *Channel) GetJointed() *Channel_Joined {
	if x != nil {
		return x.Jointed
	}
	return nil
}

type Channel_Joined struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Year          int32                  `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month         int32                  `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Date          int32                  `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Channel_Joined) Reset() {
	*x = Channel_Joined{}
	mi := &file_entity_channel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Channel_Joined) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel_Joined) ProtoMessage() {}

func (x *Channel_Joined) ProtoReflect() protoreflect.Message {
	mi := &file_entity_channel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel_Joined.ProtoReflect.Descriptor instead.
func (*Channel_Joined) Descriptor() ([]byte, []int) {
	return file_entity_channel_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Channel_Joined) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Channel_Joined) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Channel_Joined) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

var File_entity_channel_proto protoreflect.FileDescriptor

var file_entity_channel_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xcb, 0x03,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x73, 0x61, 0x66,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x53, 0x61, 0x66, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x3f, 0x0a, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x07, 0x6a, 0x6f, 0x69, 0x6e,
	0x74, 0x65, 0x64, 0x1a, 0x46, 0x0a, 0x06, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x42, 0x5e, 0x0a, 0x21, 0x69,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x75, 0x6a, 0x69, 0x6e, 0x31, 0x31, 0x33, 0x35, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_entity_channel_proto_rawDescOnce sync.Once
	file_entity_channel_proto_rawDescData []byte
)

func file_entity_channel_proto_rawDescGZIP() []byte {
	file_entity_channel_proto_rawDescOnce.Do(func() {
		file_entity_channel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_entity_channel_proto_rawDesc), len(file_entity_channel_proto_rawDesc)))
	})
	return file_entity_channel_proto_rawDescData
}

var file_entity_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entity_channel_proto_goTypes = []any{
	(*Channel)(nil),        // 0: io.contents.collector.Channel
	(*Channel_Joined)(nil), // 1: io.contents.collector.Channel.Joined
}
var file_entity_channel_proto_depIdxs = []int32{
	1, // 0: io.contents.collector.Channel.jointed:type_name -> io.contents.collector.Channel.Joined
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entity_channel_proto_init() }
func file_entity_channel_proto_init() {
	if File_entity_channel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_entity_channel_proto_rawDesc), len(file_entity_channel_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_channel_proto_goTypes,
		DependencyIndexes: file_entity_channel_proto_depIdxs,
		MessageInfos:      file_entity_channel_proto_msgTypes,
	}.Build()
	File_entity_channel_proto = out.File
	file_entity_channel_proto_goTypes = nil
	file_entity_channel_proto_depIdxs = nil
}

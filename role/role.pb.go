// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: role/role.proto

package role

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// 角色model
type RoleModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 角色名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 角色创建时间
	CreateTime *timestamp.Timestamp `protobuf:"bytes,100,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 角色最后更新时间
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,101,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *RoleModel) Reset() {
	*x = RoleModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleModel) ProtoMessage() {}

func (x *RoleModel) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleModel.ProtoReflect.Descriptor instead.
func (*RoleModel) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoleModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleModel) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RoleModel) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// 删除指定的角色
type RoleDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id列表
	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *RoleDeleteReq) Reset() {
	*x = RoleDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteReq) ProtoMessage() {}

func (x *RoleDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteReq.ProtoReflect.Descriptor instead.
func (*RoleDeleteReq) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleDeleteReq) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 根据角色id获取角色信息
type RoleFindReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleFindReq) Reset() {
	*x = RoleFindReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindReq) ProtoMessage() {}

func (x *RoleFindReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindReq.ProtoReflect.Descriptor instead.
func (*RoleFindReq) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleFindReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 分页信息
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   uint32 `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{3}
}

func (x *Page) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *Page) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_role_role_proto protoreflect.FileDescriptor

var file_role_role_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x21, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0xf5, 0x01,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x0a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x11, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_role_proto_rawDescOnce sync.Once
	file_role_role_proto_rawDescData = file_role_role_proto_rawDesc
)

func file_role_role_proto_rawDescGZIP() []byte {
	file_role_role_proto_rawDescOnce.Do(func() {
		file_role_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_role_proto_rawDescData)
	})
	return file_role_role_proto_rawDescData
}

var file_role_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_role_role_proto_goTypes = []interface{}{
	(*RoleModel)(nil),           // 0: role.RoleModel
	(*RoleDeleteReq)(nil),       // 1: role.RoleDeleteReq
	(*RoleFindReq)(nil),         // 2: role.RoleFindReq
	(*Page)(nil),                // 3: role.Page
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*common.Response)(nil),     // 5: common.Response
}
var file_role_role_proto_depIdxs = []int32{
	4, // 0: role.RoleModel.create_time:type_name -> google.protobuf.Timestamp
	4, // 1: role.RoleModel.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: role.Role.Add:input_type -> role.RoleModel
	1, // 3: role.Role.Delete:input_type -> role.RoleDeleteReq
	0, // 4: role.Role.Update:input_type -> role.RoleModel
	3, // 5: role.Role.GetList:input_type -> role.Page
	2, // 6: role.Role.GetUserRole:input_type -> role.RoleFindReq
	5, // 7: role.Role.Add:output_type -> common.Response
	5, // 8: role.Role.Delete:output_type -> common.Response
	5, // 9: role.Role.Update:output_type -> common.Response
	0, // 10: role.Role.GetList:output_type -> role.RoleModel
	0, // 11: role.Role.GetUserRole:output_type -> role.RoleModel
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_role_role_proto_init() }
func file_role_role_proto_init() {
	if File_role_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleModel); i {
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
		file_role_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteReq); i {
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
		file_role_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindReq); i {
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
		file_role_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_role_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_role_proto_goTypes,
		DependencyIndexes: file_role_role_proto_depIdxs,
		MessageInfos:      file_role_role_proto_msgTypes,
	}.Build()
	File_role_role_proto = out.File
	file_role_role_proto_rawDesc = nil
	file_role_role_proto_goTypes = nil
	file_role_role_proto_depIdxs = nil
}

// Code generated by protoc-gen-go.
// source: steammessages_unified_base.steamworkssdk.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EProtoExecutionSite int32

const (
	EProtoExecutionSite_k_EProtoExecutionSiteUnknown     EProtoExecutionSite = 0
	EProtoExecutionSite_k_EProtoExecutionSiteSteamClient EProtoExecutionSite = 3
)

var EProtoExecutionSite_name = map[int32]string{
	0: "k_EProtoExecutionSiteUnknown",
	3: "k_EProtoExecutionSiteSteamClient",
}
var EProtoExecutionSite_value = map[string]int32{
	"k_EProtoExecutionSiteUnknown":     0,
	"k_EProtoExecutionSiteSteamClient": 3,
}

func (x EProtoExecutionSite) Enum() *EProtoExecutionSite {
	p := new(EProtoExecutionSite)
	*p = x
	return p
}
func (x EProtoExecutionSite) String() string {
	return proto.EnumName(EProtoExecutionSite_name, int32(x))
}
func (x *EProtoExecutionSite) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoExecutionSite_value, data, "EProtoExecutionSite")
	if err != nil {
		return err
	}
	*x = EProtoExecutionSite(value)
	return nil
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "dota.description",
	Tag:           "bytes,50000,opt,name=description",
}

var E_ServiceDescription = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "dota.service_description",
	Tag:           "bytes,50000,opt,name=service_description",
}

var E_ServiceExecutionSite = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*EProtoExecutionSite)(nil),
	Field:         50008,
	Name:          "dota.service_execution_site",
	Tag:           "varint,50008,opt,name=service_execution_site,enum=dota.EProtoExecutionSite,def=0",
}

var E_MethodDescription = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "dota.method_description",
	Tag:           "bytes,50000,opt,name=method_description",
}

var E_EnumDescription = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "dota.enum_description",
	Tag:           "bytes,50000,opt,name=enum_description",
}

var E_EnumValueDescription = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "dota.enum_value_description",
	Tag:           "bytes,50000,opt,name=enum_value_description",
}

func init() {
	proto.RegisterEnum("dota.EProtoExecutionSite", EProtoExecutionSite_name, EProtoExecutionSite_value)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_ServiceDescription)
	proto.RegisterExtension(E_ServiceExecutionSite)
	proto.RegisterExtension(E_MethodDescription)
	proto.RegisterExtension(E_EnumDescription)
	proto.RegisterExtension(E_EnumValueDescription)
}

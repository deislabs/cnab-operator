// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/column_spec.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A representation of a column in a relational table. When listing them, column specs are returned in the same order in which they were
// given on import .
// Used by:
//   *   Tables
type ColumnSpec struct {
	// Output only. The resource name of the column specs.
	// Form:
	//
	// `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/tableSpecs/{table_spec_id}/columnSpecs/{column_spec_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The data type of elements stored in the column.
	DataType *DataType `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	// Output only. The name of the column to show in the interface. The name can
	// be up to 100 characters long and can consist only of ASCII Latin letters
	// A-Z and a-z, ASCII digits 0-9, underscores(_), and forward slashes(/), and
	// must start with a letter or a digit.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Stats of the series of values in the column.
	// This field may be stale, see the ancestor's
	// Dataset.tables_dataset_metadata.stats_update_time field
	// for the timestamp at which these stats were last updated.
	DataStats *DataStats `protobuf:"bytes,4,opt,name=data_stats,json=dataStats,proto3" json:"data_stats,omitempty"`
	// Output only. Top 10 most correlated with this column columns of the table,
	// ordered by
	// [cramers_v][google.cloud.automl.v1beta1.CorrelationStats.cramers_v] metric.
	// This field may be stale, see the ancestor's
	// Dataset.tables_dataset_metadata.stats_update_time field
	// for the timestamp at which these stats were last updated.
	TopCorrelatedColumns []*ColumnSpec_CorrelatedColumn `protobuf:"bytes,5,rep,name=top_correlated_columns,json=topCorrelatedColumns,proto3" json:"top_correlated_columns,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag                 string   `protobuf:"bytes,6,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnSpec) Reset()         { *m = ColumnSpec{} }
func (m *ColumnSpec) String() string { return proto.CompactTextString(m) }
func (*ColumnSpec) ProtoMessage()    {}
func (*ColumnSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_column_spec_5bd50af362cf7442, []int{0}
}
func (m *ColumnSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnSpec.Unmarshal(m, b)
}
func (m *ColumnSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnSpec.Marshal(b, m, deterministic)
}
func (dst *ColumnSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnSpec.Merge(dst, src)
}
func (m *ColumnSpec) XXX_Size() int {
	return xxx_messageInfo_ColumnSpec.Size(m)
}
func (m *ColumnSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnSpec proto.InternalMessageInfo

func (m *ColumnSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnSpec) GetDataType() *DataType {
	if m != nil {
		return m.DataType
	}
	return nil
}

func (m *ColumnSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ColumnSpec) GetDataStats() *DataStats {
	if m != nil {
		return m.DataStats
	}
	return nil
}

func (m *ColumnSpec) GetTopCorrelatedColumns() []*ColumnSpec_CorrelatedColumn {
	if m != nil {
		return m.TopCorrelatedColumns
	}
	return nil
}

func (m *ColumnSpec) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// Identifies the table's column, and its correlation with the column this
// ColumnSpec describes.
type ColumnSpec_CorrelatedColumn struct {
	// The column_spec_id of the correlated column, which belongs to the same
	// table as the in-context column.
	ColumnSpecId string `protobuf:"bytes,1,opt,name=column_spec_id,json=columnSpecId,proto3" json:"column_spec_id,omitempty"`
	// Correlation between this and the in-context column.
	CorrelationStats     *CorrelationStats `protobuf:"bytes,2,opt,name=correlation_stats,json=correlationStats,proto3" json:"correlation_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ColumnSpec_CorrelatedColumn) Reset()         { *m = ColumnSpec_CorrelatedColumn{} }
func (m *ColumnSpec_CorrelatedColumn) String() string { return proto.CompactTextString(m) }
func (*ColumnSpec_CorrelatedColumn) ProtoMessage()    {}
func (*ColumnSpec_CorrelatedColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_column_spec_5bd50af362cf7442, []int{0, 0}
}
func (m *ColumnSpec_CorrelatedColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnSpec_CorrelatedColumn.Unmarshal(m, b)
}
func (m *ColumnSpec_CorrelatedColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnSpec_CorrelatedColumn.Marshal(b, m, deterministic)
}
func (dst *ColumnSpec_CorrelatedColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnSpec_CorrelatedColumn.Merge(dst, src)
}
func (m *ColumnSpec_CorrelatedColumn) XXX_Size() int {
	return xxx_messageInfo_ColumnSpec_CorrelatedColumn.Size(m)
}
func (m *ColumnSpec_CorrelatedColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnSpec_CorrelatedColumn.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnSpec_CorrelatedColumn proto.InternalMessageInfo

func (m *ColumnSpec_CorrelatedColumn) GetColumnSpecId() string {
	if m != nil {
		return m.ColumnSpecId
	}
	return ""
}

func (m *ColumnSpec_CorrelatedColumn) GetCorrelationStats() *CorrelationStats {
	if m != nil {
		return m.CorrelationStats
	}
	return nil
}

func init() {
	proto.RegisterType((*ColumnSpec)(nil), "google.cloud.automl.v1beta1.ColumnSpec")
	proto.RegisterType((*ColumnSpec_CorrelatedColumn)(nil), "google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/column_spec.proto", fileDescriptor_column_spec_5bd50af362cf7442)
}

var fileDescriptor_column_spec_5bd50af362cf7442 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0xc6, 0xe9, 0x7f, 0x7f, 0x87, 0xcb, 0x86, 0xcc, 0x20, 0x52, 0x36, 0xc1, 0x29, 0x2a, 0xbb,
	0x70, 0x29, 0x9b, 0x37, 0x82, 0x57, 0xdb, 0x14, 0xf1, 0x42, 0x91, 0x4e, 0xbc, 0x18, 0x83, 0x72,
	0x96, 0x86, 0x52, 0x48, 0x73, 0x42, 0x9b, 0x0a, 0xbb, 0xf7, 0x11, 0x7c, 0x2a, 0xdf, 0xc4, 0xb7,
	0x90, 0x26, 0xb1, 0x93, 0x21, 0x9d, 0x77, 0xa7, 0x27, 0xbf, 0x7c, 0xdf, 0xc9, 0xd7, 0x43, 0x16,
	0x19, 0x62, 0x26, 0x45, 0xc4, 0x25, 0xd6, 0x69, 0x04, 0xb5, 0xc1, 0x42, 0x46, 0xdf, 0x96, 0x47,
	0x61, 0x60, 0x19, 0x71, 0x94, 0x75, 0xa1, 0x92, 0x4a, 0x0b, 0xce, 0x74, 0x89, 0x06, 0xe9, 0xd4,
	0xe1, 0xcc, 0xe2, 0xcc, 0xe1, 0xcc, 0xe3, 0x93, 0x47, 0x5e, 0x0b, 0x74, 0x1e, 0x81, 0x52, 0x68,
	0xc0, 0xe4, 0xa8, 0x2a, 0x77, 0x75, 0xf2, 0xb2, 0xcb, 0x29, 0x05, 0x03, 0x49, 0x65, 0xc0, 0xfc,
	0x3f, 0x6d, 0x4e, 0x5a, 0x78, 0xfa, 0xe9, 0xaf, 0x1e, 0x21, 0x5b, 0x3b, 0xec, 0x4e, 0x0b, 0x4e,
	0x29, 0xb9, 0x55, 0x50, 0x88, 0x30, 0x98, 0x05, 0xf3, 0x41, 0x6c, 0x6b, 0xba, 0x21, 0x83, 0xf6,
	0x5a, 0x78, 0x33, 0x0b, 0xe6, 0xc3, 0xd5, 0x73, 0xd6, 0xf1, 0x1a, 0xf6, 0x16, 0x0c, 0x7c, 0x39,
	0x69, 0x11, 0xdf, 0x4d, 0x7d, 0x45, 0x9f, 0x90, 0x51, 0x9a, 0x57, 0x5a, 0xc2, 0x29, 0xb1, 0xfa,
	0x3d, 0xab, 0x3f, 0xf4, 0xbd, 0x4f, 0x8d, 0xcd, 0x3b, 0x42, 0xce, 0x6f, 0x09, 0x6f, 0xad, 0xcf,
	0x8b, 0xab, 0x3e, 0xbb, 0x86, 0x8e, 0xed, 0x80, 0xb6, 0xa4, 0x8a, 0x3c, 0x34, 0xa8, 0x13, 0x8e,
	0x65, 0x29, 0x24, 0x18, 0x91, 0x26, 0xee, 0x5f, 0x54, 0xe1, 0x9d, 0x59, 0x6f, 0x3e, 0x5c, 0xbd,
	0xee, 0x94, 0x3c, 0x47, 0xc1, 0xb6, 0xad, 0x82, 0x6b, 0xc6, 0x0f, 0x0c, 0xea, 0xcb, 0x66, 0xd5,
	0x24, 0x26, 0x0c, 0x64, 0x61, 0xdf, 0x25, 0xd6, 0xd4, 0x93, 0x1f, 0x01, 0x19, 0x5f, 0x92, 0xf4,
	0x19, 0xb9, 0xf7, 0xd7, 0x56, 0x24, 0x79, 0xea, 0x43, 0x1e, 0xf1, 0xd6, 0xf3, 0x43, 0x4a, 0xf7,
	0xe4, 0xfe, 0x9f, 0xd1, 0x73, 0x54, 0x3e, 0x0c, 0x17, 0xfa, 0xe2, 0xca, 0xe4, 0xed, 0x2d, 0x97,
	0xc9, 0x98, 0x5f, 0x74, 0x36, 0xdf, 0x03, 0xf2, 0x98, 0x63, 0xd1, 0x25, 0xf3, 0x39, 0xd8, 0xaf,
	0xfd, 0x71, 0x86, 0x12, 0x54, 0xc6, 0xb0, 0xcc, 0xa2, 0x4c, 0x28, 0xbb, 0x2d, 0x91, 0x3b, 0x02,
	0x9d, 0x57, 0xff, 0x5c, 0xaf, 0x37, 0xee, 0xf3, 0xe7, 0xcd, 0xf4, 0xbd, 0x05, 0x0f, 0xdb, 0x06,
	0x3a, 0xac, 0x6b, 0x83, 0x1f, 0xe5, 0xe1, 0xab, 0x83, 0x8e, 0x7d, 0xab, 0xf5, 0xea, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x06, 0x4f, 0x07, 0x5d, 0x41, 0x03, 0x00, 0x00,
}

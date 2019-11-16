// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FoodEnforcement.proto

package v1

import (
	fmt "fmt"
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

type FoodEnforcement struct {
	FoodEnforcementID        string   `protobuf:"bytes,1,opt,name=FoodEnforcementID,proto3" json:"FoodEnforcementID,omitempty"`
	Classification           string   `protobuf:"bytes,2,opt,name=classification,proto3" json:"classification,omitempty"`
	CenterClassificationDate string   `protobuf:"bytes,3,opt,name=center_classification_date,json=centerClassificationDate,proto3" json:"center_classification_date,omitempty"`
	ReportDate               string   `protobuf:"bytes,4,opt,name=report_date,json=reportDate,proto3" json:"report_date,omitempty"`
	PostalCode               string   `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	TerminationDate          string   `protobuf:"bytes,6,opt,name=termination_date,json=terminationDate,proto3" json:"termination_date,omitempty"`
	RecallInitiationDate     string   `protobuf:"bytes,7,opt,name=recall_initiation_date,json=recallInitiationDate,proto3" json:"recall_initiation_date,omitempty"`
	RecallNumber             string   `protobuf:"bytes,8,opt,name=recall_number,json=recallNumber,proto3" json:"recall_number,omitempty"`
	City                     string   `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	MoreCodeInfo             string   `protobuf:"bytes,10,opt,name=more_code_info,json=moreCodeInfo,proto3" json:"more_code_info,omitempty"`
	EventId                  string   `protobuf:"bytes,11,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	DistributionPattern      string   `protobuf:"bytes,12,opt,name=distribution_pattern,json=distributionPattern,proto3" json:"distribution_pattern,omitempty"`
	Openfda                  string   `protobuf:"bytes,13,opt,name=openfda,proto3" json:"openfda,omitempty"`
	RecallingFirm            string   `protobuf:"bytes,14,opt,name=recalling_firm,json=recallingFirm,proto3" json:"recalling_firm,omitempty"`
	VoluntaryMandated        string   `protobuf:"bytes,15,opt,name=voluntary_mandated,json=voluntaryMandated,proto3" json:"voluntary_mandated,omitempty"`
	State                    string   `protobuf:"bytes,16,opt,name=state,proto3" json:"state,omitempty"`
	ReasonForRecall          string   `protobuf:"bytes,17,opt,name=reason_for_recall,json=reasonForRecall,proto3" json:"reason_for_recall,omitempty"`
	InitialFirmNotification  string   `protobuf:"bytes,18,opt,name=initial_firm_notification,json=initialFirmNotification,proto3" json:"initial_firm_notification,omitempty"`
	Status                   string   `protobuf:"bytes,19,opt,name=status,proto3" json:"status,omitempty"`
	ProductType              string   `protobuf:"bytes,20,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	Country                  string   `protobuf:"bytes,21,opt,name=country,proto3" json:"country,omitempty"`
	ProductDescription       string   `protobuf:"bytes,22,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	CodeInfo                 string   `protobuf:"bytes,23,opt,name=code_info,json=codeInfo,proto3" json:"code_info,omitempty"`
	Address_1                string   `protobuf:"bytes,24,opt,name=address_1,json=address1,proto3" json:"address_1,omitempty"`
	Address_2                string   `protobuf:"bytes,25,opt,name=address_2,json=address2,proto3" json:"address_2,omitempty"`
	ProductQuantity          string   `protobuf:"bytes,26,opt,name=product_quantity,json=productQuantity,proto3" json:"product_quantity,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *FoodEnforcement) Reset()         { *m = FoodEnforcement{} }
func (m *FoodEnforcement) String() string { return proto.CompactTextString(m) }
func (*FoodEnforcement) ProtoMessage()    {}
func (*FoodEnforcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_5020c6583df3db5a, []int{0}
}

func (m *FoodEnforcement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodEnforcement.Unmarshal(m, b)
}
func (m *FoodEnforcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodEnforcement.Marshal(b, m, deterministic)
}
func (m *FoodEnforcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodEnforcement.Merge(m, src)
}
func (m *FoodEnforcement) XXX_Size() int {
	return xxx_messageInfo_FoodEnforcement.Size(m)
}
func (m *FoodEnforcement) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodEnforcement.DiscardUnknown(m)
}

var xxx_messageInfo_FoodEnforcement proto.InternalMessageInfo

func (m *FoodEnforcement) GetFoodEnforcementID() string {
	if m != nil {
		return m.FoodEnforcementID
	}
	return ""
}

func (m *FoodEnforcement) GetClassification() string {
	if m != nil {
		return m.Classification
	}
	return ""
}

func (m *FoodEnforcement) GetCenterClassificationDate() string {
	if m != nil {
		return m.CenterClassificationDate
	}
	return ""
}

func (m *FoodEnforcement) GetReportDate() string {
	if m != nil {
		return m.ReportDate
	}
	return ""
}

func (m *FoodEnforcement) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *FoodEnforcement) GetTerminationDate() string {
	if m != nil {
		return m.TerminationDate
	}
	return ""
}

func (m *FoodEnforcement) GetRecallInitiationDate() string {
	if m != nil {
		return m.RecallInitiationDate
	}
	return ""
}

func (m *FoodEnforcement) GetRecallNumber() string {
	if m != nil {
		return m.RecallNumber
	}
	return ""
}

func (m *FoodEnforcement) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *FoodEnforcement) GetMoreCodeInfo() string {
	if m != nil {
		return m.MoreCodeInfo
	}
	return ""
}

func (m *FoodEnforcement) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *FoodEnforcement) GetDistributionPattern() string {
	if m != nil {
		return m.DistributionPattern
	}
	return ""
}

func (m *FoodEnforcement) GetOpenfda() string {
	if m != nil {
		return m.Openfda
	}
	return ""
}

func (m *FoodEnforcement) GetRecallingFirm() string {
	if m != nil {
		return m.RecallingFirm
	}
	return ""
}

func (m *FoodEnforcement) GetVoluntaryMandated() string {
	if m != nil {
		return m.VoluntaryMandated
	}
	return ""
}

func (m *FoodEnforcement) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *FoodEnforcement) GetReasonForRecall() string {
	if m != nil {
		return m.ReasonForRecall
	}
	return ""
}

func (m *FoodEnforcement) GetInitialFirmNotification() string {
	if m != nil {
		return m.InitialFirmNotification
	}
	return ""
}

func (m *FoodEnforcement) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FoodEnforcement) GetProductType() string {
	if m != nil {
		return m.ProductType
	}
	return ""
}

func (m *FoodEnforcement) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *FoodEnforcement) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *FoodEnforcement) GetCodeInfo() string {
	if m != nil {
		return m.CodeInfo
	}
	return ""
}

func (m *FoodEnforcement) GetAddress_1() string {
	if m != nil {
		return m.Address_1
	}
	return ""
}

func (m *FoodEnforcement) GetAddress_2() string {
	if m != nil {
		return m.Address_2
	}
	return ""
}

func (m *FoodEnforcement) GetProductQuantity() string {
	if m != nil {
		return m.ProductQuantity
	}
	return ""
}

func init() {
	proto.RegisterType((*FoodEnforcement)(nil), "v1.FoodEnforcement")
}

func init() { proto.RegisterFile("FoodEnforcement.proto", fileDescriptor_5020c6583df3db5a) }

var fileDescriptor_5020c6583df3db5a = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0x4b, 0x6f, 0x13, 0x31,
	0x10, 0xc7, 0xd5, 0xd2, 0xe6, 0x31, 0x79, 0x35, 0x93, 0x34, 0x75, 0xda, 0x03, 0x6f, 0x04, 0x88,
	0x87, 0x52, 0x38, 0x21, 0x6e, 0x0d, 0x91, 0x72, 0xa0, 0x82, 0x8a, 0xbb, 0xe5, 0xac, 0xbd, 0xc8,
	0xd2, 0xae, 0xbd, 0x78, 0xbd, 0x91, 0xf2, 0x99, 0xf9, 0x12, 0x68, 0xc7, 0x9b, 0x64, 0x13, 0x6e,
	0x99, 0xff, 0xef, 0x37, 0xeb, 0xd1, 0xd8, 0x81, 0xcb, 0x85, 0xb5, 0xf2, 0x9b, 0x89, 0xad, 0x8b,
	0x54, 0xaa, 0x8c, 0xff, 0x90, 0x39, 0xeb, 0x2d, 0x9e, 0xae, 0x67, 0xcf, 0xfe, 0x36, 0x61, 0x70,
	0x44, 0xf1, 0x1d, 0x0c, 0x8f, 0xa2, 0xe5, 0x9c, 0x9d, 0x3c, 0x39, 0x79, 0xdd, 0x7e, 0xf8, 0x1f,
	0xe0, 0x2b, 0xe8, 0x47, 0x89, 0xc8, 0x73, 0x1d, 0xeb, 0x48, 0x78, 0x6d, 0x0d, 0x3b, 0x25, 0xf5,
	0x28, 0xc5, 0xaf, 0x70, 0x1d, 0x29, 0xe3, 0x95, 0xe3, 0x87, 0x80, 0x4b, 0xe1, 0x15, 0x7b, 0x44,
	0x3d, 0x2c, 0x18, 0x77, 0x07, 0xc2, 0x5c, 0x78, 0x85, 0x8f, 0xa1, 0xe3, 0x54, 0x66, 0x9d, 0x0f,
	0xfa, 0x19, 0xe9, 0x10, 0xa2, 0xad, 0x90, 0xd9, 0xdc, 0x8b, 0x84, 0x47, 0x56, 0x2a, 0x76, 0x1e,
	0x84, 0x10, 0xdd, 0x59, 0xa9, 0xf0, 0x0d, 0x5c, 0x78, 0xe5, 0x52, 0x6d, 0x6a, 0xa7, 0x36, 0xc8,
	0x1a, 0xd4, 0x72, 0xfa, 0xd6, 0x67, 0x98, 0x38, 0x15, 0x89, 0x24, 0xe1, 0xda, 0x68, 0xaf, 0x6b,
	0x0d, 0x4d, 0x6a, 0x18, 0x07, 0xba, 0xdc, 0x41, 0xea, 0x7a, 0x0e, 0xbd, 0xaa, 0xcb, 0x14, 0xe9,
	0x4a, 0x39, 0xd6, 0x22, 0xb9, 0x1b, 0xc2, 0x7b, 0xca, 0x10, 0xe1, 0x2c, 0xd2, 0x7e, 0xc3, 0xda,
	0xc4, 0xe8, 0x37, 0xbe, 0x80, 0x7e, 0x6a, 0x9d, 0xa2, 0xc1, 0xb9, 0x36, 0xb1, 0x65, 0x10, 0x3a,
	0xcb, 0xb4, 0x9c, 0x7d, 0x69, 0x62, 0x8b, 0x53, 0x68, 0xa9, 0xb5, 0x32, 0x9e, 0x6b, 0xc9, 0x3a,
	0xc4, 0x9b, 0x54, 0x2f, 0x25, 0xce, 0x60, 0x2c, 0x75, 0xee, 0x9d, 0x5e, 0x15, 0x34, 0x6a, 0x26,
	0xbc, 0x57, 0xce, 0xb0, 0x2e, 0x69, 0xa3, 0x3a, 0xfb, 0x11, 0x10, 0x32, 0x68, 0xda, 0x4c, 0x99,
	0x58, 0x0a, 0xd6, 0x0b, 0x1f, 0xab, 0x4a, 0x7c, 0x09, 0xfd, 0x30, 0xb1, 0x36, 0xbf, 0x79, 0xac,
	0x5d, 0xca, 0xfa, 0x24, 0xf4, 0x76, 0xe9, 0x42, 0xbb, 0x14, 0xdf, 0x03, 0xae, 0x6d, 0x52, 0x18,
	0x2f, 0xdc, 0x86, 0xa7, 0xc2, 0x94, 0xeb, 0x91, 0x6c, 0x10, 0x5e, 0xc9, 0x8e, 0x7c, 0xaf, 0x00,
	0x8e, 0xe1, 0x3c, 0xf7, 0xe5, 0x06, 0x2f, 0xc8, 0x08, 0x05, 0xbe, 0x85, 0xa1, 0x53, 0x22, 0xb7,
	0x86, 0xc7, 0xd6, 0xf1, 0x70, 0x00, 0x1b, 0x86, 0x4b, 0x09, 0x60, 0x61, 0xdd, 0x03, 0xc5, 0xf8,
	0x05, 0xa6, 0xe1, 0x36, 0x12, 0x9a, 0x8a, 0x1b, 0xeb, 0xf7, 0x4f, 0x0e, 0xa9, 0xe7, 0xaa, 0x12,
	0xca, 0x01, 0xef, 0x6b, 0x18, 0x27, 0xd0, 0x28, 0x0f, 0x2c, 0x72, 0x36, 0x22, 0xb1, 0xaa, 0xf0,
	0x29, 0x74, 0x33, 0x67, 0x65, 0x11, 0x79, 0xee, 0x37, 0x99, 0x62, 0x63, 0xa2, 0x9d, 0x2a, 0xfb,
	0xb5, 0xc9, 0x54, 0xb9, 0xa8, 0xc8, 0x16, 0xc6, 0xbb, 0x0d, 0xbb, 0x0c, 0x8b, 0xaa, 0x4a, 0xfc,
	0x08, 0xa3, 0x6d, 0xb3, 0x54, 0x79, 0xe4, 0x74, 0x46, 0xa3, 0x4c, 0xc8, 0xc2, 0x0a, 0xcd, 0xf7,
	0x04, 0x6f, 0xa0, 0xbd, 0xbf, 0xe2, 0x2b, 0xd2, 0x5a, 0xd1, 0xf6, 0x7a, 0x6f, 0xa0, 0x2d, 0xa4,
	0x74, 0x2a, 0xcf, 0xf9, 0x8c, 0xb1, 0x00, 0xab, 0x60, 0x56, 0x87, 0xb7, 0x6c, 0x7a, 0x00, 0x6f,
	0xcb, 0x87, 0xbd, 0x9d, 0xe3, 0x4f, 0x21, 0x8c, 0x2f, 0x9f, 0xd7, 0x75, 0xd8, 0x61, 0x95, 0xff,
	0xac, 0xe2, 0x55, 0x83, 0xfe, 0xf8, 0x9f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x98, 0x4f,
	0x06, 0x11, 0x04, 0x00, 0x00,
}

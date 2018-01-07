// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/transaction_pb2/transaction.proto

/*
Package transaction_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/transaction_pb2/transaction.proto

It has these top-level messages:
	TransactionHeader
	Transaction
	TransactionList
*/
package transaction_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionHeader struct {
	// Public key for the client who added this transaction to a batch
	BatcherPubkey string `protobuf:"bytes,1,opt,name=batcher_pubkey,json=batcherPubkey" json:"batcher_pubkey,omitempty"`
	// A list of transaction signatures that describe the transactions that
	// must be processed before this transaction can be valid
	Dependencies []string `protobuf:"bytes,2,rep,name=dependencies" json:"dependencies,omitempty"`
	// The family name correlates to the transaction processor's family name
	// that this transaction can be processed on, for example 'intkey'
	FamilyName string `protobuf:"bytes,3,opt,name=family_name,json=familyName" json:"family_name,omitempty"`
	// The family version correlates to the transaction processor's family
	// version that this transaction can be processed on, for example "1.0"
	FamilyVersion string `protobuf:"bytes,4,opt,name=family_version,json=familyVersion" json:"family_version,omitempty"`
	// A list of addresses that are given to the context manager and control
	// what addresses the transaction processor is allowed to read from.
	Inputs []string `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	// A random string that provides uniqueness for transactions with
	// otherwise identical fields.
	Nonce string `protobuf:"bytes,6,opt,name=nonce" json:"nonce,omitempty"`
	// A list of addresses that are given to the context manager and control
	// what addresses the transaction processor is allowed to write to.
	Outputs []string `protobuf:"bytes,7,rep,name=outputs" json:"outputs,omitempty"`
	// The payload encoding correlates to the transaction processor's payload
	// encoding that this transaction can be processed on,
	// for example ""application/cbor""
	PayloadEncoding string `protobuf:"bytes,8,opt,name=payload_encoding,json=payloadEncoding" json:"payload_encoding,omitempty"`
	// The sha512 hash of the encoded payload
	PayloadSha512 string `protobuf:"bytes,9,opt,name=payload_sha512,json=payloadSha512" json:"payload_sha512,omitempty"`
	// Public key for the client that signed the TransactionHeader
	SignerPubkey string `protobuf:"bytes,10,opt,name=signer_pubkey,json=signerPubkey" json:"signer_pubkey,omitempty"`
}

func (m *TransactionHeader) Reset()                    { *m = TransactionHeader{} }
func (m *TransactionHeader) String() string            { return proto.CompactTextString(m) }
func (*TransactionHeader) ProtoMessage()               {}
func (*TransactionHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransactionHeader) GetBatcherPubkey() string {
	if m != nil {
		return m.BatcherPubkey
	}
	return ""
}

func (m *TransactionHeader) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *TransactionHeader) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *TransactionHeader) GetFamilyVersion() string {
	if m != nil {
		return m.FamilyVersion
	}
	return ""
}

func (m *TransactionHeader) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TransactionHeader) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *TransactionHeader) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TransactionHeader) GetPayloadEncoding() string {
	if m != nil {
		return m.PayloadEncoding
	}
	return ""
}

func (m *TransactionHeader) GetPayloadSha512() string {
	if m != nil {
		return m.PayloadSha512
	}
	return ""
}

func (m *TransactionHeader) GetSignerPubkey() string {
	if m != nil {
		return m.SignerPubkey
	}
	return ""
}

type Transaction struct {
	// The serialized version of the TransactionHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature" json:"header_signature,omitempty"`
	// The payload is the encoded family specific information of the
	// transaction. Example cbor({'Verb': verb, 'Name': name,'Value': value})
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Transaction) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetHeaderSignature() string {
	if m != nil {
		return m.HeaderSignature
	}
	return ""
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// A simple list of transactions that needs to be serialized before
// it can be transmitted to a batcher.
type TransactionList struct {
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *TransactionList) Reset()                    { *m = TransactionList{} }
func (m *TransactionList) String() string            { return proto.CompactTextString(m) }
func (*TransactionList) ProtoMessage()               {}
func (*TransactionList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TransactionList) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionHeader)(nil), "TransactionHeader")
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*TransactionList)(nil), "TransactionList")
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/transaction_pb2/transaction.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x4f, 0xe2, 0x40,
	0x14, 0xc6, 0x03, 0x2c, 0xb0, 0x3c, 0xca, 0xb2, 0x3b, 0xd9, 0xdd, 0xcc, 0x6d, 0x49, 0x37, 0x26,
	0xe8, 0xa1, 0x28, 0xc6, 0x78, 0xd7, 0x98, 0x78, 0x30, 0x86, 0x14, 0xe3, 0xc1, 0x4b, 0x33, 0x6d,
	0x07, 0x5a, 0x81, 0x99, 0xa6, 0x33, 0xd5, 0xf0, 0x07, 0xfb, 0x7f, 0xd8, 0xbe, 0x99, 0x86, 0xe2,
	0xad, 0xdf, 0xaf, 0xdf, 0xf4, 0x7d, 0xf3, 0xf5, 0xc1, 0xb5, 0x62, 0xef, 0x5a, 0x4a, 0x9d, 0x04,
	0x2a, 0xde, 0xcc, 0xb2, 0x5c, 0x6a, 0x19, 0x16, 0xab, 0x99, 0xce, 0x99, 0x50, 0x2c, 0xd2, 0xa9,
	0x14, 0x41, 0x16, 0xce, 0x9b, 0xda, 0x43, 0x93, 0xfb, 0xd1, 0x86, 0x5f, 0x4f, 0x07, 0x7a, 0xcf,
	0x59, 0xcc, 0x73, 0x72, 0x02, 0x3f, 0x42, 0xa6, 0xa3, 0x84, 0xe7, 0x41, 0x56, 0x84, 0x1b, 0xbe,
	0xa7, 0xad, 0x49, 0x6b, 0x3a, 0xf0, 0x47, 0x96, 0x2e, 0x10, 0x12, 0x17, 0x9c, 0x98, 0x67, 0x5c,
	0xc4, 0x5c, 0x44, 0x29, 0x57, 0xb4, 0x3d, 0xe9, 0x94, 0xa6, 0x23, 0x46, 0xfe, 0xc1, 0x70, 0xc5,
	0x76, 0xe9, 0x76, 0x1f, 0x08, 0xb6, 0xe3, 0xb4, 0x83, 0xdf, 0x01, 0x83, 0x1e, 0x4b, 0x52, 0xcd,
	0xb2, 0x86, 0x37, 0x9e, 0xab, 0x32, 0x03, 0xfd, 0x66, 0x66, 0x19, 0xfa, 0x6c, 0x20, 0xf9, 0x0b,
	0xbd, 0x54, 0x64, 0x85, 0x56, 0xb4, 0x8b, 0x53, 0xac, 0x22, 0xbf, 0xa1, 0x2b, 0xa4, 0x88, 0x38,
	0xed, 0xe1, 0x29, 0x23, 0x08, 0x85, 0xbe, 0x2c, 0x34, 0xda, 0xfb, 0x68, 0xaf, 0x25, 0x39, 0x85,
	0x9f, 0x19, 0xdb, 0x6f, 0x25, 0x8b, 0x83, 0x32, 0xa1, 0x8c, 0x53, 0xb1, 0xa6, 0xdf, 0xf1, 0xe8,
	0xd8, 0xf2, 0x3b, 0x8b, 0xab, 0x64, 0xb5, 0x55, 0x25, 0xec, 0xea, 0x62, 0x4e, 0x07, 0x26, 0x99,
	0xa5, 0x4b, 0x84, 0xe4, 0x3f, 0x8c, 0x54, 0xba, 0x16, 0x87, 0xae, 0x00, 0x5d, 0x8e, 0x81, 0xa6,
	0x2a, 0xf7, 0x15, 0x86, 0x8d, 0x9a, 0xab, 0xdb, 0x24, 0x58, 0x35, 0x16, 0xeb, 0xf8, 0x56, 0x55,
	0xe9, 0xcc, 0x53, 0x50, 0x9d, 0x66, 0xba, 0xc8, 0x79, 0xd9, 0x2a, 0xa6, 0x33, 0x7c, 0x59, 0xe3,
	0xea, 0x8a, 0x36, 0x07, 0x96, 0xea, 0xf8, 0xb5, 0x74, 0x6f, 0x61, 0xdc, 0x98, 0xf5, 0x90, 0x2a,
	0x4d, 0xce, 0xc1, 0x69, 0xfc, 0x7b, 0x55, 0x4e, 0xed, 0x4c, 0x87, 0x73, 0xc7, 0x6b, 0xf8, 0xfc,
	0x23, 0xc7, 0xcd, 0x19, 0xfc, 0xa9, 0x77, 0xca, 0x2b, 0x77, 0xca, 0xab, 0x77, 0x6a, 0xd1, 0x7a,
	0x19, 0x7f, 0x59, 0xab, 0xb0, 0x87, 0x2f, 0x2f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x52, 0xa9,
	0x27, 0x0e, 0x86, 0x02, 0x00, 0x00,
}

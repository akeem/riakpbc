// Code generated by protoc-gen-go.
// source: riak_kv.proto
// DO NOT EDIT!

package riakpbc

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type RpbIndexReq_IndexQueryType int32

const (
	RpbIndexReq_eq    RpbIndexReq_IndexQueryType = 0
	RpbIndexReq_range RpbIndexReq_IndexQueryType = 1
)

var RpbIndexReq_IndexQueryType_name = map[int32]string{
	0: "eq",
	1: "range",
}
var RpbIndexReq_IndexQueryType_value = map[string]int32{
	"eq":    0,
	"range": 1,
}

func (x RpbIndexReq_IndexQueryType) Enum() *RpbIndexReq_IndexQueryType {
	p := new(RpbIndexReq_IndexQueryType)
	*p = x
	return p
}
func (x RpbIndexReq_IndexQueryType) String() string {
	return proto.EnumName(RpbIndexReq_IndexQueryType_name, int32(x))
}
func (x RpbIndexReq_IndexQueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RpbIndexReq_IndexQueryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpbIndexReq_IndexQueryType_value, data, "RpbIndexReq_IndexQueryType")
	if err != nil {
		return err
	}
	*x = RpbIndexReq_IndexQueryType(value)
	return nil
}

type RpbGetClientIdResp struct {
	ClientId         []byte `protobuf:"bytes,1,req,name=client_id" json:"client_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetClientIdResp) Reset()         { *m = RpbGetClientIdResp{} }
func (m *RpbGetClientIdResp) String() string { return proto.CompactTextString(m) }
func (*RpbGetClientIdResp) ProtoMessage()    {}

func (m *RpbGetClientIdResp) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

type RpbSetClientIdReq struct {
	ClientId         []byte `protobuf:"bytes,1,req,name=client_id" json:"client_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbSetClientIdReq) Reset()         { *m = RpbSetClientIdReq{} }
func (m *RpbSetClientIdReq) String() string { return proto.CompactTextString(m) }
func (*RpbSetClientIdReq) ProtoMessage()    {}

func (m *RpbSetClientIdReq) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

type RpbGetReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	R                *uint32 `protobuf:"varint,3,opt,name=r" json:"r,omitempty"`
	Pr               *uint32 `protobuf:"varint,4,opt,name=pr" json:"pr,omitempty"`
	BasicQuorum      *bool   `protobuf:"varint,5,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk       *bool   `protobuf:"varint,6,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	IfModified       []byte  `protobuf:"bytes,7,opt,name=if_modified" json:"if_modified,omitempty"`
	Head             *bool   `protobuf:"varint,8,opt,name=head" json:"head,omitempty"`
	Deletedvclock    *bool   `protobuf:"varint,9,opt,name=deletedvclock" json:"deletedvclock,omitempty"`
	Timeout          *uint32 `protobuf:"varint,10,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbGetReq) Reset()         { *m = RpbGetReq{} }
func (m *RpbGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbGetReq) ProtoMessage()    {}

func (m *RpbGetReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbGetReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbGetReq) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbGetReq) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbGetReq) GetBasicQuorum() bool {
	if m != nil && m.BasicQuorum != nil {
		return *m.BasicQuorum
	}
	return false
}

func (m *RpbGetReq) GetNotfoundOk() bool {
	if m != nil && m.NotfoundOk != nil {
		return *m.NotfoundOk
	}
	return false
}

func (m *RpbGetReq) GetIfModified() []byte {
	if m != nil {
		return m.IfModified
	}
	return nil
}

func (m *RpbGetReq) GetHead() bool {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return false
}

func (m *RpbGetReq) GetDeletedvclock() bool {
	if m != nil && m.Deletedvclock != nil {
		return *m.Deletedvclock
	}
	return false
}

func (m *RpbGetReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type RpbGetResp struct {
	Content          []*RpbContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Vclock           []byte        `protobuf:"bytes,2,opt,name=vclock" json:"vclock,omitempty"`
	Unchanged        *bool         `protobuf:"varint,3,opt,name=unchanged" json:"unchanged,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *RpbGetResp) Reset()         { *m = RpbGetResp{} }
func (m *RpbGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbGetResp) ProtoMessage()    {}

func (m *RpbGetResp) GetContent() []*RpbContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RpbGetResp) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbGetResp) GetUnchanged() bool {
	if m != nil && m.Unchanged != nil {
		return *m.Unchanged
	}
	return false
}

type RpbPutReq struct {
	Bucket           []byte      `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte      `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Vclock           []byte      `protobuf:"bytes,3,opt,name=vclock" json:"vclock,omitempty"`
	Content          *RpbContent `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	W                *uint32     `protobuf:"varint,5,opt,name=w" json:"w,omitempty"`
	Dw               *uint32     `protobuf:"varint,6,opt,name=dw" json:"dw,omitempty"`
	ReturnBody       *bool       `protobuf:"varint,7,opt,name=return_body" json:"return_body,omitempty"`
	Pw               *uint32     `protobuf:"varint,8,opt,name=pw" json:"pw,omitempty"`
	IfNotModified    *bool       `protobuf:"varint,9,opt,name=if_not_modified" json:"if_not_modified,omitempty"`
	IfNoneMatch      *bool       `protobuf:"varint,10,opt,name=if_none_match" json:"if_none_match,omitempty"`
	ReturnHead       *bool       `protobuf:"varint,11,opt,name=return_head" json:"return_head,omitempty"`
	Timeout          *uint32     `protobuf:"varint,12,opt,name=timeout" json:"timeout,omitempty"`
	Asis             *bool       `protobuf:"varint,13,opt,name=asis" json:"asis,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RpbPutReq) Reset()         { *m = RpbPutReq{} }
func (m *RpbPutReq) String() string { return proto.CompactTextString(m) }
func (*RpbPutReq) ProtoMessage()    {}

func (m *RpbPutReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbPutReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbPutReq) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbPutReq) GetContent() *RpbContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RpbPutReq) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbPutReq) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbPutReq) GetReturnBody() bool {
	if m != nil && m.ReturnBody != nil {
		return *m.ReturnBody
	}
	return false
}

func (m *RpbPutReq) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbPutReq) GetIfNotModified() bool {
	if m != nil && m.IfNotModified != nil {
		return *m.IfNotModified
	}
	return false
}

func (m *RpbPutReq) GetIfNoneMatch() bool {
	if m != nil && m.IfNoneMatch != nil {
		return *m.IfNoneMatch
	}
	return false
}

func (m *RpbPutReq) GetReturnHead() bool {
	if m != nil && m.ReturnHead != nil {
		return *m.ReturnHead
	}
	return false
}

func (m *RpbPutReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbPutReq) GetAsis() bool {
	if m != nil && m.Asis != nil {
		return *m.Asis
	}
	return false
}

type RpbPutResp struct {
	Content          []*RpbContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Vclock           []byte        `protobuf:"bytes,2,opt,name=vclock" json:"vclock,omitempty"`
	Key              []byte        `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *RpbPutResp) Reset()         { *m = RpbPutResp{} }
func (m *RpbPutResp) String() string { return proto.CompactTextString(m) }
func (*RpbPutResp) ProtoMessage()    {}

func (m *RpbPutResp) GetContent() []*RpbContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RpbPutResp) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbPutResp) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type RpbDelReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Rw               *uint32 `protobuf:"varint,3,opt,name=rw" json:"rw,omitempty"`
	Vclock           []byte  `protobuf:"bytes,4,opt,name=vclock" json:"vclock,omitempty"`
	R                *uint32 `protobuf:"varint,5,opt,name=r" json:"r,omitempty"`
	W                *uint32 `protobuf:"varint,6,opt,name=w" json:"w,omitempty"`
	Pr               *uint32 `protobuf:"varint,7,opt,name=pr" json:"pr,omitempty"`
	Pw               *uint32 `protobuf:"varint,8,opt,name=pw" json:"pw,omitempty"`
	Dw               *uint32 `protobuf:"varint,9,opt,name=dw" json:"dw,omitempty"`
	Timeout          *uint32 `protobuf:"varint,10,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbDelReq) Reset()         { *m = RpbDelReq{} }
func (m *RpbDelReq) String() string { return proto.CompactTextString(m) }
func (*RpbDelReq) ProtoMessage()    {}

func (m *RpbDelReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbDelReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbDelReq) GetRw() uint32 {
	if m != nil && m.Rw != nil {
		return *m.Rw
	}
	return 0
}

func (m *RpbDelReq) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbDelReq) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbDelReq) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbDelReq) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbDelReq) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbDelReq) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbDelReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type RpbListBucketsResp struct {
	Buckets          [][]byte `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RpbListBucketsResp) Reset()         { *m = RpbListBucketsResp{} }
func (m *RpbListBucketsResp) String() string { return proto.CompactTextString(m) }
func (*RpbListBucketsResp) ProtoMessage()    {}

func (m *RpbListBucketsResp) GetBuckets() [][]byte {
	if m != nil {
		return m.Buckets
	}
	return nil
}

type RpbListKeysReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbListKeysReq) Reset()         { *m = RpbListKeysReq{} }
func (m *RpbListKeysReq) String() string { return proto.CompactTextString(m) }
func (*RpbListKeysReq) ProtoMessage()    {}

func (m *RpbListKeysReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

type RpbListKeysResp struct {
	Keys             [][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Done             *bool    `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RpbListKeysResp) Reset()         { *m = RpbListKeysResp{} }
func (m *RpbListKeysResp) String() string { return proto.CompactTextString(m) }
func (*RpbListKeysResp) ProtoMessage()    {}

func (m *RpbListKeysResp) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RpbListKeysResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

type RpbMapRedReq struct {
	Request          []byte `protobuf:"bytes,1,req,name=request" json:"request,omitempty"`
	ContentType      []byte `protobuf:"bytes,2,req,name=content_type" json:"content_type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbMapRedReq) Reset()         { *m = RpbMapRedReq{} }
func (m *RpbMapRedReq) String() string { return proto.CompactTextString(m) }
func (*RpbMapRedReq) ProtoMessage()    {}

func (m *RpbMapRedReq) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *RpbMapRedReq) GetContentType() []byte {
	if m != nil {
		return m.ContentType
	}
	return nil
}

type RpbMapRedResp struct {
	Phase            *uint32 `protobuf:"varint,1,opt,name=phase" json:"phase,omitempty"`
	Response         []byte  `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Done             *bool   `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbMapRedResp) Reset()         { *m = RpbMapRedResp{} }
func (m *RpbMapRedResp) String() string { return proto.CompactTextString(m) }
func (*RpbMapRedResp) ProtoMessage()    {}

func (m *RpbMapRedResp) GetPhase() uint32 {
	if m != nil && m.Phase != nil {
		return *m.Phase
	}
	return 0
}

func (m *RpbMapRedResp) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RpbMapRedResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

type RpbIndexReq struct {
	Bucket           []byte                      `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Index            []byte                      `protobuf:"bytes,2,req,name=index" json:"index,omitempty"`
	Qtype            *RpbIndexReq_IndexQueryType `protobuf:"varint,3,req,name=qtype,enum=RpbIndexReq_IndexQueryType" json:"qtype,omitempty"`
	Key              []byte                      `protobuf:"bytes,4,opt,name=key" json:"key,omitempty"`
	RangeMin         []byte                      `protobuf:"bytes,5,opt,name=range_min" json:"range_min,omitempty"`
	RangeMax         []byte                      `protobuf:"bytes,6,opt,name=range_max" json:"range_max,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *RpbIndexReq) Reset()         { *m = RpbIndexReq{} }
func (m *RpbIndexReq) String() string { return proto.CompactTextString(m) }
func (*RpbIndexReq) ProtoMessage()    {}

func (m *RpbIndexReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbIndexReq) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *RpbIndexReq) GetQtype() RpbIndexReq_IndexQueryType {
	if m != nil && m.Qtype != nil {
		return *m.Qtype
	}
	return 0
}

func (m *RpbIndexReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbIndexReq) GetRangeMin() []byte {
	if m != nil {
		return m.RangeMin
	}
	return nil
}

func (m *RpbIndexReq) GetRangeMax() []byte {
	if m != nil {
		return m.RangeMax
	}
	return nil
}

type RpbIndexResp struct {
	Keys             [][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RpbIndexResp) Reset()         { *m = RpbIndexResp{} }
func (m *RpbIndexResp) String() string { return proto.CompactTextString(m) }
func (*RpbIndexResp) ProtoMessage()    {}

func (m *RpbIndexResp) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type RpbContent struct {
	Value            []byte     `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	ContentType      []byte     `protobuf:"bytes,2,opt,name=content_type" json:"content_type,omitempty"`
	Charset          []byte     `protobuf:"bytes,3,opt,name=charset" json:"charset,omitempty"`
	ContentEncoding  []byte     `protobuf:"bytes,4,opt,name=content_encoding" json:"content_encoding,omitempty"`
	Vtag             []byte     `protobuf:"bytes,5,opt,name=vtag" json:"vtag,omitempty"`
	Links            []*RpbLink `protobuf:"bytes,6,rep,name=links" json:"links,omitempty"`
	LastMod          *uint32    `protobuf:"varint,7,opt,name=last_mod" json:"last_mod,omitempty"`
	LastModUsecs     *uint32    `protobuf:"varint,8,opt,name=last_mod_usecs" json:"last_mod_usecs,omitempty"`
	Usermeta         []*RpbPair `protobuf:"bytes,9,rep,name=usermeta" json:"usermeta,omitempty"`
	Indexes          []*RpbPair `protobuf:"bytes,10,rep,name=indexes" json:"indexes,omitempty"`
	Deleted          *bool      `protobuf:"varint,11,opt,name=deleted" json:"deleted,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RpbContent) Reset()         { *m = RpbContent{} }
func (m *RpbContent) String() string { return proto.CompactTextString(m) }
func (*RpbContent) ProtoMessage()    {}

func (m *RpbContent) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RpbContent) GetContentType() []byte {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *RpbContent) GetCharset() []byte {
	if m != nil {
		return m.Charset
	}
	return nil
}

func (m *RpbContent) GetContentEncoding() []byte {
	if m != nil {
		return m.ContentEncoding
	}
	return nil
}

func (m *RpbContent) GetVtag() []byte {
	if m != nil {
		return m.Vtag
	}
	return nil
}

func (m *RpbContent) GetLinks() []*RpbLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *RpbContent) GetLastMod() uint32 {
	if m != nil && m.LastMod != nil {
		return *m.LastMod
	}
	return 0
}

func (m *RpbContent) GetLastModUsecs() uint32 {
	if m != nil && m.LastModUsecs != nil {
		return *m.LastModUsecs
	}
	return 0
}

func (m *RpbContent) GetUsermeta() []*RpbPair {
	if m != nil {
		return m.Usermeta
	}
	return nil
}

func (m *RpbContent) GetIndexes() []*RpbPair {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *RpbContent) GetDeleted() bool {
	if m != nil && m.Deleted != nil {
		return *m.Deleted
	}
	return false
}

type RpbLink struct {
	Bucket           []byte `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key              []byte `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Tag              []byte `protobuf:"bytes,3,opt,name=tag" json:"tag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbLink) Reset()         { *m = RpbLink{} }
func (m *RpbLink) String() string { return proto.CompactTextString(m) }
func (*RpbLink) ProtoMessage()    {}

func (m *RpbLink) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbLink) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbLink) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func init() {
	proto.RegisterEnum("RpbIndexReq_IndexQueryType", RpbIndexReq_IndexQueryType_name, RpbIndexReq_IndexQueryType_value)
}
package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BatchJobRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "User":
			z.User, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "User")
				return
			}
		case "Started":
			z.Started, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "Started")
				return
			}
		case "Location":
			z.Location, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Location")
				return
			}
		case "Replicate":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Replicate")
					return
				}
				z.Replicate = nil
			} else {
				if z.Replicate == nil {
					z.Replicate = new(BatchJobReplicateV1)
				}
				err = z.Replicate.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Replicate")
					return
				}
			}
		case "KeyRotate":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "KeyRotate")
					return
				}
				z.KeyRotate = nil
			} else {
				if z.KeyRotate == nil {
					z.KeyRotate = new(BatchJobKeyRotateV1)
				}
				err = z.KeyRotate.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "KeyRotate")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BatchJobRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "ID"
	err = en.Append(0x86, 0xa2, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteString(z.ID)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	// write "User"
	err = en.Append(0xa4, 0x55, 0x73, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.User)
	if err != nil {
		err = msgp.WrapError(err, "User")
		return
	}
	// write "Started"
	err = en.Append(0xa7, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Started)
	if err != nil {
		err = msgp.WrapError(err, "Started")
		return
	}
	// write "Location"
	err = en.Append(0xa8, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Location)
	if err != nil {
		err = msgp.WrapError(err, "Location")
		return
	}
	// write "Replicate"
	err = en.Append(0xa9, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	if z.Replicate == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Replicate.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Replicate")
			return
		}
	}
	// write "KeyRotate"
	err = en.Append(0xa9, 0x4b, 0x65, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	if z.KeyRotate == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.KeyRotate.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "KeyRotate")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BatchJobRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "ID"
	o = append(o, 0x86, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "User"
	o = append(o, 0xa4, 0x55, 0x73, 0x65, 0x72)
	o = msgp.AppendString(o, z.User)
	// string "Started"
	o = append(o, 0xa7, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Started)
	// string "Location"
	o = append(o, 0xa8, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Location)
	// string "Replicate"
	o = append(o, 0xa9, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65)
	if z.Replicate == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Replicate.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Replicate")
			return
		}
	}
	// string "KeyRotate"
	o = append(o, 0xa9, 0x4b, 0x65, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65)
	if z.KeyRotate == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.KeyRotate.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "KeyRotate")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchJobRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "User":
			z.User, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "User")
				return
			}
		case "Started":
			z.Started, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Started")
				return
			}
		case "Location":
			z.Location, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Location")
				return
			}
		case "Replicate":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Replicate = nil
			} else {
				if z.Replicate == nil {
					z.Replicate = new(BatchJobReplicateV1)
				}
				bts, err = z.Replicate.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Replicate")
					return
				}
			}
		case "KeyRotate":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.KeyRotate = nil
			} else {
				if z.KeyRotate == nil {
					z.KeyRotate = new(BatchJobKeyRotateV1)
				}
				bts, err = z.KeyRotate.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "KeyRotate")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BatchJobRequest) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.ID) + 5 + msgp.StringPrefixSize + len(z.User) + 8 + msgp.TimeSize + 9 + msgp.StringPrefixSize + len(z.Location) + 10
	if z.Replicate == nil {
		s += msgp.NilSize
	} else {
		s += z.Replicate.Msgsize()
	}
	s += 10
	if z.KeyRotate == nil {
		s += msgp.NilSize
	} else {
		s += z.KeyRotate.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *batchJobInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "jid":
			z.JobID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "JobID")
				return
			}
		case "jt":
			z.JobType, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "JobType")
				return
			}
		case "st":
			z.StartTime, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "lu":
			z.LastUpdate, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "ra":
			z.RetryAttempts, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "RetryAttempts")
				return
			}
		case "cmp":
			z.Complete, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Complete")
				return
			}
		case "fld":
			z.Failed, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Failed")
				return
			}
		case "lbkt":
			z.Bucket, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "lobj":
			z.Object, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Object")
				return
			}
		case "ob":
			z.Objects, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Objects")
				return
			}
		case "dm":
			z.DeleteMarkers, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "DeleteMarkers")
				return
			}
		case "obf":
			z.ObjectsFailed, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ObjectsFailed")
				return
			}
		case "dmf":
			z.DeleteMarkersFailed, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "DeleteMarkersFailed")
				return
			}
		case "bt":
			z.BytesTransferred, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "BytesTransferred")
				return
			}
		case "bf":
			z.BytesFailed, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "BytesFailed")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *batchJobInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 16
	// write "v"
	err = en.Append(0xde, 0x0, 0x10, 0xa1, 0x76)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Version)
	if err != nil {
		err = msgp.WrapError(err, "Version")
		return
	}
	// write "jid"
	err = en.Append(0xa3, 0x6a, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.JobID)
	if err != nil {
		err = msgp.WrapError(err, "JobID")
		return
	}
	// write "jt"
	err = en.Append(0xa2, 0x6a, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.JobType)
	if err != nil {
		err = msgp.WrapError(err, "JobType")
		return
	}
	// write "st"
	err = en.Append(0xa2, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteTime(z.StartTime)
	if err != nil {
		err = msgp.WrapError(err, "StartTime")
		return
	}
	// write "lu"
	err = en.Append(0xa2, 0x6c, 0x75)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LastUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastUpdate")
		return
	}
	// write "ra"
	err = en.Append(0xa2, 0x72, 0x61)
	if err != nil {
		return
	}
	err = en.WriteInt(z.RetryAttempts)
	if err != nil {
		err = msgp.WrapError(err, "RetryAttempts")
		return
	}
	// write "cmp"
	err = en.Append(0xa3, 0x63, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Complete)
	if err != nil {
		err = msgp.WrapError(err, "Complete")
		return
	}
	// write "fld"
	err = en.Append(0xa3, 0x66, 0x6c, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Failed)
	if err != nil {
		err = msgp.WrapError(err, "Failed")
		return
	}
	// write "lbkt"
	err = en.Append(0xa4, 0x6c, 0x62, 0x6b, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bucket)
	if err != nil {
		err = msgp.WrapError(err, "Bucket")
		return
	}
	// write "lobj"
	err = en.Append(0xa4, 0x6c, 0x6f, 0x62, 0x6a)
	if err != nil {
		return
	}
	err = en.WriteString(z.Object)
	if err != nil {
		err = msgp.WrapError(err, "Object")
		return
	}
	// write "ob"
	err = en.Append(0xa2, 0x6f, 0x62)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Objects)
	if err != nil {
		err = msgp.WrapError(err, "Objects")
		return
	}
	// write "dm"
	err = en.Append(0xa2, 0x64, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.DeleteMarkers)
	if err != nil {
		err = msgp.WrapError(err, "DeleteMarkers")
		return
	}
	// write "obf"
	err = en.Append(0xa3, 0x6f, 0x62, 0x66)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ObjectsFailed)
	if err != nil {
		err = msgp.WrapError(err, "ObjectsFailed")
		return
	}
	// write "dmf"
	err = en.Append(0xa3, 0x64, 0x6d, 0x66)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.DeleteMarkersFailed)
	if err != nil {
		err = msgp.WrapError(err, "DeleteMarkersFailed")
		return
	}
	// write "bt"
	err = en.Append(0xa2, 0x62, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BytesTransferred)
	if err != nil {
		err = msgp.WrapError(err, "BytesTransferred")
		return
	}
	// write "bf"
	err = en.Append(0xa2, 0x62, 0x66)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BytesFailed)
	if err != nil {
		err = msgp.WrapError(err, "BytesFailed")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *batchJobInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 16
	// string "v"
	o = append(o, 0xde, 0x0, 0x10, 0xa1, 0x76)
	o = msgp.AppendInt(o, z.Version)
	// string "jid"
	o = append(o, 0xa3, 0x6a, 0x69, 0x64)
	o = msgp.AppendString(o, z.JobID)
	// string "jt"
	o = append(o, 0xa2, 0x6a, 0x74)
	o = msgp.AppendString(o, z.JobType)
	// string "st"
	o = append(o, 0xa2, 0x73, 0x74)
	o = msgp.AppendTime(o, z.StartTime)
	// string "lu"
	o = append(o, 0xa2, 0x6c, 0x75)
	o = msgp.AppendTime(o, z.LastUpdate)
	// string "ra"
	o = append(o, 0xa2, 0x72, 0x61)
	o = msgp.AppendInt(o, z.RetryAttempts)
	// string "cmp"
	o = append(o, 0xa3, 0x63, 0x6d, 0x70)
	o = msgp.AppendBool(o, z.Complete)
	// string "fld"
	o = append(o, 0xa3, 0x66, 0x6c, 0x64)
	o = msgp.AppendBool(o, z.Failed)
	// string "lbkt"
	o = append(o, 0xa4, 0x6c, 0x62, 0x6b, 0x74)
	o = msgp.AppendString(o, z.Bucket)
	// string "lobj"
	o = append(o, 0xa4, 0x6c, 0x6f, 0x62, 0x6a)
	o = msgp.AppendString(o, z.Object)
	// string "ob"
	o = append(o, 0xa2, 0x6f, 0x62)
	o = msgp.AppendInt64(o, z.Objects)
	// string "dm"
	o = append(o, 0xa2, 0x64, 0x6d)
	o = msgp.AppendInt64(o, z.DeleteMarkers)
	// string "obf"
	o = append(o, 0xa3, 0x6f, 0x62, 0x66)
	o = msgp.AppendInt64(o, z.ObjectsFailed)
	// string "dmf"
	o = append(o, 0xa3, 0x64, 0x6d, 0x66)
	o = msgp.AppendInt64(o, z.DeleteMarkersFailed)
	// string "bt"
	o = append(o, 0xa2, 0x62, 0x74)
	o = msgp.AppendInt64(o, z.BytesTransferred)
	// string "bf"
	o = append(o, 0xa2, 0x62, 0x66)
	o = msgp.AppendInt64(o, z.BytesFailed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *batchJobInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "jid":
			z.JobID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "JobID")
				return
			}
		case "jt":
			z.JobType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "JobType")
				return
			}
		case "st":
			z.StartTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "lu":
			z.LastUpdate, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "ra":
			z.RetryAttempts, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RetryAttempts")
				return
			}
		case "cmp":
			z.Complete, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Complete")
				return
			}
		case "fld":
			z.Failed, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Failed")
				return
			}
		case "lbkt":
			z.Bucket, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "lobj":
			z.Object, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Object")
				return
			}
		case "ob":
			z.Objects, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Objects")
				return
			}
		case "dm":
			z.DeleteMarkers, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DeleteMarkers")
				return
			}
		case "obf":
			z.ObjectsFailed, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ObjectsFailed")
				return
			}
		case "dmf":
			z.DeleteMarkersFailed, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DeleteMarkersFailed")
				return
			}
		case "bt":
			z.BytesTransferred, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BytesTransferred")
				return
			}
		case "bf":
			z.BytesFailed, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BytesFailed")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *batchJobInfo) Msgsize() (s int) {
	s = 3 + 2 + msgp.IntSize + 4 + msgp.StringPrefixSize + len(z.JobID) + 3 + msgp.StringPrefixSize + len(z.JobType) + 3 + msgp.TimeSize + 3 + msgp.TimeSize + 3 + msgp.IntSize + 4 + msgp.BoolSize + 4 + msgp.BoolSize + 5 + msgp.StringPrefixSize + len(z.Bucket) + 5 + msgp.StringPrefixSize + len(z.Object) + 3 + msgp.Int64Size + 3 + msgp.Int64Size + 4 + msgp.Int64Size + 4 + msgp.Int64Size + 3 + msgp.Int64Size + 3 + msgp.Int64Size
	return
}

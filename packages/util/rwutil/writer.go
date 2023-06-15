// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package rwutil

import (
	"bytes"
	"encoding"
	"errors"
	"io"
	"math"
	"math/big"
	"time"

	"github.com/iotaledger/hive.go/serializer/v2"
)

type Writer struct {
	Err error
	w   io.Writer
}

func NewWriter(w io.Writer) *Writer {
	if w == nil {
		panic("nil io.Writer")
	}
	return &Writer{w: w}
}

func NewBytesWriter() *Writer {
	return NewWriter(new(bytes.Buffer))
}

func (ww *Writer) Bytes() []byte {
	buf, ok := ww.w.(*bytes.Buffer)
	if !ok {
		panic("writer expects bytes buffer")
	}
	if ww.Err != nil {
		// writing to bytes buffer never fails
		panic(ww.Err)
	}
	return buf.Bytes()
}

func (ww *Writer) Skip() *Reader {
	skip := &Skipper{ww: ww, w: ww.w}
	ww.w = skip
	return &Reader{r: skip}
}

func (ww *Writer) Write(writer interface{ Write(w io.Writer) error }) *Writer {
	if writer == nil {
		panic("nil writer")
	}
	if ww.Err == nil {
		ww.Err = writer.Write(ww.w)
	}
	return ww
}

func (ww *Writer) WriteN(val []byte) *Writer {
	if ww.Err == nil {
		ww.Err = WriteN(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteBool(val bool) *Writer {
	if ww.Err == nil {
		ww.Err = WriteBool(ww.w, val)
	}
	return ww
}

//nolint:govet
func (ww *Writer) WriteByte(val byte) *Writer {
	if ww.Err == nil {
		ww.Err = WriteByte(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteBytes(val []byte) *Writer {
	if ww.Err == nil {
		ww.Err = WriteBytes(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteDuration(val time.Duration) *Writer {
	return ww.WriteInt64(int64(val))
}

func (ww *Writer) WriteFromBytes(bytes interface{ Bytes() []byte }) *Writer {
	if ww.Err == nil {
		ww.WriteBytes(bytes.Bytes())
	}
	return ww
}

func (ww *Writer) WriteInt8(val int8) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt8(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteInt16(val int16) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt16(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteInt32(val int32) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt32(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteInt64(val int64) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt64(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteKind(msgType Kind) *Writer {
	return ww.WriteByte(byte(msgType))
}

func (ww *Writer) WriteMarshaled(m encoding.BinaryMarshaler) *Writer {
	if m == nil {
		panic("nil marshaler")
	}
	if ww.Err == nil {
		var buf []byte
		buf, ww.Err = m.MarshalBinary()
		ww.WriteBytes(buf)
	}
	return ww
}

type serializable interface {
	Serialize(serializer.DeSerializationMode, interface{}) ([]byte, error)
}

func (ww *Writer) WriteSerialized(s serializable, sizes ...int) *Writer {
	if s == nil {
		panic("nil serializer")
	}
	if ww.Err != nil {
		return ww
	}

	var buf []byte
	buf, ww.Err = s.Serialize(serializer.DeSeriModeNoValidation, nil)
	switch len(sizes) {
	case 0:
		ww.WriteSize16(len(buf))
	case 1:
		limit := sizes[0]
		if limit < 0 || limit > math.MaxInt32 {
			panic("invalid serialize limit")
		}
		ww.WriteSizeWithLimit(len(buf), uint32(limit))
	case 2:
		size := sizes[1]
		if size < 0 || size > math.MaxInt32 {
			panic("invalid serialize size")
		}
		if size != len(buf) && ww.Err == nil {
			ww.Err = errors.New("unexpected serialize size")
		}
	default:
		panic("too many serialize params")
	}
	ww.WriteN(buf)
	return ww
}

func (ww *Writer) WriteSize16(val int) *Writer {
	return ww.WriteSizeWithLimit(val, math.MaxUint16)
}

func (ww *Writer) WriteSize32(val int) *Writer {
	// note we cannot exceed SIGNED max
	// because if int is actually 32 bit it would be negative
	return ww.WriteSizeWithLimit(val, math.MaxInt32)
}

func (ww *Writer) WriteSizeWithLimit(val int, limit uint32) *Writer {
	if ww.Err == nil {
		if val < 0 {
			ww.Err = errors.New("write size limit underflow")
			return ww
		}
		if val > int(limit) {
			ww.Err = errors.New("write size limit overflow")
			return ww
		}
		ww.Err = WriteSize32(ww.w, uint32(val))
	}
	return ww
}

func (ww *Writer) WriteString(val string) *Writer {
	if ww.Err == nil {
		ww.Err = WriteString(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint8(val uint8) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint8(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint16(val uint16) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint16(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint32(val uint32) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint32(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint64(val uint64) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint64(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint256(val *big.Int) *Writer {
	if val == nil {
		val = new(big.Int)
	}
	if ww.Err == nil && val.Sign() < 0 {
		ww.Err = errors.New("negative uint256")
		return ww
	}
	return ww.WriteBytes(val.Bytes())
}

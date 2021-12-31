//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes Int16 rows from *Reader.
func (c *ColInt16) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]int16, rows)...)
	s := *(*slice)(unsafe.Pointer(c))
	s.Len *= 2
	s.Cap *= 2
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}

// EncodeColumn encodes Int16 rows to *Buffer.
func (c ColInt16) EncodeColumn(b *Buffer) {
	if len(c) == 0 {
		return
	}
	offset := len(b.Buf)
	const size = 16 / 8
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)

	s := *(*slice)(unsafe.Pointer(&c))
	s.Len *= 2
	s.Cap *= 2
	src := *(*[]byte)(unsafe.Pointer(&s))
	dst := b.Buf[offset:]
	copy(dst, src)
}

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
)

var _ = binary.LittleEndian // clickHouse uses LittleEndian

// ColInt256 represents Int256 column.
type ColInt256 []Int256

// Compile-time assertions for ColInt256.
var (
	_ ColInput  = ColInt256{}
	_ ColResult = (*ColInt256)(nil)
	_ Column    = (*ColInt256)(nil)
)

// Type returns ColumnType of Int256.
func (ColInt256) Type() ColumnType {
	return ColumnTypeInt256
}

// Rows returns count of rows in column.
func (c ColInt256) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt256) Row(i int) Int256 {
	return c[i]
}

// Append Int256 to column.
func (c *ColInt256) Append(v Int256) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt256) Reset() {
	*c = (*c)[:0]
}

// NewArrInt256 returns new Array(Int256).
func NewArrInt256() *ColArr {
	return &ColArr{
		Data: new(ColInt256),
	}
}

// AppendInt256 appends slice of Int256 to Array(Int256).
func (c *ColArr) AppendInt256(data []Int256) {
	d := c.Data.(*ColInt256)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

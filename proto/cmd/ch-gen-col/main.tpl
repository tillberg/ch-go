{{- /*gotype: github.com/go-faster/ch/proto/cmd/ch-gen-col.Variant*/ -}}
// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
)

var _ = binary.LittleEndian // clickHouse uses LittleEndian

// {{ .Type }} represents {{ .Name }} column.
type {{ .Type }} []{{ .ElemType }}

// Compile-time assertions for {{ .Type }}.
var (
  _ ColInput  = {{ .Type }}{}
  _ ColResult = (*{{ .Type }})(nil)
  _ Column    = (*{{ .Type }})(nil)
)

// Type returns ColumnType of {{ .Name }}.
func ({{ .Type }}) Type() ColumnType {
  return {{ .ColumnType }}
}

// Rows returns count of rows in column.
func (c {{ .Type }}) Rows() int {
  return len(c)
}

// Row returns i-th row of column.
func (c {{ .Type }}) Row(i int) {{ .ElemType }} {
  return c[i]
}

// Append {{ .ElemType }} to column.
func (c *{{ .Type }}) Append(v {{ .ElemType }})  {
  *c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *{{ .Type }}) Reset() {
  *c = (*c)[:0]
}

// NewArr{{ .Name }} returns new Array({{ .Name }}).
func NewArr{{ .Name }}() *ColArr {
  return &ColArr{
	Data: new({{ .Type }}),
  }
}

// Append{{ .Name }} appends slice of {{ .ElemType }} to Array({{ .Name }}).
func (c *ColArr) Append{{ .Name }}(data []{{ .ElemType }}) {
  d := c.Data.(*{{ .Type }})
  *d = append(*d, data...)
  c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// This file is automatically generated. Do not modify.

package gentest

import (
	"fmt"
)

var _ = fmt.Sprintf

type Binary []byte
type Int32 int32
type String string

type St struct {
	B *Binary `thrift:"1,required" json:"b"`
	S *String `thrift:"2,required" json:"s"`
	I *Int32  `thrift:"3,required" json:"i"`
}

// This file is automatically generated. Do not modify.

package gentest

import (
	"fmt"
)

var _ = fmt.Sprintf

type NestedColor struct {
	Rgb *Rgb `thrift:"1,required" json:"rgb"`
}

type Rgb struct {
	Red   *int32 `thrift:"1,required" json:"red"`
	Green *int32 `thrift:"2,required" json:"green"`
	Blue  *int32 `thrift:"3,required" json:"blue"`
}

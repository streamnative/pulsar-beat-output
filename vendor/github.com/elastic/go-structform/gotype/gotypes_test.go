// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package gotype

import (
	"runtime"
	"testing"

	structform "github.com/elastic/go-structform"
	"github.com/elastic/go-structform/sftest"
)

type mapstr map[string]interface{}

// special visitor enforcing a gc before and after every
// every operators, checking the folder/unfolder to not hold on
// invalid pointers.
type gcCheckVisitor struct {
	u structform.ExtVisitor
}

func TestFoldUnfoldToIfcConsistent(t *testing.T) {
	sftest.TestEncodeParseConsistent(t, sftest.Samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			var v interface{}
			unfolder, err := NewUnfolder(&v)
			if err != nil {
				panic(err)
			}
			return newGCCheckVisitor(unfolder), func(to structform.Visitor) error {
				return Fold(v, to)
			}
		})
}

func newGCCheckVisitor(v structform.Visitor) *gcCheckVisitor {
	return &gcCheckVisitor{structform.EnsureExtVisitor(v)}
}

func (g *gcCheckVisitor) OnObjectStart(len int, baseType structform.BaseType) error {
	runtime.GC()
	err := g.u.OnObjectStart(len, baseType)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnObjectFinished() error {
	runtime.GC()
	err := g.u.OnObjectFinished()
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnKey(s string) error {
	runtime.GC()
	err := g.u.OnKey(s)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnArrayStart(len int, baseType structform.BaseType) error {
	runtime.GC()
	err := g.u.OnArrayStart(len, baseType)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnArrayFinished() error {
	runtime.GC()
	err := g.u.OnArrayFinished()
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnNil() error {
	runtime.GC()
	err := g.u.OnNil()
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnBool(b bool) error {
	runtime.GC()
	err := g.u.OnBool(b)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnString(s string) error {
	runtime.GC()
	err := g.u.OnString(s)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt8(i int8) error {
	runtime.GC()
	err := g.u.OnInt8(i)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt16(i int16) error {
	runtime.GC()
	err := g.u.OnInt16(i)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt32(i int32) error {
	runtime.GC()
	err := g.u.OnInt32(i)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt64(i int64) error {
	runtime.GC()
	err := g.u.OnInt64(i)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt(i int) error {
	runtime.GC()
	err := g.u.OnInt(i)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnByte(b byte) error {
	runtime.GC()
	err := g.u.OnByte(b)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint8(u uint8) error {
	runtime.GC()
	err := g.u.OnUint8(u)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint16(u uint16) error {
	runtime.GC()
	err := g.u.OnUint16(u)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint32(u uint32) error {
	runtime.GC()
	err := g.u.OnUint32(u)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint64(u uint64) error {
	runtime.GC()
	err := g.u.OnUint64(u)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint(u uint) error {
	runtime.GC()
	err := g.u.OnUint(u)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnFloat32(f float32) error {
	runtime.GC()
	err := g.u.OnFloat32(f)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnFloat64(f float64) error {
	runtime.GC()
	err := g.u.OnFloat64(f)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnBoolArray(a []bool) error {
	runtime.GC()
	err := g.u.OnBoolArray(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnStringArray(a []string) error {
	runtime.GC()
	err := g.u.OnStringArray(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt8Array(a []int8) error {
	runtime.GC()
	err := g.u.OnInt8Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt16Array(a []int16) error {
	runtime.GC()
	err := g.u.OnInt16Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt32Array(a []int32) error {
	runtime.GC()
	err := g.u.OnInt32Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt64Array(a []int64) error {
	runtime.GC()
	err := g.u.OnInt64Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnIntArray(a []int) error {
	runtime.GC()
	err := g.u.OnIntArray(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnBytes(a []byte) error {
	runtime.GC()
	err := g.u.OnBytes(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint8Array(a []uint8) error {
	runtime.GC()
	err := g.u.OnUint8Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint16Array(a []uint16) error {
	runtime.GC()
	err := g.u.OnUint16Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint32Array(a []uint32) error {
	runtime.GC()
	err := g.u.OnUint32Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint64Array(a []uint64) error {
	runtime.GC()
	err := g.u.OnUint64Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUintArray(a []uint) error {
	runtime.GC()
	err := g.u.OnUintArray(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnFloat32Array(a []float32) error {
	runtime.GC()
	err := g.u.OnFloat32Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnFloat64Array(a []float64) error {
	runtime.GC()
	err := g.u.OnFloat64Array(a)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnBoolObject(o map[string]bool) error {
	runtime.GC()
	err := g.u.OnBoolObject(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnStringObject(o map[string]string) error {
	runtime.GC()
	err := g.u.OnStringObject(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt8Object(o map[string]int8) error {
	runtime.GC()
	err := g.u.OnInt8Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt16Object(o map[string]int16) error {
	runtime.GC()
	err := g.u.OnInt16Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt32Object(o map[string]int32) error {
	runtime.GC()
	err := g.u.OnInt32Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnInt64Object(o map[string]int64) error {
	runtime.GC()
	err := g.u.OnInt64Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnIntObject(o map[string]int) error {
	runtime.GC()
	err := g.u.OnIntObject(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint8Object(o map[string]uint8) error {
	runtime.GC()
	err := g.u.OnUint8Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint16Object(o map[string]uint16) error {
	runtime.GC()
	err := g.u.OnUint16Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint32Object(o map[string]uint32) error {
	runtime.GC()
	err := g.u.OnUint32Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUint64Object(o map[string]uint64) error {
	runtime.GC()
	err := g.u.OnUint64Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnUintObject(o map[string]uint) error {
	runtime.GC()
	err := g.u.OnUintObject(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnFloat32Object(o map[string]float32) error {
	runtime.GC()
	err := g.u.OnFloat32Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnFloat64Object(o map[string]float64) error {
	runtime.GC()
	err := g.u.OnFloat64Object(o)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnStringRef(s []byte) error {
	runtime.GC()
	err := g.u.OnStringRef(s)
	runtime.GC()
	return err
}

func (g *gcCheckVisitor) OnKeyRef(s []byte) error {
	runtime.GC()
	err := g.u.OnKeyRef(s)
	runtime.GC()
	return err
}

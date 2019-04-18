package bin

import (
	"errors"
	"fmt"
	"testing"
	"testing/quick"
	"unsafe"
)

func TestCast(t *testing.T) {
	tests := []struct {
		title string
		cast  func(interface{}, []byte) error
	}{
		{"unsafe", unsafeCast},
		{"safe", CastStruct},
	}

	type Product struct {
		A uint64
		B uint32
		C struct {
			D uint16
		}
	}

	eq := func(a, b Product) bool {
		return a.A == b.A && a.B == b.B && a.C.D == b.C.D
	}

	for _, test := range tests {
		castFn := test.cast
		t.Run(fmt.Sprintf("go struct: %v", test.title), func(t *testing.T) {
			err := quick.Check(func(in Product) bool {
				if eq(in, Product{}) { // tmp is all zeros -> check input is not 'zero'
					t.Log("zero input")
					return true
				}

				var buf [unsafe.Sizeof(Product{})]byte
				var tmp *Product

				if err := castFn(&tmp, buf[:]); err != nil {
					t.Error(err)
					return false
				}

				// set contents
				*tmp = in

				// check contents being written to buffer
				if countNonZero(buf[:]) == 0 {
					t.Error(errors.New("buffer is empty"))
					return false
				}

				return true
			}, nil)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run(fmt.Sprintf("enc struct: %v", test.title), func(t *testing.T) {
			err := quick.Check(func(in Product) bool {
				if eq(in, Product{}) { // tmp is all zeros -> check input is not 'zero'
					return true
				}

				var buf [unsafe.Sizeof(Product{})]byte
				tmp := &struct {
					A U64le
					B U32be
					C U16le
				}{}

				if err := castFn(&tmp, buf[:]); err != nil {
					t.Error(err)
					return false
				}

				// encode contents
				tmp.A.Set(in.A)
				tmp.B.Set(in.B)
				tmp.C.Set(in.C.D)

				// check contents being written to buffer
				if countNonZero(buf[:]) == 0 {
					t.Error(errors.New("buffer is empty"))
					return false
				}

				// decode and validate:
				decoded := Product{}
				decoded.A = tmp.A.Get()
				decoded.B = tmp.B.Get()
				decoded.C.D = tmp.C.Get()

				if !eq(in, decoded) {
					t.Error("decoding error")
					return false
				}

				return true
			}, nil)
			if err != nil {
				t.Error(err)
			}
		})
	}

}

func countNonZero(b []byte) int {
	count := 0
	for _, v := range b {
		if v != 0 {
			count++
		}
	}
	return count
}

// unsafeCast wraps UnsafeCastStruct, in order to have same interface as
// CastStruct.
func unsafeCast(to interface{}, b []byte) error {
	UnsafeCastStruct(to, b)
	return nil
}

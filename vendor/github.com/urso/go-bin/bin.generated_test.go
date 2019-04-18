// This file has been generated from 'bin_test.yml', do not edit
package bin

import (
	"bytes"
	"encoding/binary"
	"testing"
	"testing/quick"
)

func TestPrimitives(t *testing.T) {

	t.Run("int8 big endian", func(t *testing.T) {
		var v I8be
		err := quick.Check(func(in int8) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			tmp[0] = byte(in)

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int16 big endian", func(t *testing.T) {
		var v I16be
		err := quick.Check(func(in int16) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.BigEndian.PutUint16(tmp, uint16(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int32 big endian", func(t *testing.T) {
		var v I32be
		err := quick.Check(func(in int32) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.BigEndian.PutUint32(tmp, uint32(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int64 big endian", func(t *testing.T) {
		var v I64be
		err := quick.Check(func(in int64) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.BigEndian.PutUint64(tmp, uint64(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint8 big endian", func(t *testing.T) {
		var v U8be
		err := quick.Check(func(in uint8) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			tmp[0] = byte(in)

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint16 big endian", func(t *testing.T) {
		var v U16be
		err := quick.Check(func(in uint16) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.BigEndian.PutUint16(tmp, uint16(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint32 big endian", func(t *testing.T) {
		var v U32be
		err := quick.Check(func(in uint32) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.BigEndian.PutUint32(tmp, uint32(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint64 big endian", func(t *testing.T) {
		var v U64be
		err := quick.Check(func(in uint64) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.BigEndian.PutUint64(tmp, uint64(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int8 little endian", func(t *testing.T) {
		var v I8le
		err := quick.Check(func(in int8) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			tmp[0] = byte(in)

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int16 little endian", func(t *testing.T) {
		var v I16le
		err := quick.Check(func(in int16) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.LittleEndian.PutUint16(tmp, uint16(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int32 little endian", func(t *testing.T) {
		var v I32le
		err := quick.Check(func(in int32) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.LittleEndian.PutUint32(tmp, uint32(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("int64 little endian", func(t *testing.T) {
		var v I64le
		err := quick.Check(func(in int64) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.LittleEndian.PutUint64(tmp, uint64(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint8 little endian", func(t *testing.T) {
		var v U8le
		err := quick.Check(func(in uint8) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			tmp[0] = byte(in)

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint16 little endian", func(t *testing.T) {
		var v U16le
		err := quick.Check(func(in uint16) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.LittleEndian.PutUint16(tmp, uint16(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint32 little endian", func(t *testing.T) {
		var v U32le
		err := quick.Check(func(in uint32) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.LittleEndian.PutUint32(tmp, uint32(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("uint64 little endian", func(t *testing.T) {
		var v U64le
		err := quick.Check(func(in uint64) bool {
			v.Set(in)

			// check raw contents correct encoding
			tmp := make([]byte, v.Len())

			binary.LittleEndian.PutUint64(tmp, uint64(in))

			if !bytes.Equal(v[:], tmp) {
				t.Error("encoding mismatch")
				return false
			}

			// check extracted value matches original value
			return v.Get() == in
		}, nil)
		if err != nil {
			t.Error(err)
		}
	})

}

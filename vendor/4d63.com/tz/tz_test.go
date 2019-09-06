package tz_test

import (
	"fmt"
	"testing"
	"time"

	"4d63.com/tz"
)

func ExampleLoadLocation() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t)

	syd, _ := tz.LoadLocation("Australia/Sydney")
	fmt.Println(t.In(syd))

	// Output:
	// 2009-11-10 23:00:00 +0000 UTC
	// 2009-11-11 10:00:00 +1100 AEDT
}

func TestLoadLocation_Same(t *testing.T) {
	cases := []struct {
		name string
	}{
		{""},
		{"UTC"},
		{"Local"},
		{"Etc/GMT"},
		{"Etc/GMT-0"},
		{"Etc/GMT-1"},
		{"Etc/GMT-2"},
		{"Etc/GMT-4"},
		{"Etc/GMT-5"},
		{"Etc/GMT-6"},
		{"Etc/GMT-7"},
		{"Etc/GMT-8"},
		{"Etc/GMT-9"},
		{"Etc/GMT-10"},
		{"Etc/GMT-11"},
		{"Etc/GMT+0"},
		{"Etc/GMT+1"},
		{"Etc/GMT+2"},
		{"Etc/GMT+4"},
		{"Etc/GMT+5"},
		{"Etc/GMT+6"},
		{"Etc/GMT+7"},
		{"Etc/GMT+8"},
		{"Etc/GMT+9"},
		{"Etc/GMT+10"},
		{"Etc/GMT+11"},
		{"Australia/Sydney"},
		{"America/New_York"},
		{"US/Central"},
		{"US/Eastern"},
		{"US/Pacific"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pkgLoc, err := tz.LoadLocation(c.name)
			if err != nil {
				t.Fatalf("error loading package location: %s", err)
			}
			stdlibLoc, err := time.LoadLocation(c.name)
			if err != nil {
				t.Fatalf("error loading stdlib location: %s", err)
			}

			if pkgLoc.String() == stdlibLoc.String() {
				t.Logf("tz package returns same named Location as stdlib")
			} else {
				t.Fatalf("tz package returns different named Location as stdlib, got %s, want %s", pkgLoc, stdlibLoc)
			}

			utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
			t.Logf("utc: %s", utc)

			pkgTime := utc.In(pkgLoc)
			t.Logf("tz: %s", pkgTime)

			stdlibTime := utc.In(stdlibLoc)
			t.Logf("stdlib: %s", stdlibTime)

			if pkgTime.Equal(stdlibTime) {
				t.Logf("tz package location represents time the same as stdlib")
			} else {
				t.Fatalf("tz package location represents time the different to stdlib, got %s, want %s", pkgTime, stdlibTime)
			}
		})
	}
}

func TestLoadLocation_Unknown(t *testing.T) {
	name := "not-a-real-location"
	wantErr := "unknown location not-a-real-location"

	_, err := tz.LoadLocation(name)

	if err.Error() == wantErr {
		t.Logf("got %s", err)
	} else {
		t.Errorf("got %s, want %s", err.Error(), wantErr)
	}
}

func BenchmarkLoadLocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tz.LoadLocation("US/Central")
	}
}

func BenchmarkLoadLocation_StdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.LoadLocation("US/Central")
	}
}

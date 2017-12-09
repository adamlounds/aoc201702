package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLineChecksum(t *testing.T) {
	Convey("Given a tab-separated string", t, func() {
		line := "1\t2\t3\t4"
		rowChecksum, err := calcLineChecksum(line)
		Convey("We should be able to parse with no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("And the checksum should be largest - smallest", func() {
			So(rowChecksum, ShouldEqual, 3)
		})
	})
}

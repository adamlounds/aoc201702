package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLineChecksumA(t *testing.T) {
	Convey("Given a tab-separated string", t, func() {
		line := "1\t2\t3\t4"
		rowChecksumA, err := calcLineChecksumA(line)
		Convey("We should be able to parse with no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("And the checksum should be largest - smallest", func() {
			So(rowChecksumA, ShouldEqual, 3)
		})
	})
}
func TestLineChecksumB(t *testing.T) {
	checkB(t, "5\t9\t2\t8", 4)
	checkB(t, "9\t4\t7\t3", 3)
	checkB(t, "3\t8\t6\t5", 2)
}

func checkB(t *testing.T, line string, checksum uint64) {
	Convey("Given a tab-separated string", t, func() {
		rowChecksumB, err := calcLineChecksumB(line)
		Convey("We should be able to parse with no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("And the checksumB should be division result", func() {
			So(rowChecksumB, ShouldEqual, checksum)
		})
	})
}

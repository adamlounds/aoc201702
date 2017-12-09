package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var totChecksum uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		rowChecksum, err := calcLineChecksum(line)
		if err != nil {
			//fmt.Printf("bad line: %s\n", err)
			break
		}

		totChecksum += rowChecksum
	}
	fmt.Printf("total checksum: %d", totChecksum)

	if scanner.Err() != nil {
		os.Stderr.WriteString(fmt.Sprintf("scan error %s", scanner.Err()))
	}
}

func calcLineChecksum(line string) (checksum uint64, err error) {

	nums, err := strToArray(line)
	if err != nil {
		return 0, err
	}

	var smallest = nums[0]
	var largest = nums[0]
	for _, num := range nums {
		if num < smallest {
			smallest = num
		}
		if num > largest {
			largest = num
		}
	}
	checksum = largest - smallest

	return checksum, nil
}

func strToArray(line string) (nums []uint64, err error) {
	cells := strings.Split(line, "\t")
	nums = make([]uint64, len(cells))

	for idx, digit := range cells {
		num, err := strconv.ParseUint(digit, 10, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("not a number [%s]", digit))
		}
		nums[idx] = num
	}
	return nums, nil
}

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

	var totChecksumA uint64 = 0
	var totChecksumB uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		rowChecksumA, err := calcLineChecksumA(line)
		if err != nil {
			fmt.Printf("bad line A: %s\n", err)
		}
		totChecksumA += rowChecksumA

		rowChecksumB, err := calcLineChecksumB(line)
		if err != nil {
			fmt.Printf("bad line B: %s\n", err)
		}

		totChecksumB += rowChecksumB
	}
	fmt.Printf("total checksumA: %d B: %d\n", totChecksumA, totChecksumB)

	if scanner.Err() != nil {
		os.Stderr.WriteString(fmt.Sprintf("scan error %s", scanner.Err()))
	}
}

func calcLineChecksumA(line string) (checksum uint64, err error) {

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

func calcLineChecksumB(line string) (checksum uint64, err error) {

	nums, err := strToArray(line)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(nums); i++ {
		var num = nums[i]
		for j := 0; j < len(nums); j++ {
			if i == j {
				// same value is not a candidate divisor :-)
				continue
			}
			var candidateDivisor = nums[j]
			if num%candidateDivisor == 0 {
				//fmt.Printf("woo, %d and %d divide cleanly\n", num, candidateDivisor)
				return uint64(num / candidateDivisor), nil
			}
		}
	}

	return 0, errors.New(fmt.Sprintf("cannot find clean divisor in %v", nums))
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

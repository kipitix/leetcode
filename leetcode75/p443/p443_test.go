package p443

import (
	"reflect"
	"strconv"
	"testing"
)

func compress(chars []byte) int {
	resLen := 0
	curChar := chars[0]
	charCount := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == curChar {
			charCount++
		} else {
			chars[resLen] = curChar
			resLen++
			if charCount > 1 {
				for _, n := range []byte(strconv.Itoa(charCount)) {
					chars[resLen] = n
					resLen++
				}
			}
			curChar = chars[i]
			charCount = 1
		}
	}
	chars[resLen] = curChar
	resLen++
	if charCount > 1 {
		for _, n := range []byte(strconv.Itoa(charCount)) {
			chars[resLen] = n
			resLen++
		}
	}

	return resLen
}

func TestCompress(t *testing.T) {
	res := compress([]byte("aabbccc"))
	if !reflect.DeepEqual(res, 6) {
		t.Error("failed 1")
	}

	res = compress([]byte("a"))
	if !reflect.DeepEqual(res, 1) {
		t.Error("failed 2")
	}

}

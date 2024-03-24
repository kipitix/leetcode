package p1071

import (
	"reflect"
	"testing"
)

func gcdOfStrings(str1 string, str2 string) string {
	bs1 := ([]byte)(str1)
	bs2 := ([]byte)(str2)
	l1 := len(bs1)
	l2 := len(bs2)
	var str []byte
	var sub []byte
	var maxLen, minLen int

	if l1 > l2 {
		str = bs1
		sub = bs2
		maxLen = l1
		minLen = l2
	} else {
		str = bs2
		sub = bs1
		maxLen = l2
		minLen = l1
	}

	for subLen := minLen; subLen > 0; subLen-- {
		curSub := (string)(sub[0:subLen])

		if (minLen%subLen != 0) || (maxLen%subLen != 0) {
			continue
		}

		entMax := 0
		for i := 0; i < maxLen-subLen+1; i += subLen {
			if (string)(str[i:i+subLen]) == curSub {
				entMax++
			}
		}

		entMin := 0
		for i := 0; i < minLen-subLen+1; i += subLen {
			if (string)(sub[i:i+subLen]) == curSub {
				entMin++
			}
		}

		if (entMax == maxLen/subLen) && (entMin == minLen/subLen) {
			return curSub
		}

	}

	return ""
}

func TestGcdOfStrings(t *testing.T) {
	res := gcdOfStrings("ABCABC", "ABC")
	if !reflect.DeepEqual(res, "ABC") {
		t.Error("failed 1")
	}

	res = gcdOfStrings("ABABAB", "AB")
	if !reflect.DeepEqual(res, "AB") {
		t.Error("failed 2")
	}

	res = gcdOfStrings("LEET", "CODE")
	if !reflect.DeepEqual(res, "") {
		t.Error("failed 3")
	}
}

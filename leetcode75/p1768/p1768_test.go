package p1768

import (
	"reflect"
	"testing"
)

func mergeAlternately(word1 string, word2 string) string {
	w1 := ([]rune)(word1)
	w2 := ([]rune)(word2)
	l1 := len(w1)
	l2 := len(w2)
	minLen := minMax(l1, l2)
	res := make([]rune, minLen*2, l1+l2)
	for i := 0; i < minLen; i++ {
		res[i*2] = w1[i]
		res[i*2+1] = w2[i]
	}

	if l1 > l2 {
		res = append(res, w1[minLen:]...)
	} else if l2 > l1 {
		res = append(res, w2[minLen:]...)
	}

	return (string)(res)
}

func minMax(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMerge(t *testing.T) {
	res := mergeAlternately("abc", "pqr")
	if !reflect.DeepEqual(res, "apbqcr") {
		t.Error("failed 1")
	}

	res = mergeAlternately("ab", "pqrs")
	if !reflect.DeepEqual(res, "apbqrs") {
		t.Error("failed 2")
	}

	res = mergeAlternately("abcd", "pq")
	if !reflect.DeepEqual(res, "apbqcd") {
		t.Error("failed 3")
	}
}

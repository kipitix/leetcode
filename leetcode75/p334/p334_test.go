package p238

import (
	"math"
	"reflect"
	"testing"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	i := math.MaxInt32
	j := math.MaxInt32

	for _, v := range nums {
		if v <= i {
			i = v
		} else if v <= j {
			j = v
		} else {
			return true
		}
	}

	return false
}

func TestIncreasingTriplet(t *testing.T) {
	res := increasingTriplet([]int{1, 2, 3})
	if !reflect.DeepEqual(res, true) {
		t.Error("failed 1")
	}

	res = increasingTriplet([]int{3, 2, 1})
	if !reflect.DeepEqual(res, false) {
		t.Error("failed 2")
	}

	res = increasingTriplet([]int{0, 1, 0, 1, 0, 1, 2})
	if !reflect.DeepEqual(res, true) {
		t.Error("failed 3")
	}

	res = increasingTriplet([]int{2, 4, -2, -3})
	if !reflect.DeepEqual(res, false) {
		t.Error("failed 4")
	}

	res = increasingTriplet([]int{20, 100, 10, 12, 5, 13})
	if !reflect.DeepEqual(res, true) {
		t.Error("failed 5")
	}

	res = increasingTriplet([]int{100, 100, 100, 1, 2, 3})
	if !reflect.DeepEqual(res, true) {
		t.Error("failed 6")
	}

	res = increasingTriplet([]int{1, 5, 0, 4, 1, 3})
	if !reflect.DeepEqual(res, true) {
		t.Error("failed 7")
	}

	src := make([]int, 500000)
	for i := 0; i < len(src); i += 2 {
		src[i] = 1
		src[i+1] = 2
	}
	res = increasingTriplet(src)
	if !reflect.DeepEqual(res, false) {
		t.Error("failed 8")
	}
}

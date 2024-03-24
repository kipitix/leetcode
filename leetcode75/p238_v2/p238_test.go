package p238

import (
	"reflect"
	"testing"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	pref := 1
	for i := range nums {
		res[i] = pref
		pref *= nums[i]
	}
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= post
		post *= nums[i]
	}
	return res
}

func TestProductExceptSelf(t *testing.T) {
	res := productExceptSelf([]int{1})
	if !reflect.DeepEqual(res, []int{1}) {
		t.Error("failed with 1")
	}

	res = productExceptSelf([]int{1, 2})
	if !reflect.DeepEqual(res, []int{2, 1}) {
		t.Error("failed with 2")
	}

	res = productExceptSelf([]int{1, 2, 3})
	if !reflect.DeepEqual(res, []int{6, 3, 2}) {
		t.Error("failed with 3")
	}

	res = productExceptSelf([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(res, []int{24, 12, 8, 6}) {
		t.Error("failed with 4")
	}

	res = productExceptSelf([]int{1, 2, 3, 4, 5, 6, 7})
	if !reflect.DeepEqual(res, []int{5040, 2520, 1680, 1260, 1008, 840, 720}) {
		t.Error("failed with 7")
	}
}

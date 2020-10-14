package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/upbit/LeetCode/commons"
)

type inputPair struct {
	Nums   []int
	Target int
}

// 两数之和 https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	offsets := make(map[int]int)
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
		expect := target - nums[i]
		if idx, ok := offsets[expect]; ok {
			if idx < i {
				return []int{idx, i}
			}
			return []int{i, idx}
		}
		expect = target - nums[j]
		if idx, ok := offsets[expect]; ok {
			if idx < i {
				return []int{idx, j}
			}
			return []int{j, idx}
		}
		offsets[nums[i]] = i
		offsets[nums[j]] = j
	}
	return []int{}
}

func Test_twoSum(t *testing.T) {
	for _, test := range []TestCase{
		NewCase("正常输入", inputPair{[]int{2, 7, 11, 15}, 9}, []int{0, 1}),
		NewCase("无序数组", inputPair{[]int{3, 2, 4}, 6}, []int{1, 2}),
		NewCase("负数输入", inputPair{[]int{-99, 100, 2, 4, 6}, 1}, []int{0, 1}),
	} {
		actual := twoSum(test.In.(inputPair).Nums, test.In.(inputPair).Target)
		assert.Equal(t, test.Expect.([]int), actual)
	}
}

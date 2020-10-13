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

func twoSum(nums []int, target int) (results []int) {
	results = append(results, 0)
	results = append(results, 1)
	return
}

func Test_twoSum(t *testing.T) {
	for _, test := range []TestCase{
		NewCase("正常输入", inputPair{[]int{2, 7, 11, 15}, 9}, []int{0, 1}),
	} {
		actual := twoSum(test.In.(inputPair).Nums, test.In.(inputPair).Target)
		assert.Equal(t, test.Expect.([]int), actual)
	}
}

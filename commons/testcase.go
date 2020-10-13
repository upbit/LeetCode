package commons

import "fmt"

// TestCase 测试用例定义
type TestCase struct {
	Name   string
	In     interface{}
	Expect interface{}
}

// NewCase 新建一个case
func NewCase(name string, in, expect interface{}) TestCase {
	return TestCase{Name: name, In: in, Expect: expect}
}

// String 输出成字符串
func (c *TestCase) String() string {
	return fmt.Sprintf("%s(%s): expect %s", c.Name, c.In, c.Expect)
}

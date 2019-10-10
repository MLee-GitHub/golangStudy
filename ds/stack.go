package ds

import (
	"fmt"
)

type stack []interface{} // 栈结构

func test(s *[]int){
	*s = append(*s, 11)
}

func test1(s []int){
	s[1]= 11
}

// 栈当前长度
func (s *stack) Len() int{
	return len(*s)
}

// 栈当前容量
func (s *stack) Cap() int {
	return cap(*s)
}

// 栈当前是否为空
func (s *stack) IsEmpty() bool{
	return s.Len() == 0
}

// 取当前栈顶元素
func (s *stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("empty stack, thx")
	}
	return (*s)[s.Len()-1], nil
}

// 元素压栈
func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

// 元素出栈
func (s *stack) Pop() (interface{}, error){
	if s.IsEmpty() {
		return nil, fmt.Errorf("empty stack, thx")
	}
	newStack := make([]interface{}, s.Len()-1)
	copy(newStack, (*s)[:s.Len()-1])
	v, err := s.Top()
	*s = newStack
	return v, err
}

package ds

import (
	"fmt"
	"testing"
)
var s stack

func TestStack_Len(t *testing.T) {
	fmt.Println(s.Len())
}
func TestStack_Push(t *testing.T) {
	s.Push(12)
	s.Push("aaa")
	s.Push(11)
	fmt.Println(s)
}
func TestStack_Top(t *testing.T) {
	if v, err := s.Top(); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(v)
	}
}

func TestStack_Pop(t *testing.T) {
	v, err := s.Pop()
	fmt.Println(s, v, err)
}

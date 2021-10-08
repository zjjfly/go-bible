package ch2

import (
	"fmt"
	"testing"
)

func TestAssignment(t *testing.T) {
	//一些操作返回两个值,其中第二个是一个布尔类型,表示操作是否成功
	//1.map查找
	m := make(map[string]int)
	v, ok := m["a"]
	fmt.Println(v)
	fmt.Println(ok)
	//2.类型断言
	var s interface{} = "hello"
	i, ok := s.(string)
	fmt.Println(i)
	fmt.Println(ok)
	//3.channel接收
	ch := make(chan string)
	ch <- "hello"
	msg, ok := <-ch
	fmt.Println(msg)
}

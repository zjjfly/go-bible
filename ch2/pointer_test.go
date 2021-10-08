package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {
	n := 1
	//声明一个指针
	p := &n
	//获取指针指向的内容
	assert.Equal(t, 1, *p)
	//修改指针的内容
	*p = 2
	assert.Equal(t, 2, *p)
	//指针的零值是nil
	var pp *int
	assert.Nil(t, pp)
	//指针直接是可以比较的,当它们都是nil或指向是同一个内存地址的时候就认为它们是相等的
	var x, y int
	assert.True(t, &x == &x)
	assert.False(t, &x == &y)
	assert.False(t, &x == nil)
}

func TestNew(t *testing.T) {
	//另一种创建变量的方法,new
	p:=new(int)
	assert.Equal(t, 0, *p)
	*p = 2
	assert.Equal(t, 2, *p)
	//每次调用new返回的是一个新的变量的地址
	q:=new(int)
	assert.NotEqual(t,p,q)
}


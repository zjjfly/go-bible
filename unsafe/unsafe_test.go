package unsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestSizeOf(t *testing.T) {
	//Sizeof返回传入的参数在内存中的大小,单位是字节
	t.Log(unsafe.Sizeof(0))
	//Sizeof返回的大小只包括固定的那部分,如字符串中的指针和长度部分
	assert.Equal(t, unsafe.Sizeof("1"), unsafe.Sizeof("12345"))
}

func TestAlignOf(t *testing.T) {
	//Alignof返回传入的参数的类型需要的内存对齐倍数
	t.Log(unsafe.Alignof(int32(1)))
	t.Log(unsafe.Alignof("1"))
}

func TestOffetOf(t *testing.T) {
	var x struct {
		a bool
		b int16
		c []int
	}
	//Offsetof必须传入一个字段,返回的是这个字段相对于这个这个struct的起始地址的偏移量
	t.Log(unsafe.Offsetof(x.c))
	t.Log(unsafe.Sizeof(x))
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

//unsafe.Pointer返回的是一个特殊的指针,可以通过这个指针对内存直接进行操作和地址计算
func TestPointer(t *testing.T) {
	t.Logf("%#16x\n",Float64bits(1.2) )
}

func TestDeepEqual(t *testing.T) {
	got := strings.Split("a:b:c",":")
	want := []string{"a","b","c"}
	assert.True(t,reflect.DeepEqual(got,want))

	//对于DeepEqual,空的slice和nil不相等,空的map和nil也不等
	assert.False(t,reflect.DeepEqual([]string{},nil))
	assert.False(t,reflect.DeepEqual(map[string]string{},nil))
}

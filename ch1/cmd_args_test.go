package ch1

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

//如何获取命令行参数
func TestPrintCmdArgs(t *testing.T) {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func TestPrintCmdArgs2(t *testing.T) {
	//strings.Join可以使用特定的分隔符拼接字符串
	fmt.Println(strings.Join(os.Args[1:], " "))
}

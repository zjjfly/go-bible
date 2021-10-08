package ch2

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestUtf8(t *testing.T) {
	//utf-8是为了更高效的编码unicode而产生的
	//由于unicode编码是固定的32bit的长度,而常用的unicode字符少于65536个,只需要16bit就可以了
	//都是有32bit会造成极大的空间浪费,为了是编码更紧凑,就出现了utf-8
	//utf-8编码的规则:
	//由1到4个字节构成一个字符的编码
	//如果第一个字节的最高位是0,则是单字节;其他情况下第一个字节的高位有多少个连续的1就表示这个字符有几个字节,
	//余下的字节的高位都已10开头

	//要得到一个字符串中有多少个unicode字符,需要使用utf8包的方法
	s := "Hello,世界"
	assert.Equal(t, 8, utf8.RuneCountInString(s))
	//获取真实的字符,需要解码器,unicode/utf8包中提供了这个功能
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	//range会自动解码utf8字符
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	//utf8字符串作为交换格式很方便,但在程序内部使用rune序列(unicode序列)更好,它大小一致,并且支持数组索引,方便切割
	s = "世界那么大"
	fmt.Printf("% x\n", s)
	//把字符串转换成rune序列
	r := []rune(s)
	fmt.Printf("%x\n", r)
	//把rune序列转成string
	assert.Equal(t, s, string(r))
	//可以使用string把整型转成包含对应Unicode码点的字符串
	assert.Equal(t, "A", string(65))
	assert.Equal(t, "京", string(0x4eac))
	//如果对应的码点无效,则会使用\uFFFD(�)来代替
	assert.Equal(t, "�", string(12345678))
}

func TestStrings(t *testing.T) {
	//raw string,不需要转义符,适合写正则表达式
	s := `\d\x\\`
	assert.Equal(t, "\\d\\x\\\\", s)
	//strings包中的一些实用方法
	assert.Equal(t, 2, strings.LastIndex("abbc", "b"))
	assert.True(t, strings.HasPrefix("abc", "a"))
	assert.True(t, strings.HasSuffix("abc", "c"))
	assert.True(t, strings.Contains("abc", "b"))
	assert.Equal(t, 2, strings.Index("abcde", "c"))
	assert.Equal(t, 2, strings.Count("abbccc", "b"))
	assert.Equal(t, "abc", strings.TrimSpace(" abc"))
	assert.Equal(t, "abc", strings.Trim("!!abc!!!", "!"))
	assert.Equal(t, "A Big Apple", strings.Title("a big apple"))
	assert.Equal(t, []string{"a", "big", "apple"}, strings.Fields("a big apple"))
	assert.Equal(t, "apple", strings.ToLower("Apple"))
	assert.Equal(t, "APPLE", strings.ToUpper("Apple"))
}

func TestBytes(t *testing.T) {
	//string可以和字节slice之间互相转换
	s := " abbc!!"
	b := []byte(s)
	b2 := []byte("a big apple")
	b3 := []byte("Apple")
	s2 := string(b)
	assert.Equal(t, s, s2)
	//bytes包中的函数和上面的strings包中的函数类似
	assert.Equal(t, 3, bytes.LastIndex(b, []byte{'b'}))
	assert.True(t, bytes.HasPrefix(b, []byte(" a")))
	assert.True(t, bytes.HasSuffix(b, []byte("!")))
	assert.True(t, bytes.Contains(b, []byte("bb")))
	assert.Equal(t, 2, bytes.Index(b, []byte("b")))
	assert.Equal(t, 2, bytes.Count(b, []byte("b")))
	assert.Equal(t, []byte("abbc!!"), bytes.TrimSpace(b))
	assert.Equal(t, []byte(" abbc"), bytes.Trim(b, "!"))
	assert.Equal(t, []byte("A Big Apple"), bytes.Title(b2))
	assert.Equal(t, [][]byte{[]byte("a"), []byte("big"), []byte("apple")}, bytes.Fields(b2))
	assert.Equal(t, []byte("apple"), bytes.ToLower(b3))
	assert.Equal(t, []byte("APPLE"), bytes.ToUpper(b3))
	//bytes.Buffer用于字节slice的缓存,类似java的StringBuilder
	//它还实现了Writer接口
	var buf bytes.Buffer
	//写入ASCII字符使用WriteByte,如果是单个字符的UTF8编码,则使用WriteRune
	runes := []rune("数组")
	buf.WriteRune(runes[0])
	buf.WriteRune(runes[1])
	buf.WriteByte('[')
	for i, v := range []int{1, 2, 3} {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	assert.Equal(t, "数组[1, 2, 3]", buf.String())
}

func TestComma(t *testing.T) {
	comma := Comma("1234567890")
	fmt.Println(comma)
}

func Comma(s string) string {
	var length = len(s)
	if length <= 3 {
		return s
	}
	var buf bytes.Buffer
	offset := length % 3
	if 0 == offset {
		offset = 3
	}
	delimiter:=""
	for ; len(s) > 0;
	{
		x := s[:offset]
		s = s[offset:]
		buf.WriteString(delimiter)
		buf.WriteString(x)
		offset = 3
		delimiter=","
	}
	return buf.String()
}

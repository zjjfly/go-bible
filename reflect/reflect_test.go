package reflect

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReflectType(t *testing.T) {
	//使用反射获取类型
	i := reflect.TypeOf(3)
	fmt.Println(i.String())
	fmt.Println(i)
	//反射获取到的类型是具体类型
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	//go的格式化字符串中有一个占位符%T用于打印参数类型
	fmt.Printf("%T\n", 3.1)

	//反射获取真实的值
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	//Value可以获取到Type
	tt := v.Type() // a reflect.Type
	fmt.Println(tt.String())
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

var Strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,
	Color:    false,
	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},
	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
}

func TestAny(t *testing.T) {
	println(Any(1))
	println(Any("2"))
	array := [...]string{"1", "2", "3"}
	println(Any(array))
	println(Any(array[1:]))
	println(Any(Strangelove))
}

func TestDisplay(t *testing.T) {
	Display("Strangelove", Strangelove)
	Display("array", [...]int{1, 2,})
	Display("stdout", os.Stdout)
}

func TestReflectValue(t *testing.T) {
	x := 2
	//reflect.ValueOf(x)返回的Value是不可取址的
	assert.False(t, reflect.ValueOf(x).CanAddr())
	//reflect.ValueOf(x).Elem()返回的Value都是可取址的
	assert.True(t, reflect.ValueOf(&x).Elem().CanAddr())
	//一种使用反射修改原值的方法:得到具体的值的指针并修改原值
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	assert.Equal(t, 3, x)
	//另一种方法是使用可取址的reflect.Value的Set方法直接修改
	d.Set(reflect.ValueOf(5))
	assert.Equal(t, 5, x)
	//要确保修改的变量可以接受对应的值
	//d.Set(reflect.ValueOf(int64(10)))

	//还有很多用于基本类型的Set方法,它们会在修改的时候尽量转成和原值匹配的类型
	d.SetInt(12)
	assert.Equal(t, 12, x)
	//对于一个引用interface{}类型的Value,无法使用SetInt,SetString这样的方法修改
	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	//ry.Set(reflect.ValueOf(3))
	ry.Set(reflect.ValueOf(1))
	assert.Equal(t, 1, y)
	ry.Set(reflect.ValueOf("hello"))
	assert.Equal(t, "hello", y)

	//利用反射可以获得结构体的未导出成员,但不能修改
	stdout := reflect.ValueOf(os.Stdout).Elem()
	t.Log(stdout.Type())
	pfd := stdout.FieldByName("pfd")
	t.Log(pfd.Type())
	fd := pfd.FieldByName("Sysfd")
	t.Log(fd.Int())
	//使用CanSet可以知道是否被修改
	assert.False(t, fd.CanSet())
}

//通过反射获取类型的成员
func TestField(t *testing.T) {
	var data struct {
		Label     []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
		Unit      int
	}
	v := reflect.ValueOf(&data).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		t.Log(name)
	}
}

func TestMethod(t *testing.T) {
	v := reflect.ValueOf(os.Stdout)
	tt := v.Type()
	t.Logf("type %s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		methodType := method
		t.Logf("func (%s) %s%s\n",tt,tt.Method(i).Name,strings.TrimPrefix(methodType.String(),"func"))
	}
}

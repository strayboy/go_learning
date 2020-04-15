package string_fun_test

import (
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' string 是不可变的byte slice
	s = "\xE4\xB8\xA5"
	t.Log(s, len(s))
	s = "中"
	t.Log(s, len(s))

	c := []rune(s)
	t.Log(c)
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)

	d := rune('你') // rune字符，int32
	e := byte('a') // byte字符，uint8
	t.Log(d, e)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	//for _, c := range s {
	//	t.Logf("%[1]c %[1]d", c)
	//}

	t.Logf("%T",s)

	for _, c := range s {
		t.Logf("%c %[1]d", c)
		if c == rune('人'){
			t.Logf("%c",c)
		}
	}
}

func TestStringLen(t *testing.T) {
	tip := "genji is a ninja"

	t.Log(len(tip))

	tip2 := "认真"
	t.Log(len(tip2))
	t.Logf("%T", tip2)
	t.Logf("%T", tip2[0])
	t.Log(tip2)
	t.Log(utf8.RuneCountInString(tip2))

	t3 := rune(tip2[0])
	t.Log(t3)
	t.Logf("%T", t3)

}

func TestStringTravel(t *testing.T) {
	str := "快乐 everyday"

	for _, s := range str {
		t.Logf("unicode: %c %d\n", s, s)
	}
	t.Log()
	for i := 0; i < len(str); i++ {
		t.Logf("ascii: %c %d\n", str[i], str[i])
	}
}

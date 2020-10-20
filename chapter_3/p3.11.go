package main

import (
	"bytes"
	"fmt"
)

func commonStronger(s string) string {
	var buf bytes.Buffer

	count := 0
	for i := len(s)-1; i >= 0; i-- {

		count++
		if s[i] == '.' {
			buf.WriteByte(s[i])
			count = 0
			continue
		}

		buf.WriteByte(s[i])

		if count % 3 == 0 && i != 0 && s[i-1] != '+' && s[i-1] != '-' {
			buf.WriteByte(',')
		}

	}
	return reverse(buf.String())
}

func reverse(s string) string {
	str := make([]rune,len(s))

	for i, v := range s {
		str[len(s) - i - 1] = v
	}

	return string(str)
}

func main() {
	fmt.Println(commonStronger("+.3456767"))
}

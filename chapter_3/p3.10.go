package main

import (
	"bytes"
	"fmt"
)

func common(s string) string {
	var buf bytes.Buffer

	for i := len(s)-1; i >= 0; i-- {
		if (len(s) - 1 - i) % 3 == 0 && i != len(s) - 1 {
			buf.WriteByte(',')
		}

		buf.WriteByte(s[i])
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
	fmt.Println(common("1234567"))
}

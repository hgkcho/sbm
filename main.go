package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ignoreWords = []string{
	"select",
	"from",
	"where",
	"join",
	"on",
	"as",
	"=",
	" ",
	"?",
	"+",
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	t := s.Text()
	ret := prs(t)

	fmt.Fprintln(os.Stdout, ret)
}

func surround(input string) string {
	var ret string
	// if input word is ingore word, just return input
	for _, v := range ignoreWords {
		if strings.EqualFold(v, input) {
			return input
		}
	}

	var commaBuf string
	if strings.HasSuffix(input, ",") {
		input = input[:len(input)-len(",")]
		commaBuf += ","
	}

	sp2 := strings.Split(input, ".")
	for i2, v2 := range sp2 {
		if i2 == 0 {
			ret += "`" + v2 + "`"
		} else {
			ret += ".`" + v2 + "`"
		}
	}
	return ret + commaBuf
}

func prs(input string) string {
	var ret string
	// sp := strings.Split(input, " ")
	sp := strings.Fields(input)
	for i, v := range sp {
		if v == "" {
			continue
		}
		var buf string
		buf = surround(v)

		if i == 0 {
			ret += buf
		} else {
			ret += " " + buf
		}
	}
	return ret
}

package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"go/token"
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
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ret string
	for sc.Scan() {
		t := sc.Text()
		ret += fmt.Sprintln(prs(t))
	}

	fmt.Fprintln(os.Stdout, ret)
}

func prs(input string) string {
	var ret string
	fset := token.NewFileSet()
	var s scanner.Scanner
	file := fset.AddFile("", fset.Base(), len(input))
	s.Init(file, []byte(input), nil, scanner.ScanComments)

LOOP:
	for {
		_, tok, lit := s.Scan()
		switch tok {
		case token.EOF:
			break LOOP
		case token.SEMICOLON:
			continue
		case token.ASSIGN:
			ret = ret + " " + token.ASSIGN.String()
		case token.DEFINE:
			ret = ret + " " + token.DEFINE.String()
		case token.ADD:
			ret = ret + " " + token.ADD.String()
		case token.STRING:
			// remove double quotation from lit
			lit = lit[1 : len(lit)-1]
			s := prsString(lit)
			lit = "\"" + s + "\""
			ret = ret + " " + lit
		default:
			ret = ret + " " + lit
		}
	}
	return ret
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

func prsString(input string) string {
	var ret string
	var hasSpace bool
	if strings.HasPrefix(input, " ") {
		hasSpace = true
	}
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
	if hasSpace {
		ret = " " + ret
	}
	return ret
}

func tmp() error {
	// var query = "select u.name from users as u" +
	// " join post as p join p.user_id = u.id "

	// const src = `var query = "select u.name from users as u"`
	const src = `
	var query = "select u.name from users as u" +
		" join post as p join p.user_id = u.id "
	`

	var ret string
	fset := token.NewFileSet()
	var s scanner.Scanner
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, []byte(src), nil, scanner.ScanComments)

LOOP:
	for {
		_, tok, lit := s.Scan()
		// ret = ret + " " + lit
		switch tok {
		case token.EOF:
			break LOOP
		case token.ASSIGN:
			ret = ret + " " + token.ASSIGN.String()
		case token.DEFINE:
			ret = ret + " " + token.DEFINE.String()
		case token.ADD:
			ret = ret + " " + token.ADD.String()
		case token.STRING:
			tok.IsLiteral()
			// remove double quotation from lit
			lit = lit[1 : len(lit)-1]
			s := surround(lit)
			lit = "\"" + s + "\""
			ret = ret + " " + lit
		default:
			ret = ret + " " + lit
		}
	}

	fmt.Println(ret)
	return nil
}

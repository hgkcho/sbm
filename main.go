package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"strings"
)

var (
	version = "0.0.2"
	commit  = "none"
	date    = "unknown"
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
	var v bool
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		fmt.Println("version: ", version)
		os.Exit(0)
	}

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
			break LOOP
			// continue
		case token.ASSIGN:
			ret = ret + " " + token.ASSIGN.String()
		case token.DEFINE:
			ret = ret + " " + token.DEFINE.String()
		case token.ADD:
			ret = ret + " " + token.ADD.String()
		case token.STRING:
			// lit is string that is removed double quotation from original lit
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

// surround enclose input string with backquotes if it is not already enclosed
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
	for i, v := range sp2 {
		if i == 0 {
			ret += surroundWithBackQuote(v)
		} else {
			ret += "." + surroundWithBackQuote(v)
		}
	}
	return ret + commaBuf
}

func surroundWithBackQuote(input string) string {
	var ret string
	if !strings.HasPrefix(input, "`") {
		ret += "`"
	}
	ret += input
	if !strings.HasSuffix(input, "`") {
		ret += "`"
	}
	return ret
}

// prsString parse input string into escaped SQL
func prsString(input string) string {
	var ret string
	var hasSpace bool
	if strings.HasPrefix(input, " ") {
		hasSpace = true
	}
	// sp substring of input divided by white space
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

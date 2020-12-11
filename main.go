package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s '<EXPRESSION>'\n", os.Args[0])
		os.Exit(1)
	}

	// Get a mathematical expression from the first argument
	input := flag.Arg(0)

	// Calculate the expression
	output, err := calc(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print the calculation result to standard output
	if output != nil {
		fmt.Println(output)
	}
}

func calc(input string) (output interface{}, err error) {
	// Tokenize the input
	tokens, err := tokenize(input)
	if err != nil {
		return nil, err
	}

	// Parse the token sequence
	nd, err := parse(tokens)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Input: %q =>\nNode:\n", input)
	// printNode(nd, 0)

	// Evaluate the tree
	return eval(nd)
}

type tokenKind string

const (
	tokenNum   tokenKind = "NUM"   // Number
	tokenPunct tokenKind = "PUNCT" // Punctuator
	tokenEOF   tokenKind = "EOF"   // End of file
)

type token struct {
	kind tokenKind
	lit  string // Literal
	val  int    // Integer value; Used if kind == tokenNum
}

func tokenize(input string) (tokens []token, err error) {
	for pos := 0; pos < len(input); pos++ {
		c := rune(input[pos])

		// Skip whitespaces
		if unicode.IsSpace(c) {
			continue
		}

		// Tokenize an integer literal
		if unicode.IsDigit(c) {
			start := pos
			for pos+1 < len(input) && unicode.IsDigit(rune(input[pos+1])) {
				pos++
			}
			// 123456789
			// ^       ^
			// start   pos
			s := input[start : pos+1]
			// s == "123456789"
			v, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, token{kind: tokenNum, lit: s, val: v})
			continue
		}

		if unicode.IsPunct(c) || c == '+' {
			tokens = append(tokens, token{kind: tokenPunct, lit: string(c)})
			continue
		}
	}

	tokens = append(tokens, token{kind: tokenEOF})

	return tokens, nil
}

type nodeKind string

const (
	nodeNum nodeKind = "NUM"
	nodeAdd nodeKind = "ADD"
	nodeSub nodeKind = "SUB"
	nodeMul nodeKind = "MUL"
	nodeDiv nodeKind = "DIV"
)

type node struct {
	kind nodeKind

	// Number node
	val int // Integer value; used if kind == nodeNum

	// Binary operator node
	left, right *node
}

func printNode(nd *node, level int) {
	if nd == nil {
		return
	}

	prefix := ""
	if level > 0 {
		prefix += strings.Repeat(" ", (level-1)*4) + "--> "
	}

	if nd.kind == nodeNum {
		fmt.Printf("%s(%v %v)\n", prefix, nd.kind, nd.val)
	} else {
		fmt.Printf("%s(%v)\n", prefix, nd.kind)
	}

	printNode(nd.left, level+1)
	printNode(nd.right, level+1)
}

func parse(tokens []token) (nd *node, err error) {
	nd, _, err = expr(tokens)
	return nd, err
}

// expr = product ("+" product | "-" product)*
func expr(tokens []token) (nd *node, rest []token, err error) {
	nd, rest, err = product(tokens)
	if err != nil {
		return nil, nil, err
	}

	for {
		var right *node

		switch rest[0].lit {
		case "+":
			right, rest, err = product(rest[1:])
			if err != nil {
				return nil, nil, err
			}

			nd = &node{kind: nodeAdd, left: nd, right: right}
		case "-":
			right, rest, err = product(rest[1:])
			if err != nil {
				return nil, nil, err
			}

			nd = &node{kind: nodeSub, left: nd, right: right}
		default:
			return nd, rest, nil
		}
	}
}

// product = num ("*" num | "/" num)*
// num = ("0"..."9")+
func product(tokens []token) (nd *node, rest []token, err error) {
	nd = &node{kind: nodeNum, val: tokens[0].val}
	rest = tokens[1:]

	for {
		var right *node

		switch rest[0].lit {
		case "*":
			right = &node{kind: nodeNum, val: rest[1].val}
			nd = &node{kind: nodeMul, left: nd, right: right}
			rest = rest[2:]
		case "/":
			right = &node{kind: nodeNum, val: rest[1].val}
			nd = &node{kind: nodeDiv, left: nd, right: right}
			rest = rest[2:]
		default:
			return nd, rest, nil
		}
	}
}

func eval(nd *node) (output interface{}, err error) {
	if nd.kind == nodeNum {
		return nd.val, nil
	}

	l, err := eval(nd.left)
	if err != nil {
		return nil, err
	}
	left := l.(int)

	r, err := eval(nd.right)
	if err != nil {
		return nil, err
	}
	right := r.(int)

	switch nd.kind {
	case nodeAdd:
		return left + right, nil
	case nodeSub:
		return left - right, nil
	case nodeMul:
		return left * right, nil
	case nodeDiv:
		return left / right, nil
	default:
		return nil, fmt.Errorf("unknown node kind: %s", nd.kind)
	}
}

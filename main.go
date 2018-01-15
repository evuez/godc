package main

import (
	"fmt"
	"github.com/prataprc/goparsec"
	dcVm "godc/vm"
	"os"
)

var vm = dcVm.NewVM()

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	parse([]byte(os.Args[1]))
}

func parse(dc []byte) {
	ast := parsec.NewAST("dc", 1024)
	scanner := parsec.NewScanner([]byte(dc))

	parser := newParser(ast)

	node, scanner := ast.Parsewith(parser, scanner)
	_, scanner = scanner.SkipWS()

	for node != nil {
		node, scanner = ast.Parsewith(parser, scanner)
		_, scanner = scanner.SkipWS()
	}
}

func newParser(ast *parsec.AST) parsec.Parser {
	var dc parsec.Parser

	register := parsec.TokenExact("[A-Za-z]", "REG") // Should be 256 registers, not 52

	pop := ast.And("pop", pop, parsec.AtomExact("s", "POP"), register)
	cond := ast.And("cond", cond, parsec.TokenExact("[<>]", "OP"), register)
	duplicate := ast.Many("duplicate", duplicate, parsec.AtomExact("d", "DUP"))
	eval := ast.Many("eval", eval, parsec.AtomExact("x", "EVL"))
	print := ast.Many("print", print, parsec.AtomExact("p", "PRT"))
	value := ast.Many("value", value, parsec.TokenExact("[0-9]+", "VAL"))
	operator := ast.Many("operator", operator, parsec.TokenExact("[-+*/]", "OP"))
	string_literal := ast.And(
		"string_literal",
		string_literal,
		parsec.AtomExact("[", "OT"),
		parsec.TokenExact("[A-Za-z0-9-<+]+", "ANY"),
		parsec.AtomExact("]", "CT"),
	)

	dc = ast.OrdChoice("dc", nil, string_literal, duplicate, pop, print, eval, cond, value, operator)

	return dc
}

func value(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	vm.Stack.Push(node.GetValue())
	return node
}

func string_literal(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	vm.Stack.Push(node.GetValue())
	return node
}

func duplicate(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	top, _ := vm.Stack.Top()
	vm.Stack.Push(top)
	return node
}

func eval(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	value, _ := vm.Stack.Pop()

	parse([]byte(value[1 : len(value)-1]))

	return node
}

func operator(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	v1, _ := vm.Stack.PopInt()
	v2, _ := vm.Stack.PopInt()

	switch node.GetValue() {
	case "-":
		vm.Stack.Push(v2 - v1)
	case "+":
		vm.Stack.Push(v2 + v1)
	case "*":
		vm.Stack.Push(v2 * v1)
	case "/":
		vm.Stack.Push(v2 / v1)
	}
	return node
}

func cond(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	value := node.GetValue()
	operator := string(value[0])
	register := string(value[1])

	var shouldEval bool

	v1, _ := vm.Stack.PopInt()
	v2, _ := vm.Stack.PopInt()

	switch operator {
	case "<":
		shouldEval = v1 < v2
	case ">":
		shouldEval = v1 > v2
	}

	if shouldEval {
		macro := vm.Registers[register]

		parse([]byte(macro[1 : len(macro)-1]))
	}

	return node
}

func print(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	top, _ := vm.Stack.Top()
	fmt.Println(top)
	return node
}

func pop(n string, s parsec.Scanner, node parsec.Queryable) parsec.Queryable {
	top, _ := vm.Stack.Pop()
	vm.Registers[node.GetValue()[1:]] = top
	return node
}

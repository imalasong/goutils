package eval

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type stringHelper struct {
	scan  scanner.Scanner
	token rune
}

func (s *stringHelper) next() {
	s.token = s.scan.Scan()
}

func (s *stringHelper) text() string {
	return s.scan.TokenText()
}

func (s *stringHelper) describe() string {
	switch s.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", s.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", s.text())
	}
	return fmt.Sprintf("%q", rune(s.token)) // any other rune
}

type typePanic string

func Parse(exprStr string) (expreR Expr, errR error) {

	defer func() {
		switch e := recover().(type) {
		case nil:
		case typePanic:
			expreR = nil
			errR = fmt.Errorf("%s", e)
		default:
			panic(e)
		}
	}()

	if exprStr == "" {
		return nil, errors.New("express not allow empty")
	}
	sh := new(stringHelper)
	sh.scan.Init(strings.NewReader(exprStr))
	sh.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	sh.next()

	//解析
	expr := parseExpr(sh)
	if sh.token != scanner.EOF {
		panic(typePanic("unexpected " + sh.describe()))
	}
	return expr, nil
}

func parseExpr(sh *stringHelper) Expr {
	return parseBinary(sh, 1)
}

func parseBinary(sh *stringHelper, level int) Expr {
	e := parseUnary(sh)
	for lc := precedence(sh.token); lc >= level; lc-- {
		for precedence(sh.token) == lc {
			op := sh.token
			sh.next()
			y := parseBinary(sh, lc+1)
			e = binary{op: op, x: e, y: y}
		}
	}
	return e
}

func precedence(token rune) int {
	switch token {
	case '%':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func parseUnary(sh *stringHelper) Expr {
	op := sh.token
	if op == '-' || op == '+' {
		sh.next()
		return unary{op: op, x: parseUnary(sh)}
	}
	return parseElement(sh)
}

func parseElement(sh *stringHelper) Expr {
	switch sh.token {
	case scanner.Int | scanner.Float:
		float, err := strconv.ParseFloat(sh.text(), 64)
		if err != nil {
			panic(typePanic(err.Error()))
		}
		sh.next()
		return literal(float)
	case scanner.Ident:
		id := sh.text()
		sh.next()
		if '(' != sh.token {
			return Var(id)
		}
		sh.next()
		var args []Expr
		if sh.token != ')' {
			for {
				args = append(args, parseExpr(sh))
				if sh.token != ',' {
					break
				}
				sh.next()
			}
			if sh.token != ')' {
				msg := fmt.Sprintf("end of %s function missing ')'", id)
				panic(typePanic(msg))
			}
		}
		sh.next() //consumer )
		return call{fn: id, args: args}
	case '(':
		sh.next()
		plan := parseExpr(sh)
		if sh.token != ')' {
			panic(typePanic("lack of ')'"))
		}
		sh.next() //consumer )
		return plan
	}
	msg := fmt.Sprintf("unexpected %s", sh.describe())
	panic(typePanic(msg))
}

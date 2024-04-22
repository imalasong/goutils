// Package eval provides an expression evaluator.
package eval

import (
	"bytes"
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	//return fmt.Sprintf("%.5f", l)
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%v %s %v", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(c.fn)
	buffer.WriteByte('(')

	for i, e := range c.args {
		buffer.WriteString(e.String())
		if i < len(c.args)-1 {
			buffer.WriteByte(',')
		}
	}

	buffer.WriteByte(')')
	return buffer.String()
}

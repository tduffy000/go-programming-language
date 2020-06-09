// Exercise 7.13
package eval

import (
	"bytes"
	"fmt"
	"reflect"
)

func (v Var) String() string {
	return reflect.ValueOf(v).String()
}

func (l literal) String() string {
	return fmt.Sprintf("%v", reflect.ValueOf(l).Float())
}

func (u unary) String() string {
	return fmt.Sprintf("%c%v", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%v%c%v", b.x, b.op, b.y)
}

func (c call) String() string {
	var b bytes.Buffer
	b.Write([]byte(c.fn))
	b.WriteRune('(')
	for i, arg := range c.args {
		b.Write([]byte(arg.String()))
		if i < len(c.args)-1 {
			b.WriteRune(',')
		}
	}
	b.WriteRune(')')
	return b.String()
}

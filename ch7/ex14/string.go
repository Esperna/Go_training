package eval

import (
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', 0, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	str := c.fn + "("
	for i := 0; i < len(c.args); i++ {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprintf("%s", c.args[i])
	}
	str += ")"
	return str
}

func (vf variadicfunc) String() string {
	str := vf.fn + "("
	for i := 0; i < len(vf.args); i++ {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprintf("%s", vf.args[i])
	}
	str += ")"
	return str
}

package eval

import (
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.Itoa(int(float64(l)))
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	str := c.fn
	str += "("
	for i := 0; i < len(c.args); i++ {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprintf("%s", c.args[i])
	}
	str += ")"
	return str
}

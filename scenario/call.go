package scenario

import (
	"fmt"
	"strings"
)

type Call struct {
	args         []interface{}
	returnValues []interface{}
	description  string
}

func (c *Call) Returns(values ...interface{}) *Call {
	c.returnValues = values
	return c
}

func (c *Call) ReturnsNothing() *Call {
	return c.Returns()
}

func (c *Call) Describe(description string) *Call {
	c.description = description
	return c
}

func (_ *Call) formatSlice(s []interface{}) string {
	argsAsString := make([]string, 0, len(s))
	for _, a := range s {
		argsAsString = append(argsAsString, fmt.Sprintf("`%v`", a))
	}

	return strings.Join(argsAsString, ", ")
}

func (c *Call) String() string {
	argsAsString := make([]string, 0, len(c.args))
	for _, a := range c.args {
		argsAsString = append(argsAsString, fmt.Sprint(a))
	}
	var returnPart, descriptionPart string
	if len(c.returnValues) > 0 {
		returnPart = fmt.Sprintf("->(%v)", c.formatSlice(c.returnValues))
	}
	if c.description != "" {
		descriptionPart = fmt.Sprintf(" //%v", c.description)
	}
	return fmt.Sprintf(
		"Call[(%v)%v%v]",
		c.formatSlice(c.args),
		returnPart,
		descriptionPart,
	)
}

func (c *Call) Equal(o *Call) bool {
	if len(c.args) != len(o.args) ||
		len(c.returnValues) != len(o.returnValues) {
		return false
	}

	for i := range c.args {
		if c.args[i] != o.args[i] {
			return false
		}
	}

	for i := range c.returnValues {
		if c.returnValues[i] != o.returnValues[i] {
			return false
		}
	}

	return true
}

func NewCall(args ...interface{}) *Call {
	return &Call{args: args}
}

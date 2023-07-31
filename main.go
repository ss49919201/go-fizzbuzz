package main

import (
	"fmt"
	"strconv"
)

type replacer interface {
	apply(carry string, n int) string
	match(carry string, n int) bool
}

type cyclicNumber struct {
	base        int
	replacement string
}

func newCyclicNumber(base int, replacement string) *cyclicNumber {
	return &cyclicNumber{
		base:        base,
		replacement: replacement,
	}
}

func (c *cyclicNumber) apply(carry string, _ int) string {
	return carry + c.replacement
}

func (c *cyclicNumber) match(carry string, n int) bool {
	return n%c.base == 0
}

type passThrough struct{}

func (p *passThrough) apply(carry string, n int) string {
	return strconv.Itoa(n)
}

func (p *passThrough) match(carry string, n int) bool {
	return carry == ""
}

type converter struct {
	replacerList []replacer
}

func newConverter(replacerList ...replacer) *converter {
	return &converter{
		replacerList: replacerList,
	}
}

func (c *converter) do(n int) string {
	var result string
	for _, v := range c.replacerList {
		if v.match(result, n) {
			result = v.apply(result, n)
		}
	}
	return result
}

func main() {
	c := newConverter(
		newCyclicNumber(3, "Fizz"),
		newCyclicNumber(5, "Buzz"),
		new(passThrough),
	)
	fmt.Println(c.do(3))
	fmt.Println(c.do(5))
	fmt.Println(c.do(15))
	fmt.Println(c.do(22))
}

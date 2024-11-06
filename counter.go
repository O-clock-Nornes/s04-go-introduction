package main

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value = c.Value + 1
}

func incr(c *Counter) {
	c.Value = c.Value + 1
}

package main

type Clock struct{ h, m int }

func (c Clock) Adjust() Clock {
	totalmin := c.h*60 + c.m

	c.m = totalmin % 60
	c.h = (totalmin / 60) % 24

	if c.m < 0 {
		c.m = 60 + c.m
		c.h--
	}

	if c.h < 0 {
		c.h += 24
	}

	return c
}

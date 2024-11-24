package utils

type Cursor struct {
	index  int
	length int
}

func NewCursor(index, length int) Cursor {
	return Cursor{
		index:  index,
		length: length,
	}
}

// Get current index
func (c Cursor) Index() int {
	return c.index
}

// Increment cursor position
func (c *Cursor) Next() {
	if c.index < c.length-1 {
		c.index++
	} else if c.index >= c.length-1 {
		c.index = 0
	}
}

// Decrement cursor position
func (c *Cursor) Previous() {
	if c.index > 0 {
		c.index--
	} else if c.index <= 0 {
		// Go to bottom of list
		c.index = c.length - 1
	}
}

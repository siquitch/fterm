package utils

// A menu navigation object
type Navigator struct {
	index    int
	length   int
	showHelp bool
}

const CursorChar = ">"

func NewNavigator(index, length int) Navigator {
	return Navigator{
		index:    index,
		length:   length,
		showHelp: false,
	}
}

// Toggle help message
func (c *Navigator) ToggleHelp() {
	c.showHelp = !c.showHelp
}

// Whether to show help
func (c *Navigator) ShouldShowHelp() bool {
	return c.showHelp
}

// Get current index
func (c *Navigator) Index() int {
	return c.index
}

// Increment cursor position
func (c *Navigator) Next() {
	if c.index < c.length-1 {
		c.index++
	} else if c.index >= c.length-1 {
		c.index = 0
	}
}

// Reset with new length. 0 keeps current length
func (c *Navigator) Reset(length int) {
	c.index = 0
	if length > 0 {
		c.length = length
	}
}

// Decrement cursor position
func (c *Navigator) Previous() {
	if c.index > 0 {
		c.index--
	} else if c.index <= 0 {
		// Go to bottom of list
		c.index = c.length - 1
	}
}

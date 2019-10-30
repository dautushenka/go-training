package practice2

import (
	"fmt"
	"errors"
)

type Collection struct {
	elements []Element
	index int
};

func (c *Collection) Add(e Element) {
	c.elements = append(c.elements, e);
}

func (c *Collection) validateIndex(index int) error {
	if (index < 0 || index >= len(c.elements)) {
		return errors.New("test");
	}
	return nil
}

func (c *Collection) Get(index int) (Element, error) {
	if e := c.validateIndex(index); e != nil {
		return Element(0), e;
	}
		
	return c.elements[index], nil;
}

func (c *Collection) Remove(index int) {
	if e := c.validateIndex(index); e == nil {
		copy(c.elements[index:], c.elements[index + 1:]);
		c.elements = c.elements[:len(c.elements) - 1]
	}
}

func (c Collection) First()  Element {
	elm, _ := c.Get(0);
	
	return elm;
}

func (c Collection) Last()  Element {
	elm, _ := c.Get(len(c.elements) - 1);
	
	return elm;
}

func (c Collection) Length() int {
	return len(c.elements);
}

func (c Collection) Print() {
	fmt.Println(c.elements);
}

func (c *Collection) Value() Element {
	elm, _ := c.Get(c.index);
	
	return elm;
}

func (c *Collection) Next() Element {
	c.index += 1;
	if error := c.validateIndex(c.index); error != nil {
		c.index = 0;
	}

	return c.Value();
}

func (c *Collection) Prev() Element {
	c.index -= 1;
	if error := c.validateIndex(c.index); error != nil {
		c.index = len(c.elements) - 1;
	}

	return c.Value();
}
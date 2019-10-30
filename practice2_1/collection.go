package practice2_1

import (
	"fmt"
	"errors"
)

type Collection struct {
	first *Element
	last *Element
	current *Element
	length int
};

func (c *Collection) getNodeByIndex(index int) (*Element, error) {
	if (index < 0 || index >= c.length) {
		return &Element{}, errors.New("Index is out of collection range");
	}

	elm := c.first;
	for i := 0; i < c.length; i++ {
		if (i == index) {
			return elm, nil;
		}
		elm = elm.next;
	}

	panic("Error in finding node by index in collection");
}

func (c *Collection) Add(value int) {
	elm := &Element{value: value}
	if (c.last == nil) {
		c.first = elm;
		c.last = elm;
		c.current = elm;
	} else {
		elm.prev = c.last;
		c.last.next = elm;
		c.last = elm;
	}
	c.length++;
}

func (c *Collection) Get(index int) (int, error) {
	elm, error := c.getNodeByIndex(index);
	if (error != nil) {
		return 0, error;
	}

	return elm.value, nil;
}

func (c *Collection) Remove(index int) error {
	elm, error := c.getNodeByIndex(index);
	if (error != nil) {
		return error;
	}
	elm.next.prev = elm.prev;
	elm.prev.next = elm.next;
	
	if (c.first == elm) {
		c.first = elm.next;
	}
	if (c.last == elm) {
		c.last = elm.prev;
	}
	if (c.current == elm) {
		if (elm.next == nil) {
			c.current = c.first;
		} else {
			c.current = elm.next;
		}
	}

	c.length--;
	return nil;
}

func (c *Collection) First() (int, error) {
	if (c.first == nil) {
		return 0, errors.New("Collection is empty");
	}
	return c.first.value, nil;
}

func (c *Collection) Last()  (int, error) {
	if (c.last == nil) {
		return 0, errors.New("Collection is empty");
	}
	return c.last.value, nil;
}

func (c *Collection) Length() int {
	return c.length;
}

func (c *Collection) Print() {
	current := c.first;
	if (current == nil) {
		fmt.Println("Collection is empty");
	}

	fmt.Print("[");
	fmt.Printf("%d", current.value);
	for current.next != nil {
		current = current.next;
		fmt.Printf(" %d", current.value);
	}
	fmt.Println("]");
}

func (c *Collection) Value() (int, error) {
	if (c.current == nil) {
		return 0, errors.New("Collection is empty");
	}
	return c.current.value, nil;
}

func (c *Collection) Next() (int, error) {
	if (c.current == nil) {
		return 0, errors.New("Collection is empty");
	}
	if (c.current.next == nil) {
		return 0, errors.New("This is a last element in the collection");
	}

	c.current = c.current.next;
	return c.current.value, nil;
}

func (c *Collection) Reset() {
	c.current = c.first;
}

func (c *Collection) Prev() (int, error) {
	if (c.current == nil) {
		return 0, errors.New("Collection is empty");
	}
	if (c.current.prev == nil) {
		return 0, errors.New("This is a fist element in the collection");
	}

	c.current = c.current.prev;
	return c.current.value, nil;
}
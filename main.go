package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}
type linkedList struct {
	head *node
	len  int
}

func main() {
	l := linkedList{}
	l.add(5)
	l.add(6)
	l.add(7)
	fmt.Println(*(l.get(5)))
	l.remove(5)
	fmt.Println(l)
}
func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value
	if l.head == nil {
		l.head = newNode
	} else {
		itrator := l.head
		for ; itrator.next != nil; itrator = itrator.next {

		}
		itrator.next = newNode
	}
}

func (l *linkedList) remove(value int) {
	var previous *node
	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if previous != nil {
				previous.next = current.next
				return
			} else {
				l.head = current.next
			}
		}
		previous = current
	}
}
func (l linkedList) get(value int) *node {
	for itrator := l.head; itrator != nil; itrator = itrator.next {
		if itrator.value == value {
			return itrator
		}
	}
	return nil
}
func (l linkedList) String() string {
	sb := strings.Builder{}
	for itrator := l.head; itrator != nil; itrator = itrator.next {
		sb.WriteString(fmt.Sprintf("%d ", itrator.value))
	}
	return sb.String()
}

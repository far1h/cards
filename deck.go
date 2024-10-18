package main

import "fmt"

type deck []string

// set methods on variables
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

package main

import "fmt"

type Oecophylla struct{}

func (o *Oecophylla) sting() {
	fmt.Println("Oecophylla can't sting, because it doesn't have a stinger!")
}

func (o *Oecophylla) describe() {
	fmt.Println("Oecophylla ants live in Asia, Australia and Africa. Australian ones have green abdomens and heads.")
}

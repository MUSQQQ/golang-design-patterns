package main

type AbsurdismCheckout struct {
	Checkout
}

func (a *AbsurdismCheckout) AddBook() {
	a.ChosenBook = &AbsurdistBook{
		Author: "Albert Camus",
		Title:  "Stranger",
	}
}

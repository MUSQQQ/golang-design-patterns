package main

type MangaCheckout struct {
	Checkout
}

func (m *MangaCheckout) AddBook() {
	m.ChosenBook = &Manga{
		Author: "Tatsuki Fujimoto",
		Title:  "Chainsaw Man",
	}
}

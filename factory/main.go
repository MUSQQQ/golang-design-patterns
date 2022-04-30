package main

import "fmt"

type Book interface {
	GetAuthor() string
	GetTitle() string
}

type Checkout struct {
	ChosenBook Book
}

func (c *Checkout) IsEmpty() bool {
	return c.ChosenBook == nil
}

func (c *Checkout) BookInfo() {
	fmt.Println(c.ChosenBook.GetAuthor())
	fmt.Println(c.ChosenBook.GetTitle())
}

func main() {
	checkout1 := MangaCheckout{}
	checkout2 := AbsurdismCheckout{}

	fmt.Println("checkout 1 (manga)")
	fmt.Println("is checkout empty:")
	fmt.Println(checkout1.IsEmpty())
	checkout1.AddBook()
	fmt.Println("book added")
	fmt.Println("is checkout empty:")
	fmt.Println(checkout1.IsEmpty())
	checkout1.BookInfo()

	fmt.Println("checkout 2 (absurdism)")
	fmt.Println("is checkout empty:")
	fmt.Println(checkout2.IsEmpty())
	checkout2.AddBook()
	fmt.Println("book added")
	fmt.Println("is checkout empty:")
	fmt.Println(checkout2.IsEmpty())
	checkout2.BookInfo()
}

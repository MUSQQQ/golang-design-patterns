package main

type AbsurdistBook struct {
	Title  string
	Author string
}

func (a *AbsurdistBook) GetAuthor() string {
	return a.Author
}

func (a *AbsurdistBook) GetTitle() string {
	return a.Title
}

package main

type Manga struct {
	Title  string
	Author string
}

func (m *Manga) GetAuthor() string {
	return m.Author
}

func (m *Manga) GetTitle() string {
	return m.Title
}

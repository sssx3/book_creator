package main

import (
	"flag"
	"fmt"
	markdown "gitlab.com/golang-commonmark/markdown"
	_ "io/ioutil"
	"os"
)

var (
	pagesDir = flag.String("pgs", "", "bcreator -pgs [Path to pages dir]")
	output   = flag.String("o", "book.html", "bcreator -o [Path to output the book]")
)

func main() {
	flag.Parse()
	// delete
	fmt.Println(os.Args)
	fmt.Println(*pagesDir, *output)
	fmt.Println(flag.Args())
}

type Page struct {
	number uint
	title  string
	source string
}

func (p *Page) generate() {
	// TODO(sssx3)
	fmt.Println("Generated")
}

type Navigation struct {
	pages []Page
}

func (n *Navigation) generate() {
	// TODO(sssx3)
	fmt.Println("Generated")
}

type Book struct {
	name       string
	author     string
	navigation Navigation
	pages      []Page
	pagesLen   int
}

func (b *Book) generate() {
	// TODO(sssx3)
	fmt.Println("Generated")
}

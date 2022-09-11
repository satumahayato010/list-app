package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type BookList struct {
	Books []string
}

func New(books []string) *BookList {
	return &BookList{Books: books}
}

func fileRead(fileName string) []string {
	var bookList []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	defer file.Close()
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		bookList = append(bookList, scaner.Text())
	}
	return bookList
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	bookList := fileRead("reading.txt")
	html, err := template.ParseFiles("view.html")
	if err != nil {
		log.Fatal(err)
	}
	getBooks := New(bookList)
	if err := html.Execute(w, getBooks); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/view", viewHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

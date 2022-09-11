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

func createHandler(w http.ResponseWriter, r *http.Request) {
	formValue := r.FormValue("value")
	file, err := os.OpenFile("reading.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(file, formValue)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/view", http.StatusFound)
}

func main() {
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/view/create", createHandler)
	// 追加
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

package main

import  (
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"strings"
)


type Page struct{
	Title string
	Body []byte
}


func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	// change .txt to file type you wish to open
    filename := title + ".txt"
    body, err := ioutil.ReadFile( filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func processinghttp(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage(strings.TrimLeft( r.URL.Path , "/"))
	if p != nil {
		// page found
		io.WriteString(w, "Content: " + string(p.Body) + "\n")
	} else {
		//404 page not found
		  io.WriteString(w, string("  ") +"\n")	
	}
}

func main() {
	fmt.Println("starting server ...")
	http.HandleFunc("/", processinghttp)
	http.ListenAndServe(":8080", nil)
	
	
}

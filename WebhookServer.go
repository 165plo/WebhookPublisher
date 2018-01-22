package main

import(
	"net/http"
	"fmt"
	"html"
	"log"
	"io/ioutil"
)

func main(){
	http.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request){
		content, err := ioutil.ReadFile("Senders.txt")
		if(err != nil){
			log.Fatal(err)
		}
		s := string(content)
		fmt.Fprintf(w, "%q", s)
	})

	log.Fatal(http.ListenAndServe(":222", nil))
}
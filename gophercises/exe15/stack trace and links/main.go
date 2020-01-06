package main

import (
	"net/url"
	"strings"
	"github.com/alecthomas/chroma/quick"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"bytes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {

	//?path=/home/gslab/Documents/Go/gophercises/exe15/highlight/trace/main.go
	
	path := r.FormValue("path")
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = quick.Highlight(w, b.String(), "go", "html", "monokai")
	io.Copy(w, file)
}

func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}


// goroutine 10 [running]:
// runtime/debug.Stack(0xc4202bbb48, 0x1, 0x1)
// 	/usr/lib/go-1.10/src/runtime/debug/stack.go:24 +0xa7
// main.devMw.func1.1(0xaa3960, 0xc4205761c0)
// 	/home/gslab/Documents/Go/gophercises/exe15/highlight/main.go:50 +0xac
// panic(0x8719c0, 0xa9a290)
// 	/usr/lib/go-1.10/src/runtime/panic.go:502 +0x229
// main.funcThatPanics()
// 	/home/gslab/Documents/Go/gophercises/exe15/highlight/main.go:70 +0x39
// main.panicDemo(0xaa3960, 0xc4205761c0, 0xc4205a0000)
// 	/home/gslab/Documents/Go/gophercises/exe15/highlight/main.go:61 +0x20
// net/http.HandlerFunc.ServeHTTP(0x9b5a50, 0xaa3960, 0xc4205761c0, 0xc4205a0000)
// 	/usr/lib/go-1.10/src/net/http/server.go:1947 +0x44
// net/http.(*ServeMux).ServeHTTP(0xc420252f00, 0xaa3960, 0xc4205761c0, 0xc4205a0000)
// 	/usr/lib/go-1.10/src/net/http/server.go:2340 +0x130
// main.devMw.func1(0xaa3960, 0xc4205761c0, 0xc4205a0000)
// 	/home/gslab/Documents/Go/gophercises/exe15/highlight/main.go:56 +0x95
// net/http.HandlerFunc.ServeHTTP(0xc4200662a0, 0xaa3960, 0xc4205761c0, 0xc4205a0000)
// 	/usr/lib/go-1.10/src/net/http/server.go:1947 +0x44
// net/http.serverHandler.ServeHTTP(0xc420456dd0, 0xaa3960, 0xc4205761c0, 0xc4205a0000)
// 	/usr/lib/go-1.10/src/net/http/server.go:2697 +0xbc
// net/http.(*conn).serve(0xc4200dcaa0, 0xaa3ca0, 0xc4204a3040)
// 	/usr/lib/go-1.10/src/net/http/server.go:1830 +0x651
// created by net/http.(*Server).Serve
// 	/usr/lib/go-1.10/src/net/http/server.go:2798 +0x27b


func makeLinks(stack string) string{
	lines := strings.Split(stack, "\n")
	for li, line := range lines{
		if len(line) == 0 || line[0] != '\t'{
			continue
		}
		file := ""
		for i, ch := range line{
			if ch == ':'{
				file = line[1:i]
				break
			}
		}
		//for encode the url 
		v := url.Values{}
		v.Set("path", file)
		lines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" +file+ "</a>" + line[len(file)+1:]
	}
	return strings.Join(lines, "\n")
}
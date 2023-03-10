package main
import "fmt"
import "net/http"

func helloHandler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "Hello, world!")
}

func main() {
 fmt.Println("Launching server...")

 hh := http.HandlerFunc(helloHandler)
 http.Handle("/hello", hh)
//  http.Handle("/helloo", hh)

 fs := http.FileServer(http.Dir("."))
 http.Handle("/", http.StripPrefix("/", fs))
 http.ListenAndServe(":12004", nil) // set no handler, but there's a default handler.
}
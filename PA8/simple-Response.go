package main
import "fmt"
import "net/http"

func main() {
 fmt.Println("Launching server...")
 http.ListenAndServe(":12004", http.FileServer(http.Dir(".")))
}
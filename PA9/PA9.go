package main
import "fmt"
import "net/http"
import "strings"
import "os"

func sPrefix(prefix string, h http.Handler) http.Handler {
	if (prefix==""){
		return h
	}
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		p:=strings.TrimPrefix(r.URL.Path,prefix)

		if _,err := os.Stat(p); !os.IsNotExist(err){
			
			h.ServeHTTP(w,r)
		}else{
			fmt.Fprintln(w,"File not found")
		}
	})
}

func main() {
 fmt.Println("Launching server...")
 fs := http.FileServer(http.Dir("."))
 http.Handle("/", sPrefix("/", fs))
 http.ListenAndServeTLS(":12004", "server.cer", "server.key", nil)
}
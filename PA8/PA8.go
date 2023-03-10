package main
import "fmt"
import "net/http"
import "strings"
import "os"

func sPrefix(prefix string, h http.Handler) http.Handler { //http.Handler為輸出格式
	if (prefix==""){
		return h
	}

	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){ //return這個Func 但型態為http.Handler
		//包在http.HandlerFunc裏面的會吃到http.responseWriter 以及*http.Request

		p:=strings.TrimPrefix(r.URL.Path, prefix) //URL刪掉prefix

		if _,err := os.Stat(p); !os.IsNotExist(err){ //傳回檔案內容
			h.ServeHTTP(w,r) //讀取以及傳回URL的內容
		}else{
			fmt.Fprintln(w,"File not found") //找不到
		}
	})
}


func main() {
 fmt.Println("Launching server...")
 fs := http.FileServer(http.Dir(".")) //在當前資料夾找尋的 http.handler
 http.Handle("/", sPrefix("/", fs)) //fs為http.handler, prefix 為 "/"
 
 http.ListenAndServe(":12004", nil) 
}
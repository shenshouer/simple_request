package main
import (
	"net/http"
	"log"
)

func main() {
	log.SetFlags(log.Flags()|log.Lshortfile)
	startHttpServer()
}

func startHttpServer() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", defaultHandler)
	log.Println("Starting http server on port 30001!")
	http.ListenAndServe(":30001", serverMux)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("========> 请求来源:",r.RemoteAddr)
	w.WriteHeader(200)
	w.Write([]byte(r.RemoteAddr))
}
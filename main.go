package main
import (
	"net/http"
	"log"
	"encoding/json"
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

	sourceip := r.RemoteAddr
	url := r.RequestURI
	referer := r.Referer()
	host := r.Host

	data := r.Header
	data.Add("sourceip", sourceip)
	data.Add("requestUrl", url)
	data.Add("referer", referer)
	data.Add("host", host)

	byteData , err := json.Marshal(data)
	if err != nil{
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}else{
		w.WriteHeader(200)
		w.Write(byteData)
	}
}
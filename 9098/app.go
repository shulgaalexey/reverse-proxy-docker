package main

import (
        "net/http"
	"fmt"
)

func main() {

	const port = "9098"

	println("Start listening on port: " + port + "...")


        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
                println("======>", port, req.URL.String())
		title := "Http Server on port " + port
		body := "This is a demo of http server, run on port " + port
		body += "</p><button type='button'  onclick='myClick()'>C'mon</button>"
		body += "<script>function myClick() { alert('wow') }</script>"
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, body)
        })
        http.ListenAndServe(":"+port, nil)
}

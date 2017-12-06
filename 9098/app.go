package main

import (
        "net/http"
)

func main() {

	const port = "9098"

	println("Start listening on port: " + port + "...")


        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
                println("===>", port, req.URL.String())
        })
        http.ListenAndServe(":"+port, nil)
}

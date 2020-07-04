package main

import(
	"net/http"
	"github.com/fly0c8/consensus"
)

func main() {
	consensus.Config(3)
	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}



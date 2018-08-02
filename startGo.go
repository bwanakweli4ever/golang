package main

import (
	"net/http"
	"fmt"
)

func main(){


	http.HandleFunc("/",makeRouting)
	http.ListenAndServe("8080",nil)

}

func makeRouting(wReq http.ResponseWriter,req *http.Request){


	fmt.Fprint(wReq,"Hello World of Devs")

}
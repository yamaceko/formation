package main

import (
	"fmt"
	"net/http"
	 "github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReadBufferSIze:1024,
	WriteBufferSize:1024,
	CheckOrigin: func(r *http.Request)bool {return true},
	
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":4000",nil)

}

func handler(w http.ResponseWriter, r *http.Request){
	socket,err:=upgrader.Upgrader(w,r,nil)
}
package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){

	resp,err := http.Get("http://google.com")
	if err!=nil{
		println("Error:", err)
		os.Exit(1)
	}else{
		defer resp.Body.Close()
        contents, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))

	}


}
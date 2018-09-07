package main

import (
	"fmt"
)

func main(){


	colors:=map[string]string{
		"red":"#ff0000",
		"green":"#4bf45",
		"white":"#fffff",
	}

	printMap(colors)
	
}

func printMap(c map[string]string){
	for color,hex:=range c{
		println("hex code for", color, "is ", hex)
	}
}
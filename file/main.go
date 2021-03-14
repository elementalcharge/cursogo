package main

import (
	"fmt"
	"io/ioutil"
)

func main (){
	data, err := ioutil.ReadFile("leeme.txt")

	if err !=nil {
		panic(err.Error())
	}

	fmt.Println(string(data))

}
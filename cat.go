package main

import "fmt"

func cats(s string,ss string)string{
	var a string;
	a=s+ss;
	return a
}


func main(){
	fmt.Printf("%s\n",cats("hello ","world"))



}
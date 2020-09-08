package main

import "fmt"

func main(){
	var i int = 0 
	var a int = 0
	for  a = 0 ; a < 49 ; a++  {
		fmt.Printf("=%d\n",i);
		i=i+8;
	}
}
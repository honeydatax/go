package main

import "fmt"

func main(){
	var i [50] int 
	var a int = 0
	for a = 0 ; a < 49 ; a++ {
		i[a]=a*8
	}

	for  a = 0 ; a < 49 ; a++  {
		fmt.Printf("=%d\n",i[a]);
	}
}
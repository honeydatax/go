package main

import "fmt"

func bar(ss string ,ii int) string{
	var s string=""
	var i int = 0
	for i=0;i<ii;i++{
		s=s+ss;
	}
	return s
}

func barsize(ss string ,sss string,ii int) string{
	var s string=""
	var i int = 0
	var sz int = 0
	s=ss
	sz=ii-len(ss)
	for i=0;i<sz;i++{
		s=s+sss;
	}
	return s
}


func main(){
	var a int = 0
	for  a = 0 ; a < 16 ; a++  {
		fmt.Printf("%s |\n",barsize(bar(".",a)," ",20))
	}
}
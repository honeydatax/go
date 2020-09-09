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


func center(s string,size int)string{
	var i int =0
	var s1 string=""
	i=(size/2)-(len(s)/2)
	s1=bar(" ",i)
	s1=s1+s
	return s1
}

func barrigth(ss string ,sss string,ii int) string{
	var s string=""
	var i int = 0
	var sz int = 0
	s=ss
	sz=ii-len(ss)
	for i=0;i<sz;i++{
		s=sss+s;
	}
	return s
}

func drawtop(ss string ,ii int,iii int){
	fmt.Printf("%s%s\n",bar(ss,ii),bar("-",iii))
}
func drawbox(ss string ,sss string,ii int,iii int){
	fmt.Printf("%s|%s|\n",bar(sss,ii),barsize(ss,sss,iii-2))
}

func drawleft(ss string ,sss string,ii int,iii int){
	fmt.Printf("%s|%s|\n",bar(sss,ii),barrigth(ss,sss,iii-2))
}


func main(){
	var a int = 0
	drawtop(" ",5,35)
	for  a = 0 ; a < 16 ; a++  {

		drawleft(bar("*",a+1)," ",5,35)	

	}
	drawtop(" ",5,35)
}
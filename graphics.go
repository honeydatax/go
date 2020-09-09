package main

import "fmt"

func ggraphics(s string,ss string,lines int,size int){
	var i int=0;
	var ii int=0;
	var count=0
	for i=0;i<lines;i++{
		for ii=0;ii<count;ii++{
			fmt.Printf("%s",s);
		}
		count=count+size

		for ii=0;ii<size;ii++{
			fmt.Printf("%s",ss);
		}
		fmt.Printf("\n");
	}

}

func main(){
		ggraphics(" ","o",16,4)	



}
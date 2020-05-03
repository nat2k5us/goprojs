package main

import ( 
	"fmt"
	sw "./somewhere"	
)

func main(){
	sm := sw.Something()
	fmt.Println(sm)

	fmt.Println(sw.Nothing())
}
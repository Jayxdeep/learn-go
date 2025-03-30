package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	reader:=bufio.NewReader(os.Stdin)//to read the user input
	fmt.Println("Enter the rating")
	//comma ok syntax or comma err syntax basically its kind of try and catch blcks
	inp,_:=reader.ReadString('\n')
	fmt.Println("Thanks fr rating,",inp)
	fmt.Printf("Type of this rating is %T \n",inp)

	read:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter name:")
	inpt,_:=read.ReadString('\n')
	fmt.Println("Hello",inpt)
}
package main

import (
	"bufio"
	"strings"
	// "errors" removed it cause in go when the imported is not used it has to be removed or used
	"fmt"
	"os"
	"strconv"//string converstion
)
func main(){
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate our app")

	reader:=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	fmt.Println("Thanks for rating",input)

	numrating,err:=strconv.ParseFloat (strings.TrimSpace(input),64)//converts int to float so the trimspace is used to remove whitespace characters include " ",\n,/t,/r[escape sequence moves cursor to begin of current line] 4\r\n they are trailing characters so need to remove them to add +1 in the input
	
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 in rating",numrating+1)
	}

}
//concatenation
package main
import "fmt"
func main(){
	a:="hello"
	b:="World"
	c:=(a+" "+b)//concatenation of a and b
	fmt.Println(c)//print concatenated string
	
	var num1 string="20"
	var num2 string="21"
	fmt.Println(num1+num2)//concatenation of string and int

	//length of string
	txt:="hello world"
	fmt.Println(len(txt))//length of string

	//indexin 
	word:="Programming"//i need the char o and r then its starts form 0 
	fmt.Println(string(word[2]))
	fmt.Println(string(word[4]))

	
}
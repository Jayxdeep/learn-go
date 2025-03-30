package main
import "fmt"
func main(){
	num1:=1
	num2:=2
	fmt.Println(num1+num2)//simple addition of two numbers

	// num3:=3
	// num4:=4
	// fmt.Println("num3+num4")// gives error 
	num3:="4"
	num4:="5"
	fmt.Println(num3+num4)//returns 45 because num3 and num4 are strings 45 op
	num5:="2 "
	num6:="4"
	fmt.Println(num5+num6)//returns 24 because num5 and num6 are strings  2 4 op

	a:="12345"
	fmt.Println(string(a[1]))//returns 2 because a[1] is 2 and its index starts from 0
}
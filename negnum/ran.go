package main
import "fmt"
func main(){
	a:=-50
	b:=40
	sum:=a+b//sum of a and b
	prod:=a*b//product of a and b
	quo:=a/b//quotient of a and b
	fmt.Println("Sum:",sum)//print sum
	fmt.Println("Product:",prod)//print product
	fmt.Println("Quotient:",quo)//print quotient
	x:=10
	y:=9
	res:=float32(x)/ float32(y)//float division
	fmt.Println("Result:",res)//print result
}
package main
import "fmt"
var a int
var(
	b bool
	c float32
	d string
)
func main(){
	a=42
	b,c=true,32.43
	d="jaydeep"
	fmt.Println(a,b,c,d)
}
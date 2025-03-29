package main
import "fmt"
func main(){
	var length int=11
	var width int =13
	
	fmt.Println("Current area:",length*width)//current area
	fmt.Println("Current perimeter:",2*length+width)//perimeter cuurent
	width=15//changed value
	fmt.Println("New area:",length*width)//new area
	fmt.Println("New perimeter:",2*length+width)//new perimeter	
	//declare variables
	Num:=19//short declaration
	fmt.Print(Num)//print value of Num
	a:=34
	b:=45
	c:=a+b
	fmt.Print((c))
	// fmt.Print((a+b))
	//float32 using
	var pi float32 = 3.14
	var radius float32 = 5.0
	area:=pi*radius*radius//area of circle	
	fmt.Print(area)//print area of circle 78.5
	

}
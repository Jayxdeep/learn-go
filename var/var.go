// package main
// import "fmt"
// var a int
// var(//grp of variables
// 	b bool
// 	c float32
// 	d string
// )
// func main(){
// 	a=42
// 	b,c=true,32.43
// 	d="jaydeep"
// 	fmt.Println(a,b,c,d)
// }
package main
import "fmt"
const LoginToken string="damnnnn"//public variable


func main(){
	var name string="Jaydeep"
	fmt.Println(name)
	fmt.Printf("variable is of type: %T \n",name)//%T is used to print the type of variable
	var isStudent bool=true
	fmt.Println(isStudent)
	fmt.Printf("variable is of type: %T \n",isStudent)//%T is used to print the type of variable

	var smallval uint8=255 //256 is not possible because it is out of range of uint8
	fmt.Println(smallval)
	fmt.Printf("variable is of type: %T \n",smallval)//%T is used to print the type of variable

	var smallfloat float32=255.455555555
	fmt.Println(smallfloat)//takes upto 5 places after decimal
	var smallfloat1 float64=255.455334234
	fmt.Println(smallfloat1)//takes upto 15 places after decimal more precision
	fmt.Printf("variable is of type: %T \n",smallfloat)//%T is used to print the type of variable

	//alias default values
	var anotherval int
	fmt.Println(anotherval)//default value of int is 0 thers no garbage value in go
	fmt.Printf("variable is of type: %T \n",anotherval)//%T is used to print the type of variable)

	var anotherString string
	fmt.Println(anotherString)//default value of string is "" empty string
	fmt.Printf("variable is of type: %T \n",anotherString)//%T is used to print the type of variable

	var anotherbool bool
	fmt.Println(anotherbool)//default value of bool is false

	var anotherfloat float32
	fmt.Println(anotherfloat)//default value of float is 0.00

	//implicit type conversion[due to lexar rules] := cant used outisde the function only inside func/method 
	fmt.Println(LoginToken)//calling the public variable LoginToken 
	fmt.Printf("variable is of type: %T \n",LoginToken)//%T is used to print the type of variable
	
}
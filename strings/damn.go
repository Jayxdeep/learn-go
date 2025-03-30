package main
import "fmt"
func main(){
	//use of rune
	// rune is a data type in Go that represents a Unicode code point. It is an alias for int32 and is used to represent characters in a string.
	word:="ocygen"
	words:=[]rune(word)
	words[1]='x'
	word=string(words)// The code converts the string word to a slice of runes, modifies the second rune to 'x', and then converts it back to a string.
	fmt.Println(word)//prints oxygen
}
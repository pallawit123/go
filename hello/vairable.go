// variable Declection
// redeclarations and shadowed variables
// visible variables
// naming conventions
// type of conversions

// printing statements: fmt.Println, fmt.Printf, fmt.Sprintf
// printing variables: fmt.Printf("%T", variable)
// printing constants: fmt.Printf("%T", constant)

// Print statements:
// package main 
// import(
//     "fmt"
// )
// func main() {
//     fmt.Printf("hello World!")
// } 



// varibale declarations:
// Declaring varibale itself

// package main

// import (
//     "fmt"
// )
// func main() {
//     var i int
//     i = 20
//     fmt.Println(i)
// }
// the above code can be decleared as follow
// package main
// import (
//     "fmt"
// )
// func main() {
//     var i int = 20
//     fmt.Println(i)
// }

// package main

// import (
//     "fmt"
// )

// func main() {
//     i := 20
//     fmt.Println(i)
// }
// package main

// import (
//     "fmt"
// )

// func main() {
//     // var i int = 0
//     var j float64 = 27.3 // Correct type declaration
//     k := 99              // Variable names should start with a lowercase letter
//     fmt.Printf("%v, %T\n", j, j)
//     fmt.Println(k)       // Use the variable to avoid unused variable error
// }


// Declearing variables in package level
// While declearing the variable outside the function scope := is not allowed to use for this we should actually write full declearation sy
// package main
// import(
//     "fmt"
// )
// var i int = 20
// func main() {
//     fmt.Printf("%v, %T",i,i)
// }


// In package level block of a variable can be decleared together
// package main
// import(
//     "fmt"
// )
// var(
//     codewriter_name="coder"
//      codewriter_age=20
// )
// var (
//     counter int =0
//     counter_name="counter"
// )
// func main(){   
//      fmt.Printf("%v, %T",codewriter_name,codewriter_name)

// }

// if you declare the variable outside the scope and inside the function it will be overwrite but if you write the two value in same scope for same variable it will be an error you cannot declare/assign two variables in same function in same scope

// package main
// import(
//     "fmt"
// )
// var i int = 20
// func main(){
//     var i int =42
//     fmt.Println(i)
// }


// package main
// import(
//     "fmt"
// )
// var i int = 20
// func main(){
//     fmt.Println(i)
//     var i int =42
//     fmt.Println(i)
// } 


// Naming conventions Rule
// 1 The first is that lenght of a variable name should reflect the lifeline of a varibale.
//  eg If i use i it should represent the small life line
//  if i use longer name like seasonalflu it will represent the long life line

// 2. All acronym should be in uppercase
// eg. var URL string = "https://www.google.com"



// converting the variable type
// package main
// import(
// 	"fmt"
// )
// func main() {
// 	var i int = 30
// 	fmt.Printf("%v, %T\n", i, i)
// 	var j float64
// 	j = float64(i)
// 	fmt.Printf("%v, %T\n", j, j)
// }

// string conversion function
// package main
// import(
// 	"fmt"
// 	"strconv"
// )
// func main() {
// 	var i int = 30
// 	fmt.Printf("%v, %T\n", i, i)
// 	var j string
// 	j = strconv.Itoa(i)
// 	fmt.Printf("%v, %T\n", j, j)
// }

// using the complex function also change the varibale in complex format. 
// package main
// import(	
// 	"fmt"
// )
// func main() {
// 	var n complex128 = complex(1,2)
// 	fmt.Printf("%v,%T\n", n, n)
// 	}

// text types
// there are two types of text types in go language
// 1. string
// 2. rune

// string
// package main
// import(
// 	"fmt"
// )
// func main() {	
// 	s := "hi my name is ......."
// 	fmt.Printf("%v,%T\n", s, s)
// }

// string can be treated like an arry
// package main
// import(
// 	"fmt"
// )
// func main() {	
// 	s := "hi my name is ......."
// 	fmt.Printf("%v,%T\n", string(s[1]), s[1])
// }

// string can perform string concatenation. This is done by using the + operator. This is the only operator that can be used with strings. 
// package main
// import(
// 	"fmt"
// )
// func main() {	
// 	s1 := "hi my name is ......."
// 	s2 := "Pallawit timalsina"
// 	fmt.Printf("%v,%T\n", s1+s2, s1+s2)
// }

// It can be conver into slice of byte
// package main
// import(
// 	"fmt"
// )
// func main() {	
// 	s := "hi my name is ......."
// 	b  :=	[]byte(s)
// 	fmt.Printf("%v,%T\n", b, b)
// }	

// rune: if we use single quote in string it will be rune. Otherwise it will be string. This is the only difference between string and rune. The rune is the alias of int32. 
// package main
// import(
// 	"fmt"
// )
// func main() {	
// 	 r :='a'
// 	fmt.Printf("%v,%T\n", r,r)
// }
// Information that can be stored in go
// 1> Boolean type
// 2> Numeric types
    // .Int,
	//  .Float, 
	// .Complex
// 3> Text types


// Boolean data: It is the 


// Boolean

// package main
// import (
// 	"fmt"
// )
// func main(){
// 	var n bool = true
// 	fmt.Printf("%v,%T\n",n,n)
// }


// package main
// import (
// 	"fmt"
// )
// func main(){
// 	var n bool =  false
// 	fmt.Printf("%v,%T\n",n,n)
// }


// package main
// import (
// 	"fmt"
// )
// func main(){
// 	n :=1==1
// 	m :=1==2
// 	fmt.Printf("%v,%T\n",n,n)
// 	fmt.Printf("%v,%T",m,m)	
// }

// if we donot initialize the value variable then it takes the default value as false. This is the behavior of the go language. the value it will take will be 0. This is the default value of the boolean type in go language. 
// package main
// import (
// 	"fmt"
// )	
// func main(){
// 	var n bool
// 	fmt.Printf("%v,%T\n",n,n)
// }

// Numeric types
// it is little bit complex in go language beacuse we are working with different numeric types
// package main
// import (
// 	"fmt"
// )	
// 	func main(){
// 		n :=42
// 		fmt.Printf("%v,%T\n",n,n)
// 	}



// package main
// import (
// 	"fmt"
// )	
// 	func main(){
// 		var n uint16 =42
// 		fmt.Printf("%v,%T\n",n,n)
// 	}

// Note that in go language you cannot add integer and float values. You have to convert the integer value to float value and then add them.
// Note that as well we cannot add two intger type i.e adding int and int8 value. This is not possible in go language. to add them we have to convert them to the same type and then add them.


// package main
// import (
// 	"fmt"

// )	
// 	func main(){
// 		a := 1
// 		b :=42
// 		fmt.Println(a+b)
// 		fmt.Println(a-b)
// 		fmt.Println(a*b)
// 		fmt.Println(a/b)
// 		fmt.Println(a%b)

// 	}

// type conversion
// package main
// import (
//     "fmt"
// )	
// func main() {
//     a := 1
//     b := 42.0
//     fmt.Println(a + int(b))
//     fmt.Println(float64(a) - float64(b)) // Corrected to use float64
//     fmt.Println(a * int(b))
//     fmt.Println(a / int(b)) // Added missing closing parenthesis
//     fmt.Println(a % int(b))
// }


// Bit operators
// package main
// import (
//     "fmt"
// )	
// func main() {
//     a := 10
//     b := 3
// 	fmt.Println(a &b) // 1010 & 0011 = 0010
// 	fmt.Println(a | b) // 1010 | 0011 = 1011
// 	fmt.Println(a ^ b) // 1010 ^ 0011 = 1001
// 	fmt.Println(a &^ b) // 1010 &^ 0011 = 0100

// }

// Bit shifting
// package main
// import (
//     "fmt"
// )	
// func main() {
   
// 	a := 10 // 1010
// 	fmt.Println(a << 3) // 10*2^3 = 10*8 = 80
// 	fmt.Println(a >> 3) // 10/2^3 = 10/8 = 1
// }


// float type
// It stores decimal points values.
// It follows the IEEE 754 standard. there are two type of standard in IEEE 754 standard. The first one is 16 bit and the second one is 32 bit.
// in folating value we cannot directly use exponent notation. We have to use the float64 type to use the exponent notation.
// package main
// import (
//     "fmt"
// )	
// func main() {
//     n := 1.23456789
// 	n =13.7e72
// 	n = 2.1E14
// 	fmt.Printf("%v,%T\n",n,n)
// }

// Airthmetic operation in float64
// Reminder operator is not allowed in float64 type. only used in integer type.
// bitwise operator or bit shift is not allowed in float64 type. only used in integer type.
// package main
// import (
//     "fmt"
// )	
// func main() {
//     a := 1.00
//     b := 42.000
//     fmt.Println(a + b)
//     fmt.Println(a -b)
//     fmt.Println(a * b)
//     fmt.Println(a / b) 
// }


// Complex type
// there are two type of complex number complex 64 and complex 128
// package main
// import (
// 	"fmt" // fmt is the package that is used to print the output on the console. It is a built-in package in Go. 
// )
// 	func main(){
// 		var n complex64 = 1+2i
// 		fmt.Printf("%v,%T\n",n,n) // %v is used to print the value of the variable and %T is used to print the type of the variable.
// 	}


// package main
// import (
// 	"fmt" // fmt is the package that is used to print the output on the console. It is a built-in package in Go. 
// )
// 	func main(){
// 		var n complex64 = 2i
// 		fmt.Printf("%v,%T\n",n,n) // %v is used to print the value of the variable and %T is used to print the type of the variable.
// 	}


// Airthmetic Operation
// package main
// import(
// 	"fmt" 
// )
// 	func main() {
// 		a :=1+2i
// 		b :=2+3i
// 		fmt.Println(a+b)
// 		fmt.Println(a-b)
// 		fmt.Println(a*b)
// 		fmt.Println(a/b)
// 	}

// in complex numbers if we want to have a real number and imaginary number then we have to use the real and imag function which is built-in function in go language.

package main
import(
	"fmt"
)
	func main() {
		a := 1 + 2i
		fmt.Println(real(a))
		fmt.Println(imag(a))
	}
	
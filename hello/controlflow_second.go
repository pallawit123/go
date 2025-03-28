// Defer 
// panic
// recover



// defer function
// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	fmt.Println("First line")
// 	defer fmt.Println("Second line")
// 	fmt.Println("Third line")
	
// }
// defer function is used to delay the execution of a function until the surrounding function returns.
// defer execute any function that are passed into it after the function finished it's final stament but before it actually returns.

// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	defer fmt.Println("First line")
// 	defer fmt.Println("Second line")
// 	defer fmt.Println("Third line")
	
// }

// if we use the defer function in all the statements it will execute the last statement first and the first statement last.
// defer work on LIFO manner. Last in first out.

// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	fmt.Println("First line")
// 	fmt.Println("Second line")
// 	fmt.Println("Third line")
	
// }


// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	a:= "start"
// 	defer fmt.Println(a)
// 	a = "end"
	
// }


// Panic function

// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	a, b := 1 ,0
// 	ans := a / b
// 	fmt.Println("The answer is: ", ans)
	
// }
// panic function is used to stop the execution of the program and print the error message. It is used to handle the run time error. It will stop the execution of the program and print the error message. It will not execute any statement after the panic statement.
// It will not execute the defer statement. It will not execute the return statement. It will not execute the main function.

// package main
// import (
//     "fmt"
// )

// func main() {
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("Recovered from panic:", r)
//         }
//     }()

//     fmt.Println("First line")
//     panic("Second line")
//     fmt.Println("Third line") // This will not execute unless panic is recovered.
// }


// using defer and panic on the same code
// package main
// import (
//     "fmt"
// )

// func main() {
//     fmt.Printf("starting the program\n")
//     defer fmt.Printf("Programme is deffered\n")
//     panic("Something went wrong")
//     fmt.Println("end") // This will not execute because of the panic.
// }


package main
import (
    "fmt"
    
)

func main() {
	fmt.Printf("starting the program\n")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}() //invoking function
	panic("Something went wrong")
	fmt.Println("end") // This will not execute because of the panic.
}
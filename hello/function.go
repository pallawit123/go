// package main
// import (
// 	"fmt"
// 	)

// func main() {
// 	sayMessage("Hello go!")

// }
// func sayMessage(msg string) {
// 	fmt.Println(msg)

// }

// we can pass multiple parameters to a function that returns multiple values but it should be in the same order as the function signature
// package main
// import (
// 	"fmt"
// 	)

// func main() {
// 	for  i := 0; i < 5; i++ {
// 		sayMessage("Hello go!",i)
// 	}

// }
// func sayMessage(msg string, index int) {
// 	fmt.Println(msg)
// 	fmt.Println("the value of the string is:",index)
// }




// passing multiple parameters of same type
//package main
// import (
// 	"fmt"
// 	)

// func main() {
// 	sayGreeting("Hello" ,"go!")

// }
// func sayGreeting(message string, name string) {   // on this same type we can pass multiple parameters so we have to use different name for the parameters
// 	// we can also use the same name for the parameters but it will not be a good practice.
// 	fmt.Println(message , name)
// }
// while passing the same parameter we should not have to declare the parameter twice we can use comma and a parameter name at last

//  passing value in a function through a pointer

// package main
// import (
// 	"fmt"
// 	)

// func main() {
// 	greeting := "Hello go!"
// 	name := "go!"
// 	saygreeting(&greeting, &name)
// 	fmt.Println(name)

// }
// func saygreeting(greeting ,name *string) {
// 	fmt.Println(*greeting, *name)
// 	// we can also use the same name for the parameters but it will not be a good practice.
// 	// fmt.Println(greeting , name)
// 	*name = "Goings"
// 	fmt.Println(*name)
// }

// VARIADIC PARAMETERS: Name of the parameter should be in the last and we can pass multiple values to the function

// package main
// import (
//     "fmt"
//     )
// func main() {
//     sum(1,2,3,4,5)
//     }
// func sum(values ...int) {   
// 	fmt.Println(values)
// 	result :=0
// 	for _, value := range values {
// 		result += value
// 	}
// 	fmt.Println("The sum of the numbers is:", result)
// }

// calling function
// package main
// import (
//     "fmt"
//     )
// func main() {
//     s:=sum(1,2,3,4,5)
// 	fmt.Println("The sum of the number is :", s)
//     }
// func sum(values ...int) int{//returning number type  
// 	fmt.Println(values)
// 	result :=0
// 	for _, value := range values {
// 		result += value
// 	}
// 	return result
// }


// reutnring local variable as a pointer
package main
import (
    "fmt"
    )
// func main() {
//     s:=sum(1,2,3,4,5)
// 	fmt.Println("The sum of the number is :", * s)
//     }
// func sum(values ...int) *int{//returning number type  
// 	fmt.Println(values)
// 	result :=0
// 	for _, value := range values {
// 		result += value
// 	}
// 	return &result
// }

// sending pointer as a local variable is rare but possible in  go language.  In above example we had return value but it was copied to  other variable. But in this case we are  returning the address of the resulting variable. So we are not copying the value but we are copying the address of the variable. So it is a good practice to use pointer in this case.

// Name returns value


// func main() {
//     s:=sum(1,2,3,4,5)
// 	fmt.Println("The sum of the number is :", s)
//     }
// func sum(values ...int) (result int){// name and returning value
	
// 	fmt.Println(values)
// 	for _, value := range values {
// 		result += value
// 	}
// 	return 
// }

// returning multiple value
// package main
// import (
//     "fmt"
//     )
// func main() {
// 	d := divide(5.0,3.0)
// 	fmt.Println("The division of the number is :", d)
//     }
// func divide(a,b float64) float64 {
// 	return a/b
// }


// func main() {
//     d, err := divide(5.0, 0.0) // Capture both the result and the error
//     if err != nil {
//         fmt.Println("Error:", err) // Handle the error
//         return
//     }
//     fmt.Println("The division of the number is:", d) // Print the result if no error
// }
// func divide(a, b float64) (float64, error) {
//     if b == 0.0 {
//         return 0.0, fmt.Errorf("division by zero") // Return an error if b is 0
//     }
//     return a / b, nil // Return the result and no error
// }

// in the above code multiple value are returned but in this code we changes 0.0 the output value will ne +INF. tol solve we had done panic,but it is not a best way instead of this we return the second value and check if it is 0.0 or not. If it is 0.0 then we can return the error and we can handle it in the main function. So we can use the error handling in go language.
// function it self can be treated a  type they can be passes around variable they can be passed ad argument into function

// // FUnction as a type:
// func main() {
//     d, err := divide(5.0, 0.0) // Capture both the result and the error
//     if err != nil {
//         fmt.Println("Error:", err) // Handle the error
//         return
//     }
//     fmt.Println("The division of the number is:", d) // Print the result if no error
// }
// func divide(a, b float64) (float64, error) {
//     if b == 0.0 {
//         return 0.0, fmt.Errorf("division by zero") // Return an error if b is 0
//     }
//     return a / b, nil // Return the result and no error
// }

// func main(){   
// 	func(){           // anonymous function
// 	fmt.Println("Hello go!") 
// }() //this small bracket is used to invoke the function.
// }

// this type of anonymous function is used to when  you don't want to send the message to the main function but you want to print it in the same function. So we can use this type of function.
// using loop
// func main(){
//     for i:=0;i<5;i++{
//         func(){
//             fmt.Println(i)
//         }()
//     }
// }
// in the above code we are using the loop and we are using the anonymous function to print the value of i. So it will print the value of i from 0 to 4. But if we use the normal function then it will print the value of i as 5 because it will take the last value of i. So we can use this type of function to print the value of i in the loop.
// in this code we get the value of i because we are at the scope of the main function and but the use of anonymous function is not send messange to the main function, but as  we are at the scope of the main function,so to solve this use can do:

// func main(){
//     for i:=0;i<5;i++{
//         func(i int){
//             fmt.Println(i)
//         }(i)
//     }

//     //we initialize the i in the anonmymous function and we passed that i variable to the function. Doing this we are not actually not using the main function scope but we are using the anonymous function scope.
    
// }

// working in funciton like a varibale
// func main() {
//     f:= func(){    //this function is like this var f func() = func()
//         fmt.Println("Hello go!")

//     }
//     f()
// }

// func main(){
//     var divide func(float64, float64) (float64, error) // declare the function variable
//     divide = func(a, b float64) (float64, error) { // define the function
//         if b==0.0{
//             return 0.0, fmt.Errorf("division by zero") // Return an error if b is 0
//         } else{
//             return a/b,nil // Return the result and no error
//         }   
//     }
//     d,err:=  divide(5.0,3.0) // Capture both the result and the error
//     if err != nil {
//         fmt.Println("Error:", err) // Handle the error
//         return
//     }
//     fmt.Println("The division of the number is:", d) // Print the result if no error
// }


// function using structured representation

// func main() {
//     g := greeter{ // Use the correct struct name
//         greeting: "hello",
//         name:     "Go",
//     }
//     g.greet()
// }

// type greeter struct { // Struct name is greeter
//     greeting string
//     name     string
// }

// func (g greeter) greet() { // g greeter is used as value type.it is a value reciever
//     fmt.Println(g.greeting, g.name) // Use g.greeting to access the field
// }
//this code used as a copy if we put the name in last function and copy in main func it will not manipulate anything in the code

// using pointer reciever
// func main() {
//     g := greeter{ // Use the correct struct name
//         greeting: "hello",
//         name:     "Go",
//     }
//     g.greet()
//     fmt.Println(g.name)
// }

// type greeter struct { // Struct name is greeter
//     greeting string
//     name     string
// }

// func (g * greeter) greet() { // g greeter is used as value type.it is a value reciever
//     fmt.Println(g.greeting, g.name) // Use g.greeting to access the field
//     g.name = "World" // Changing the value of g.name
// }
// here actually we can maniputale the data.
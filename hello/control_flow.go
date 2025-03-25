// focus on if and else statements

// if statement
// package main
// import (
// 	"fmt"
// )	

// func main() {
// 	if true  {
// 		fmt.Println("This is true")
// 	}
// }
//  if there is false { } the code will not run
// in go there will be no single block statement like in other languages. if we use a block statement we will have to use curly braces. for eg. if true 
// {
//  fmt.Println("This is true")
 //}
// if we donot use curly braces the code will not run. we will get an error when we run the code.
 
// Initializer Syntax
//  calling into a map that is pulling out a value and a boolean. if the value is not there it will return false.
// package main
// import (
// 	"fmt"
// )	

// func main() {
// 	statePopulations := map[string]int{
// 		"California": 39250017,
// 		"Texas": 27862596,
// 		"Florida": 20612439,
// 		"New York": 19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois": 12801539,
// 		"Ohio": 11614373,
// 	}
// 		if pop, ok := statePopulations["Florida"]; ok {
// 			fmt.Println(pop)
// 		}
// }


// comparison operators
// package main
// import (
// 	"fmt"
// )

// func main() {
// 	number :=5
// 	guess :=6
// 	if guess <number {
// 		fmt.Println("Too low")
// 	}
// 	if guess > number {
// 		fmt.Println("Too high")
// 	}
// 	if guess == number {
// 		fmt.Println("You guessed correctly")
//     }
// 	fmt.Println(number <= guess, number>=guess , number!=guess)
// }


// package main
// import (
// 	"fmt"
// )

// func main() {
// 	number :=5
// 	guess :=6
// 	if guess < 1|| guess >10 {
// 		fmt.Println("The guess must be between 1 and 10")
// 	}
// 	if guess < number {
// 		fmt.Println("Too low")
// 	}
// 	if guess > number {
// 		fmt.Println("Too high")
// 	}
// 	if guess == number {
// 		fmt.Println("You guessed correctly")
// 	}
// 	fmt.Println(number <= guess, number>=guess , number!=guess)	
// fmt.Println(!true, "You guessed correctly")
// } 

// package main
// import (
// 	"fmt"
// )	

// func main() {
// 	number :=5
// 	guess :=-6
// 	if guess < 1|| returnTrue() ||guess >10 {
// 		fmt.Println("The guess must be between 1 and 10")
// 		return
// 	}
// 	 if guess < 1 || guess > 10 {
// 	 	fmt.Println("The guess must be between 1 and 10")
// 	 	return
// 	 }
// 	 if guess < number {
// 	 	fmt.Println("Too low")
// 	 }
// 	 if guess > number {
// 	 	fmt.Println("Too high")
// 	 }
// 	 if guess == number {
// 	 	fmt.Println("You guessed correctly")
// 	 }
// 	 fmt.Println(number <= guess, number>=guess , number!=guess)	
// 	 fmt.Println(!true, "You guessed correctly")
// }

// func returnTrue() bool{
// 	fmt.Println("Returning true")
// 	return true
// }

//  the above code is the example of short-circuit evaluation. if the first condition is true the other conditions will not be checked.


// Else If statement

// package main
// import (
// 	"fmt"
// )

// func main() {
// 	switch 2 {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	default:
// 		fmt.Println("not one two or three")
// }
// }

// the above function is the example of switch statement. the switch statement is used to compare the value of the expression to the values specified in the case statements. if the value of the expression matches the value of the case statement the code will run. if the value of the expression does not match the value of the case statement the code will not run.
// if we want to run the code we will have to use the fallthrough statement. the fallthrough statement is used to run the code of the next case statement.
// package main
// import (
// 	"fmt"
// )

// func main() {
// 	switch 6 {
// 	case 1,5,6,10:
// 		fmt.Println("one five six or ten...")
// 	case 2,3,0,9:
// 		fmt.Println("two three zero or nine...")
//     default:
// 		fmt.Println("not one two or three")
// 	}
// }
// the above code is the example of multiple case statements. we can use multiple case statements in a single case statement. if the value of the expression matches the value of the case statement the code will run. if the value of the expression does not match the value of the case statement the code will not run.

// we can also use initializer in case statements

// package main
// import (
// 	"fmt"
// )

// func main() {
// 	switch i:=2+3; i {
// 	case 1,5,6,10:
// 		fmt.Println("one five six or ten...")
// 	case 2,3,0,9:
// 		fmt.Println("two three zero or nine...")
// 	default:
// 		fmt.Println("not one two or three")
// 	}
// }

// the above code is the example of initializer in case statements. we can use initializer in case statements. the initializer is used to initialize the value of the variable. the value of the variable is then compared to the value of the case statement. if the value of the variable matches the value of the case statement the code will run. if the value of the variable does not match the value of the case statement the code will not run.

// tagless syntax
// package main
// import (
// 	"fmt"
// )

// func main() {
// 	i:=10
// 	switch {
// 	case i<=10:
// 		fmt.Println("less than or equal to ten")
// 		fallthrough
// 	case i<=20:
// 		fmt.Println("less than or equal to twenty")
// 	default:
// 		fmt.Println("greater than twenty")
// 	}
// }
//falltrough is logicless.  using this user is taking the responsibility of the code. it is not recommended to use fallthrough. it is better to use the tagless syntax

// type switch
// package main
// import (
// 	"fmt"
// )	

// func main() {
// 	var i interface{} = 1  // interface is a type that can hold any type of value. we are using interface to hold the value of i. This is an example of type assertion.
// 	switch i.(type) { // type assertion is used to get the type of the value of i.
// 	case int:
// 		fmt.Println("i is an integer")
// 		break
// 	case float64:
// 		fmt.Println("i is a float64")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i is another type")
// 	}
// }


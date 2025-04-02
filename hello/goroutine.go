// creting go routine
// package main
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go sayHello() // Call the function to print "Hello, World!"
// 	time.Sleep(1 * time.Millisecond) // Sleep for 100 milliseconds to allow the goroutine to finish

// }

// func sayHello() {
// 	fmt.Println("Hello, World!")
// }

// using  anonymous function
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     var msg = "hello"
//     go func() { // Start a goroutine with an anonymous function
//         fmt.Println(msg) // Print the message
//     }() // immediately call the anonymous function
//     time.Sleep(1 * time.Millisecond) // Sleep for 1 millisecond to allow the goroutine to finish
// }
// although var msg is at the main function, it is still accessible in the goroutine because of closure which means the anonmymous function has a scope of using the var located outside the outerscope

// the goroutine has access to the variables in the scope where it was created	

//creating the variable dependency variable in go routine

// package main
// import (
// 	"fmt"
// 	"time"
// )
// func main() {
// 	var msg = "hello"
// 	go func() { // Start a goroutine with an anonymous function
// 		fmt.Println(msg) // Print the message
// 	}() // immediately call the anonymous function
// 	msg = "hello world" // Change the value of msg after starting the goroutine
// 	time.Sleep(1 * time.Millisecond) // Sleep for 1 millisecond to allow the goroutine to finish
// }
// in this case the goroutine will print "hello world" because the value of msg was changed after the goroutine was started. This is because the goroutine has access to the variable msg and it will use the latest value of msg when it is executed. This is a common pitfall in Go when using goroutines, as the value of the variable may change before the goroutine is executed. To avoid this, you can pass the value of the variable as an argument to the goroutine, like this:

	// package main
	// import (
	// 	"fmt"
	// 	"time"
	// )
	// func main() {
	// 	var msg = "hello"
	// 	go func(msg string) { // Start a goroutine with an anonymous function
	// 		fmt.Println(msg) // Print the message
	// 	}(msg) // immediately call the anonymous function
	// 	msg = "hello world" // Change the value of msg after starting the goroutine
	// 	time.Sleep(1 * time.Millisecond) // Sleep for 1 millisecond to allow the goroutine to finish
	// }


// in this case the goroutine will print "hello" because the value of msg was passed as an argument to the goroutine and it will use the value of msg at the time the goroutine was started. This is a common pattern in Go when using goroutines, as it allows you to avoid the pitfalls of variable dependencies and closures.
// this is a common pattern in Go when using goroutines, as it allows you to avoid the pitfalls of variable dependencies and closures. It is also a good practice to use this pattern when passing variables to goroutines, as it makes the code more readable and easier to understand.

// weight group

// package main
// import (
// 	"fmt"
// 	"sync"
// )

// // var wg = sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish // weight group synchronizes the goroutines and wait for them to finish before exiting the program
// // func main() {
// // 	var msg = "hello"
// // 	wg.Add(1) // Add 1 to the WaitGroup counter
// // 	go func() { // Start a goroutine with an anonymous function
// // 		fmt.Println(msg) // Print the message
// // 		wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// // 		// wg.Wait() // Wait for the goroutine to finish	
// // 	}() // immediately call the anonymous function
// // 	msg = "hello world" // Change the value of msg after starting the goroutine
// // 	//time.Sleep(1 * time.Millisecond) // Sleep for 1 millisecond to allow the goroutine to finish // using wait group instead of sleep
// // 	wg.Wait() // Wait for the goroutine to finish
// // }

// // working with multiple data

// var wg = sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish
// var counter = 0 // Create a counter variable to keep track of the number of goroutines
// func main(){
// 	for i := 0; i<10;i++ { // Loop 10 times
// 		wg.Add(2) // Add 1 to the WaitGroup counter
// 		go sayHello() // Start a goroutine with an anonymous function
// 		go increment () // Start a goroutine with an anonymous function

// 	}
// 	wg.Wait() // Wait for the goroutine to finish
// }
// func sayHello() { // Function to print "Hello, World!"
// 	fmt.Println("Hello, World!") // Print the message
// 	wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// }
// func increment() { // Function to increment the counter variable
// 	counter++ // Increment the counter variable
// 	wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// }
// func printCounter() { // Function to print the counter variable
// 	fmt.Println("Counter:", counter) // Print the counter variable
// 	wg.Done() // Decrement the WaitGroup counter when the goroutine is done
//}
// in this case the goroutine will print "hello world" because the value of msg was changed after the goroutine was started. This is because the goroutine has access to the variable msg and it will use the latest value of msg when it is executed. This is a common pitfall in Go when using goroutines, as the value of the variable may change before the goroutine is executed. To avoid this, you can pass the value of the variable as an argument to the goroutine,


// Using mutex
// mutex is a lock that the application going to honor
// package main

// import (
//     "fmt"
//     "sync"
//     "runtime"
// )

// var wg = sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish
// var counter = 0           // Create a counter variable to keep track of the number of goroutines
// var m = sync.Mutex{}      // Create a mutex to protect the counter variable

// func main() {
//     runtime.GOMAXPROCS(1) // Set the number of OS threads to 1
//     for i := 0; i < 10; i++ { // Loop 10 times
//         wg.Add(2) // Add 2 to the WaitGroup counter (one for each goroutine)
// 		m.Lock()         // Lock the mutex to protect the counter variable
//         go sayHello() // Start a goroutine to print "Hello, World!"
//         go increment() // Start a goroutine to increment the counter
//     }
//     wg.Wait() // Wait for all goroutines to finish
// }

// func sayHello() { // Function to print "Hello, World!"
//     fmt.Println("1") // Print the message
//     wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// }

// func increment() { // Function to increment the counter variable
   
//     counter++        // Increment the counter variable
//     m.Unlock()       // Unlock the mutex
//     wg.Done()        // Decrement the WaitGroup counter when the goroutine is done
// }
// in this case the goroutine will print "hello world" because the value of msg was changed after the goroutine was started. This is because the goroutine has access to the variable msg and it will use the latest value of msg when it is executed. This is a common pitfall in Go when using goroutines, as the value of the variable may change before the goroutine is executed. To avoid this, you can pass the value of the variable as an argument to the goroutine,
	
	

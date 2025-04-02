// mutex is a lock that the application going to honor
package main
import (
	"fmt"
	"sync"
	"runtime"
)
var wg = sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish
var counter = 0 // Create a counter variable to keep track of the number of goroutines
var m = sync.Mutex{} // Create a mutex to protect the counter variable
func main(){
	runtime.GOMAXPROCS(1) // Set the number of OS threads to 1
	for i := 0; i<10;i++ { // Loop 10 times
// 		wg.Add(2) // Add 1 to the WaitGroup counter
		go sayHello() // Start a goroutine with an anonymous function
		go increment () // Start a goroutine with an anonymous function
	}
	wg.Wait() // Wait for the goroutine to finish
}

func sayHello() { // Function to print "Hello, World!"
	fmt.Println("Hello, World!") // Print the message
	wg.Done() // Decrement the WaitGroup counter when the goroutine is done
}
func increment() { // Function to increment the counter variable
	counter++ // Increment the counter variable
	wg.Done() // Decrement the WaitGroup counter when the goroutine is done
}
func printCounter() { // Function to print the counter variable
	
	fmt.Println("Counter:", counter) // Print the counter variable
	wg.Done() // Decrement the WaitGroup counter when the goroutine is done
}
// in this case the goroutine will print "hello world" because the value of msg was changed after the goroutine was started. This is because the goroutine has access to the variable msg and it will use the latest value of msg when it is executed. This is a common pitfall in Go when using goroutines, as the value of the variable may change before the goroutine is executed. To avoid this, you can pass the value of the variable as an argument to the goroutine,
	
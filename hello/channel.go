// // while working with the channel it is almost always going to work with go routine because channel are designed to synchronize the goroutines

// package main

// import (
//     "fmt"
//     "sync"
// )

// var wg = sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish

// func main() {
//     ch := make(chan int) // Create a channel of type int
//     wg.Add(2)            // Add 2 to the WaitGroup counter (one for each goroutine)

//     go func() { // Start a goroutine to receive data from the channel
//         i := <-ch // Receive data from the channel
//         fmt.Println(i) // Print the received value
//         wg.Done()      // Decrement the WaitGroup counter when the goroutine is done
//     }()

//     go func() { // Start a goroutine to send data to the channel
//         ch <- 42 // Send data to the channel //this is the process of sending the data to the channel by <- operator. in this channel is data is sending to the channel so the head of operator is toward at the channel where as when data is receiving from the channel the head of operator is toward at the variable where data is receiving.
		 
//         wg.Done() // Decrement the WaitGroup counter when the goroutine is done
//     }()

//     wg.Wait() // Wait for all goroutines to finish
// }

// since in this we are sending the copies of the data so we can maniputale the data

// creating multiple  go routine that works in only one channel
// package main
// import(
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// func main(){
// 	ch:= make(chan int) // Create a channel of type int
// 	for j:=0;j<5;j++{{
// 		wg.Add(2) // Add 2 to the WaitGroup counter for each goroutine
// 		go func(){ // Start a goroutine to receive data from the channel
// 			i := <-ch // Receive data from the channel
// 			fmt.Println(i) // Print the received value
// 			wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// 		}()
// 		go func(){ // Start a goroutine to send data to the channel
// 			ch <-j 
// 			wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// 		}()
// 		wg.Wait() // Wait for all goroutines to finish
// 	}
// }
// }
	


// package main
// import(
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// func main(){
// 	ch := make(chan int) // Create a channel of type int
// 	go func(){ //anonymous function // Start a goroutine to receive data from the channel
// 		i := <-ch // Receive data from the channel
// 		fmt.Println(i) // Print the received value
// 		wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// 	}()
// 	for j :=0;j<5;j++{ // Start a goroutine to send data to the channel
// 		wg.Add(2) // Add 1 to the WaitGroup counter for each goroutine
// 		go func(){ // Start a goroutine to send data to the channel
// 			ch<-42 // Send data to the channel
// 			wg.Done() // Decrement the WaitGroup counter when the goroutine is done
// 		}()
// 		wg.Wait() // Wait for all goroutines to finish
// 		}
// }
// it is causing deadlock problem because the channel is not buffered and the goroutine is waiting for the data to be sent to the channel and the main goroutine is waiting for the goroutine to finish. so we need to use buffered channel.
// buffered channel is a channel that can hold a fixed number of values before it blocks. this means that the sender can send data to the channel without waiting for the receiver to receive it, up to the buffer size. once the buffer is full, the sender will block until the receiver receives some data from the channel. this allows for more efficient communication between goroutines, as they can work independently without blocking each other.

// sending and reaciving betwen go func
// package main

// import (
//     "fmt"
//     "sync"
// )

// var wg = sync.WaitGroup{}


// func main(){
// 	ch := make(chan int)
// 	wg.Add(2)
	
	
// 	go func(){
// 		i :=<-ch
// 		fmt.Println(i)
// 		ch<-27
// 		wg.Done()
// 	}()
	
	
// 	go func(){
// 		ch<-42
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
	
	
// 	wg.Wait()


// }

// Dedicate channel to write or read we do bias on the direction of the channel. It is done by accepting the variable in the go routine

// package main

// import(

// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main(){

// 	ch := make(chan int)
// 	wg.Add(2)


// 	go func(ch<-chan int){ // recieving data from the channel
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
	
// 	go func(ch chan<- int){ // sending data to the channel
// 		ch<-42
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }


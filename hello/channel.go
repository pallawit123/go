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
// 		ch <-27
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// in this code we got deadlock because  there is nothing to  do with the msg ch <-27 so application blowsup and the go routine never completes so tosolvethis we use buffer while defining channel like in this way

// package main

// import(

// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main(){

// 	ch := make(chan int , 50) // Create a buffered channel of type int with a buffer size of 50
// 	// this means that the channel can hold up to 50 values before it blocks. this allows for more efficient communication between goroutines, as they can work independently without blocking each other.	
// 	wg.Add(2)


// 	go func(ch<-chan int){ // recieving data from the channel
// 		 i := <-ch
// 		fmt.Println(i)
// 		i = <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
	
// 	go func(ch chan<- int){ // sending data to the channel
// 		ch<-42
// 		ch  <-27 // this is the process of sending the data to the channel by <- operator. in this channel is data is sending to the channel so the head of operator is toward at the channel where as when data is receiving from the channel the head of operator is toward at the variable where data is receiving.
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// doing buffer the message ch <-27 is lost. this problem is handles above but this  is not the best way to handle this problem so so solve this using looping contruct
package main

import(

	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){

	ch := make(chan int , 50) // Create a buffered channel of type int with a buffer size of 50
	// this means that the channel can hold up to 50 values before it blocks. this allows for more efficient communication between goroutines, as they can work independently without blocking each other.	
	wg.Add(2)
	go func(ch<-chan int){ // recieving data from the channel
		for i := range ch { // using range to receive data from the channel until it is closed
			fmt.Println(i) // Print the received value
		}
		wg.Done() // Decrement the WaitGroup counter when the goroutine is done
	}(ch)


	go func(ch<-chan int){ // recieving data from the channel
		 i  := <-ch            // is we use looping contruct the first data will be index and second will be variable
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	
	go func(ch chan<- int){ // sending data to the channel
		ch<-42
		ch  <-27 // this is the process of sending the data to the channel by <- operator. in this channel is data is sending to the channel so the head of operator is toward at the channel where as when data is receiving from the channel the head of operator is toward at the variable where data is receiving.
		close(ch) // Close the channel to signal that no more data will be sent
		// this is important because the receiving goroutine will block until the channel is closed, so it will not exit until the channel is closed. if the channel is not closed, the receiving goroutine will block forever and the program will hang.
		wg.Done()
	}(ch)
	wg.Wait()
}

// if we close all botht he channel we will get a panic in the program because the channel is closed and we are trying to send data to the channel. this is not allowed in go. so we need to close the channel only when we are done sending data to the channel. this is important because the receiving goroutine will block until the channel is closed, so it will not exit until the channel is closed. if the channel is not closed, the receiving goroutine will block forever and the program will hang.
// so we need to close the channel only when we are done sending data to the channel. this is important because the receiving goroutine will block until the channel is closed, so it will not exit until the channel is closed. if the channel is not closed, the receiving goroutine will block forever and the program will hang.


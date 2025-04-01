// interference is a typejust like structure. type is a keyword in interface

// package main

// import (
//     "fmt"

// )





// func main() {
// 	var w Writer = consoleWriter{} // creating a variable of type writer and assigning it to consolewriter
// 	w.Write([]byte("Hello, GO!")) // calling the write method of the interface and passing the string as byte array


// }
// type  Writer interface {  // type name of interface is writer interface
// 	Write([]byte) (int, error)
// }
// type consoleWriter struct {}
// func (cw consoleWriter) Write(data[]byte)(int,error){
// 	n,err:=fmt.Println(string(data))
// 	return n ,err
// }

// in single method of naming convention it should be called at what it done like if it read it should be called reader and if it writer it should be writer
// any type that have a method can be implemented  here.


// func main() {
//     myInt := intcounter(0) // Declare and initialize myInt
//     var inc Increment = &myInt // Assign the address of myInt to the Increment interface
//     for i := 0; i < 10; i++ {
//         fmt.Println(inc.Increment()) // Call the Increment method and print the result
//     }
// }

// type Increment interface { // Define the Increment interface
//     Increment() int  //method that only increment an integer value
// 	// Decrement() int // Uncomment if you want to add a Decrement method
// }

// type intcounter int // Define intcounter as a type alias for int

// func (ic *intcounter) Increment() int {
//     *ic = *ic + 1 // Increment the value of ic
//     return int(*ic) // Return the incremented value
// }


// using multiple interfaces
package main

import (
    "fmt"
	"bytes"
)
//  func main(){
// 	var wc writerCloser = NewBufferdWriter()
// 	we.Write([]byte("Hello, GO!")) // calling the write method of the interface and passing the string as byte array
// 	wc.Close() // calling the close method of the interface and passing the string as byte array
// }
// type Writer interface{
// 	write([]byte) (int, error)
// }
// type Closer interface{
// 	close() error
// }
// type writerCloser interface{
// 	Writer
// Closer // embedding the writer and closer interface
// }
// type bufferedWriterCloser interface {
// 	buffer *bytes.BUfferbuffer
	
// }


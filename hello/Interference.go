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
// package main

// import (
//     "bytes"
//     "fmt"
// )

// func main() {
//     var wc writerCloser = NewBufferedWriter() // Corrected function name
//     wc.Write([]byte("Hello, GO!"))           // Use wc instead of we
//     wc.Close()
// }

// type Writer interface {
//     Write([]byte) (int, error) // Corrected method name to uppercase
// }

// type Closer interface {
//     Close() error // Corrected method name to uppercase
// }

// type writerCloser interface { // composing two inteface
//     Writer
//     Closer // Embedding the Writer and Closer interfaces
// }

// type bufferedWriterCloser struct { // Changed to struct
//     buffer *bytes.Buffer
// }

// // Implement the Write method for bufferedWriterCloser
// func (bwc *bufferedWriterCloser) Write(data []byte) (int, error) {
//     n, err := bwc.buffer.Write(data)
//     if err == nil {
//         fmt.Println(string(data)) // Print the written data
//     }
//     return n, err
// }

// // Implement the Close method for bufferedWriterCloser
// func (bwc *bufferedWriterCloser) Close() error {
//     fmt.Println("Closing the buffer")
//     return nil
// }

// // Define the NewBufferedWriter function
// func NewBufferedWriter() *bufferedWriterCloser {
//     return &bufferedWriterCloser{
//         buffer: bytes.NewBuffer([]byte{}),
//     }
// }

// type conversion

// package main

// import (
//     "bytes"
//     "fmt"
// )

// func main() {
//     var wc writerCloser = NewBufferedWriter() // Corrected to use writerCloser
//     wc.Write([]byte("Hello, GO!"))           // Use wc to call Write
//     wc.Close()                               // Use wc to call Close

//     // Perform type assertion to convert writerCloser to bufferedWriterCloser
//     bwc := wc.(*bufferedWriterCloser)
//     fmt.Printf("%T\n", bwc) // Print the type of bwc
// }

// type Writer interface {
//     Write([]byte) (int, error) // Corrected method name to uppercase
// }

// type Closer interface {
//     Close() error // Corrected method name to uppercase
// }

// type writerCloser interface { // Composing two interfaces
//     Writer
//     Closer // Embedding the Writer and Closer interfaces
// }

// type bufferedWriterCloser struct { // Changed to struct
//     buffer *bytes.Buffer
// }

// // Implement the Write method for bufferedWriterCloser
// func (bwc *bufferedWriterCloser) Write(data []byte) (int, error) {
//     n, err := bwc.buffer.Write(data)
//     if err == nil {
//         fmt.Println(string(data)) // Print the written data
//     }
//     return n, err
// }

// // Implement the Close method for bufferedWriterCloser
// func (bwc *bufferedWriterCloser) Close() error {
//     fmt.Println("Closing the buffer")
//     return nil
// }

// // Define the NewBufferedWriter function
// func NewBufferedWriter() *bufferedWriterCloser {
//     return &bufferedWriterCloser{
//         buffer: bytes.NewBuffer([]byte{}),
//     }
// }


// empty interface
// package main

// import (
//     "fmt"
// )

// func main() {
// 	var i interface{} =0
// 	switch i.(type) { // var name i . type this is called a swtich type. in this type of cases switch works aas type.
// 	case int:
// 		fmt.Println("i is an integer number")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i is of a different type")
// 	}
// }


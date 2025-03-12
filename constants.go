// CONSTANT
// Naming convention
// type constants
// untyped constants
// Enumerated constants
// Enumerated expressions

 // NAming convetion
// package main
// import(
// 	"fmt"
// )
// func main() {	
	

// type constants const keyword name of the constant and type constants and values
// value of the contant never be changed	
// package main
// import(
// 	"fmt"
// )
// func main() {	
// 	 const myconst int = 50
// 	 fmt.Printf("%v, %T\n", myconst, myconst)
// }

// constant can be shadowed
// package main
// import(
// 	"fmt"
// )
// 	const a int = 14
// func main() {	
// 	const  a int = 50
// 	fmt.Printf("%v, %T\n", a, a)
// }
// At above inner constant shadowed the outer constant


// Adding two variable
// if constant are of different types then they can't be added like int and float int and int16 etc.
// package main
// import(
// 	"fmt"
// )
// func main() {
// 	const a int = 47
// 	const b = 27
// 	fmt.Printf("%v, %T\n", a+b, a+b)
// }


// package main
// import(
// 	"fmt"
// )
// func main() {
// 	const a  = 47
// 	const b int16 = 27
// 	fmt.Printf("%v, %T\n", a+b, a+b)
// }
// in the code constant a is not treated as an int because it is untyped constant. so whenever the a+b is executing the value is a is 42 and the addition is executed as 42+27=69.

// Enumerated constants
// package main
// import(
// 	"fmt"
// )
// const (
// 	a = iota
// 	b = iota
// 	c = iota 
// )
// 	func main() {
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// }

// package main
// import(
// 	"fmt"
// )
// const (
// 	a = iota
// 	b 
// 	c 
// )
// 	func main() {
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// }
	

// package main
// import(
// 	"fmt"
// )
// const (
// 	a = iota
// 	b 
// 	c 
// )
// const (
// 	a1= iota
// )
// 	func main() {
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%v\n", a1)
// }

// _ is the blank identifier character that is used to skip the value of the iota.
// package main
// import(
// 	"fmt"
// )
// const (
// 	_ = iota // ignore the first value by assigning to blank identifier
// 	Kb = 1 << (iota * 10)
// 	Mb 
// 	Gb
// 	Tb
// 	Pb
// 	Eb
// 	Zb
// 	Yb
// )
// func main() {
// 	fileSize := 4000000000.
// 	fmt.Printf("%.2fGB\n", fileSize/Gb)
// }
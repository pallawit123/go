// declaration of an array. The size of the array is fixed.
// array can hold single type of data. so while initialing the array we use a=[index(number of elemnt that array can hold array)] type of data {data}

// package main
// import(
// 	"fmt"
// )

// func main() {
// 	grades := [4]int{97, 85, 93, 74}
// 	fmt.Printf("grades %v ", grades)

// }
//  In above code we decalre the size of array twice which does not really require  so for that we use [...] instead of [4]
// package main
// import(
// 	"fmt"
// )

// func main() {
// 	grades := [...]int{97, 85, 93, 74}
// 	fmt.Printf("grades %v ", grades)

// }

// // we can also  declare the array without any value and then assign the value to it.
// package main
// import(
// 	"fmt"
// )

// func main() {
// 	var students [3]string
// 	fmt.Printf("students %v ", students)

// }
//   In above example we declare the array without any value. to specify the value in an array we use index number of the array.
// we can also  declare the array without any value and then assign the value to it.
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	var students [3]string
// 	fmt.Printf("students %v \n ", students)
// 	students[0] = "Lisa"
// 	fmt.Printf("students %v ", students)
// }

// package main
// import(
// 	"fmt"
// )

// func main() {

// 	var students [3]string
// 	fmt.Printf("students %v \n ", students)
// 	students[0] = "Lisa"
// 	students[1] = "Ahmed"
// 	students[2] = "Arnold"
// 	fmt.Printf("students %v ", students)
// 	fmt.Printf("number of students %v ", len(students))
// }

// in other language array point to the first element of the array but in go array is a value type so when we assign the array to another array it will create a copy of the array.
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := [...]int{1,2,3}
// 	b := a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// & operator is used to get the address of the variable.
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := [...]int{1,2,3}
// 	b := &a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// in the aboev code if we change the data of a it will change the data of b also because b is pointing to the address of a.

// slice
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := [...]int{1,2,3}
// 	b := &a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println("Length: ",len(a))
// 	fmt.Println("Capacity: ",cap(a))
// }

// In program length and capacity are different. length is the number of elements in the array and capacity is the number of elements that the array can hold.
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := [...]int{1,2,3,4,5,6,7,8,9,10}
// 	 b := a[:] //slice all the element of the array
// 	c := a[3:] // slice the element from 3 to end
// 	d := a[:6] // slice the element from 0 to 6
// 	e := a[3:6] // slice the element from 3 to 6
// 	f := a[0:0] // slice the element from 0 to 0
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// 	fmt.Println(f)
// }

// make function: this function is used to create a slice. make function is used to create a slice with a specific length and capacity.

// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := make ([] int, 3)
// 	fmt.Println(a)
// 	fmt.Println("Length: ",len(a))
// 	fmt.Println("Capacity: ",cap(a))
// }

// capacity is the number of elements that the slice can hold. if we add more element to the slice then the capacity of the slice will be doubled.

// slice doesnot have a fixed size. we can add element and remove elements from the slice.
// to add more elements we use append. append function is used to add elements to the slice.
// to remove elements from

// append function
// syntax = append(source slice, element)
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := []int{}
// 	fmt.Printf("The length and the capacity of the slice before appending  is :\n")
// 	fmt.Println("Length: ",len(a))
// 	fmt.Println("Capacity: ",cap(a))
// 	a = append(a, 1)
// 	fmt.Println(a)
// 	fmt.Printf("\n the length and the capacity of the slice after appending is \n")
// 	fmt.Println("Length: ",len(a))
// 	fmt.Println("Capacity: ",cap(a))
// }

// append function can take multiple elements at a time.
// package main
// import(
// 	"fmt"
// )

// func main() {

// 	a := []int{}
// 	fmt.Printf("The length and the capacity of the slice before appending  is :\n")
// 	fmt.Println("Length: ",len(a))
// 	fmt.Println("Capacity: ",cap(a))
// 	a = append(a, 1)
// 	fmt.Println(a)
// 	fmt.Printf("\n the length and the capacity of the slice after appending is \n")
// 	a = append(a,2,3,4)
// 	fmt.Println(a)
// 	fmt.Println("Length: ",len(a))
// 	fmt.Println("Capacity: ",cap(a))
// }

// poping the elements of the slice

// package main
// import(
// 	"fmt" //

// )

// func main() {
// 	a := []int{1,2,3,4,5}
// 	b := a[1:] // poping the first element of the slice
// 	fmt.Println(b)
// }

// trim elements of the end

// package main

// import (
// 	"fmt" //
// )

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := a[:len(a)-1] // poping the first element of the slice
// 	fmt.Println(b)
// }

// trim from the middle
package main

import (
	"fmt" //
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := append (a[:2],a[3:]...) // poping the first element of the slice
	fmt.Println(b)
	fmt.Println(a)
}

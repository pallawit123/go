// Mapping a struct to a map of pointers
// map is going to take some kind of key and some kind of value.

// literal syntax methods
// literal syntax methods are used to create a map with some key-value pairs.
// syntax = map[key_type]value_type{key1: value1, key2: value2, key3: value3}
// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Printf("product_details %v ", product_details)
// }

// A slice cannot be a key in a map. A map is a collection of key-value pairs, where the key must be unique. A slice is a variable-length sequence which is not comparable.

// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	 m := map[[2]int]string{}
// 	fmt.Printf("map %v ", product_details)
// 	fmt.Printf("map %v ", m)
// }

// mapping using make function 
// syntax: map_name := make(map[key_type]value_type)
// used while executing loops.

// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Printf("product_details %v ", product_details)
// }


// by creating second parameter

// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int, 10)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Printf("product_details %v ", product_details)
// }

// manipulation of maps
// it can be done by pulling out the value of the key and then changing the value of the key.
 
// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Println(product_details)
// 	// adding in map
// 	product_details["headphone"] = 678901
// 	fmt.Println(product_details ["headphone"])
// 	fmt.Println(product_details)

// }

// deleting from map
// it can be done by using the built in function delete(map_name, key_name)

// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Println(product_details)
// 	product_details["headphone"] = 678901
// 	fmt.Println(product_details ["headphone"])
// 	fmt.Println(product_details)
// 	// deleting from map
// 	delete(product_details, "headphone")
// 	fmt.Println(product_details)

// } 

// it can be done by popping out the value of the key and then changing the value of the key.

// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	pop, ok := product_details["computer"]
// 	fmt.Println(pop,ok)
// 	pop, ok = product_details["lapto"]
// 	fmt.Println(pop,ok)
// }
// pop,ok is used to check if the key is present in the map or not. if the key is present in the map then it will return the value of the key and true. if the key is not present in the map then it will return 0 and false.

// It can be used to find out the total number of keys in the map.
// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Println(len(product_details))

// }

// manipulating the one element in map might affect the other element in the map. to avoid this we do
// package main

// import (
//     "fmt"
// )
// func main() {
// 	product_details := make(map[string]int)
// 	product_details = map[string]int{
// 	"computer" : 234567,
// 	"laptop" : 123456,
// 	"mobile" : 345678,
// 	"tablet" : 456789,
// 	"smartwatch" : 567890,
// 	}
// 	fmt.Println(product_details)
// 	product_details1 := product_details
// 	delete(product_details1, "computer")
// 	fmt.Println(product_details1)
// 	fmt.Println(product_details)

// }

// mapping a struct to a map of pointers
// package main

// import (
//     "fmt"
// )

// type student struct {
//     name      string
//     roll_no   int
//     companion []string
// }

// func main() {
//     astudents := student{
//         name:    "Lisa,",
//         roll_no: 123,
//         companion: []string{
//             "Ahmed",
//             "Arnold",
//         },
//     }

//     // Use fmt to print the astudents variable
//     fmt.Println( astudents)
// }


// package main

// import (
//     "fmt"
// )

// type student struct {
//     name      string
//     roll_no   int
//     companion []string
// }

// func main() {
//     astudents := student{
//         name:    "Lisa,",
//         roll_no: 123,
//         companion: []string{
//             "Ahmed",
//             "Arnold",
//         },
//     }

//     // Use fmt to print the astudents variable
//     fmt.Println( astudents.roll_no)
// 	fmt.Println( astudents.companion)
// }

// the above code is of postional syntax. Here all the position should be correct. If we change the position of the companion then it will give the error.


// Naming convention
// If we start with the capital leteer then it is going to be exported from the package / public and if we start witht he lowercase it's going to be internal to the package/ private package

// package main

// import (
//     "fmt"
// )

// type student struct {
//     Name      string
//     Roll_no   int
//     Companion []string
//     Attend    []string
// }

// func main() {
//     astudents := student{
//         Name:    "Lisa",
//         Roll_no: 123,
//         Attend:  []string{}, // Correctly initializing the `Attend` field
//         Companion: []string{
//             "Ahmed",
//             "Arnold",
//         },
//     }
//     fmt.Println(astudents.Roll_no)
// }
// // In the above code we have changed the name of the struct fields to the capital letter.
// uppercase is going to import the struct fields from the package and lowercase is going to be internal to the package.

// we can declare the struct fields in different ways
// package main

// import (
//     "fmt"
// )
// func main(){
//     aDoctor := struct{name string}{name: "doctor"}
//     fmt.Println(aDoctor)
// }

//first curly braces are used to declare the struct and the second curly braces are used to declare the fields of the struct.
// it is an anonymous struct.


// unlike map sturct is value type. unlike map or slices it represent imndependent data type. eg
// package main

// import (
//     "fmt"
// )
// func main(){
//     aDoctor := struct{name string}{name: "doctor"}
//     anotherDoctor := aDoctor
//     anotherDoctor.name = "harry"
//     fmt.Println(aDoctor)
//     fmt.Println(anotherDoctor)
// }

// in the above code we have created a struct aDoctor and then we have created another struct anotherDoctor and then we have assigned the value of aDoctor to anotherDoctor. then we have changed the value of anotherDoctor. It will not change the value of aDoctor because struct is a value type.
 
// just like arryas if we want to point same underlying data then we can use the pointer.
// package main

// import (
//     "fmt"
// )
// func main(){
//     aDoctor := struct{name string}{name: "doctor"}
//     anotherDoctor := &aDoctor
//     anotherDoctor.name = "harry"
//     fmt.Println(aDoctor)
//     fmt.Println(anotherDoctor)
// }

// in above code, the pointeris pointing to the underlying data. ADoctor is structureditself and anotherDcotor is pointer to the struct. so when we manipulate the feild, we are actually manipulating the name feild of aDoctor.

//Embedding
// go doesn't support traditional OOP principles like inheritance. Instead it uses composition. like go doesnot support the inheritance but it supports the composition.

// package main

// import (
//     "fmt"
// )
// type animal struct{
//     name string

//     origin string
// }

// type bird struct{
//     animal  //it is embedding the animal struct in the bird struct. if we had write animal animal ithen it will have a name feild in the bird struct.
//     sppedKPH float64
//     canFly bool
// }
// func main(){
//     b := bird{}
//     b.name = "Emu"
//     b.origin = "Australia"
//     b.sppedKPH = 48
//     b.canFly = false
//     fmt.Println(b.name, b.origin, b.sppedKPH, b.canFly)
// }

// above code feel like it is using inheritance but it is not using inheritance. it is using composition. it is embedding the animal struct in the bird struct. it is not inheriting the animal struct. it is just using the animal struct in the bird struct.

// tag: it is used to distinguish the fields of the struct

// package main

// import (
//     "fmt"
//     "reflect"
// )

// type Animal struct {
//     Name   string `required max:"100"`
//     Origin string
// }

// func main() {
//     t := reflect.TypeOf(Animal{}) // Ensure it's a struct and not a pointer to a struct
//     field, _ := t.FieldByName("Name") // Correct spelling of `field`
//     fmt.Println(field.Tag) // Use `field` to print the tag
// }
// in the above code we have used the tag to distinguish the fields of the struct. we have used the tag required and max to distinguish the fields of the struct.
// syntax of tag should be like this `required max:"100"`
// to gt the value of the struct we can use the reflect package.
// there is no way from object to get the value of the tag. we can only get the value of the tag from the reflect package.
// reflect package is used to get the value of the tag.

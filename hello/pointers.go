


// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	a :=42
// 	b := a
// 	fmt.Println(a,b)
// }


// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	a :=42
// 	b := a
// 	fmt.Println(a,b)
// 	a = 27
// 	fmt.Println(a,b)
// }

// in above code the value of a is copied to b so when we change the value of a it will not change the value of b. Now to change some behavior of b we use something called pointers.
// pointers are used to store the address of a variable. In go we use & operator to get the address of a variable and * operator to dereference the pointer. Dereferencing means getting the value at the address stored in the pointer. So if we change the value at that address it will change the value of the variable also.
// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var a int=42
// 	var b *int = &a // b is a pointer to a  
// 	fmt.Println(a,b)

// }

// when you run this code you get a output like this:
// 42 0xc00000a0b8 the address of a in memory.
// to know if actually b is holding the address of a we can use & in a while printing : fmt.Println(a,&a,b)
// so when we run this code we get the output like this: 
// 42 0xc00000a0b8 0xc00000a0b8

// dereferencing the pointer means getting the value at the address stored in the pointer. So if we change the value at that address it will change the value of the variable also.

// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var a int = 42
// 	var b *int = &a // * before type is declaring data to the type
// 	fmt.Println(a,*b) // * before to the pointer means we are going to dereference the pointer and get the value at that address.
// }

// in this code since b is pointing to the address of a so when we dereference b we get the value of a. So if we change the value at that address it will change the value of the variable also.

// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var a int =42
// 	var b *int= &a
// 	fmt.Println(a,*b)
// 	a =27
// 	fmt.Println(a,*b)
// }


// we can use dereference the pointer and use that to change the value . so using deference to b and assign the value it will change both the value of a nad b

// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var a int =42
// 	var b *int = &a
// 	fmt.Println(a,*b)
// 	a =27
// 	fmt.Println(a,*b)
// 	*b = 27
// 	fmt.Println(a,*b)
// }

// in the above code we are using dereference to b and assign the value it will change both the value of a nad b. so when we run this code we get the output like this:
// 42 42
// 27 42
// 27 27

// pointer pointing to the array

// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	a := [3]int{1,2,3}
// 	b := &a[0] // b is a pointer to the first element of the array a
// 	//fmt.Println(a,b)
// 	c := &a[1] // c is a pointer to the second element of the array a
// 	fmt.Printf("%v %p %p \n",a,b,c)
// }

// in the above code we are using & operator to get the address of the first element of the array a and assign it to b. so when we run this code we get the output like this:
// [1 2 3] 0xc00000a0b8 0xc00000a0c0
// the address of the first element of the array a is 0xc00000a0b8 and the address of the second element of the array a is 0xc00000a0c0. so we can see that the address of the second element is 4 bytes ahead of the first element. this is because each int takes 4 bytes in memory. so we can say that the address of the second element is equal to the address of the first element + 4 bytes. so we can use this to iterate through the array using pointers.
// we can see the memory address of both the element [1 2 3] 0xc000014120 0xc000014128  now if we substract -8 from the memory address from b it should give the value of a. If  wedo this type  of thing then go will show erorr because go doesnot allowed to perform such kind of mathematical operations on pointers. so we have to use the pointer arithmetic interface to do this.
// we can use the unsafe package to do this because we don't want to use the pointer arithmetic interface in go because it is not safe


// creating pointer type

// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var ms *mystruct 
// 		ms = &mystruct{foo:42}
// 		fmt.Println(ms.foo)
// 	}
// type mystruct struct {
// 	foo int
// 	}
// in the above code we are creating a pointer to a struct and assigning it to ms. so when we run this code we get the output like this:
// this is not only the way to initialize a pointer to a struct. we can also create a pointer using a new function that 


// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var ms *mystruct 
// 		ms = new(mystruct) // this will create a pointer to a struct and assign it to ms. so when we run this code we get the output like this:
// 		fmt.Println(ms.foo)
// 	}
// type mystruct struct {
// 	foo int
// 	}

// in the above code we  created a new  struct  using the new function and assigned it to ms. so when we run this code we get the output like this:
// 0 
// the value of foo is 0 because we donot have assigned any value to foo. and ever value declared in go has an initial value . so right after loc 147 ms is holding something called  nill we  can see this by printing the value of ms like this:

// package main
// import (
// 	"fmt"
//  )
 
// func main() {
// 	var ms *mystruct 
// 	fmt.Println(ms) // this will print the value of ms which is nil because we have not assigned any value to ms yet.
// 	ms = new(mystruct) // this will create a pointer to a struct and assign it to ms. so when we run this code we get the output like this:
// 	fmt.Println(ms.foo)
// 	}
// type mystruct struct {
// 	foo int
// 	}
//  we get the nill value so the pointer that user doesnot initalize is nill and it is going to hold the nill value. 


// now to work with the underlying feild  we had to deference the ms pointer in order to fet that struct and then the user can change the feild like this:
// package main
// import (
// 	"fmt"
// 	 )

// func main() {
// 	var ms *mystruct 
// 	ms = new(mystruct) // this will create a pointer to a struct and assign it to ms. so when we run this code we get the output like this:
// 	(*ms).foo = 42 // this will assign the value of foo to 42. 
// 	fmt.Println((*ms).foo)
// 	}
// type mystruct struct {
// 	foo int
// 	}

// without using pointer we can also work on underlying feild like this:
// package main
// import (
// 	"fmt"
//     )


// func main() {
// 	var ms *mystruct 
// 	ms = new(mystruct) // this will create a pointer to a struct and assign it to ms. so when we run this code we get the output like this:
// 	ms.foo = 42 // this will assign the value of foo to 42.
// 	fmt.Println(ms.foo)
// 	}
// type mystruct struct {
// 	foo int
// 	}

// pointer handling the operator that are assigned one to another

// package main
// import (
// 	"fmt"
// 	)

// func main() {
// 	a := [3]int{1,2,3}
// 	b := a
// 	a[1] = 42
// 	fmt.Println(a,b) // this will print [1 42 3] [1 2 3] because a and b are pointing to the same memory location. so when we change the value in a the value in b also gets changed.
// 	}

// if we do same code and do slice it will behave slightly differently

// package main
// import (
// 	"fmt"
// 	)

// func main() {
// 	a := []int{1,2,3}
// 	b := a
// 	a[1] = 42
// 	fmt.Println(a,b) 
// 	}

// in this both a and b both are changes  because slice is the underlying projection of an array and it doesnot contain data on itself 

// so when we change the value in a the value in b also gets changed. 

// map also doesn't contain data on itself. so when we change the value in a the value in b also gets changed.

package main
import (
	"fmt"
	)

func main() {
	a := map[string]string{"foo":  "bar", "baz": "buzz"}
	b := a
	fmt.Println(a,b) // this will print map[foo:bar baz:buzz] map[foo:bar baz:buzz] because a and b are pointing to the same memory location. so when we change the value in a the value in b also gets changed.
	a["foo"] ="quux"
	fmt.Println(a,b) // this will print map[foo:quux baz:buzz] map[foo:quux baz:buzz] because a and b are pointing to the same memory location. so when we change the value in a the value in b also gets changed.
	// so when we change the value in a the value in b also gets changed.
}
	
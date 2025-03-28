// there is only one loop in a go language which is for loop.
// we can use for loop in different ways. we can use for loop as a while loop and also as a for each loop.

// package main
// import(
// 	"fmt"
// )

// func main() {
// 	for i :=0 ; i<5;i++{
// 		fmt.Println(i)
// 	}
// }

// Note there is no , (comma) in go language which is used in other languages to separate the initialization and condition in for loop. in go language we use ; (semicolon) to separate the initialization and condition in for loop.
// we can also use for loop as a while loop in go language. in other languages we use while loop but in go language we can use for loop as a while loop.


// package main
// import (
//     "fmt"
// )

// func main() {
//     for i, j := 0, 1; i < 5; i, j = i+1, j+1 {        // loop through i and j times until we reach the condition variable that we set in the for loop. use use i+1 for j we didnot use j++ because we are using i as the condition variable.
//         fmt.Println(i, j)
//     }
// }



// package main
// import(
// 	"fmt"
// )

// func main() {
// 	for i :=0 ; i<5;i++{
// 		fmt.Println(i)
// 		if i%2==0 {
// 			i/=2	
// 		}else{
// 			i = 2*i+4
// 		}
// 	}
// }




// package main
// import(
// 	"fmt"
// )

// func main() {
// 	for i :=0 ; i<5;i++{
// 		fmt.Println(i)
// 	}
// }

//  the main difference between declaring the i in this was in previous way is that it is scoped to the main function like if you use i direclty in the loop then it is scoped to the for loop.
// 	// but in this case it is scoped to the main function. so we can use i outside the for loop.
// package main
// import(
// 	"fmt"
// )

// func main() {
// 	i:=0  // declaring i in main function scope.
// 	for ; i<5;{
// 		fmt.Println(i)
// 	}
// 	fmt.Println(i)
// }

// if we remove the increment value from the loop in go it will give an error and it will not compile. so we have to use the increment value in the loop. if we remove the increment value then it will give an error.
// if the user is using in local variables then the output will be 0000000000000000000000... and so on.

// we can decalre the increment value inside the loop
// package main
// import(
// 	"fmt"
// )

// func main() {
// 	for i :=0 ; i<5;{
// 		fmt.Println(i)
// 		i++ // incrementing the value of i inside the loop.
// 	}
// }
// this is the example of using the do wile loop using for loop

// when we do a code in loop and when it enter in the infinite loop then we can use the break statement to break the loop. and we can also use the continue statement to skip the current iteration of the loop and move to the next iteration of the loop.
// package main
// import(
// 	"fmt"
// )

// func main() {
// 	i:=0
// 	for  {
// 		fmt.Println(i)
// 		i++ // incrementing the value of i inside the loop.	
// 		if i == 5 {
// 			break // break the loop when i is equal to 5.
// 		}

// 	}
// }

// using continue statemnt to skip the current iteration of the loop and move to the next iteration of the loop.

// package main
// import(
// 	"fmt"
// )

// func main() {
// 	for i :=0 ; i<10;i++{
// 		if i%2==0 {
// 			continue // skip the current iteration of the loop when i is even.
// 		}
// 		fmt.Println(i) // print the value of i when i is odd.
// 	}
// }

// using nested loop
// package main
// import(
// 	"fmt"
// )
// func main(){
// 	for i,j:=0,1;i<5;i,j=i+1,j+1{
// 		fmt.Println(i*j)
// }
// }

// // it can be declared as:
// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	for i:=0;i<5;i++{
// 		for j:=0;j<5;j++{
// 			fmt.Println(i*j)
//         }
// 	}

// }
 
// using break statement inside nested loop
// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	for i:=0;i<5;i++{
// 		for j:=0;j<5;j++{
// 			fmt.Println(i*j)
// 			if i*j >=3{
// 				break // break the loop when i*j is greater than or equal to 3.
// 			}
//         }
// 	}

// }

// in above program we use break statement in the nested loop. Break will break the break the loop which it finds closest. in above code break find the j loop which it break and the i loop is not breaked it is reset and worked as usual.
// to break the outer loop we use the level to break the outer loop
// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	Loop:
// 		for i:=0;i<5;i++{
// 			for j:=0;j<5;j++{
// 				fmt.Println(i*j)
// 				if i*j >=3{
// 					break Loop // break the loop when i*j is greater than or equal to 3.
// 				}
// 			}
// 		}

// }

// working with collection of for loop
// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	s:= [] int{1,2,3}
// for k , v := range s{         //k v is key and value
// 	fmt.Println(k,v) // print the index and value of the slice.
// 	}
// }

// package main
// import(
// 	"fmt" //

// )
// func main(){
// 	s:= "hello world"
// for k , v := range s{         //k v is key and value
// 	fmt.Println(k,v) // print the index and value of the string.
// }

// }

// if user want only key then printing the the key will give the key but if user want only value then user can use the underscore _ to ignore the key and only print the value.
// package main
// import(
// 	"fmt" //
// )
// func main(){
// 	s:= "hello world"
// for _, v := range s{         //k v is key and value
// 	fmt.Println(v) // print the value of the string.
// }
// }
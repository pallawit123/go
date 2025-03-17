package main

import (
    "fmt"
    "reflect"
)
type Animal struct {
    Name   string `required max:"100"`
    Origin string
}
func main(){
    t := reflect.TypeOf(Animal{}) // make sure  it's a struct and not a pointer to a struct or pointer
    field, _ := t.FieldByName("Name") // make sure it's a struct and not a pointer to a struct or pointer
    fmt.Println(feild.Tag)
}
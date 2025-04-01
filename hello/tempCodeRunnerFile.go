package main
import (
	"fmt"
	)

func main() {
	sayMessage("Hello go!")

}
func sayMessage(msg string) {
	fmt.Println(msg)
}
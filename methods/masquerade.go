// You can assign an anonymous function to a variable and then use that variable like any other function,
// as shown in the following listing.

package main

import "fmt"

var f = func() {
	fmt.Println("Dress up for the masquerade")
}

func main() {
	f()
}

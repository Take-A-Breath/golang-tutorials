// The variable you declare can be in the scope of the package or within a function, as shown in the next listing.
package main

import "fmt"

func main() {
	f := func(message string) {
		fmt.Println(message)
	}
	f("Go to the masquerade")
}

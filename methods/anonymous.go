// You can even declare and invoke an anonymous function in one step, as shown in the following listing
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Functions anonymous")
	}()
}

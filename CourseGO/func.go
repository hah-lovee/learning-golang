package coursego

import "fmt"

func main() {

	defer func() {
		fmt.Println("я defer")
	}()
	foo()

}

func foo() {
	// stack!!
	defer func() {
		fmt.Println("я defer1")
	}()

	defer func() {
		fmt.Println("я defer2")
	}()
	defer func() {
		fmt.Println("я defer3")
	}()
}

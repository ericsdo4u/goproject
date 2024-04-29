

import "fmt"

var x string = "Hello world"

func main() {
	fmt.Println(x)

	const name string = "Solomon"
	fmt.Println(name)

}

func check() {
	fmt.Println("Enter a Number")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

}


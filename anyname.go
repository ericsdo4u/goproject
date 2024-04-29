
import "fmt"

func main() {
	fmt.Println("Welcome to Credit Limit calculator")
	for {
		var beginingBalance, totlaCharges, totalCredits, creditLimit float64
		fmt.Println("Enter beginning balance (or -1 to quit): ")
		fmt.Scanf("%f", &beginingBalance)
		if beginingBalance == -1 {
			break
		}
		fmt.Println("Enter total charges: ")
		fmt.Scanf("%f", &totlaCharges)

		fmt.Println("Enter total credits: ")
		fmt.Scanf("%f", &totalCredits)

		fmt.Println("Enter credit limit: ")
		fmt.Scanf("%f", &creditLimit)

		newBalance := beginingBalance + totlaCharges - totalCredits
		if newBalance > creditLimit {
			fmt.Printf("Credit limit exceeded New balance: $%.2f\n", newBalance)
		} else {
			fmt.Println("Credit limit not exceeded.")
		}
		fmt.Println("Tanks for Using our app. Goodbye")

	}

}


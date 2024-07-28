package expenseSplit

/*
import (
	"bufio"
	"math"
	"os"
	"strings"
)

var balance float64 = 0

type User struct {
	Id string
	// Name    string
	// Email   string
	// Mobile  string
	TotalPaid  float64
	TotalSpent float64
}

type ExpenseType int

const (
	EQUAL ExpenseType = iota
	EXACT
	PERCENT
)

type Expense struct {
	ExpenseId   int
	PaidBy      *User
	Amount      float64
	ExpenseType ExpenseType
	Splits      map[*User]float64
}

// func NewUser(id string) *User {
// 	return &User{
// 		Id: id,
// 		// Name:name,
// 		// Email: email,
// 		// Mobile: mobile,
// 		Owes: map[string]int,
// 	}
// }

func FormatDecimalsUpTo(round int, value float64) float64 {
	divisor := 1
	for i := 0; i < round; i++ {
		divisor *= 10
	}
	rounded := math.Round(value*float64(divisor)) / float64(divisor)
	return float64(rounded) // Ensure precise decimal representation
}

/*
func Show(userId string, users map[string]*User) {
	if strings.TrimSpace(userId) != "" {

		for key, val := range u.Expense {

		}
	}else{
		if balance !=0{
			for _,user := users{
				fmt.Printf("User %s owes ", user.Id)
				for key,val := range user.Owes{
					fmt.Printf("")
				}
			}
		}else{
			fmt.Println("No balances")
		}

	}

}
*/
/*

func parseInput(users map[string]*User) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if strings.Contains(input, "SHOW") {
			parts := strings.Split(input, " ")

			// if len(parts) > 1 {

			// } else {

			// }
		}

		if strings.Contains(input, "EXPENSE") {
			parts := strings.Split(input, " ")
			userIdPaid := parts[1]
			amount := parts[2]
			noUsers := parts[3]
			var usersList []string
			var mode string
			for i := 4; i < len(parts); i++ {
				switch strings.TrimSpace(parts[i]) {
				case "EQUAL", "PERCENT", "EXACT":
					mode = parts[i]
				default:
					usersList = append(usersList, parts[i])
				}
			}

		}

	}

}

func Init() {
	users := map[string]*User{
		"u1": NewUser("u1", "User1"),
		"u2": NewUser("u2", "User2"),
		"u3": NewUser("u3", "User3"),
		"u4": NewUser("u4", "user4"),
	}
	parseInput(users)
}

*/

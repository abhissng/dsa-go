package dsago

import "fmt"

func Init() []int {
	s := []int{5, 4, 9, 7, 3}
	targetValue := 11

	var values map[int]int

	for i := 0; i < len(s); i++ {
		complement = targetValue - s[i]
		if _, ok := values[complement]; ok {
			return []int{currentIndex, values[complement]}
		}
		values[s[i]] = i

	}
	return []int{-1, -1}

}

type Human interface {
	Eat()
}

type PersonA struct {
	Name   string
	Mobile int
}

type PersonB struct {
	Email string
}

func (p *PersonA) Eat() {
	fmt.Println("person eats A")
}

func (p *PersonB) Eat() {
	fmt.Println("person eats B")
}

func main() {
	var t Human

	var p PersonA
	p.Name = "A"

	t = p

	t.Eat()

	var p1 PersonB
	p.Email = "B"
	t = p1

	t.Eat()

}

// Shallow and deep copy

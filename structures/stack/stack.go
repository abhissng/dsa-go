package stack

import "fmt"

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]int, 0),
	}
}

func (s *Stack) Push(n int) {
	s.elements = append(s.elements, n)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true

}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true

}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
func (s *Stack) Traverse() {
	for i := len(s.elements) - 1; i >= 0; i-- {
		fmt.Println(s.elements[i])
	}
}

func Run() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(30)
	stack.Push(40)
	stack.Push(5)

	fmt.Println("Stack elements (top to bottom):")
	stack.Traverse()

	top, _ := stack.Peek()
	fmt.Printf("\nTop element: %d\n", top)

	popped, _ := stack.Pop()
	fmt.Printf("Popped element: %d\n", popped)

	fmt.Println("\nStack after pop (top to bottom):")
	stack.Traverse()

	isEmpty := stack.IsEmpty()
	fmt.Printf("\nIs stack empty? %v\n", isEmpty)

}

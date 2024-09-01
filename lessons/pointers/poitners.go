package pointers

import "fmt"

type Gorila struct {
	Atype string
	Legs  int
}
type Mouse struct {
	Atype string
	Legs  int
}

func (m Mouse) speak() {
	fmt.Println("chuu")
}
func (m Gorila) speak() {
	fmt.Println("ki ki ki")
}

type Animal interface {
	speak()
}

func add(value int) {
	fmt.Println("a fucntion")
}

func sub(value int) {
	fmt.Println("b fucntion")
}

func square(a *[3]int16) [3]int16 {
	for i := range a {
		a[i] = a[i] * a[i]
	}
	return *a
}
func PrintPointer() {

	var p1 *int = new(int)
	var i int = 10
	var fn func(int)
	fn = add

	fn(10)
	fn = sub
	fn(20)
	p1 = &i
	fmt.Printf("\nvalue %v", *p1)
	fmt.Printf("\naddre %p", p1)
	*p1 = *p1 + 1
	fmt.Printf("\n new value %v", *p1)
	var ar1 = [3]int16{1, 2, 3}
	var result [3]int16 = square(&ar1)
	fmt.Printf("\n the result : %v", result)
	fmt.Printf("\n the original : %v", ar1)
	fmt.Printf("\n the addr : %p  == %p", &ar1, &result)

	var p *int16 = &ar1[0]
	fmt.Printf("\n the addr : %v ", *p)
}

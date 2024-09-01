package hello

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Car struct {
	model  string
	number int
	price  float64
}

func (c Car) milstoneleLft() float64 {
	return c.price
}
func (e Car) printing() {
	fmt.Println(e)
}

func printMe(s string) (int, float32) {
	fmt.Println(s)
	return 1, 12.12
}

func intDivision(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	var result int = a / b
	var remender int = a % b
	return result, remender, err
}

func list() {

	intr := [...]int{1, 2, 34}
	fmt.Println((intr[1]))
	fmt.Println(&intr[0])
	fmt.Println(&intr[1])
	fmt.Println(&intr[2])

	var intSlice []int32 = []int32{4, 3, 4, 5}

	fmt.Printf("the lenght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nthe lenght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, intSlice...)
	fmt.Printf("\nthe lenght is %v with capacity %v", len(intSlice), cap(intSlice))

	var sl []int32 = make([]int32, 4)
	fmt.Printf("\nthe lenght is %v with capacity %v", len(sl), cap(sl))
	var mp2 map[string]int32 = map[string]int32{"priya": 23, "sarha": 32}
	fmt.Println(mp2)
	var age, ok = mp2["priya"]
	if ok {
		fmt.Println(age)

	} else {
		fmt.Println("key not exist")
	}
	delete(mp2, "priya")
	for name, age := range mp2 {
		fmt.Printf("\n key : %v value: %v", name, age)
	}
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for {
		if i > 10 {
			break
		}
		fmt.Print(i)
		i++
	}

}
func test() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation %v \n", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func stringExmaple() {

	var myString string = "Š🤣cret"
	var indexed = myString[1]
	fmt.Printf("%v , %T\n", indexed, indexed)
	for i, v := range []rune(myString) {
		fmt.Println(i, v)
	}
	var s []string = []string{"a", "b", "v", "c"}
	var catstr = ""
	for v := range s {
		catstr += s[v]
	}
	fmt.Println(catstr)

	var strbuilder strings.Builder
	for i := range s {
		strbuilder.WriteString(s[i])
	}
	fmt.Println(strbuilder.String())
}

func structExample() {

	var myCar = Car{model: "tesla", number: 12312, price: 30000}
	fmt.Println(myCar.model)
	a := myCar.milstoneleLft()
	myCar.printing()
	fmt.Println(a)
}

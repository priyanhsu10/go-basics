package channels

import (
	"fmt"
	"time"
)

func someFunc(i int) {
	fmt.Println(i)
}
func Exp4() {

	//f1()
	// f2()
	// f3()
	// f5()
	pipeline()
}

func f1() {
	go someFunc(1)
	go someFunc(2)
	go someFunc(3)
	fmt.Println("hi")
	time.Sleep(time.Second * 3)
}
func f2() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "data"
	}()
	go func() {
		ch2 <- "cow"
	}()

	select {
	case mch1 := <-ch1:
		fmt.Println(mch1)
	case mch2 := <-ch2:
		fmt.Println(mch2)
	}

}
func f3() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, v := range chars {
		select {
		case charChannel <- v:

		}
	}
	close(charChannel)
	for c := range charChannel {
		fmt.Println(c)
	}

}
func f4() {
	go func() {
		for {
			select {
			default:
				fmt.Println("do work")
			}
		}
	}()

	time.Sleep(time.Second * 3)
}

func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK !")
		}
	}
}
func f5() {
	done := make(chan bool)
	go dowork(done)
	time.Sleep(time.Second * 5)
	done <- true
	close(done)

}
func getchan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range nums {
			out <- i
		}
		close(out)
	}()
	return out
}

func sq(ch <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for v := range ch {
			out <- v * v
		}
		close(out)
	}()
	return out
}
func pipeline() {
	// create slice for number
	nums := []int{1, 2, 4, 5, 6}

	ch := getchan(nums)
	//squure it
	result := sq(ch)

	//print it
	for r := range result {
		fmt.Println(r)
	}

}

package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var MAX_CHICKET_PRICE float32 = 5
var MAX_TOFOO_PRICE float32 = 4
var wg = sync.WaitGroup{}

func Exp2() {

	var chickenChannel = make(chan string)
	var tofooChannel = make(chan string)
	var websites = []string{"walmart.com", "amazsonecom", "woholefood.com"}
	for i := range websites {
		wg.Add(1)
		go checkChickenPrices(websites[i], chickenChannel)
		wg.Add(1)
		go checkTofooPrices(websites[i], tofooChannel)
	}
	sendMessage(chickenChannel, tofooChannel)
	go func() {
		wg.Wait()
		close(chickenChannel)
		close(tofooChannel)
	}()

}
func sendMessage(chickenChannel chan string, tofooChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound deal on chicken on bewesite %s", website)
	case website := <-tofooChannel:
		fmt.Printf("\nFound deal on tofoo on bewesite %s", website)
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKET_PRICE {
			chickenChannel <- website

		}
	}
}
func checkTofooPrices(website string, tofooChannel chan string) {
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		var tofooPrice = rand.Float32() * 30
		if tofooPrice <= MAX_TOFOO_PRICE {
			tofooChannel <- website
			break
		}
	}
}

func Exp() {

	var c = make(chan int, 5)
	go process(c)
	for i := range c {

		fmt.Printf(" \nreading from channel %v", i)
		time.Sleep(time.Second * 1)
		break
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("\nputting i in channel %v", i)
		c <- i
	}
	fmt.Println("\nexit process function")
}

func Exmp3() {

	wg := sync.WaitGroup{}

	ch := make(chan int, 4)
	wg.Add(4)
	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)

	for i := 0; i < 30; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
func reader(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("channel is close")
			return
		}
		fmt.Printf("Reader %v recieve %d \n", i, val)
	}

}

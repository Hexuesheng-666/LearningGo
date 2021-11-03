package main

import (
	"fmt"
	"time"
)

// ch <-,<- ch, close(ch)
// buffer
// range
// wait group
// Select

//var wg sync.WaitGroup
//
//func square(ch chan int, x int){
//	defer wg.Done()
//	ch <- x * x
//}
//
//func main(){
//	ch := make(chan int, 2)
//	arr := []int{1,2}
//	for _,num := range arr{
//		wg.Add(1)
//		go square(ch,num)
//	}
//	wg.Wait()
//	close(ch)
//	//go func() {
//	//	for i :=0; i < 10; i++{
//	//		ch <- i * i
//	//	}
//	//	close(ch)
//	//}()
//	for val := range ch{
//		fmt.Println(val)
//	}
//	time.Sleep(1 * time.Second)
//
//}

func foo1(ch chan string){
	time.Sleep(1 * time.Second)
	ch <- "success"
}

func main(){
	ch := make(chan string)
	go foo1(ch)
	select {
	case val := <- ch :
		fmt.Println(val)
		case <- time.After(1 * time.Second):
			fmt.Println("Time Out")
	}
}
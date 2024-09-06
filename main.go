package main

import (
	"fmt"
	//"sync"
	//"time"
)

func main() {

	var ch chan string = make(chan string)

	for i :=1;i<=5;i++{
		go func(){
			ch <- "welcome"
			close(ch)
		}()
	}
	for val := range ch{
		fmt.Println(val)
	}
	fmt.Println("completed")








// 	var gm sync.WaitGroup
// 	gm.Add(1)

// ch1 := make(chan int)
// 	go func(gm *sync.WaitGroup){
// 		defer gm.Done()
// 		ch1 <- 10
// 		close(ch1)



// 	}(&gm)
// 	gm.Add(1)
// 	go func(gm *sync.WaitGroup){
// 		defer gm.Done()

// 		ch1 <- 20
// 		//close(ch1)

// 	}(&gm)
// 	for val := range ch1{
// 		fmt.Println(val)
// 		//close(ch1)
// 	}

	//val := <- ch1
	//fmt.Println(val)
	//gm.Wait()
	fmt.Println("completed")




	//Empty select
	// select{}
	// fmt.Println("Empty select")


//  var ch1 chan int = make(chan int)
//  var ch2 chan int = make(chan int)

//  go func(){
// 	ch1 <- 10
// 	fmt.Println("sending data ")
//  }()
//  go func(){
// 	val := <- ch2
// 	fmt.Println("Data received ",val)
//  }()
//  select{
//  case <- ch1: fmt.Println("Data received")

//  case ch2 <- 10 :
// 	fmt.Println("data send ch2")

//  case val := <- time.After(time.Second *2):
// 	fmt.Println(val)

	// default : fmt.Println("default")
 }

	//Receive from channel
	//var ch1 chan int = make(chan int)
	// ch2 := make(chan int)

	// go func(){
	// 	ch1 <- 10
	// 	fmt.Println("sending data ch1")
	// }()
	
		
	
	// go func()  {
	// 	ch2 <- 20
	// 	fmt.Println("send data ch2")
	// 	}()


	// select {
	// case val1 := <- ch1:
	// 	{
	// 		fmt.Println("sending a data",val1)
	// 	}
	// case ch1 <- 20:{
	// 	fmt.Println("sending data")
	// } 
	// default : fmt.Println("default")
	// }


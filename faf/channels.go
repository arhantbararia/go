package main

import "fmt"

//channel actss as a input/output queues to distribute data in go routines

func main() {
	//1
	// var c = make(chan int)
	// c <- 1 //this a unbufferred channel (only space for 1). c : [1]
	// // the standard go routine will wait till somefunction reads from this channel.
	// // here it will wait forever and return a deadlock error!
	// var i = <-c //here '1' gets popped out and assigned to i and channel is now empty
	// fmt.Println(i)

	//2
	// var c = make(chan int)
	// go process(c)
	// fmt.Println(<-c)
	// //fmt.Println(<-c) //raises deadlock! reading from an empty channel
	// go process(c)
	// fmt.Println(<-c)

	//3
	// var c = make(chan int)
	// go goprocess(c)
	// for i := 0; i < 5; i++ { //if the number of inserts in channel == outserts , no deadlock error
	// 	fmt.Println(<-c)
	// }

	//4
	// var c = make(chan int)
	// go goprocess(c)
	// for i := range c { //now since c's range is now not fixed. range will wait for next value in empty channel! Raises Deadlock!
	// 	fmt.Println(i)
	// }

	

}

func process(c chan int) {
	c <- 123
}

func goprocess(c chan int) {
	//set defer processes are executed just before the function exit.
	defer close(c) //this will close the channel same as make function. listening functions will stop listening to this channel
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting GoProcess")
	
}

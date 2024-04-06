package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"1", "2", "3", "4"}

var results = []string{}

func main() {
	to := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1) //2. this will increment group member by 1 adding the thread(subtask) to the group
		// go dbCall(i)
		// go timed() // to complete all timed() goroutines is independent of number of threads(or calls) and constant
		go iterative() // time to complete all iterative() goroutines is directly proportional to the number of threads
		// this is dependent of resource use. time.Sleep doesn't use any resources. While iterative's for loop do use CPU time.

		//1.  using go alone will put all threads in background and will not wait for them to finish
		// here it will just exit, therefore we need to define some breakpoints, between which
		// execution of all subtasks must be completed

	}

	wg.Wait() //4. this is the endpoint before which waitgroup member count must be zero
	fmt.Printf("\n Total Execution time: %v", time.Since(to))
	fmt.Printf("\nThe results are: %v", results)

}

func timed() {
	time.Sleep(time.Duration(float32(2000)) * time.Millisecond)
	wg.Done()

}

func iterative() {
	var res int32
	t := time.Now()

	for i := 0; i <= 10000000; i++ {
		res += 1
	}
	fmt.Printf("time taken by the loop number is: %v \n", time.Since(t))
	wg.Done()

}

func dbCall(i int) {
	t := time.Now()
	// var delay float32 = float32((5 - i) * 2000)
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\n execution time for query: %v is %v", i, time.Since(t))

	fmt.Println("The results from the database is: ", dbData[i])
	// 5. Now this is a problem, multiple subtasks
	//	running concurrently and accessing and modifying
	//	the same data (give inconsistent results),
	//  We need to define mutex. to control writes to data, TO make it safe

	// 6. from mutex we use lock and unlock before and after the resource
	mtx.Lock()
	results = append(results, dbData[i])
	mtx.Unlock()

	wg.Done() //3. this will decrement the counter.

}

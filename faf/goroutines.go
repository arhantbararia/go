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
	for i := 0; i < len(dbData); i++ {
		wg.Add(1) //2. this will increment group member by 1 adding the thread(subtask) to the group
		go dbCall(i)
		//1.  using go alone will put all threads in background and will not wait for them to finish
		// here it will just exit, therefore we need to define some breakpoints, between which
		// execution of all subtasks must be completed

	}

	wg.Wait() //4. this is the endpoint before which waitgroup member count must be zero
	fmt.Printf("\n Total Execution time: %v", time.Since(to))
	fmt.Printf("\nThe results are: %v", results)

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

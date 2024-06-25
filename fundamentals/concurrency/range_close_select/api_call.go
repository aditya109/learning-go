//package main
//
//import (
//	"errors"
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//type Result struct {
//	Data interface{}
//	Err  error
//}
//
//func fetchData(id int, resultChan chan<- Result, quit <-chan bool) {
//	for {
//		select {
//		case <-quit:
//			fmt.Println("API call forcefully stopped\n", id)
//			return
//		default:
//			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
//
//			if rand.Float32() < 0.2 {
//				resultChan <- Result{
//					Err: errors.New("fetch Data Error"),
//				}
//			} else {
//				resultChan <- Result{
//					Data: fmt.Sprintf("%d", id),
//				}
//			}
//		}
//	}
//}
//
//func main() {
//	resultChan := make(chan Result)
//	quit := make(chan bool)
//
//	for i := 0; i < 10; i++ {
//		go fetchData(i+1, resultChan, quit)
//	}
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			result := <-resultChan
//			if result.Err != nil {
//				fmt.Printf("Error fetching data: %s\n", result.Err)
//			} else {
//				fmt.Printf("Data fetched for id: %d\n", result.Data)
//			}
//		}
//		close(quit)
//	}()
//
//}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Result struct to hold either data or an error from the API call
type Result struct {
	Data string
	Err  error
}

// fetchData simulates an API call that takes some time to complete and may produce an error
func fetchData(id int, resultChan chan<- Result, quit <-chan bool) {
	for {
		select {
		case <-quit:
			fmt.Printf("API call %d stopped\n", id)
			return
		default:
			// Simulate a network call with random delay
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

			// Simulate random error occurrence
			if rand.Float32() < 0.2 { // 20% chance of error
				resultChan <- Result{Err: fmt.Errorf("API call %d failed", id)}
			} else {
				resultChan <- Result{Data: fmt.Sprintf("Data from API call %d", id)}
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	resultChan := make(chan Result)
	quit := make(chan bool)

	// Start 3 goroutines simulating concurrent API calls
	for i := 1; i <= 3; i++ {
		go fetchData(i, resultChan, quit)
	}

	// Collect data and errors from the resultChan and print them
	go func() {
		for i := 0; i < 10; i++ {
			result := <-resultChan
			if result.Err != nil {
				fmt.Println("Error:", result.Err)
			} else {
				fmt.Println("Success:", result.Data)
			}
		}
		// Signal all fetchData goroutines to stop
		close(quit)
	}()

	// Give some time for goroutines to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Program finished")
}

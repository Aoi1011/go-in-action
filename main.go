package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create unbuffered channel
	court := make(chan int)
	// Add a count of 2, one for goroutine
	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg.Wait()

	// allocate 1 logical processor for the scheduler to use
	// runtime.GOMAXPROCS(1)

	// wg.Add(2)

	// fmt.Println("Start Gorountines")

	// go printPrime("A")
	// go printPrime("B")

	// go func() {
	// 	// Schedule the call to Done to tell main we are done
	// 	defer wg.Done()

	// 	// Display the alphabet 3 times
	// 	for count := 0; count < 3; count++ {
	// 		for char := 'A'; char < 'A'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}

	// }()

	// go func() {
	// 	// Schedule the call to Done to tell main we are done
	// 	defer wg.Done()

	// 	// Display the alphabet 3 times
	// 	for count := 0; count < 3; count++ {
	// 		for char := 'a'; char < 'a'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}

	// }()

	// Wait for the gorountines to finish
	// fmt.Println("Waiting to finish")
	// wg.Wait()

	// fmt.Println("\nTerminating Program")
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s win\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}

// printPrime displays prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)

	}
	fmt.Println("Completed", prefix)
}

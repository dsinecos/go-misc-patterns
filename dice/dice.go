// Package dice provides functionality to Roll a dice and retrieve a value
package dice

import "fmt"

// Throw TODO
func Throw(out chan<- int) {

	/*
	 Shutdown channel is used to exit all the goroutines once
	 one of the goroutines publishes a dice value
	*/
	shutdown := make(chan int)

	/*
		For loop is used to initialize 6 goroutines to publish one
		of the dice's face values. Refer [Asynchronous execution inside for loops (Golang and Nodejs)](https://dsinecos.github.io/blog/Asynchronous-execution-inside-for-loops)
	*/
	for i := 1; i <= 6; i++ {

		/*
			Goroutine executes an anonymous function that writes one of the
			six values of the dice to an output channel
		*/
		go func(i int) {

			defer (func() {
				fmt.Println("Exiting goroutine")
			})()

			/*
				The select statement is used to ensure that a goroutine either publishes a dice
				value to the output channel or shuts down if another goroutine has already
				published a value
			*/
			select {
			case out <- i:
				close(shutdown)
			case <-shutdown: // Since reading from a closed channel is non-blocking, this allows to close the goroutine when a value has been published to the output channel and `shutdown` is closed
				return
			}
		}(i)
	}
}

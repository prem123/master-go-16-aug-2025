package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("Goodbye! Thanks for using the timer.") // defer demo

	fmt.Println(" Welcome to GoTimer - Countdown Timer App!")
	fmt.Println("==========================================")

	for {
		var seconds int
		fmt.Print("\nEnter number of seconds for countdown (0 to exit): ")
		fmt.Scanln(&seconds)

		if seconds <= 0 {
			return // exit program
		}

		// run the countdown
		runCountdown(seconds, alertMessage)
	}
}

// runCountdown handles the countdown logic
func runCountdown(seconds int, alert func()) {
	fmt.Printf("\n🚀 Starting countdown from %d seconds...\n", seconds)
	
	for i := seconds; i > 0; i-- {
		fmt.Printf("%d\n", i)
		time.Sleep(1 * time.Second)
	}
	
	// call the passed function
	alert()
}

// alertMessage is the function to notify when timer ends
func alertMessage() {
	fmt.Println("Time's up!")
	fmt.Println("Timer completed successfully!")
}

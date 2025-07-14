package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func logWithContext(ctx context.Context, message string) {
	requestIdVal := ctx.Value("requestID")
	log.Printf("Request ID: %v - %v", requestIdVal, message)
}

// === CONTEXT WITH CANCEL
func ContextRun3() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	
	go func(){
		time.Sleep(2 * time.Second) // simulating a heavy task
		cancel() // manually cancalling only after the task is finished
	}()

	ctx = context.WithValue(ctx, "requestID", "qwehergd12343231")
	ctx = context.WithValue(ctx, "name", "Bekarys")
	ctx = context.WithValue(ctx, "IP", "dnd.12.3141.56")
	go doWorkWithContext(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}	

	logWithContext(ctx, "This is a test log message")
}

func doWorkWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// === CONTEXT WITH TIMEOUT
func ContextRun2() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second) // timer of the ctx starts here
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "qwehergd12343231")
	go doWorkWithContext(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}
}

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation cancelled"
	default:
		if num&1 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func ContextRun1() {
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result from timeout context:", result)

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result after timeout", result)
}

// === DIFFERENCE BETWEEN CONTEXT.TODO AND CONTEXT.BACKGROUND
func ContextTest() {
	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "name", "Bekarys")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx1 := context.WithValue(contextBkg, "city", "Almaty")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("city"))
}

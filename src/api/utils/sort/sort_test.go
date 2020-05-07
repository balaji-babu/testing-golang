package sort

import (
	"fmt"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	// Init
	values := []int{32, 54, 67, 87, 91, 07, 65, 32}

	// Execution
	timeoutChan := make(chan bool, 1)

	go func() {
		BubbleSort(values)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		fmt.Println("It took more than 500 second for Bubble sort")
		return
	}

	// Validation
	if values[0] != 7 {
		t.Error("Something went wrong")
	}

	if values[len(values)-1] != 91 {
		t.Error("Something went wrong")
	}

}

func TestSort(t *testing.T) {
	// Init
	values := []int{32, 54, 67, 87, 91, 07, 65, 32}

	// Execution
	Sort(values)

	// Validation
	if values[0] != 7 {
		t.Error("Something went wrong")
	}

	if values[len(values)-1] != 91 {
		t.Error("Something went wrong")
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	values := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}

}

func BenchmarkSort(b *testing.B) {
	values := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}

}

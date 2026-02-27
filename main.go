package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE              = 100_000_000
	CHUNKS            = 8
	MAX_RANDOM_NUMBER = 1000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = rand.Intn(MAX_RANDOM_NUMBER)
	}
	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return -1
	}
	size := len(data)
	maxValue := data[0]
	for i := 1; i < size; i++ {
		if maxValue < data[i] {
			maxValue = data[i]
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	maximums := []int{}
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		sectionLength := len(data) / CHUNKS

		startSectionIdx := i * sectionLength
		endSectionIdx := startSectionIdx + sectionLength

		wg.Add(1)
		go func() {
			defer wg.Done()
			maxValue := maximum(data[startSectionIdx:endSectionIdx])
			maximums = append(maximums, maxValue)
		}()
	}

	wg.Wait()
	return maximum(maximums)
}

func main() {
	var (
		max          int
		startMoment  time.Time
		finishMoment time.Time
		elapsed      int64
	)

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	// one thread
	fmt.Println("Ищем максимальное значение в один поток")
	startMoment = time.Now()
	max = maximum(elements)
	finishMoment = time.Now()

	elapsed = finishMoment.Sub(startMoment).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	// multithreading
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMoment = time.Now()
	max = maxChunks(elements)
	finishMoment = time.Now()

	elapsed = finishMoment.Sub(startMoment).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

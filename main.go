package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	rand.Seed(time.Now().UnixNano())

	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = 1 + rand.Intn(2147483647) // Взял край диапазона
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) <= 0 {
		return 0
	}
	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) <= 0 {
		return 0
	}

	var wg sync.WaitGroup
	localMaxs := make([]int, CHUNKS)

	// Находим длину чанка с округлением вверх
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func(start, end int) {
			defer wg.Done()

			if start >= len(data) || end - start <= 0 {
				return
			}
			if end >= len(data) {
				end = len(data)
			}
	
			localMaxs[i] = data[start]

			for j := start + 1; j < end; j++ {
				if data[j] > localMaxs[i] {
					localMaxs[i] = data[j]
				}
			}
		}(i*chunkSize, (i+1)*chunkSize)
	}

	wg.Wait()

	// В качестве финального подсчёта воспользуемся уже реализованной функцией maximum
	return maximum(localMaxs)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	
	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

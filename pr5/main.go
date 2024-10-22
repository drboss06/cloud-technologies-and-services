package main

import (
	"fmt"
	"sync"
)

// Функция для подсчета вхождений символа b в часть строки
func countOccurrences(s string, b byte, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			count++
		}
	}
	ch <- count
}

func main() {
	// Исходные данные
	C := "abacabadabacaba"
	b := byte('a')

	// Количество потоков
	numGoroutines := 4

	// Длина каждой части строки
	partLen := len(C) / numGoroutines

	wg := &sync.WaitGroup{}
	ch := make(chan int, numGoroutines)

	// Запуск горутин
	for i := 0; i < numGoroutines; i++ {
		start := i * partLen
		end := start + partLen

		// Для последней горутины захватим остаток строки
		if i == numGoroutines-1 {
			end = len(C)
		}

		wg.Add(1)
		go countOccurrences(C[start:end], b, wg, ch)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
	close(ch)

	// Подсчет общего количества вхождений
	totalCount := 0
	for count := range ch {
		totalCount += count
	}

	fmt.Printf("Количество вхождений символа '%c' в строке: %d\n", b, totalCount)
}

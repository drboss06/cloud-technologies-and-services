package main

import (
	"fmt"
	"sync"
)

// Размер матрицы 3x3
const N = 3

var (
	A = [N][N]int{{1, -2, 0}, {4, 6, 2}, {-3, 4, -2}}
	B = [N][N]int{{0, 2, 0}, {1, 1, 1}, {5, -3, 10}}
	C [N][N]int // результирующая матрица
)

// Структура элемента очереди
type Result struct {
	Message string
	Next    *Result
}

var (
	head    *Result // указатель на первый элемент очереди
	current *Result // указатель на текущий последний элемент очереди
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

// Функция для выполнения потоков
func calculateRow(p int) {
	defer wg.Done()
	for i := 0; i < N; i++ {
		C[p][i] = 0
		for j := 0; j < N; j++ {
			C[p][i] += A[p][j] * B[j][i]
		}

		// Защита критической секции
		mutex.Lock()
		newResult := &Result{
			Message: fmt.Sprintf("Thread %d: Calculated element [%d][%d] = %d", p, p, i, C[p][i]),
		}

		if head == nil {
			head = newResult
			current = newResult
		} else {
			current.Next = newResult
			current = newResult
		}
		mutex.Unlock()
	}
}

// Функция для вывода очереди
func printResults() {
	fmt.Println("Results queue:")
	for r := head; r != nil; r = r.Next {
		fmt.Println(r.Message)
	}
}

func main() {
	// Инициализация мьютекса
	wg.Add(N)

	// Создание потоков для вычисления каждой строки
	for p := 0; p < N; p++ {
		go calculateRow(p)
	}

	// Ожидание завершения всех потоков
	wg.Wait()

	// Вывод результатов
	printResults()
}

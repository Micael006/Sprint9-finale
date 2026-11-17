package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	// Так как функция не подразумевает передачу фиксированной случайности, 
	// проверяю только длину и принадлеждность элементов диапазону
	const maxVal int = 2147483647

	tests := []struct {
		name string
		size int
	}{
		{"positive small", 5},
		{"positive large", 1000},
		{"zero", 0},
		{"negative", -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			// Ожидаемая длина слайса
			expectedLen := 0
			if tt.size > 0 {
				expectedLen = tt.size
			}

			// Проверка на совпадение длины
			require.Equal(t, expectedLen, len(data),
					"для size=%d: ожидалась длина %d, получено %d", tt.size, expectedLen, len(data))
			
			// Проверяем, что все сгенерированные элементы не выходят за пределы диапазона
			for i, v := range data {
				assert.True(t, v >= 1 && v < maxVal,
						"индекс %d: значение %d вне диапазона [1, %d]", i, v, maxVal)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		answer int
	}{
		{"zero", []int{}, 0},
		{"one", []int{1}, 1},
		{"ten", []int{2, 3, 4, 5, 6, 7, 8, 1, 9, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)

			// Проверка на совпадение ответов
			require.Equal(t, tt.answer, result,
					"для теста %s: ожидался максимум %d, получен %d", tt.name, tt.answer, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		answer int
	}{
		{"zero", []int{}, 0},
		{"one", []int{1}, 1},
		{"eight", []int{8, 8, 8, 8, 8, 8, 8, 8}, 8},
		{"ten", []int{2, 3, 4, 5, 6, 7, 8, 1, 9, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.data)

			// Проверка на совпадение ответов
			require.Equal(t, tt.answer, result,
					"для теста %s: ожидался максимум %d, получен %d", tt.name, tt.answer, result)
		})
	}
}
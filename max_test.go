package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 5, 7, 10, 11, 59, 93, 21, 54, 12, -12, -1},
			expected: 93,
		},
		{
			numbers:  []int{1, 2, 5, 7, 10, 11, 59, 96, 21, 54, 12, -12, -1},
			expected: 96,
		},
		{
			numbers:  []int{1, 2, 5, 7, 10, 11, 59, 21, 54, 12, -12, -1},
			expected: 59,
		},
	}

	for _, testCase := range testTable {

		result := max(testCase.numbers)

		t.Logf("Вызвали максимальное(%v) результат %d\n", testCase.numbers,
			result)

		if result != testCase.expected {
			t.Errorf("В коде есть ошибка! Ожидалось %d, Получили %d",
				testCase.expected, result)
		}
	}
}

func TestMaxIndex(t *testing.T) {
	numbers := []int{1, 2, 5, 7, 10, 11, 59, 93, 21, 54, 12, -12, -1}
	expect := 7

	result := maxIndex(numbers)

	if result != expect {
		t.Errorf("В коде есть ошибка! Ожидалось %d, Получили %d",
			expect, result)
	}
}

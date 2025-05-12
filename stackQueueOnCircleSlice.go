package main

import (
	"errors"
	"fmt"
)

// QueueOnCircleSlice - очередь на базе
type QueueOnCircleSlice struct {
	items      []int // слайс на безе которого строится очередь
	size       int   // кол-во значений в очереди
	head, tail int   // индексы головы и хвоста очереди
}

// NewQueueOnCircleSlice - создает очередь
func NewQueueOnCircleSlice() *QueueOnCircleSlice {
	queue := &QueueOnCircleSlice{}
	queue.items = make([]int, 2)
	return queue
}

// Print - возвращает строковое представление очереди
func (q *QueueOnCircleSlice) Print() string {
	if q.size == 0 {
		return ""
	}
	result := ""
	if q.head < q.tail { // случай, когда голова слева, хвост справа
		for i := q.head; i < q.tail; i++ {
			result += fmt.Sprintf("%d ", q.items[i])
		}
	} else { // случай, когда голова справа, хвост слева
		// добавляем в результат элементы, начиная с головы, до конца слайса
		for i := q.head; i < len(q.items); i++ {
			result += fmt.Sprintf("%d ", q.items[i])
		}
		// добавляем в результат элементы, начиная с начала слайса, до хвоста
		for i := 0; i < q.tail; i++ {
			result += fmt.Sprintf("%d ", q.items[i])
		}
	}
	return result
}

// Empty - возвращает признак, пустой ли слайс
func (q *QueueOnCircleSlice) Empty() bool {
	return q.size == 0
}

// Push - добавляет элемент в очередь
func (q *QueueOnCircleSlice) Push(item int) {
	// если кол-во значений в очереди равно кол-ву выделенного места
	if q.size == cap(q.items) {
		newItems := make([]int, cap(q.items)*2) // создаем новый слайс вдвое больше текущего
		// если голова очереди слева и хвост справа, другими словами если мы добавляем в конец слайса
		if q.head < q.tail {
			copy(newItems, q.items[q.head:q.tail]) // копируем в новый слайс, все значения из старого начиная (включая) с индекса q.head и до индекса q.tail (не включая)
		} else { // если хвост слева а голова справа
			copy(newItems, q.items[q.head:])                       // копируем в новый слайс из старого с q.head и до конца
			copy(newItems[len(q.items)-q.head:], q.items[:q.tail]) // копируем в новый слайс, в начиная со следующей пустой позиции после предыдущего копирования, все значения из старого слайса с 0 позиции и до q.tail
		}

		q.items = newItems
		q.head = 0
		q.tail = q.size
	}
	q.items[q.tail] = item // добавляем новое значение
	q.tail = (q.tail + 1) % cap(q.items)
	q.size++
}

// Pop - возвращает первый элемент очереди с удалением
func (q *QueueOnCircleSlice) Pop() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}
	removed := q.items[q.head]
	q.head = (q.head + 1) % cap(q.items)
	q.size--
	return removed, nil
}

// Peek - возвращает первый элемент очереди без удаления
func (q *QueueOnCircleSlice) Peek() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.items[q.head], nil
}

// Clear - очищает очередь
func (q *QueueOnCircleSlice) Clear() {
	q.head = 0
	q.tail = 0
	q.size = 0
}

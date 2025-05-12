package main

import (
	"container/list"
	"fmt"
)

// QueueOnLinkedList - очередь на основе связного списка
type QueueOnLinkedList struct {
	items *list.List
}

// NewQueueOnLinkedList - создает пустую очередь
func NewQueueOnLinkedList() *QueueOnLinkedList {
	return &QueueOnLinkedList{
		items: list.New(),
	}
}

// Print - возвращает строковое представление очереди
func (q *QueueOnLinkedList) Print() string {
	result := ""
	for e := q.items.Front(); e != nil; e = e.Next() {
		result += fmt.Sprintf("%d ", e.Value.(int))
	}
	return result
}

// Empty - возвращает признак пустая ли очередь
func (q *QueueOnLinkedList) Empty() bool {
	return q.items.Len() == 0
}

// Push - добавляет значение в очередь
func (q *QueueOnLinkedList) Push(item int) {
	q.items.PushBack(item)
}

// Pop - извлекаем значение из очереди
func (q *QueueOnLinkedList) Pop() (int, error) {
	if q.items.Len() == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	firstItem := q.items.Front()
	q.items.Remove(firstItem)
	return firstItem.Value.(int), nil
}

// Peek - возвращает голову очереди без извлечения
func (q *QueueOnLinkedList) Peek() (int, error) {
	if q.items.Len() == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items.Front().Value.(int), nil
}

// Clear - очищает очередь
func (q *QueueOnLinkedList) Clear() {
	q.items.Init()
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root.next = &l.root // голове списка присваиваем root-овый элемент списка
	l.root.prev = &l.root // хвосту списка присваиваем root-овый элемент списка
	l.len = 0             // сбрасываем размер списка в 0
	return l
}

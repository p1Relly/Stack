package main

import (
	"container/list"
	"errors"
	"fmt"
)

// StackOnLinkedList - стек на базе списка
type StackOnLinkedList struct {
	items *list.List
}

// NewStackOnLinkedList - создает стек
func NewStackOnLinkedList() *StackOnLinkedList {
	return &StackOnLinkedList{items: list.New()}
}

// Print - возвращает строковое представление списка
func (s *StackOnLinkedList) Print() string {
	var result string
	for e := s.items.Back(); e != nil; e = e.Prev() {
		result += fmt.Sprintf("%v ", e.Value)
	}
	return result
}

// Empty - проверяет пустой ли стек
func (s *StackOnLinkedList) Empty() bool {
	return s.items.Len() == 0
}

// Push - добавляет значение в стек
func (s *StackOnLinkedList) Push(item int) {
	s.items.PushBack(item)
}

// Pop - выталкивает вершину стека
func (s *StackOnLinkedList) Pop() (int, error) {
	if s.items.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	lastElement := s.items.Back()
	s.items.Remove(lastElement)
	return lastElement.Value.(int), nil
}

// Peek - возвращает вершину стека
func (s *StackOnLinkedList) Peek() (int, error) {
	if s.items.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items.Back().Value.(int), nil
}

// Clear - очищает стек
func (s *StackOnLinkedList) Clear() {
	s.items.Init()
}

func main() {
	stack := NewStackOnLinkedList()

	stack.Push(3)              // добавить 3
	stack.Push(4)              // добавить 4
	stack.Push(6)              // добавить 6
	fmt.Println(stack.Print()) // выводит 6 4 3

	top, err := stack.Peek() // получить последний элемент
	fmt.Println(err == nil)  // выводит true
	fmt.Println(top)         // выводит 6

	deletedItem, err := stack.Pop() // удалить последний элемент
	fmt.Println(err == nil)         // true
	fmt.Println(deletedItem)        // выводит 6
	fmt.Println(stack.Print())      // выводит 4 3

	stack.Clear()
	fmt.Println(stack.Empty()) // выводит true
	fmt.Println(stack.Print()) // выводит пустую строку
}

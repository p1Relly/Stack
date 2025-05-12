package main

import (
	"errors"
	"fmt"
)

// StackOnSlice - стек на базе slice-а
type StackOnSlice struct {
	items []int // элементы стека
	size  int   // размер стека
}

// NewStackOnSlice - конструктор стека
func NewStackOnSlice() *StackOnSlice {
	return &StackOnSlice{
		items: make([]int, 0), // в момент создания стек пустой
		size:  0,              // в момент создания размер стека равен 0
	}
}

// Print - возвращает строковое представление стека
func (s *StackOnSlice) Print() string {
	result := ""
	for i := s.size - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d ", s.items[i])
	}
	return result
}

// Empty - возвращает признак пустой ли стек
func (s *StackOnSlice) Empty() bool {
	return s.size == 0
}

// increaseSlice - расширяет слайс
func (s *StackOnSlice) increaseSlice() {
	newCount := s.size * 2 // в newCount кладем новый размер слайса
	if s.size == 0 {       // если текущий размер 0, то по умолчанию задаем размер 4
		newCount = 4
	}
	newSlice := make([]int, s.size, newCount) // создаем новый слайс
	copy(newSlice, s.items)                   // копируем элементы из старого слайса в новый
	s.items = newSlice                        // обновляем ссылку на элементы стека новым слайсом
}

// Push - добавляет новый элемент
func (s *StackOnSlice) Push(item int) {
	if s.size == len(s.items) { // если размер стека равен кол-ву узлов слайса, на базе которого организован стек, то расширяем слайс
		s.increaseSlice()
	}

	s.items[s.size] = item // добавляем новый элемент
	s.size++               // увиличиваем размер
}

// Pop - удаляет элемент из стека
func (s *StackOnSlice) Pop() (int, error) {
	// если стек пустой возвращаем ошибку
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}
	s.size--                    // уменьшаем размер стека не единицу, косвенно сдвигаем вершину стека вниз
	return s.items[s.size], nil // возвращаем вершинку стека
}

// Peek - возвращает вершину стека
func (s *StackOnSlice) Peek() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.items[s.size-1], nil
}

// Clear - очищает стек
func (s *StackOnSlice) Clear() {
	s.size = 0
}

func main() {
	stack := NewStackOnSlice()

	stack.Push(3)              // добавить 3
	stack.Push(4)              // добавить 4
	stack.Push(6)              // добавить 6
	fmt.Println(stack.Print()) // выводит 6 4 3

	top, err := stack.Peek() // получить последний элемент
	fmt.Println(err == nil)  // выводит true
	fmt.Println(top)         // выводит 6

	deletedItem, err := stack.Pop() // удалить последний элемент
	fmt.Println(err == nil)         // выводит true
	fmt.Println(deletedItem)        // выводит 6
	fmt.Println(stack.Print())      // выводит 4 3

	deletedItem, err = stack.Pop() // удалить последний элемент
	fmt.Println(err == nil)        // выводит true
	fmt.Println(deletedItem)       // выводит 4
	fmt.Println(stack.Print())     // выводит 3

	stack.Clear()
	fmt.Println(stack.Empty()) // выводит true
	fmt.Println(stack.Print()) // выводит пустую строку

	deletedItem, err = stack.Pop() // удалить последний элемент
	fmt.Println(err == nil)        // выводит true
	fmt.Println(deletedItem)       // выводит 3
	fmt.Println(stack.Print())     // выводит пустую строку

	deletedItem, err = stack.Pop() // удалить последний элемент
	fmt.Println(err == nil)        // выводит false
	fmt.Println(err)               // выводит stack is empty
}

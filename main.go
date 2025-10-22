package main

import (
	"container/list"
	"fmt"
)

// OptimalList реализует список с оптимальным внутренним представлением данных
// в зависимости от количества элементов. Автоматически переключается между
// различными структурами хранения для оптимизации производительности и памяти.
type OptimalList struct {
	size int
	data interface{}
}

// NewOptimalList создает и возвращает новый пустой список.
func NewOptimalList() *OptimalList {
	return &OptimalList{
		size: 0,
		data: nil,
	}
}

// Add добавляет элемент в конец списка.
// Автоматически изменяет внутреннее представление при изменении размера.
func (ol *OptimalList) Add(value interface{}) {
	ol.size++

	switch {
	case ol.size == 1:
		ol.data = value

	case ol.size >= 2 && ol.size <= 5:
		if ol.size == 2 {
			arr := [5]interface{}{ol.data, value}
			ol.data = arr
		} else {
			arr := ol.data.([5]interface{})
			arr[ol.size-1] = value
			ol.data = arr
		}

	case ol.size == 6:
		arr := ol.data.([5]interface{})
		l := list.New()
		for i := 0; i < 5; i++ {
			l.PushBack(arr[i])
		}
		l.PushBack(value)
		ol.data = l

	default:
		l := ol.data.(*list.List)
		l.PushBack(value)
	}
}

// Remove удаляет и возвращает последний элемент списка.
// Возвращает nil если список пуст. Автоматически изменяет
// внутреннее представление при изменении размера.
func (ol *OptimalList) Remove() interface{} {
	if ol.size == 0 {
		return nil
	}

	var value interface{}

	switch {
	case ol.size == 1:
		value = ol.data
		ol.data = nil

	case ol.size >= 2 && ol.size <= 5:
		arr := ol.data.([5]interface{})
		value = arr[ol.size-1]
		arr[ol.size-1] = nil

		if ol.size == 2 {
			ol.data = arr[0]
		} else {
			ol.data = arr
		}

	case ol.size == 6:
		l := ol.data.(*list.List)
		last := l.Back()
		value = last.Value
		l.Remove(last)

		arr := [5]interface{}{}
		i := 0
		for e := l.Front(); e != nil; e = e.Next() {
			arr[i] = e.Value
			i++
		}
		ol.data = arr

	default:
		l := ol.data.(*list.List)
		last := l.Back()
		value = last.Value
		l.Remove(last)
	}

	ol.size--
	return value
}

// Get возвращает элемент по указанному индексу.
// Возвращает nil если индекс выходит за границы списка.
func (ol *OptimalList) Get(index int) interface{} {
	if index < 0 || index >= ol.size {
		return nil
	}

	switch {
	case ol.size == 1:
		return ol.data

	case ol.size >= 2 && ol.size <= 5:
		arr := ol.data.([5]interface{})
		return arr[index]

	default:
		l := ol.data.(*list.List)
		i := 0
		for e := l.Front(); e != nil; e = e.Next() {
			if i == index {
				return e.Value
			}
			i++
		}
		return nil
	}
}

// Set устанавливает новое значение элемента по указанному индексу.
// Возвращает false если индекс выходит за границы списка.
func (ol *OptimalList) Set(index int, value interface{}) bool {
	if index < 0 || index >= ol.size {
		return false
	}

	switch {
	case ol.size == 1:
		ol.data = value
		return true

	case ol.size >= 2 && ol.size <= 5:
		arr := ol.data.([5]interface{})
		arr[index] = value
		ol.data = arr
		return true

	default:
		l := ol.data.(*list.List)
		i := 0
		for e := l.Front(); e != nil; e = e.Next() {
			if i == index {
				e.Value = value
				return true
			}
			i++
		}
		return false
	}
}

// Size возвращает текущее количество элементов в списке.
func (ol *OptimalList) Size() int {
	return ol.size
}

// String возвращает строковое представление списка для отладки.
func (ol *OptimalList) String() string {
	result := "["

	switch {
	case ol.size == 0:

	case ol.size == 1:
		result += fmt.Sprintf("%v", ol.data)

	case ol.size >= 2 && ol.size <= 5:
		arr := ol.data.([5]interface{})
		for i := 0; i < ol.size; i++ {
			if i > 0 {
				result += ", "
			}
			result += fmt.Sprintf("%v", arr[i])
		}

	default:
		l := ol.data.(*list.List)
		first := true
		for e := l.Front(); e != nil; e = e.Next() {
			if !first {
				result += ", "
			}
			result += fmt.Sprintf("%v", e.Value)
			first = false
		}
	}

	result += "]"
	return result
}

// Модульные тесты для OptimalList
func testOptimalList() {
	fmt.Println("=== Тестирование OptimalList ===")

	// Тест 1: Создание пустого списка
	fmt.Println("\nТест 1: Создание пустого списка")
	ol := NewOptimalList()
	fmt.Printf("Размер: %d, Данные: %s\n", ol.Size(), ol.String())

	// Тест 2: Добавление одного элемента
	fmt.Println("\nТест 2: Добавление одного элемента")
	ol.Add("первый")
	fmt.Printf("Размер: %d, Данные: %s\n", ol.Size(), ol.String())

	// Тест 3: Добавление элементов для перехода к массиву
	fmt.Println("\nТест 3: Добавление элементов для перехода к массиву")
	for i := 2; i <= 5; i++ {
		ol.Add(fmt.Sprintf("элемент%d", i))
		fmt.Printf("После добавления %d: Размер: %d, Данные: %s\n", i, ol.Size(), ol.String())
	}

	// Тест 4: Добавление элементов для перехода к связному списку
	fmt.Println("\nТест 4: Добавление элементов для перехода к связному списку")
	for i := 6; i <= 8; i++ {
		ol.Add(fmt.Sprintf("элемент%d", i))
		fmt.Printf("После добавления %d: Размер: %d, Данные: %s\n", i, ol.Size(), ol.String())
	}

	// Тест 5: Получение элементов по индексу
	fmt.Println("\nТест 5: Получение элементов по индексу")
	for i := 0; i < ol.Size(); i++ {
		fmt.Printf("Get(%d) = %v\n", i, ol.Get(i))
	}

	// Тест 6: Установка значений по индексу
	fmt.Println("\nТест 6: Установка значений по индексу")
	ol.Set(0, "измененный_первый")
	ol.Set(4, "измененный_пятый")
	ol.Set(7, "измененный_восьмой")
	fmt.Printf("После Set операций: %s\n", ol.String())

	// Тест 7: Удаление элементов с переходом между представлениями
	fmt.Println("\nТест 7: Удаление элементов с переходом между представлениями")
	for ol.Size() > 0 {
		removed := ol.Remove()
		fmt.Printf("Удален: %v, Размер: %d, Данные: %s\n", removed, ol.Size(), ol.String())
	}

	// Тест 8: Граничные случаи
	fmt.Println("\nТест 8: Граничные случаи")
	fmt.Printf("Get(-1) = %v\n", ol.Get(-1))
	fmt.Printf("Get(100) = %v\n", ol.Get(100))
	fmt.Printf("Set(-1, 'value') = %v\n", ol.Set(-1, "value"))
	fmt.Printf("Remove из пустого списка = %v\n", ol.Remove())

	// Тест 9: Производительность с числами
	fmt.Println("\nТест 9: Производительность с числами")
	numList := NewOptimalList()
	for i := 0; i < 10; i++ {
		numList.Add(i * 10)
	}
	fmt.Printf("Числовой список: %s\n", numList.String())

	fmt.Println("\n=== Тестирование завершено ===")
}

func main() {
	testOptimalList()

	// Дополнительные примеры использования
	fmt.Println("\n=== Дополнительные примеры ===")

	// Пример 1: Работа со строками
	fmt.Println("\nПример 1: Работа со строками")
	stringList := NewOptimalList()
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}
	for _, fruit := range fruits {
		stringList.Add(fruit)
	}
	fmt.Printf("Список фруктов: %s\n", stringList.String())

	// Пример 2: Работа с разными типами данных
	fmt.Println("\nПример 2: Работа с разными типами данных")
	mixedList := NewOptimalList()
	mixedList.Add(42)
	mixedList.Add("hello")
	mixedList.Add(3.14)
	mixedList.Add(true)
	fmt.Printf("Смешанный список: %s\n", mixedList.String())

	// Пример 3: Интенсивные операции
	fmt.Println("\nПример 3: Интенсивные операции добавления/удаления")
	dynamicList := NewOptimalList()
	for i := 0; i < 3; i++ {
		fmt.Printf("Цикл %d:\n", i+1)
		for j := 0; j < 7; j++ {
			dynamicList.Add(j + i*10)
		}
		fmt.Printf("  После добавления: %s\n", dynamicList.String())

		for j := 0; j < 4; j++ {
			dynamicList.Remove()
		}
		fmt.Printf("  После удаления: %s\n", dynamicList.String())
	}
}

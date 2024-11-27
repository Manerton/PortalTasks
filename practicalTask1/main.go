package main

import "fmt"

func task4PopInSliceByIndex(slice *[]int, index int) (int, error) {
	if index >= len(*slice) || index < 0 {
		return 0, fmt.Errorf("index out of range")
	}
	result := (*slice)[index]
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
	*slice = (*slice)[:len(*slice)]
	return result, nil
}

func task4FilterSlice(slice []int, filter func(int) bool) []int {
	result := []int{}
	for _, value := range slice {
		if filter(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {

	// task1
	var sliceEmpty []int
	fmt.Println("Пустрой slice:", sliceEmpty)

	var sliceWithValues []int = []int{0, 1, 2, 3, 4}
	fmt.Println("Слайс со значениями:", sliceWithValues)

	literalSlice := []int{1, 23, 4}
	fmt.Println("Слайс созданый литералов:", literalSlice)

	makeSlice := make([]int, 3, 6)
	fmt.Println("Слайс созданый make:", makeSlice)
	newSlice := makeSlice[3:6]
	fmt.Println(newSlice)

	// task2
	sliceEmpty = append(sliceEmpty, 13)
	fmt.Println("Добавление в пустой слайс:", sliceEmpty)

	sliceWithValues = append(sliceWithValues[:2], sliceWithValues[3:]...)
	fmt.Println("Удаление 2 из слайсла со значениями:", sliceWithValues)

	sliceWithValues = append(sliceWithValues[:2], append([]int{2}, sliceWithValues[2:]...)...)
	fmt.Println("Возвращение 2 в слайсла со значениями:", sliceWithValues)
	// task3 (На малых значения при превышении cap происходит увееничение cap в 2 раза)
	nums := make([]int, 0, 1)
	for i := 0; i < 8; i++ {
		nums = append(nums, i)
		fmt.Println("Длина:", len(nums), "Емкость:", cap(nums))
	}
	// task4
	value, err := task4PopInSliceByIndex(&sliceWithValues, 2)
	if err != nil {
		return
	}
	fmt.Println("Pop по index = 2 из слайса со значениями: ", sliceWithValues, value)
	plusMinusSlice := []int{-1, 2, -3, -4, 5, -6}
	plusSlice := task4FilterSlice(plusMinusSlice, func(i int) bool {
		if i > 0 {
			return true
		}
		return false
	})
	fmt.Println("Слайс положительных и отрицательных чисел:", plusMinusSlice)
	fmt.Println("Слайс положительных чисел:", plusSlice)

}

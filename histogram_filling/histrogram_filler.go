package main

import (
	"fmt"
)

func getLeftMaxList(list []int) []int {
	leftMax := make([]int, len(list))
	leftMax[0] = -1
	max := list[0]
	for i := 1; i < len(list); i += 1 {
		if list[i] > max {
			max = list[i]
		}
		leftMax[i] = max
	}
	return leftMax
}

func getRightMaxList(list []int) []int {
	size := len(list)
	rightMax := make([]int, size)
	rightMax[size-1] = -1
	max := list[size-1]
	for i := size - 2; i >= 0; i -= 1 {
		if list[i] > max {
			max = list[i]
		}
		rightMax[i] = max
	}
	return rightMax
}
func calculateArea(list []int) int {
	if len(list) <= 0 {
		return 0
	}
	leftMax := getLeftMaxList(list)
	rightMax := getRightMaxList(list)
	area := make([]int, len(list))
	for i := 0; i < len(list); i += 1 {
		min := leftMax[i]
		if rightMax[i] < leftMax[i] {
			min = rightMax[i]
		}
		if list[i] < min {
			area[i] = min - list[i]
		} else {
			area[i] = 0
		}
	}

	maxArea := 0
	sum := 0
	for i := 0; i < len(area); i += 1 {
		if area[i] <= 0 {
			sum = 0
		} else {
			sum += area[i]
		}
		if sum > maxArea {
			maxArea = sum
		}
	}

	return maxArea
}

func main() {
	histogram := []int{1, 5, 3, 4, 6, 4}
	fmt.Println(histogram)
	fmt.Println(calculateArea(histogram))
}

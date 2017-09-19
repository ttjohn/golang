package main

import (
	"testing"
)

func TestEmptyArray(t *testing.T) {
	var list []int
	result := calculateArea(list)
	if result != 0 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func TestSingle(t *testing.T) {
	list := []int{5}
	result := calculateArea(list)
	if result != 0 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func TestDouble(t *testing.T) {
	list := []int{5, 9}
	result := calculateArea(list)
	if result != 0 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func Test3Same(t *testing.T) {
	list := []int{5, 5, 5}
	result := calculateArea(list)
	if result != 0 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func Test3Assending(t *testing.T) {
	list := []int{5, 6, 7}
	result := calculateArea(list)
	if result != 0 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func Test3Dessending(t *testing.T) {
	list := []int{5, 4, 3}
	result := calculateArea(list)
	if result != 0 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func Test3Tank(t *testing.T) {
	list := []int{5, 4, 6}
	result := calculateArea(list)
	if result != 1 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func Test4Tank(t *testing.T) {
	list := []int{5, 4, 3, 6}
	result := calculateArea(list)
	if result != 3 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

func DoubleTank(t *testing.T) {
	list := []int{5, 4, 0, 6, 1}
	result := calculateArea(list)
	if result != 2 {
		t.Errorf("Should be %d Received %d", 0, result)
	}
}

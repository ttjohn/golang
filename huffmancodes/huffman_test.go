package main

import (
	"testing"
	"bytes"

)

func TestSingleNode(t *testing.T) {  
	var ch code
    min, max :=  findDepth(ch)
    if min != 0 {
       t.Errorf("Min Depth Should be %d and got %d", 0, min)
	}
	if max != 0 {
		t.Errorf("Max Depth should be %d and got %d", 0,max)
	}
}

func TestDoubleNodeLeft(t *testing.T) {  
	var ch, child code
	ch.left = &child
	min, max :=  findDepth(ch)
    if min != 0 {
       t.Errorf("Min Depth Should be %d and got %d", 0, min)
	}
	if max != 1 {
		t.Errorf("Max Depth should be %d and got %d", 1,max)
	}
}

func TestDoubleNodeRight(t *testing.T) {  
	var ch, child code
	ch.right = &child
	min, max :=  findDepth(ch)
    if min != 0 {
       t.Errorf("Min Depth Should be %d and got %d", 0, min)
	}
	if max != 1 {
		t.Errorf("Max Depth should be %d and got %d", 1,max)
	}
}

func TestMinMoreThanZero(t *testing.T) {  
	var ch, child, child2 code
	ch.left = &child
	ch.right = &child2
	min, max :=  findDepth(ch)
    if min != 1 {
       t.Errorf("Min Depth Should be %d and got %d", 1, min)
	}
	if max != 1 {
		t.Errorf("Max Depth should be %d and got %d", 1,max)
	}
}

func TestSingleSideTrees(t *testing.T) {  
	var ch, child, child2 code
	ch.left = &child
	child.right = &child2
	min, max :=  findDepth(ch)
    if min != 0 {
       t.Errorf("Min Depth Should be %d and got %d", 0, min)
	}
	if max != 2 {
		t.Errorf("Max Depth should be %d and got %d", 2,max)
	}
}

func TestCodeForSingleNode(t *testing.T) {  
	var ch code
	ch.nodeName = "tiju"
	buf := bytes.NewBufferString("")
	printCode(buf,"",ch)
	if buf.String() != "tiju = \n"{
		t.Errorf("Expected vale not correct %s",buf.String())
	}
	
}

func TestCodeWith2Nodes(t *testing.T) {  
	var ch, child, child2 code
	ch.left = &child
	ch.right = &child2
	child.nodeName = "A"
	child2.nodeName = "B"
	
	buf := bytes.NewBufferString("")
	printCode(buf,"",ch)
	if buf.String() != "A = 0\nB = 1\n"{
		t.Errorf("Expected vale not correct %s",buf.String())
	}
	
}

func TestCodeRightSide(t *testing.T) {  
	var ch, child2 code
	ch.right = &child2
	child2.nodeName = "B"
	
	buf := bytes.NewBufferString("")
	printCode(buf,"",ch)
	if buf.String() != "B = 1\n"{
		t.Errorf("Expected vale not correct %s",buf.String())
	}
}

func TestCodeRightSideMultiNodes(t *testing.T) {  
	var ch, child, child2 code
	ch.right = &child
	child.right = &child2
	child.nodeName = "A"
	child2.nodeName = "B"
	
	buf := bytes.NewBufferString("")
	printCode(buf,"",ch)
	if buf.String() != "B = 11\n"{
		t.Errorf("Expected vale not correct %s",buf.String())
	}
}

func TestCodeLeftSideMultiNodes(t *testing.T) {  
	var ch, child, child2 code
	ch.left = &child
	child.left = &child2
	child.nodeName = "A"
	child2.nodeName = "B"
	
	buf := bytes.NewBufferString("")
	printCode(buf,"",ch)
	if buf.String() != "B = 00\n"{
		t.Errorf("Expected vale not correct %s",buf.String())
	}
}








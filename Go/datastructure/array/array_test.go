package array

import (
	"testing"
)

func TestInsert(t *testing.T){
	a := NewArray(10)
	for i:=0;i<10;i++{
		a.Insert(i,i+1)
	}
	a.Insert(1,3)
	a.Print()
}

func TestArray_Delete(t *testing.T){
	a := NewArray(10)
	for i:=0;i<10;i++{
		a.Insert(i,i+1)
	}
	a.Delete(2)
	a.Delete(20)
	a.Print()
}

func TestArray_Find(t *testing.T){
	a := NewArray(10)
	for i:=0;i<10;i++{
		a.Insert(i,i+1)
	}
	r,err := a.Find(2)
	t.Log(r,err)
	r,err =a.Find(20)
	t.Log(r,err)
}
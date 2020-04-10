package array

import (
	"errors"
	"fmt"
	"strings"
)


type Array struct{
	data [] int
}

func NewArray(c int)*Array{
	if c ==0{
		return  nil
	}
	return &Array{
	}
}

func (a *Array)Len() int{
	return len(a.data)
}

func(a *Array)Insert(index int,value int){
	a.data = append(a.data,value)
	copy(a.data[index+1:],a.data[index:])
	a.data[index] =value
}

func(a *Array)Delete(index int)error{
	if len(a.data)<index {
		return errors.New("out of index range")
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return nil
}

func(a *Array)Find(index int)(int,error){
	if len(a.data)<index {
		return 0,errors.New("out of index range")
	}
	return a.data[index],nil
}


//打印数列
func (a *Array) Print() {
	var format strings.Builder
	for i := 0; i <len(a.data); i++ {
		format.WriteString( fmt.Sprintf("|%+v", a.data[i]))
	}
	fmt.Println(format.String())
}

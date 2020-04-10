package sort

func BubbleSort(a []int){
	// 一层代表总共比较的次数
	for i :=0;i<len(a);i++{
		var flag = false
		// 代表当前比较的次数
		for j:=0;j<len(a)-i-1;j++{
			if a[j]>a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 插入排序
func InsertionSort(a []int){
      for i:=1;i<len(a);i++{
      	v := a[i]
      	j :=i-1
      	// 当前已经排好的队列
      	for ;j>=0;j--{
			if a[j]>v {
				a[j+1]=a[j]// 移动数据
			}else {
				break
			}
		}
		a[j+1]= v
	  }
}


func SelectionSort(a []int){
	for i:=0;i<len(a);i++{
		minIndex :=i
		for j:=i+1;j<len(a);j++{
			if a[j]>a[minIndex]{
				a[j],a[minIndex] = a[minIndex],a[j]
			}
		}
	}
}

func SelectionSort1(a []int){
	for i:=0;i<len(a);i++{
		minIndex :=i
		for j:=i+1;j<len(a);j++{
			if a[j]>a[minIndex]{
				// 减少交换次数
				//a[j],a[min] = a[min],a[j]
				j = minIndex // 找到坐标点
			}
		}
		if i!=minIndex {
			a[i],a[minIndex] = a[minIndex],a[i]
		}
	}
}
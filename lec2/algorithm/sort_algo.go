package algorithm

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		isSwap := false
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return array
}

//triển khai quick sort
func Partion(array []int, low int, height int) int {
	p := array[height]
	i := low - 1
	for j := low; j < height; j++ {
		if array[j] < p {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[height] = array[height], array[i+1]
	return i + 1
}
func QuickSort(array []int, left int, right int) []int {
	if left < right {
		pivot := Partion(array, left, right)
		QuickSort(array, left, pivot-1)
		QuickSort(array, pivot+1, right)
	}

	return array
}

//triển khai merge sort
func MergeSort(items []int) []int {
    var num = len(items)
     
    if num == 1 {
        return items
    }

    middle := int(num / 2)

    var (
        left = make([]int, middle)
        right = make([]int, num-middle)
    )

    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
    
    return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
     
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
     
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }

    return
}



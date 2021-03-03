func getLeastNumbers(arr []int, k int) []int {
    if k==0{
        return []int{}
    }

    // 初始化堆[:k:k]的意思是引用arr的第0个元素至k个元素（不包括k个元素），同时heap的容量为k
    heap := arr[:k:k]
    for i:=(k-1)/ 2; i >=0;i--{
        heapify(heap,i, k-1)
    }

    // 从输入的第k个元素开始进行比较，如果小于最大堆的堆顶则
    // 认为这个元素是TopK小的元素
	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]

            // 取代之后要重新heapify
			heapify(heap, 0, k-1)
		}
	}
    return heap
}


// 官方sort库的siftDown方法，建立最大堆
func heapify (data []int, root, end int) {
    for {
        child := 2 * root + 1
        if child > end {
            return
        }
        for ;child < end && data[child] < data[child+1]; child++ {

        }
        if data[root] > data[child] {
            return
        }
        data[root], data[child] = data[child], data[root]
        root = child
    }
}

作者：evaccino
链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/golang-dui-pai-xu-jie-fa-shi-jian-fu-za-du-onlogk-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
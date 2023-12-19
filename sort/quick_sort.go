package sort

func QuickSort(datas []int) {
	quick_sort(datas, 0, len(datas)-1)
}

func quick_sort(datas []int, l, r int) {

	if l >= r {
		return
	}
	x, y := l, r
	base := datas[l]

	for x < y {
		for x < y && datas[y] >= base {
			y = y - 1
		}
		if x < y {
			datas[x] = datas[y]
			x = x + 1
		}
		for x < y && datas[x] <= base {
			x = x + 1
		}
		if x < y {
			datas[y] = datas[x]
			y = y - 1
		}

	}
	datas[x] = base
	quick_sort(datas, l, x-1)
	quick_sort(datas, x+1, r)
}

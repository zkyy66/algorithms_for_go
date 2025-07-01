package array_code

func findPeaks(mountain []int) []int {
	var res []int
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i-1] < mountain[i] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}
	return res
}

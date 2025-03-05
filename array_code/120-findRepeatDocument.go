package array_code

func findRepeatDocumentVersionOne(documents []int) int {
	if len(documents) <= 0 {
		return 0
	}
	lens := len(documents)
	for i := 0; i < lens; i++ {
		for documents[i] != i {
			if documents[documents[i]] == documents[i] {
				return documents[i]
			}
			documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
		}
	}
	return 1
}

func findRepeatDocumentVersionTwo(documents []int) int {
	res := 0
	if len(documents) <= 0 {
		return res
	}
	hashMap := make(map[int]int)
	for v := range documents {
		if _, ok := hashMap[documents[v]]; ok {
			res = documents[v]
			break
		}
		hashMap[documents[v]]++
	}
	return res
}

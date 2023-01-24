package chunk

func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 || len(collection) == 0 {
		return [][]T{}
	}

	// we get the chunk number by dividing the len of the collection divided by the size
	chunkNum := len(collection) / size
	result := make([][]T, 0, chunkNum)

	for i := 0; i < chunkNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last])
	}
	return result
}

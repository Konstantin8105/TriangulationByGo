package main

// Convert index to diaposone [0...size)
func normalizeBySize(index uint8, size uint8) uint8 {
	if 0 <= index && index < size {
		return index
	}
	if index < 0 {
		return normalizeBySize(index+size, size)
	}
	return normalizeBySize(index-size, size)
}

// Convert index to diaposone [0...3)
func normalizeBy3(index uint8) uint8 {
	return normalizeBySize(index, 3)
}

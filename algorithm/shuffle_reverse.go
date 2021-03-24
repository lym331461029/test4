package algorithm

func equalSlice(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type PathNode struct {
	Path   []int8
	Length int8
}

func cpSlice(src []int8) []int8 {
	dst := make([]int8, len(src))
	copy(dst, src)
	return dst
}

func ShuffleReverse(cards []int8, reverseCards []int8) int8 {
	if len(cards)%2 != 0 {
		return -1
	}
	n := len(cards) / 2

	pathNodes := []PathNode{PathNode{Path: cpSlice(cards), Length: 0}}

	for {
		if len(pathNodes) <= 0 {
			return -1
		}

		node := pathNodes[0]

		if equalSlice(node.Path, reverseCards) {
			return node.Length
		} else {
			for j := 1; j < n+1; j++ {
				subPath := cpSlice(node.Path)
				subPath = append(subPath[0:0], node.Path[j:n+j]...)
				subPath = append(subPath, node.Path[0:j]...)
				subPath = append(subPath, node.Path[j+n:]...)
				pathNodes = append(pathNodes, PathNode{
					Path:   subPath,
					Length: node.Length + 1,
				})
			}
		}

		pathNodes = pathNodes[1:]
	}

}

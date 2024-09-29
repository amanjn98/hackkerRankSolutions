package Graph

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.size[rootX] < uf.size[rootY] { // Union by rank
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
		} else {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		}
	}
}

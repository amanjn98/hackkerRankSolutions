package Graph

func OrangesRotting(grid [][]int) int {
	queue:=[]int{}
	fresh:=0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j]==2{
				queue=append(queue,i)
				queue=append(queue,j)
			}
			if grid[i][j]==1{
				fresh++
			}
		}
	}
	if len(queue)==0 && fresh>0{
		return -1
	}
	if fresh==0{
		return 0
	}
	if len(queue)==0{
		return 0
	}
	minTime:=0
	for {
		c:=len(queue)/2
		minTime++
		for c>0 {
			i := queue[0]
			j := queue[1]
			queue = queue[2:]
			l := i - 1
			o := j - 1
			count := 0
			for count < 2 {
				if o >= 0 && o < len(grid[i]) {
					if grid[i][o] == 1 {
						if c == 0 {
							c++
						}
						queue = append(queue, i)
						queue = append(queue, o)
						grid[i][o] = 2
					}
				}
				o = j + 1
				count++
			}
			count = 0
			for count < 2 {
				if l >= 0 && l < len(grid) {
					if grid[l][j] == 1 {
						if c == 0 {
							c++
						}
						queue = append(queue, l)
						queue = append(queue, j)
						grid[l][j] = 2
					}
				}
				l = i + 1
				count++
			}
			c--
		}
		if len(queue) == 0 {
			break
		}

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return minTime-1
}

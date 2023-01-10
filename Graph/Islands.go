package Graph

// find the number of islands
//Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//
//An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

func NumIslands(grid [][]byte) int {
  count:=0
  var queue [][]int
  for i:=0;i<len(grid);i++{
    for j:=0;j<len(grid[i]);j++{
        if grid[i][j]=='1' {
            queue=append(queue,[]int{i,j})
            for len(queue)>0{
                row,col:=queue[0][0],queue[0][1]
                if row-1>=0 && grid[row-1][col]=='1'{
                    queue=append(queue,[]int{row-1,col})
                    grid[row-1][col]='0'
                }
                if row+1<len(grid) && grid[row+1][col]=='1'{
                    queue=append(queue,[]int{row+1,col})
                    grid[row+1][col]='0'
                }
                if col-1>=0 && grid[row][col-1]=='1'{
                    queue=append(queue,[]int{row,col-1})
                    grid[row][col-1]='0'
                }
                if col+1<=len(grid[row]) && grid[row][col+1]=='1'{
                    queue=append(queue,[]int{row,col+1})
                    grid[row][col+1]='0'
                }
                grid[row][col]=0
                queue=queue[1:]
            }
            count++
        }
    }
  }
  return count
}

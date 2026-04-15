class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        m = len(grid)
        n = len(grid[0])
        count = 0
        for i in range(0, m):
            for j in range(0, n):
                if grid[i][j] < 0:
                    count += 1
        return count
func satisfiesConditions(grid [][]int) bool {
  m, n := len(grid), len(grid[0])

  for i := 0; i < m-1; i++ { // Check up to m-1 rows (exclude last row in main loop)
    for j := 0; j < n; j++ {
      if grid[i][j] != grid[i + 1][j] && (i + 1 == m || grid[i][j] != grid[i + 1][j]) {
        return false
      }
    }
  }

  // Check last row separately (always)
  for j := 0; j < n-1; j++ { // Check up to n-1 (exclude last column)
    if grid[m-1][j] == grid[m-1][j+1] {
      return false
    }
  }

  return true
}

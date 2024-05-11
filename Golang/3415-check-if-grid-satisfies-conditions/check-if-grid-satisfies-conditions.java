class Solution {
    public boolean satisfiesConditions(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        for (int i = 0; i < m - 1; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] != grid[i + 1][j]) {
                    return false;
                }
                if (j < n - 1 && grid[i][j] == grid[i][j + 1]) {
                    return false;
                }
            }
        }
        for (int j = 0; j < n - 1; j++) {
            if (grid[m - 1][j] == grid[m - 1][j + 1]) {
                return false;
            }
        }
        return true;
    }
}

class Solution:
    def areaOfMaxDiagonal(self, dimensions: List[List[int]]) -> int:
        max_area = max_diagonal = 0
        for i in range(0, len(dimensions)):
            diagonal = (dimensions[i][0] * dimensions[i][0] + dimensions[i][1] * dimensions[i][1]) ** 0.5
            area = dimensions[i][0] * dimensions[i][1]
            if diagonal > max_diagonal or (diagonal == max_diagonal and area > max_area):
                max_area = area
                max_diagonal = diagonal
        return max_area
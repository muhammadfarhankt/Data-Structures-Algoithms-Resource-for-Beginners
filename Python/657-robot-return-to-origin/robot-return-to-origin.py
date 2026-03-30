class Solution:
    def judgeCircle(self, moves: str) -> bool:
        X = Y = 0
        for char in moves:
            if char == 'U':
                Y += 1
            elif char == 'D':
                Y -= 1
            elif char == 'R':
                X += 1
            elif char == 'L':
                X -= 1
        if X == 0 and Y == 0:
            return True
        else:
            return False
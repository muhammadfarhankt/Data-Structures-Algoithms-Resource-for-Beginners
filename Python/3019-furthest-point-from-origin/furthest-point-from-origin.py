class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        location = count = 0
        for char in moves:
            if char == 'L':
                location -= 1
            elif char == 'R':
                location += 1
            else:
                count += 1
        if location > 0:
            return location + count
        else:
            return -(location - count)
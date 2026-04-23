class Solution:
    def interpret(self, command: str) -> str:
        output = ""
        i = 0
        while i < len(command):
            if command[i] == 'G':
                output += 'G'
                i += 1
            elif command[i] == '(' and command[i+1] == ')':
                output += 'o'
                i += 2
            else:
                output += "al"
                i += 4
        return "".join(output)
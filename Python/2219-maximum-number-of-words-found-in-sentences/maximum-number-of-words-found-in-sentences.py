class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        max_words = count_words = 0
        for sentence in (sentences):
            count_words = 1
            for char in sentence:
                if char == ' ':
                    count_words += 1
            if count_words > max_words:
                max_words = count_words
        return max_words
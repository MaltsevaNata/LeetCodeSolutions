from collections import defaultdict

class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        chars_dict = defaultdict(lambda: 0)
        result = 0
        
        for word in words:
            count = 0
            for char in chars:
                chars_dict[char] += 1 
            for char in word:
                if chars_dict[char] > 0:
                    chars_dict[char] -= 1
                    count += 1
                else:
                    break
            if count == len(word):
                result += count
            chars_dict.clear()
        return result
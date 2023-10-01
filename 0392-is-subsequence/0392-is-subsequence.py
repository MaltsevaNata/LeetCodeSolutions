class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) == 0 :
            return True
        if len(t) < len(s):
            return False
        p = 0

        for char in t:
            if char == s[p]:
                p += 1
                
            if p == len(s):
                return True
            
        return False
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        occur = {}
        for val in s:
            occur[val] = occur.get(val,0) + 1
        for val in t:
            if val in occur:
                occur[val] = occur.get(val,0) - 1
            else:
                return False
        return all(v == 0 for v in occur.values())
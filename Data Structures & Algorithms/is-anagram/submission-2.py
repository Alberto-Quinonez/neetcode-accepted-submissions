class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        sourceArr = []
        tarArr = []
        for c in s:
            sourceArr.append(c)
        sourceArr.sort()
        for c in t:
            tarArr.append(c)
        tarArr.sort()

        #loop through characters of second string. If any not found, return false
        
        for i in range(0, len(s)):
            if sourceArr[i] != tarArr[i]:
                return False
        #all found
        return True
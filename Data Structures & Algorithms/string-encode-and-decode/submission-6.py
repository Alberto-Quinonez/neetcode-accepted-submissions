class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            tmp = f"{len(s)}#{s}"
            res += tmp
        return res
    
    def decode(self, s: str) -> List[str]:
        res = []
        i = 0
        j = 0
        while j < len(s):
             # found end
            if s[j] == '#':
                strLength = int(s[i:j])
                decoded = s[j+1:j+1+strLength]
                res.append(decoded)
                i = j+1+strLength
                j = i
            j += 1
        return res






            



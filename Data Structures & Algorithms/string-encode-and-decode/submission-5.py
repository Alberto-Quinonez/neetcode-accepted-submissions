class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            res = res + str(len(s)) + "|" + s
        return res

    def decode(self, s: str) -> List[str]:
        res = []
        i = 0
        # loop entire string
        while i < len(s):
            j = i
            # start reading at i until we hit the delimeter
            while s[j] != "|":
                #increment j until we get there
                j = j + 1
            # we hit delimeter         
            lenght = int(s[i:j])
            res.append(s[j+1:lenght+j+1])
            i = lenght+j+1
        return res


            



class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        res = ""
        #check each character of 1st string 
        for i,val in enumerate(strs[0]):
            for string in strs[1:]:
                if i >= len(string):
                    return res
                if string[i] != val:
                    return res
            #if we went through all strings and they all have same character
            res += val
        return res




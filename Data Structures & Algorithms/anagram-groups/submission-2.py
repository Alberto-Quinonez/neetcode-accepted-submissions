class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groupedDict = {}
        for val in strs:
            s = "".join(sorted(val))
            if s in groupedDict:
                groupedDict[s].append(val)
            else:
                groupedDict[s] = []
                groupedDict[s].append(val)
        res = []
        return list(groupedDict.values())
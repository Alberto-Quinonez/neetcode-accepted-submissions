class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)
        #loop through list, sort each string.
        for s in strs:
            sortedStr = "".join(sorted(s))
            res[sortedStr].append(s)
        result = []
        for s in res.values():
            result.append(s)
        return result


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freqMap = {}
        for val in nums:
            freqMap[val] = freqMap.get(val, 0) + 1
        sortList = sorted(freqMap, key=lambda x: freqMap[x], reverse=True)
        return sortList[:k]
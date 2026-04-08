class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        freqMap = {}
        for n in nums:
            if n in freqMap:
                cnt = freqMap[n]
                freqMap[n] = cnt + 1
            else:
                freqMap[n] = 1
        return max(freqMap, key=lambda x: freqMap[x])
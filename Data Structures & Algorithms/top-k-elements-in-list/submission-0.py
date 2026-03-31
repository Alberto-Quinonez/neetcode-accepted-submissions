class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:    
        count = {}
        freqList = [[] for i in range(len(nums) + 1)]
        # for each number in list, 
        # check dictionary, add 1 to current. If none, use 0 as current
        for num in nums:
            count[num] = count.get(num, 0) + 1
        # fill in freqList based on counted
        for key, val in count.items():
            freqList[val].append(key)
        result = []
        taken = k
        for i in range(len(freqList) - 1, 0, -1):
            for freq in freqList[i]:
                if taken == 0:
                    break;
                result.append(freq)
                taken = taken - 1
        return result
        


        
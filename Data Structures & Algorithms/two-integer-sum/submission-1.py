class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numsDict = {}
        for i,val in enumerate(nums):
            if target - val in numsDict:
                res = [numsDict[target - val],i]
                return res
            else:
                numsDict[val] = i
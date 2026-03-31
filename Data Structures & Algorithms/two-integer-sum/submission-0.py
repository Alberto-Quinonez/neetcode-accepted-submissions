class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # target = nums[x] + nums[y]
        # difference = target - nums[i]
        foundMap = {}
        # iterate through list, check if difference exists in map.
        # if not add to hash set
        for i in range(0, len(nums)):
            diff = target - nums[i]
            # found difference, this is our result pair
            # index for foundMap will be smaller as we pass through only once
            if diff in foundMap:
                return [foundMap[diff],i]
            # not found, store pair in map for next iteration    
            foundMap[nums[i]] = i
            
        
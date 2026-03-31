class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        numSet = set()
        for num in nums:
            numSet.add(num)
        return len(nums) > len(numSet)
         
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        res = [1] * len(nums)
        pre = [1] * len(nums)
        suf = [1] * len(nums)
        for i in range(1, len(nums)):
            pre[i] = pre[i-1] * nums[i-1]
        for i in range(len(nums)-2,-1,-1):
            suf[i] = suf[i+1]*nums[i+1]
        for i,val in enumerate(nums):
            res[i] = pre[i]*suf[i]
        return res



# nums = [1,2,4,6]

# pre = [1,1,2,8]
# suf = [48,24,6,1]

# output = [48,24,12,8]
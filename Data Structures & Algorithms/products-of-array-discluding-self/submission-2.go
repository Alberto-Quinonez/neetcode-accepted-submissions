func productExceptSelf(nums []int) []int {
    n := len(nums)
    //intermediate array
    out := make([]int, n)

    // insert for left side
    left := 1
    for i := range nums {
        out[i] = left
        left = left * nums[i]
    }
    //create final product
    right := 1
    for i := n - 1; i >= 0; i-- {
        out[i] = out[i] * right
        right = right * nums[i]
    }
    return out
}

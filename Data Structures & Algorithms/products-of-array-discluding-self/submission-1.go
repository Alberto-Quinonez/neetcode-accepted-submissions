func productExceptSelf(nums []int) []int {
    n := len(nums)
    //intermediate array
    out := make([]int, n)

    // insert for left side
    left := 1
    for i := range nums {
        // check for bounds
        if i == 0 {
           left = 1 
        } else {
            // multiply the previous index's total product by the previous number
            left = nums[i-1] * left
        }
        out[i] = left
    }
    //create final product
    right := 1
    for i := n - 1; i >= 0; i-- {
        if i == n - 1{
            right = 1
        } else {
            // multiply the next index's total product by the next number
            right = nums[i+1] * right 
        }
        out[i] = out[i] * right
    }
    return out
}

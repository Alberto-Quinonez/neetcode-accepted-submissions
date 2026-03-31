func productExceptSelf(nums []int) []int {
    n := len(nums)
    //intermediate array
    left := make([]int, n)

    // insert for left side
    for i := range nums {
        // check for bounds
        if i == 0 {
           left[i] = 1 
        } else {
            // multiply the previous index's total product by the previous number
            left[i] = nums[i-1] * left[i-1]
        } 
    }

    //create final product
    out := make([]int, n)
    right := 1
    for i := n - 1; i >= 0; i-- {
        if i == n - 1{
            right = 1
        } else {
            // multiply the next index's total product by the next number
            right = nums[i+1] * right 
        }
        out[i] = left[i] * right
    }
    return out
}

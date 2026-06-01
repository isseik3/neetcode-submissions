func productExceptSelf(nums []int) []int {
    // output[i] should be pref[i] * suff[i], where pref[i] and suff[i] are the products of everything before index i and everything after index i, respectively.
    n := len(nums)
    pref := make([]int, n)
    suff := make([]int, n)
    res := make([]int, n)

    pref[0] = 1
    suff[n-1] = 1
    for i := 1; i < n; i++ {
        pref[i] = pref[i-1] * nums[i-1]
    }
    for i := n - 2; i >= 0; i-- {
        suff[i] = suff[i+1] * nums[i+1]
    }
    for i := 0; i < n; i++ {
        res[i] = pref[i] * suff[i]
    }
    return res

}

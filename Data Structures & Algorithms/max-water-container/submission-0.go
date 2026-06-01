func maxArea(heights []int) int {
    left, right := 0, len(heights)-1
    max := 0
    for left < right {
        width := right - left
        height := 0
        if heights[left] < heights[right] {
            height = heights[left]
            left++
        } else {
            height = heights[right]
            right--
        }
        product := width * height
        if max < product {
            max = product
        }
    }
    return max
}

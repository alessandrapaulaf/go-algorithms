func removeElement(nums []int, val int) int {
    var num int
    i := 0

    for i < len(nums) {
        if nums[i] != val {
            nums[num] = nums[i]
            num++
        }

        i++
    }

    return num
}
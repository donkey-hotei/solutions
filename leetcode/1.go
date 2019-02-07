func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int, 0)
    
    for i, n := range nums {
        if _, ok := hashmap[n]; ok {
            return []int{hashmap[n], i}
        } else {
            hashmap[target - n] = i
        }
    }
    
    return []int{}
}

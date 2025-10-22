package org.example.leetcode

class `001TwoSum` {

    fun twoSum(nums: IntArray, target: Int): IntArray {
        val memory = mutableMapOf<Int, Int>()
        for ((idx, value) in nums.withIndex()) {
            val substract = target - value
            if (memory.contains(substract)) {
                return intArrayOf(idx, memory[substract]!!)
            } else {
                memory[value] = idx
            }
        }
        return intArrayOf()
    }

    fun twoSum2(nums: IntArray, target: Int): IntArray {
        for (i in nums.indices) {
            for (j in i + 1..nums.size - 1) {
                if (nums[i] + nums[j] == target) {
                    return intArrayOf(i, j)
                }
            }
        }
        return intArrayOf()
    }


}
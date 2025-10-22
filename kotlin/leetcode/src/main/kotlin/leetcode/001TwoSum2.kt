package org.example.leetcode

class `001TwoSum2` {

    /*
    Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
    You may assume that each input would have exactly one solution, and you may not use the same element twice.
    You can return the answer in any order.
     */
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val memory = mutableMapOf<Int, Int>()

        /*
        On every iteration
            check if target - current = value in map
            if not - add value to map as value -> index
         */
        for ((i, value) in nums.withIndex()) {
            val substraction = target - value
            if (memory.contains(substraction)) {
                return intArrayOf(memory[substraction]!!, i)
            } else {
                memory[value] = i
            }
        }
        return intArrayOf()
    }


}
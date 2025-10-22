package org.example.leetcode

import kotlin.math.max

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.
 */
class `128LongestConsecutiveSequence` {

    // 4 5 2 8 5 1 3 7
    fun longestConsecutive(nums: IntArray): Int {
        if (nums.isEmpty()) return 0

        val numSet = nums.toSet()
        var maxLength = 0

        for (num in numSet) {
            // Only start counting if this is the beginning of a sequence
            if ((num - 1) !in numSet) {
                var currentNum = num
                var currentLength = 1

                // Count consecutive numbers
                while ((currentNum + 1) in numSet) {
                    currentNum++
                    currentLength++
                }

                maxLength = max(maxLength, currentLength)
            }
        }

        return maxLength
    }

}
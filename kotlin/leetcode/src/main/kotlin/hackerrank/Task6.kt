package org.example.hackerrank

class Task6 {

    fun findFirstOccurrence(nums: Array<Int>, target: Int): Int {
        return nums.indexOfFirst { it == target }
    }

}
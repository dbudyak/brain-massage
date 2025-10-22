package org.example.hackerrank

class Task5 {

    fun binarySearch(nums: Array<Int>, target: Int): Int {
        if (nums.isEmpty()) return -1

        fun search(low: Int, high: Int): Int {
            if (low > high) return -1

            val mid = (low + high) / 2

            return when {
                nums[mid] == target -> mid
                nums[mid] < target -> search(mid + 1, high)
                else -> search(low, mid - 1)
            }
        }

        return search(0, nums.size - 1)
    }
}
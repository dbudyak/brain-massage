package leetcode

import org.example.leetcode.`56MergeIntervals`
import kotlin.test.Test
import kotlin.test.assertEquals

class `56MergeIntervalsTest` {

    val task = `56MergeIntervals`()

    @Test
    fun testExample() = assertEquals(
        expected = arrayOf(intArrayOf(1, 6), intArrayOf(8, 10), intArrayOf(15, 18)),
        actual = task.merge(arrayOf(intArrayOf(1, 3), intArrayOf(2, 6), intArrayOf(8, 10), intArrayOf(15, 18)))
    )

}
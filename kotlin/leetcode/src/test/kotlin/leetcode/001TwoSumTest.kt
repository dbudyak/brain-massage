package leetcode

import org.example.leetcode.`001TwoSum`
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class `001TwoSumTest` {

    val task = `001TwoSum`()

    @Test fun testExample() = assertEquals(
        listOf(0, 1),
        task.twoSum(intArrayOf(2, 7, 11, 15), 9).toList().sorted()
    )

    @Test fun testExample2() = assertEquals(
        listOf(1, 2),
        task.twoSum(intArrayOf(3,2,4), 6).toList().sorted()
    )

    @Test fun testExample3() = assertEquals(
        listOf(0, 1),
        task.twoSum(intArrayOf(3, 3), 6).toList().sorted()
    )

}
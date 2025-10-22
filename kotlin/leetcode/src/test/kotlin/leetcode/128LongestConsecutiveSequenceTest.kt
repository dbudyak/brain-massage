package leetcode

import org.example.leetcode.`128LongestConsecutiveSequence`
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class `128LongestConsecutiveSequenceTest` {

    val task = `128LongestConsecutiveSequence`()

    @Test fun testExample() = assertEquals(
        4, task.longestConsecutive(intArrayOf(100,4,200,1,3,2))
    )

    @Test fun testExample2() = assertEquals(
        9, task.longestConsecutive(intArrayOf(0,3,7,2,5,8,4,6,0,1))
    )

    @Test fun testExample3() = assertEquals(
        3, task.longestConsecutive(intArrayOf(1,0,1,2))
    )

}
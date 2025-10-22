package leetcode

import org.example.leetcode.`011ContainerWithMostWater`
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class `011ContainerWithMostWaterTest` {

    val task = `011ContainerWithMostWater`()

    @Test fun testExample() = assertEquals(49, task.maxArea(intArrayOf(1,8,6,2,5,4,8,3,7)))
    @Test fun testExample2() = assertEquals(1, task.maxArea(intArrayOf(1,1)))
}
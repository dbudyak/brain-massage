package leetcode

import org.example.leetcode.`121BestTimeToBuyAndSellStock`
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class `121BestTimeToBuyAndSellStockTest` {
    val task = `121BestTimeToBuyAndSellStock`()

    @Test fun testExample() {
        assertEquals(5, task.maxProfit(intArrayOf(7,1,5,3,6,4)))
    }

    @Test fun testExample2() {
        assertEquals(0, task.maxProfit(intArrayOf(7,6,4,3,1)))
    }

    @Test fun testExample3() {
        assertEquals(2, task.maxProfit(intArrayOf(2,4,1)))
    }

}
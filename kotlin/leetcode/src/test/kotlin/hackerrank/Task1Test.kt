package hackerrank

import org.example.hackerrank.Task1
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class Task1Test {
    @Test
    fun testExample() = assertTrue { Task1().countResponseTimeRegressions(arrayOf(100, 200, 150, 300)) == 2 }

    @Test
    fun testEmpty() = assertTrue { Task1().countResponseTimeRegressions(emptyArray()) == 0 }

    @Test
    fun testSingle() = assertTrue { Task1().countResponseTimeRegressions(arrayOf(1)) == 0 }

}
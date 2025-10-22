package hackerrank

import org.example.hackerrank.Task2
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class Task2Test {
    @Test fun testExample() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(3, 4, -1, 1)) == 2 }
    @Test fun testEmpty() = assertTrue { Task2().findSmallestMissingPositive(emptyArray()) == 1 }
    @Test fun testOneHasZero() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(0)) == 1 }
    @Test fun testOneHasNeg() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(-4)) == 1 }
    @Test fun testOneHasNegs() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(-4, -3, -5)) == 1 }
    @Test fun testOneHasOne() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(1)) == 2 }
    @Test fun testOneHas() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(5)) == 1 }
    @Test fun testOneHasOnes() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(1, 1)) == 2 }
    @Test fun testStreet() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(0,1,2,3,4)) == 5 }
    @Test fun testStreet2() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(2,3,4)) == 1 }
    @Test fun testJump() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(1,1000)) == 2 }
    @Test fun testNegStreet() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(-4,-2, -1, 0, 1,3)) == 2 }
    @Test fun testPair() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(1,4)) == 2 }
    @Test fun testMirror() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(-1,1)) == 2 }
    @Test fun testLargerMirror() = assertTrue { Task2().findSmallestMissingPositive(arrayOf(-5,5)) == 1 }
}
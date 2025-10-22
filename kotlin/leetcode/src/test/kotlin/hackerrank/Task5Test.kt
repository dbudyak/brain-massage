package hackerrank

import org.example.hackerrank.Task5
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class Task5Test {
    @Test fun testExample() = assertTrue { Task5().binarySearch(arrayOf(2,4,6,8,10,12,14,16), 16) == 7 }
    @Test fun testExample2() = assertTrue { Task5().binarySearch(arrayOf(2,4,6,8,10,12,14,16), 4) == 1 }
    @Test fun testExample3() = assertTrue { Task5().binarySearch(arrayOf(), 4) == -1 }
}
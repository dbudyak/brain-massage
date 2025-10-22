package hackerrank

import org.example.hackerrank.Task7
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class Task7Test {
    @Test
    fun testExample() = assertEquals(
        3, Task7().maximizeNonOverlappingMeetings(
            arrayOf(
                arrayOf(2, 3),
                arrayOf(1, 2),
                arrayOf(3, 4),
                arrayOf(1, 3),
            )
        )
    )

    @Test
    fun testExample2() = assertEquals(
        4, Task7().maximizeNonOverlappingMeetings(
            arrayOf(
                arrayOf(0, 5),
                arrayOf(0, 1),
                arrayOf(1, 2),
                arrayOf(2, 3),
                arrayOf(3, 5),
                arrayOf(4, 6),
            )
        )
    )
}
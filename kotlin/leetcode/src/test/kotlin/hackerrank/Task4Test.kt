package hackerrank

import org.example.hackerrank.Task4
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class Task4Test {
    @Test fun testExample() = assertTrue { Task4().isNonTrivialRotation("abcde", "cdeab") }
    @Test fun testExample2() = assertFalse { Task4().isNonTrivialRotation("a", "a") }
    @Test fun testExample3() = assertFalse { Task4().isNonTrivialRotation("a", "b") }
}
package hackerrank

import org.example.hackerrank.Task3
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class Task3Test {
    @Test fun testExample() = assertTrue { Task3().isAlphabeticPalindrome("A1b2B!a") }
    @Test fun testExample2() = assertTrue { Task3().isAlphabeticPalindrome("A1be2c3eB!a") }
    @Test fun testEmpty() = assertTrue { Task3().isAlphabeticPalindrome("") }
    @Test fun testOne() = assertTrue { Task3().isAlphabeticPalindrome("a!") }
}
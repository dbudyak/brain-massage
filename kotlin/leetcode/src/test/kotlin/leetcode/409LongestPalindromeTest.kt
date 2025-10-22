package leetcode

import org.example.leetcode.`409LongestPalindrome`
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class `409LongestPalindromeTest` {

    var task: `409LongestPalindrome` =`409LongestPalindrome`()

    @Test fun testExample() = assertEquals(7, task.longestPalindrome("abccccdd"))
    @Test fun testExample2() = assertEquals(1, task.longestPalindrome("a"))

}
package org.example.leetcode

/**
 * Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
 * Letters are case sensitive, for example, "Aa" is not considered a palindrome.
 */
class `409LongestPalindrome` {

    fun longestPalindrome(s: String): Int {

        val eachCount: Map<Char, Int> = s.groupingBy { it }.eachCount()
        return 0

    }

}
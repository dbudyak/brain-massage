package org.example.hackerrank

class Task3 {
    fun isAlphabeticPalindrome(code: String): Boolean {
        val (left, right) = code
            .filter { it.isLetter() }
            .lowercase()
            .let {
                when (it.length % 2) {
                    0 -> it.substring(0, it.length / 2) to it.substring(it.length / 2, it.length)
                    else -> it.substring(0, it.length / 2) to it.substring(it.length / 2 + 1, it.length)
                }
            }
        return left == right.reversed()
    }
}
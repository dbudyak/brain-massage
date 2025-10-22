package org.example.leetcode

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 */
class `020ValidParenthess` {
    fun isValid(s: String): Boolean {
        val stack = ArrayDeque<Char>()

        for (char in s) {
            when (char) {
                '(', '[', '{' -> stack.addLast(char)
                ')' -> {
                    if (stack.isEmpty() || stack.removeLast() != '(') return false
                }
                ']' -> {
                    if (stack.isEmpty() || stack.removeLast() != '[') return false
                }
                '}' -> {
                    if (stack.isEmpty() || stack.removeLast() != '{') return false
                }
            }
        }

        return stack.isEmpty()
    }}
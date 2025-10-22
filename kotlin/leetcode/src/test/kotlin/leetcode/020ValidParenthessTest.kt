package leetcode

import org.example.leetcode.`020ValidParenthess`
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test

class `020ValidParenthessTest` {

    val task = `020ValidParenthess`()

    @Test fun testExample() = assertTrue { task.isValid("({[]})") }
    @Test fun testExample1() = assertTrue { task.isValid("()") }
    @Test fun testExample2() = assertTrue { task.isValid("()[]") }
    @Test fun testExample3() = assertFalse { task.isValid(")") }
    @Test fun testExample4() = assertFalse { task.isValid(")[") }
    @Test fun testExample5() = assertFalse { task.isValid("(]") }
    @Test fun testExample6() = assertFalse { task.isValid("((]]") }
    @Test fun testExample7() = assertFalse { task.isValid("(") }
    @Test fun testExample8() = assertFalse { task.isValid("([)]") }
    @Test fun testExample9() = assertFalse { task.isValid(")))") }

}
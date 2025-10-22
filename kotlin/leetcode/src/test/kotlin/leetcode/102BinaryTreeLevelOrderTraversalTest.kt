package leetcode

import org.example.leetcode.`102BinaryTreeLevelOrderTraversal`
import org.example.leetcode.TreeNode
import org.junit.jupiter.api.Assertions.*
import kotlin.test.Test
import kotlin.test.assertEquals

class `102BinaryTreeLevelOrderTraversalTest` {
    val task = `102BinaryTreeLevelOrderTraversal`()

    @Test
    fun testExample() = assertEquals(
        expected = listOf(listOf(3), listOf(9,20), listOf(15, 7)),
        actual = task.levelOrder(
            TreeNode(3).apply {
                left = TreeNode(9)
                right = TreeNode(20).apply {
                    left = TreeNode(15)
                    right = TreeNode(7)
                }
            }
        )
    )

}
package org.example.leetcode

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
 */
class `102BinaryTreeLevelOrderTraversal` {
    fun levelOrder(root: TreeNode?): List<List<Int>> {
        if (root == null) return emptyList()

        val traversalList = mutableListOf<List<Int>>()
        val queue = ArrayDeque(listOf(root))
        while (queue.isNotEmpty()) {
            var levelSize = queue.size
            val currentLevel = mutableListOf<Int>()

            for (i in 0 until levelSize) {
                val node = queue.removeFirst()
                currentLevel.add(node.`val`)

                if (node.left != null) queue.add(node.left!!)
                if (node.right != null) queue.add(node.right!!)
            }

            traversalList.add(currentLevel)

        }
        return traversalList
    }
}

class TreeNode(var `val`: Int) {
    var left: TreeNode? = null
    var right: TreeNode? = null
}
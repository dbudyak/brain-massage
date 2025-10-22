package org.example.leetcode

import kotlin.math.max
import kotlin.math.min

class `011ContainerWithMostWater` {
    fun maxArea2(height: IntArray): Int {
        var maxWater = 0
        for (i1 in height.indices) {
            for (i2 in height.indices) {
                if (i2 > i1) {
                    val waterAmount = waterAmount(height[i1], height[i2], i1, i2)
                    maxWater = max(maxWater, waterAmount)
                }
            }
        }
        return maxWater
    }

    //a = height (x) = min(height[i1], height[i2]), b = width (y) = i2 - i1
    fun waterAmount(height1: Int, height2: Int, i1: Int, i2: Int) = min(height1, height2) * (i2 - i1)

    fun maxArea(height: IntArray): Int {
        var maxWater = 0
        var left = 0
        var right = height.size - 1

        while (left < right) {
            maxWater = max(maxWater, waterAmount(height[left], height[right], left, right))

            if (height[left] < height[right]) {
                left++
            } else {
                right--
            }
        }

        return maxWater
    }
}
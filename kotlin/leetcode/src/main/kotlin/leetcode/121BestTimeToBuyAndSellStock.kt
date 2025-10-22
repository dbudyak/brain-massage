package org.example.leetcode

import kotlin.math.max
import kotlin.math.min

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
 */
class `121BestTimeToBuyAndSellStock` {
    fun maxProfit(prices: IntArray): Int {
        var maxProfit = 0
        var minPrice = Int.MAX_VALUE
        for (price in prices) {
            minPrice = min(minPrice, price)
            maxProfit = max(maxProfit, price - minPrice)
        }
        return maxProfit
    }

    fun maxProfit2(prices: IntArray): Int {
        var maxProfit: Int = 0
        for ((i, value) in prices.withIndex()) {
            for ((j, value2) in prices.withIndex()) {
                if (j > i) {
                    val profit = value2 - value
                    if (profit > maxProfit) {
                        maxProfit = profit
                    }
                }
            }
        }

        return maxProfit
    }
}
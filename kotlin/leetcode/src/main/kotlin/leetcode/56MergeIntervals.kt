package org.example.leetcode

import kotlin.math.max

/*
Given an array of intervals where intervals[i] = [starti, endi],
merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
 */
class `56MergeIntervals` {

    fun merge(intervals: Array<IntArray>): Array<IntArray> {
        // "First, I'll handle the edge case of empty input"
        if (intervals.isEmpty()) return emptyArray()

        // "Now I'll sort by start time"
        intervals.sortBy { it[0] }

        // "I'll use a list to build the result"
        val merged = mutableListOf<IntArray>()

        // "Start with the first interval"
        var current = intervals[0]

        // "Now iterate through the rest"
        for (i in 1 until intervals.size) {
            val next = intervals[i]

            // "Check if they overlap"
            if (next[0] <= current[1]) {
                // "They overlap, so extend the end"
                current[1] = max(current[1], next[1])
            } else {
                // "No overlap, add current to result and move to next"
                merged.add(current)
                current = next
            }
        }

        // "Don't forget to add the last interval"
        merged.add(current)

        return merged.toTypedArray()
    }

}
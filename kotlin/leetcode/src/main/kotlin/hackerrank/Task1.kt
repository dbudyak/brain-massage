package org.example.hackerrank

class Task1 {
    fun countResponseTimeRegressions(responseTimes: Array<Int>): Int {
        var count = 0
        for (i in 1 until responseTimes.size) {
            val average = responseTimes.copyOfRange(0, i).average()
            if (responseTimes[i] > average) {
                count++
            }
        }
        return count
    }
}
package org.example.hackerrank

class Task7 {

    fun maximizeNonOverlappingMeetings(meetings: Array<Array<Int>>): Int {
        var i = 0
        val comparator = compareBy<Array<Int>> { it[1] }.thenBy { it[0] }

        meetings
            .sortedWith(comparator)
            .also { it.forEach { println("${it[0]}..${it[1]}") } }
            .zipWithNext { left, right ->
                when {
                    right[1] != left[1] -> i++
                }
            }

        return i
    }

}
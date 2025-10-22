package org.example.hackerrank

class Task4 {

    fun isNonTrivialRotation(s1: String, s2: String): Boolean {
        var deque = ArrayDeque(s1.toList())
        return when {
            s1 == s2 -> false
            s1.length != s2.length -> false
            s1.length == 1 -> false
            else -> {
                deque.indices.forEach { _ ->
                    deque = run {
                        deque.addFirst(deque.removeLast())
                        deque
                    }
                    if (deque.toList() == s2.toList()) {
                        return true
                    }
                }
                return false
            }
        }
    }
}
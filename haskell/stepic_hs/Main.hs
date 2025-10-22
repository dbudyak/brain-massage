import Data.Char
import Debug.Trace
import Data.Function

twoDigits2Int :: Char -> Char -> Int
twoDigits2Int x y  | isDigit x && isDigit y = read (x:(y:[]))::Int | otherwise = 100

dist :: (Double, Double) -> (Double, Double) -> Double
dist p1 p2 = sqrt ((fst p1 - fst p2)^2 + (snd p1 - snd p2)^2)

fact2 :: Integer -> Integer
fact2 0 = 1
fact2 1 = 1
fact2 n = n * fact2 (n - 2)

fib :: Integer -> Integer
fib 1 = 1
fib 0 = 0
fib (-1) = 1
fib n 
	| n > 1 = fib (n - 1) + fib (n - 2) 
	| n < -1 = fib (n + 2) - fib (n + 1)

fib' :: Integer -> Integer
fib' n = helper 0 1 n

helper :: Integer -> Integer -> Integer -> Integer
helper curr prev n 
	| n == 0 = curr
	| n > 0  = helper (curr + prev) curr (n - 1)
	| n < 0	 = helper prev (curr - prev) (n + 1)

--[1,2,3,3,2,-1,-5,-10,-13,-13,-6,7,27,46,59,51,18,-49,-133,-218,-253]
seqA :: Integer -> Integer
seqA n = seqA' 0 3 2 1 n 0
seqA' :: Integer -> Integer -> Integer -> Integer -> Integer -> Integer -> Integer
seqA' curr prev1 prev2 prev3 n k | k < n = seqA' (prev1 + prev2 - 2 * prev3) curr prev1 prev2 n (k+1) | n == k = curr

getSecondFrom :: t1 -> t2 -> t3 -> t2
getSecondFrom x y z = y

on3 :: (b -> b -> b -> c) -> (a -> b) -> a -> a -> a -> c
on3 op f x y z = 	

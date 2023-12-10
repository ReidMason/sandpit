module Main where

fib :: Integer -> Integer
fib n
  | n <= 1 = n
  | otherwise = fib (n - 1) + fib (n - 2)

fibs1 = 0 : 1 : zipWith (+) fibs1 (tail fibs1)

numbers1 = 1 : map (+ 1) numbers1

main :: IO ()
main = do
  print (foldl (+) 2 [1 .. 5])

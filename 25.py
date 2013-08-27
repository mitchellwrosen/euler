#!/bin/python
def fib(n, memo={0: 0, 1: 1}):
   if n not in memo:
      memo[n] = fib(n-1) + fib(n-2)
   return memo[n]

n = 0
length = 1
while length < 1000:
   n += 1
   length = len(str(fib(n)))
print('%dth fibonacci number has length >= 1000' % n)

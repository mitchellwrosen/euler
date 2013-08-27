#!/bin/python
import itertools

ONE_MILLION = 1000000

iters = itertools.permutations('0123456789')
cur = 1
for i in iters:
   if cur == ONE_MILLION:
      print(i)
      break
   cur += 1

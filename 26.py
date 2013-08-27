import decimal as dec

PREC = 10000 # Arbitrary precision
dec.getcontext().prec = PREC

which = 1
longest = 1
for d in range(1,1001):
   decimals = str(dec.Decimal(1) / dec.Decimal(d))[2:]

   cycle = 1
   for i in range(1, PREC//2 + 1):
      repeated = decimals[:i] * (PREC//i)
      if repeated == decimals[:len(repeated)]:
         cycle = i
         break

   if (cycle > longest):
      longest = cycle
      which = d

print('%d - longest recurring cycle: %d digits' % (which, longest))

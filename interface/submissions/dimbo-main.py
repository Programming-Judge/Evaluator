from math import log2
y = int(input())
# Output is expecting 2 * x * x + int(log2(x))
for x in range(1, y):
    print(2 * x * x + 1 + int(log2(x)))
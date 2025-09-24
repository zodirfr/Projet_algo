from random import randint

tableau = [randint(0, 4) for _ in range(10)]
occu = [0] * 5
for v in tableau:
    occu[v] += 1
print(tableau)
print(occu)
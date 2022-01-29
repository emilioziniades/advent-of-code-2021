import numpy as np
import matplotlib.pyplot as plt

fig = plt.figure()
ax = plt.axes(projection= '3d')
ax.set_xlabel('x')
ax.set_ylabel('y')
ax.set_zlabel('z')

points = [[0, 0, 0]]
with open('19.ex2') as f:
    for line in f:
        line = line.strip()
        print(line)
        if line == "":
            continue
        if line == "--- scanner 0 ---":
            continue
        if line == "--- scanner 1 ---":
            break
        points.append([int(v) for v in line.split(',')])

print(points)

x = [point[0] for point in points]
y = [point[1] for point in points]
z = [point[2] for point in points]

ax.scatter(x, y, z)
plt.show()



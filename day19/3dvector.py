from mpl_toolkits import mplot3d
import numpy as np
import matplotlib.pyplot as plt

fig = plt.figure()
ax = plt.axes(projection='3d')

ax.set_xlabel('x')
ax.set_ylabel('y')
ax.set_zlabel('z')

point = [[1,2,3], [-1,2,3]]
origin = [0,0,0]

V = np.array([[1,2,3], [-1,2,3]])
origin = np.array([[0, 0], [0, 0], [0, 0]])

x = [p[0] for p in point]
y = [p[1] for p in point]
z = [p[2] for p in point]
ax.quiver(*origin, V[:,0], V[:,1], V[:,2], color = ['r'])
#ax.quiver(*origin, x, y, z, color = ['r'])
ax.set_xlim3d(-5,5)
ax.set_ylim3d(-5,5)
ax.set_zlim3d(-5,5)
plt.show()

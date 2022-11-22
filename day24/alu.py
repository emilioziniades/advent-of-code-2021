from datetime import datetime
from itertools import product

changing_values = [
    (10, 2, 1),
    (14, 13, 1),
    (14, 13, 1),
    (-13, 9, 26),
    (10, 15, 1),
    (-13, 3, 26),
    (-7, 6, 26),
    (11, 5, 1),
    (10, 16, 1),
    (13, 1, 1),
    (-4, 6, 26),
    (-9, 3, 26),
    (-13, 7, 26),
    (-9, 9, 26),
]


def step(inp, p, q, r):

    w, x, y, z = 0, 0, 0, 0

    # put input in w
    w = inp

    # zero x
    x *= 0
    # put z in x
    x += z
    # x mod 26
    x %= 26
    # divide z by r
    z //= r
    # add p to x
    x += p
    # 0 -> x if x == w
    # 1 -> x if x != w
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0

    # zero y
    y *= 0
    # add 25 to y
    y += 25
    # multiply y by x
    y *= x
    # add 1 to y
    y += 1
    # multiply z by y
    z *= y

    # zero y
    y *= 0
    # put w in y
    y += w
    # add q to y
    y += q
    # multiply y by x
    y *= x
    # add y to z
    z += y

    return z


def step_readable(z, inp, addX, addY, divZ):

    # step 1 - x becomes 0 or 1
    x = int(z % 26 + addX != inp)

    # step 2 - update z
    dz = z // divZ
    return dz + x * (25 * dz + inp + addY)


def remake_instructions():
    instruction_step = """inp w
mul x 0
add x z
mod x 26
div z {}
add x {}
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y {}
mul y x
add z y"""

    for a, b, c in changing_values:
        print(instruction_step.format(c, a, b))


def step_all(inputs):
    z = 0
    for i in range(14):
        inp = inputs[i]
        addX, addY, divZ = changing_values[i]
        z = step_readable(z, inp, addX, addY, divZ)

    return z


def step_all_interactive():
    z = 0
    for i in range(14):
        addX, addY, divZ = changing_values[i]
        print(f"z: {z}\taddX: {addX}\taddY: {addY}\tdivZ: {divZ}")

        # show options
        for j in range(1, 10):
            print(j, "\t", step_readable(z, j, addX, addY, divZ))

        inp = int(input("-> "))
        z = step_readable(z, inp, addX, addY, divZ)

    return z


def main():
    print(step_all_interactive())

    """
    now = datetime.now()
    inputs = [9] * 14

    for inputs in product(range(9, 0, -1), repeat=14):
        print("".join(str(i) for i in inputs), "   ", step_all(inputs))
        if datetime.now().second > now.second:
            break
    """


if __name__ == "__main__":
    main()

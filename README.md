# Advent of Code 2021

Solutions to (Advent of Code 2021)[https://adventofcode.com/2021/] done in Go. This was part of my journey to learn Go; some solutions are messy, others are quite neat.

Apart from each day, there are also utility packages which include Go 1.18 generic implementations of a stack, queue and priority queue, as well as some other utility functions.

Some solutions which I found quite interesting:

* (Day 15: Chiton)[https://github.com/emilioziniades/AdventOfCode2021/tree/main/day15] - Pathfinding problem solved by implementing Djikstra and A\*. I started with Djikstra but had to switch to A\* for part 2 of the puzzle.
* (Day 16: Packet Decoder)[https://github.com/emilioziniades/AdventOfCode2021/tree/main/day16] - Packet decoding using recursion
* (Day 18: Snailfish) [https://github.com/emilioziniades/AdventOfCode2021/tree/main/day18] - Parse expressions into strange maths syntax involving pairs of numbers '[x,y]', with its own operations and which could be arbitrarily nested.
* (Day 21: Dirac Dice)[https://github.com/emilioziniades/AdventOfCode2021/tree/main/day21] - Board game with 'quantum dice' - each roll splits the universe into multiple copies. Keep track of these possible universes and determine which player wins in more universes.

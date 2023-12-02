import timeit

from year2023.day2.part_1 import Part1, Cubes, Game
from year2023.day2.part_2 import Part2


def solution():
    part1 = Part1(Cubes(
        red_cube=12,
        green_cube=13,
        blue_cube=14,
    ))

    execution_time = timeit.timeit(stmt=lambda: print(part1.calculate("input")), number=1)
    print(f"Execution time part 1: {execution_time} seconds")

    execution_time = timeit.timeit(stmt=lambda: print(Part2().calculate("input")), number=1)
    print(f"Execution time part 2: {execution_time} seconds")

import timeit

from year2023.day1.part_1 import Part1
from year2023.day1.part_2 import Part2


def solution():
    execution_time = timeit.timeit(stmt=lambda: print(Part1().calculate("input")), number=1)
    print(f"Execution time part 1: {execution_time} seconds")

    execution_time = timeit.timeit(stmt=lambda: print(Part2().calculate("input")), number=1)
    print(f"Execution time part 2: {execution_time} seconds")

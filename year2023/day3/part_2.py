"""
--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life,
you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately,
the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer,
holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station.
You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is
adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which
gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears.
The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345.
The second gear is in the lower right; its gear ratio is 451490.
(The * adjacent to 617 is not a gear because it is only adjacent to one part number.)
Adding up all of the gear ratios produces 467835.
"""
from functools import reduce


class Part2:
    number_sum = 0
    current_line = ""

    def calculate(self, file_name: str) -> int:
        with open(f"year2023/day3/input/{file_name}.txt", "r") as document:
            lines: list[str] = []
            for line in document.readlines():
                line = line.replace("\n", "")

                lines.append(line)
                if len(lines) != 2 and self.current_line == "":
                    self.current_line = line
                    continue

                if len(lines) > 3:
                    del lines[0]

                self.read_details_number(lines)
                self.current_line = lines[lines.index(self.current_line) + 1]

            del lines[0]
            self.read_details_number(lines)

        return self.number_sum

    def read_details_number(self, lines: list[str]):
        numbers = []
        for idx, char in enumerate(self.current_line):
            if char == "*":
                for line in lines:
                    if line[idx].isnumeric():
                        raw_number = line[idx]
                        left_number = self._left_number(line, idx-1)
                        if left_number:
                            raw_number = str(left_number) + raw_number

                        right_number = self._right_number(line, idx+1)
                        if right_number:
                            raw_number += str(right_number)

                        numbers.append(int(raw_number))
                        continue

                    left_number = self._left_number(line, idx-1)
                    if left_number:
                        numbers.append(int(left_number))

                    right_number = self._right_number(line, idx+1)
                    if right_number:
                        numbers.append(int(right_number))

            if len(numbers) == 2:
                self.number_sum += numbers[0] * numbers[1]

            numbers.clear()

    def _left_number(self, line: str, idx: int) -> str:
        number = ""
        while idx >= 0 and line[idx].isdigit():
            number = line[idx] + number
            idx -= 1

        return number

    def _right_number(self, line: str, idx: int) -> str:
        number = ""
        while idx < len(line) and line[idx].isdigit():
            number += line[idx]
            idx += 1

        return number

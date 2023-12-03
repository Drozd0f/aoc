"""
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source,
but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise.
"Sorry, I wasn't expecting anyone! The gondola lift isn't working right now;
it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one.
If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine.
There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol,
even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

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

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and
58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.
"""
import re


NUMBER_SYMBOL_REGEX = r"[^\w.\n]"


class Part1:
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
        prev_char = ""
        number = ""
        number_idx = 0
        for idx, char in enumerate(self.current_line):
            if char.isnumeric():
                number += char
                if not prev_char.isnumeric():
                    number_idx = idx
            elif prev_char.isnumeric():
                if number_idx != 0:
                    number_idx -= 1
                for line in lines:
                    number_symbol = re.search(NUMBER_SYMBOL_REGEX, line[number_idx:number_idx+len(number) + 2])
                    if number_symbol:
                        self.number_sum += int(number)
                        break
                number = ""

            prev_char = char

        if number != "":
            if number_idx != 0:
                number_idx -= 1
            for line in lines:
                number_symbol = re.search(NUMBER_SYMBOL_REGEX, line[number_idx:number_idx+len(number) + 2])
                if number_symbol:
                    self.number_sum += int(number)
                    break

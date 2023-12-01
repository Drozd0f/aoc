"""
--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters:
one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are
29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.
"""


class Part2:
    _valid_digits = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    def calculate(self, file_name: str) -> int:
        values = 0
        with open(f"year2023/day1/input/{file_name}.txt", "r") as document:
            for line in document.readlines():
                value = self.calculate_calibration_value(line.replace("\n", ""))
                values += value

        return values

    def calculate_calibration_value(self, line: str) -> int:
        idx_dict = {}

        for name_digit, digit in self._valid_digits.items():
            digit_idx = line.find(str(digit))
            if not digit_idx == -1:
                idx_dict[digit_idx] = digit
                idx_dict[line.rfind(str(digit))] = digit

            digit_idx = line.find(name_digit)
            if not digit_idx == -1:
                idx_dict[digit_idx] = digit
                idx_dict[line.rfind(name_digit)] = digit

        try:
            res = idx_dict[min(idx_dict)] * 10 + idx_dict[max(idx_dict)]
        except ValueError:
            return 0

        return res

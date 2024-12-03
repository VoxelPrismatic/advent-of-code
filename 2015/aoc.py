import argparse
import cal.day01
import cal.day02
import cal.day03

parser = argparse.ArgumentParser()
parser.add_argument("stage", type=str)
parser.add_argument("input", type=str)


if __name__ == "__main__":
    args = parser.parse_args()

    data: str = open(args.input).read()
    data.strip()
    lines: list[str] = data.split("\n")
    while lines[-1] == "":
        lines.pop()

    match args.stage:
        case "1a" | "01a":
            cal.day01.part1(lines)
        case "1b" | "01b":
            cal.day01.part2(lines)

        case "2a" | "02a":
            cal.day02.part1(lines)
        case "2b" | "02b":
            cal.day02.part2(lines)

        case "3a" | "03a":
            cal.day03.part1(lines)
        case "3b" | "03b":
            cal.day03.part2(lines)

        case "4a" | "04a":
            cal.day04.part1(lines)
        case "4b" | "04b":
            cal.day04.part2(lines)

        case "5a" | "05a":
            cal.day05.part1(lines)
        case "5b" | "05b":
            cal.day05.part2(lines)

        case "6a" | "06a":
            cal.day06.part1(lines)
        case "6b" | "06b":
            cal.day06.part2(lines)

        case "7a" | "07a":
            cal.day07.part1(lines)
        case "7b" | "07b":
            cal.day07.part2(lines)

        case "8a" | "08a":
            cal.day08.part1(lines)
        case "8b" | "08b":
            cal.day08.part2(lines)

        case "9a" | "09a":
            cal.day09.part1(lines)
        case "9b" | "09b":
            cal.day09.part2(lines)

        case "10a":
            cal.day10.part1(lines)
        case "10b":
            cal.day10.part2(lines)

        case "11a":
            cal.day11.part1(lines)
        case "11b":
            cal.day11.part2(lines)

        case "12a":
            cal.day12.part1(lines)
        case "12b":
            cal.day12.part2(lines)

        case "13a":
            cal.day13.part1(lines)
        case "13b":
            cal.day13.part2(lines)

        case "14a":
            cal.day14.part1(lines)
        case "14b":
            cal.day14.part2(lines)

        case "15a":
            cal.day15.part1(lines)
        case "15b":
            cal.day15.part2(lines)

        case "16a":
            cal.day16.part1(lines)
        case "16b":
            cal.day16.part2(lines)

        case "17a":
            cal.day17.part1(lines)
        case "17b":
            cal.day17.part2(lines)

        case "18a":
            cal.day18.part1(lines)
        case "18b":
            cal.day18.part2(lines)

        case "19a":
            cal.day19.part1(lines)
        case "19b":
            cal.day19.part2(lines)

        case "20a":
            cal.day20.part1(lines)
        case "20b":
            cal.day20.part2(lines)

        case "21a":
            cal.day21.part1(lines)
        case "21b":
            cal.day21.part2(lines)

        case "22a":
            cal.day22.part1(lines)
        case "22b":
            cal.day22.part2(lines)

        case "23a":
            cal.day23.part1(lines)
        case "23b":
            cal.day23.part2(lines)

        case "24a":
            cal.day24.part1(lines)
        case "24b":
            cal.day24.part2(lines)

        case "25a":
            cal.day25.part1(lines)
        case "25b":
            cal.day25.part2(lines)
        case _:
            print(f"Stage `{args.stage}' not implemented")

def part1(lines: list[str]):
    for line in lines:
        print(line.count("(") - line.count(")"))


def part2(lines: list[str]):
    for line in lines:
        i, floor = 0, 0
        for i in range(len(line)):
            if line[i] == "(":
                floor += 1
            elif line[i] == ")":
                floor -= 1
            if floor < 0:
                print(i + 1)
                break
        else:
            print("<never reaches basement>")

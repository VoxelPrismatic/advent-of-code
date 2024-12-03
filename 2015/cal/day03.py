def part1(lines: list[str]):
    for line in lines:
        x, y = 0, 0
        houses: set[tuple[int, int]] = set()
        houses.add((x, y))
        for direction in line:
            match direction:
                case "^":
                    y += 1
                case "v":
                    y -= 1
                case ">":
                    x += 1
                case "<":
                    x -= 1
            houses.add((x, y))
        print(len(houses))


def part2(lines: list[str]):
    for line in lines:
        robo_x, robo_y, real_x, real_y = 0, 0, 0, 0
        do_robo = False
        houses: set[tuple[int, int]] = set()
        houses.add((robo_x, robo_y))
        houses.add((real_x, real_y))
        for direction in line:
            match (direction, do_robo):
                case ("^", False):
                    real_y += 1
                case ("^", True):
                    robo_y += 1
                case ("v", False):
                    real_y -= 1
                case ("v", True):
                    robo_y -= 1
                case (">", False):
                    real_x += 1
                case (">", True):
                    robo_x += 1
                case ("<", False):
                    real_x -= 1
                case ("<", True):
                    robo_x -= 1

            houses.add((robo_x, robo_y))
            houses.add((real_x, real_y))
            do_robo = not do_robo

        print(len(houses))

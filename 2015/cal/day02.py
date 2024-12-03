def part1(lines: list[str]):
    ttl = 0
    for line in lines:
        if line == "":
            print(ttl)
            ttl = 0
            continue

        w, l, h = [int(x) for x in line.split("x")]
        m = min(w * l, l * h, h * w)
        ttl += 2 * (l * w + w * h + h * l) + m
    print(ttl)


def part2(lines: list[str]):
    ttl = 0
    for line in lines:
        if line == "":
            print(ttl)
            ttl = 0
            continue

        w, l, h = [int(x) for x in line.split("x")]
        front, side, top, vol = l * w, w * h, h * l, l * w * h
        sides = {
            front: 2 * (l + w),
            side: 2 * (w + h),
            top: 2 * (h + l),
        }

        ttl += vol + sides[min(sides)]

    print(ttl)

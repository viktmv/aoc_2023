from collections import Counter

filename = "input.txt"
target = "XMAS"
directions = [
    (0, 1),  # up
    (1, 1),  # up right
    (1, 0),  # right
    (1, -1),  # down right
    (0, -1),  # down
    (-1, -1),  # down left
    (-1, 0),  # left
    (-1, 1),  # up left
]


class Node:
    used = False
    char = ""
    x = 0
    y = 0

    def __init__(self, char, x, y):
        self.char = char
        self.x = x
        self.y = y

    def __repr__(self) -> str:
        return f"{self.char}; x: {self.x}; y: {self.y}; used: {self.used}"


grid = []
row_idx = 0
for line in open(filename).readlines():
    grid.append([Node(line[idx], x=idx, y=row_idx) for idx in range(len(line.strip()))])
    row_idx += 1


def next_char(prev):
    return target[target.index(prev) + 1]


def get_node(x, y):
    if y > len(grid) or y < 0:
        return None

    if x > len(grid[0]) or x < 0:
        return None

    try:
        return grid[y][x]
    except IndexError:
        return None


def try_direction(prev: Node, dir, word):
    cx, cy = dir
    node = get_node(prev.x + cx, prev.y + cy)

    if not node:
        return False

    curr = word + node.char
    if curr == target:
        node.used = True
        return True

    if curr in target:
        if try_direction(node, dir, curr):
            node.used = True
            print(node)
            return True

    return False


# pt 1
def try_all_directions(node):
    count = 0
    for dir in directions:
        if try_direction(node, dir, "X"):
            node.used = True
            count += 1
    return count


def part1():
    count = 0
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            node = grid[y][x]
            if node.char == "X":
                count += try_all_directions(node)
    print(count)


def part2():
    counter = 0
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            node = grid[y][x]

            if node.char != "A":
                continue

            results = []
            for cx, cy in [(-1, -1), (1, 1), (1, -1), (-1, 1)]:
                n = get_node(x=node.x + cx, y=node.y + cy)
                if n and n.char in ("S", "M"):
                    results.append(n.char)

            if len(results) != 4:
                continue

            if results[0] == results[1]:
                continue

            c = Counter(results)
            if c.get("S") == 2 and c.get("M") == 2:
                counter += 1

    print(counter)


# Print
def print_grid():
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            node = get_node(x=x, y=y)
            if node and node.used:
                print(node.char, end="")
            else:
                print(".", end="")
        print()


part2()

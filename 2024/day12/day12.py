from collections import deque
from grid.main import Grid


def make_grid():
    with open("input.txt") as datafile:
        return Grid(datafile.read())


directions = [
    (0, 1),  # up
    (1, 0),  # right
    (0, -1),  # down
    (-1, 0),  # left
]

grid = make_grid()


def find_plot(start):
    curr = start
    queue = deque([curr])

    val = curr.val
    perimeter = 0
    area = 0

    while len(queue):
        curr = queue.popleft()
        area += 1

        if not curr.visited:
            curr.visited = True

        for cx, cy in directions:
            next = grid.get_point(x=curr.x + cx, y=curr.y + cy)
            if not next or next.val != val:
                perimeter += 1
            elif next and next.visited:
                continue
            else:
                next.visited = True
                queue.append(next)

    return area, perimeter


def part_1():
    price = 0
    for y in range(len(grid.rows)):
        for x in range(len(grid.rows[0])):
            point = grid.rows[y][x]
            if point.visited:
                continue
            area, perimeter = find_plot(point)
            price += area * perimeter

    print("total price:", price)


part_1()

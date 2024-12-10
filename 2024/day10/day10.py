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


def part_one(grid):
    trailheads = grid.filter(lambda p: p.val == "0")
    score = 0

    for trailhead in trailheads:
        queue = deque([trailhead])
        peaks = []

        while len(queue) > 0:
            current = queue.popleft()

            for cx, cy in directions:
                point = grid.get_point(x=current.x + cx, y=current.y + cy)
                if not point:
                    continue

                if int(point.val) - int(current.val) == 1:
                    if point.val == "9" and not point in peaks:
                        peaks.append(point)
                        continue
                    queue.append(point)

        score += len(peaks)
        print(f"score for {trailhead} = {len(peaks)}")

    print(f"total score: {score}")


def part_two(grid):
    trailheads = grid.filter(lambda p: p.val == "0")
    score = 0

    for trailhead in trailheads:
        rating = 0
        queue = deque([trailhead])

        while len(queue) > 0:
            current = queue.popleft()

            for cx, cy in directions:
                point = grid.get_point(x=current.x + cx, y=current.y + cy)
                if not point:
                    continue

                if int(point.val) - int(current.val) == 1:
                    if point.val == "9":
                        rating += 1
                        continue

                    queue.append(point)

        score += rating
        print(f"score for {trailhead} = {rating}")

    print(f"total score: {score}")


part_two(grid)

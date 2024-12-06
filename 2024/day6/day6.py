class Point:
    x = -1
    y = -1
    val = None
    visited = False

    def __init__(self, x: int, y: int, val: str) -> None:
        self.x = x
        self.y = y
        self.val = val

    def __repr__(self) -> str:
        return f"Point: x={self.x}; y={self.y}; val={self.val}"


class Grid:
    rows = []

    def __init__(self, input: str) -> None:
        lists = [list(line) for line in input.split("\n") if line]

        self.rows = [
            [Point(x=x, y=y, val=col) for x, col in enumerate(row)]
            for y, row in enumerate(lists)
        ]

    def __repr__(self) -> str:
        stringified = "\n".join(
            [
                "".join([str("X" if col.visited else col.val) for col in row])
                for row in self.rows
            ]
        )
        return stringified

    def filter(self, fn):
        filtered = []
        for row in self.rows:
            for p in row:
                if fn(p):
                    filtered.append(p)
        return filtered

    def get_point(self, x, y) -> Point | None:
        if not self.is_within_bounds(Point(x=x, y=y, val="")):
            return None

        return self.rows[y][x]

    def is_within_bounds(self, p: Point) -> bool:
        if len(self.rows) == 0:
            return False

        if p.x < 0 or p.x >= len(self.rows[0]):
            return False

        if p.y < 0 or p.y >= len(self.rows):
            return False

        return True


def read_file():
    with open("input_test.txt") as datafile:
        return datafile.read()


input = read_file()
grid = Grid(input)

dirs = {"v": (0, 1), ">": (1, 0), "^": (0, -1), "<": (-1, 0)}
next_dirs = {"v": "<", ">": "v", "^": ">", "<": "^"}


def find_guard(grid: Grid):
    for row in grid.rows:
        for col in row:
            if col.val in dirs.keys():
                return col


start_pos = find_guard(grid)
current_pos = start_pos
counter = 0
current_dir = start_pos.val

while current_pos and grid.is_within_bounds(current_pos):
    current_pos.visited = True
    cx, cy = dirs[str(current_dir)]
    next_pos = grid.get_point(x=current_pos.x + cx, y=current_pos.y + cy)

    if not next_pos:
        break

    if next_pos.val == "#":
        current_dir = next_dirs[current_dir]
        cx, cy = dirs[current_dir]
        next_pos = grid.get_point(x=current_pos.x + cx, y=current_pos.y + cy)
        current_pos = next_pos
    else:
        current_pos = next_pos

pt1_results = len(grid.filter(lambda p: p.visited == True))
print(pt1_results)

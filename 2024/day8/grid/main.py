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
        return f"Point: x={self.x}; y={self.y}; val={self.val}; visited={self.visited}"


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
            ["".join([str(col.val) for col in row]) for row in self.rows]
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

    def insert(self, p: Point) -> bool:
        if not self.is_within_bounds(p):
            return False

        self.rows[p.y][p.x] = p
        return True

    def is_within_bounds(self, p: Point) -> bool:
        if len(self.rows) == 0:
            return False

        if p.x < 0 or p.x >= len(self.rows[0]):
            return False

        if p.y < 0 or p.y >= len(self.rows):
            return False

        return True

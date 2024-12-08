from grid.main import Grid, Point
from itertools import groupby, permutations
import collections


def make_grid():
    with open("input.txt") as datafile:
        return Grid(datafile.read())


# stolen from the internets
def groupby_unsorted(seq, key=lambda x: x):
    indexes = collections.defaultdict(list)
    for i, elem in enumerate(seq):
        indexes[key(elem)].append(i)
    for k, idxs in indexes.items():
        yield k, (seq[i] for i in idxs)


grid = make_grid()
points = grid.filter(lambda p: p.val != ".")
grouped = groupby_unsorted(points, key=lambda p: p.val)


def get_delta(p1, p2):
    x = p2.x - p1.x
    y = p2.y - p1.y

    return x, y


# Part 1
def generate_antinodes(group):
    counter = 0
    for p1, p2 in list(permutations(group, 2)):
        x, y = get_delta(p1, p2)
        antinode_1 = Point(x=p1.x - x, y=p1.y - y, val="#")
        antinode_2 = Point(x=p2.x + x, y=p2.y + y, val="#")

        for anode in (antinode_1, antinode_2):
            if grid.is_within_bounds(anode):
                p = grid.get_point(x=anode.x, y=anode.y)
                if p.val != "#":
                    counter += 1
                    grid.insert(anode)
    return counter


def part_1():
    counter = 0
    for _, group in grouped:
        counter += generate_antinodes(list(group))

    print("result:", counter)


#  Part 2
def generate_antinodes_adjusted(group):
    for p1, p2 in list(permutations(group, 2)):
        x, y = get_delta(p1, p2)
        antinode_1 = Point(x=p1.x - x, y=p1.y - y, val="#")
        antinode_2 = Point(x=p2.x + x, y=p2.y + y, val="#")

        while antinode_1 or antinode_2:
            if antinode_1 and grid.is_within_bounds(antinode_1):
                p = grid.get_point(x=antinode_1.x, y=antinode_1.y)
                if p.val != "#":
                    grid.insert(antinode_1)
                antinode_1 = Point(x=antinode_1.x - x, y=antinode_1.y - y, val="#")
            else:
                # out of bounds
                antinode_1 = None

            if antinode_2 and grid.is_within_bounds(antinode_2):
                p = grid.get_point(x=antinode_2.x, y=antinode_2.y)
                if p.val != "#":
                    grid.insert(antinode_2)
                antinode_2 = Point(x=antinode_2.x + x, y=antinode_2.y + y, val="#")
            else:
                # out of bounds
                antinode_2 = None

        if p1.val != "#":
            p1.val = "#"

        if p2.val != "#":
            p2.val = "#"


def part_2():
    for _, group in grouped:
        nodes = list(group)
        generate_antinodes_adjusted(nodes)

    print("result: ", len(grid.filter(lambda p: p.val != ".")))


part_2()

import re

line = open("input.txt").read()

mul_instructions = re.findall(r"mul\((\d+),(\d+)\)", line)

# Pt1
print("Part one:", sum([int(a) * int(b) for a, b in mul_instructions]))


# Pt2
should_apply = True
sum = 0
for match in re.findall(r"(don't\(\))|(do\(\))|mul\((\d+),(\d+)\)", line):
    dont, do, a, b = match
    if dont:
        should_apply = False
    elif do:
        should_apply = True
    else:
        if should_apply:
            sum += int(a) * int(b)

print("Part 2:", sum)

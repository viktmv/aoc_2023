import re

line = open("input.txt").read()

mul_instructions = re.findall(r"mul\((\d+),(\d+)\)", line)
# Pt1
# print("Part one:", sum([int(a) * int(b) for a, b in mul_instructions]))

# Pt2
def get_action(pos):
    instruction_pattern = re.compile(r"(don't\(\))|(do\(\))")
    return instruction_pattern.search(line, pos)


def get_operation(pos):
    mul_pattern = re.compile(r"mul\((\d+),(\d+)\)")
    return mul_pattern.search(line, pos)


queue = []
pos = 0
match = get_operation(pos)
last_action = get_action(pos)
next_action = get_action(pos)
should_execute = True

while match:
    # case where no action at all
    if not last_action:
        queue.append(match.groups())
        pos = match.end()
        match = get_operation(pos)
        continue

    dont, do = last_action.groups()

    # apply most recent action
    if last_action.end() <= match.start():
        if do:
            should_execute = True
        elif dont:
            should_execute = False

    # handle end of input after last action
    if not next_action:
        if should_execute:
            queue.append(match.groups())

        pos = match.end()
        match = get_operation(pos)
        continue

    # determine execution between two actions
    while match and match.end() <= next_action.start():
        if should_execute:
            queue.append(match.groups())

        pos = match.end()
        match = get_operation(pos)

    last_action, next_action = next_action, get_action(next_action.end())

print("Part 2:", sum([int(a) * int(b) for a, b in queue]))

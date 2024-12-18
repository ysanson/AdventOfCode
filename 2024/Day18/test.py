### PART ONE
bytes = [tuple(map(int, line.strip().split(','))) for line in open('input').readlines()]
grid = {}
for x in range(71):
    for y in range(71):
        grid[(x,y)] = '.'

for i in range(1024):
    grid[bytes[i]] = '#'

seen = {}
q = [(0, (0,0))]
while len(q) > 0:
    score, pos = q.pop(0)
    if pos in seen and seen[pos] <= score:
        continue
    seen[pos] = score
    for dx, dy in [(0,1), (0,-1), (1,0), (-1,0)]:
        nx, ny = pos[0]+dx, pos[1]+dy
        if (nx, ny) in grid and grid[(nx,ny)] != '#':
            q.append((score+1, (nx,ny)))

print(seen[(70,70)])



### PART TWO
bytes = [tuple(map(int, line.strip().split(','))) for line in open('input').readlines()]
grid = {}
for x in range(71):
    for y in range(71):
        grid[(x,y)] = '.'

for i, byte in enumerate(bytes):
    grid[byte] = '#'

    seen = set()
    q = [(0,0)]
    while len(q) > 0:
        pos = q.pop(0)
        if pos in seen:
            continue
        seen.add(pos)
        for dx, dy in [(0,1), (0,-1), (1,0), (-1,0)]:
            nx, ny = pos[0]+dx, pos[1]+dy
            if (nx, ny) in grid and grid[(nx,ny)] != '#':
                q.append((nx,ny))
    if (70,70) not in seen:
        print('found it', byte)
        break
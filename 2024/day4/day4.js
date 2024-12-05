import fs from 'fs/promises'

let target = 'XMAS'
let filename = 'input.txt'

let file = await fs.open(filename)
let { buffer } = await file.read()

let grid = buffer
    .toString()
    .split('\n')
    .map(row => row.split(''))


grid.pop() // rm last weird row

let directions = [
    [0, 1],  // up
    [1, 1],  // up right
    [1, 0],  // right
    [1, -1],  // down right
    [0, -1],  // down
    [-1, -1],  // down left
    [-1, 0],  // left
    [-1, 1],  // up left
]

const getPoint = (x, y) => {
    if (y >= grid.length || y < 0) return
    if (x >= grid[0].length || x < 0) return
    return grid[y][x] 
}

let mainCount = 0
for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[0].length; x++) {
        let point = grid[y][x]
        if (point != 'X') continue
        for (let [cx, cy] of directions) {
            let count = 1
            let word = point

            while (count < 4) {
                console.log(x + (cx * count), y + (cy * count))
                let next = getPoint(x + (cx * count), y + (cy * count)) 
                if (!next) break
                word += next
                count++
            }
            console.log(word, { y, x }, word == target)
            if (word == target) {
                mainCount++
            }
        }
    }
}

console.log(mainCount)
// console.log(grid)

file.close()

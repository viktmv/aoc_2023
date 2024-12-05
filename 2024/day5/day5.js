import fs from 'fs/promises'

let pathname = 'input.txt'
let data = await fs.readFile(pathname, { encoding: 'utf8' })

let [rules, reports] =  data.split('\n\n')

let rulebook = rules.split('\n').reduce((acc, line) => {
    let [prev, next] = line.split('|')
    if (acc[prev]) acc[prev].push(next)
    else acc[prev] = [next]
    return acc
}, {})

let reportsList = reports
    .split('\n')
    .map(line => line.split(','))
    .filter(list => list.every(Boolean))

let invalidReports = []
let validReports = reportsList.filter(report => {
    for (let i = 0; i < report.length; i++) {
        let page = report[i]
        let allowedNextPages = rulebook[page]

        if (!allowedNextPages) continue

        for (let j = 0; j < report.length; j++) {
            if (j == i) continue

            let currentPage = report[j]

            if (allowedNextPages.includes(currentPage) && j < i) {
                invalidReports.push(report)
                return false
            }
        }
    }
    return true
})


function part1() {
    console.log('valid reports', validReports)
    let sum = 0
    for (let report of validReports) {
        sum += +report[Math.floor(report.length / 2)]
    }

    console.log('sum', sum)
}

for (let report of invalidReports) {
    for (let i = 0; i < report.length; i++) {
        let page = report[i]
        let allowedNextPages = rulebook[page]

        if (!allowedNextPages) continue

        for (let j = 0; j < report.length; j++) {
            if (j == i) continue

            let currentPage = report[j]

            if (allowedNextPages.includes(currentPage) && j < i) {
                let page = report[i]
                report[i] = currentPage 
                report[j] = page
            }
        }
    }
}


function part2() {
    console.log('invalid reports', invalidReports)

    let sum = 0
    for (let report of invalidReports) {
        sum += +report[Math.floor(report.length / 2)]
    }

    console.log('sum', sum)
}

part2()

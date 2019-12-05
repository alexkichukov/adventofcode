const fs = require('fs')

// Read the file
const file = fs.readFileSync(__dirname + '/input.txt', { encoding: 'utf8' })

// Generate an array of all masses
const masses = file.split(/\n/g).map(mass => parseInt(mass))

// First part
function firstPart() {
  let totalFuel = 0
  for (const mass of masses) {
    const fuel = Math.floor(mass / 3) - 2
    totalFuel += fuel
  }
  return totalFuel
}

// Second part
function secondPart() {
  let totalFuel = 0
  for (const mass of masses) {
    let fuel = Math.floor(mass / 3) - 2

    // Calculate the fuel needed for the fuel itself
    let additional = fuel
    while (additional > 0) {
      additional = Math.floor(additional / 3) - 2

      if (additional <= 0) break

      fuel += additional
    }

    // Add to the total fuel
    totalFuel += fuel
  }
  
  return totalFuel
}

console.log(secondPart())

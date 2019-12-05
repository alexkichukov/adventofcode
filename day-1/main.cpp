#include <iostream>
#include <vector>
#include <fstream>
#include <cmath>

// Forward declarations
std::vector<int> get_masses();
int first_part();
int second_part();

int main() {
  std::cout << "First part result:\n" << first_part() << "\n\n";
  std::cout << "Second part result:\n" << second_part() << "\n";
}

// The first part
int first_part() {
  std::vector<int> masses = get_masses();
  
  int total_fuel = 0;
  for (int mass : masses) {
    int fuel = std::floor(mass / 3) - 2;
    total_fuel += fuel;
  }

  return total_fuel;
}

int second_part() {
  std::vector<int> masses = get_masses();
  
  // Store the total fuel here
  int total_fuel = 0;

  // Loop through the masses of each module
  for (int mass : masses) {
    int fuel = std::floor(mass / 3) - 2;

    // Calculate the fuel needed for the fuel itself
    int additional = fuel;
    while (additional > 0) {
      // Calculate the additional fuel for the previous additioanl fuel
      // If it is <= 0 break out of the loop
      additional = std::floor(additional / 3) - 2;
      if (additional <= 0) break;

      // Add the additional fuel to the total fuel for the module
      fuel += additional;
    }

    // Add the total fuel for the module to the total fuel for all modules
    total_fuel += fuel;
  }

  return total_fuel;
}

// Get masses vector from input.txt
std::vector<int> get_masses() {
  std::ifstream input("input.txt");
  std::vector<int> masses;

  // Push each line to the vector
  int mass;
  while (input >> mass) {
    masses.push_back(mass);
  }

  return masses;
}

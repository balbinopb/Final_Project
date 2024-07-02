2.Lift or elevator scheduling
This is a simulator to see the effects of different rules of how elevators serve passengers.
The input will be multilines, each line contains three integers: the time, the starting floor, and the destination floor
Based on various consideration, such as current elevator position, current request floor(s) and destionation floor(s), the program takes a strategy determine the next floor the elevator has to move to.
Several strategies (choose either b or c):
a. Always serve the first request until it finished (from the starting floor to the destination floor)
b. Serve the closest to the current position of the elevator first, regardless of the elevator direction
c. Serve the closest to the current position of the elevator in the elevator direction.
Note:
a. Initially the elevator starts at first (ground floor)
b. If there is no request, the elevator moves to the first floor
c. At each time tick, the elevator either stands still, or moves up one floor, or moves down one floor
d. When stops (reaching requesting floor/destination floor), it should stops at least for 5 ticks

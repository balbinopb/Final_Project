/* 2.Lift or elevator scheduling

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
*/

// Member of group:
// 	Balbino/103012350551
// 	Razaqa/103012340376

package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	NMAX         int = 100
	NO_REQUEST   int = -1
	MAX_DISTANCE int = 99999999
)

type Request struct {
	time       int
	startFloor int
	endFloor   int
}

type Elevator struct {
	currentFloor int
	standTicks   int
	destination  int
}

type RequestList [NMAX]Request

// displays the introduction message
func Introduction() {
	clearScreen()
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("     WELCOME TO THE ELEVATOR APPLICATION!     ")
	fmt.Println("This application simulates an elevator system.")
	fmt.Println("==============================================")
	fmt.Println()
	fmt.Println()
	fmt.Println("----------------------------------------------")
	fmt.Println("         Enjoy using the application!         ")
	fmt.Println("----------------------------------------------")
	fmt.Println("         Press Enter to continue...           ")
	fmt.Scanln() // Wait for the user to press Enter
}

// clears the terminal screen
func clearScreen() {
	c := exec.Command("clear") // for Linux and macOS
	c.Stdout = os.Stdout
	c.Run()
}

// clears the terminal screen on Windows
func clearScreenWindows() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

// display the main menu and handles user choice
func mainMenu() {
	var choice int
	var n int

	clearScreen()
	fmt.Println()
	fmt.Println("----------------------------------------------")
	fmt.Println("               CHOICE :                       ")
	fmt.Println("----------------------------------------------")
	fmt.Println("                 1. Elevator                  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("                 2. CreatedBy                 ")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		var req RequestList
		userRequests(&req, &n)
		simulateElevator(req, n)
	case 2:
		CreatedBy()
	default:
		clearScreen()
		mainMenu()
	}
}

// display creator/developer information
func CreatedBy() {
	fmt.Println("--Balbino/103012350551--")
	fmt.Println("--Razaqa/103012340376---")
	fmt.Println("Press Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
	mainMenu()
}

// displays the input guidelines
func inputNote() {
	fmt.Println()
	fmt.Println("----------------------------------------------")
	fmt.Println("               Please note:                   ")
	fmt.Println("----------------------------------------------")
	fmt.Println("          The input must greater than 1       ")
	fmt.Println("          Otherwise will stop                 ")
	fmt.Println("----------------------------------------------")
	fmt.Println()
}

// user input for elevator requests
func userRequests(req *RequestList, n *int) {
	inputNote()
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scan(&req[*n].time, &req[*n].startFloor, &req[*n].endFloor)
		if req[*n].time < 1 || req[*n].startFloor < 1 || req[*n].endFloor < 1 {
			break
		}
	}
}

// find the first request that matches or exceeds the given time
func findFirstRequestAtOrAfterTime(reqs RequestList, currentTime int, n int) int {
	left := 0
	right := n
	for left <= right {
		middle := (left + right) / 2
		if reqs[middle].time >= currentTime {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

// finds the closest request to the current floor of the elevator
func findClosestRequest(reqs RequestList, currentFloor, currentTime, n int) int {
	startIndex := findFirstRequestAtOrAfterTime(reqs, currentTime, n)
	closestIndex := NO_REQUEST
	minDistance := MAX_DISTANCE

	for i := startIndex; i < n; i++ {
		if reqs[i].time <= currentTime && reqs[i].startFloor > 0 {
			distance := abs(reqs[i].startFloor - currentFloor)
			if distance < minDistance {
				minDistance = distance
				closestIndex = i
			}
		}
	}
	return closestIndex
}

// calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// sorts the requests based on their time using selection sort
func sortRequestsByTime(reqs *RequestList, n int) {
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if reqs[j].time < reqs[minIdx].time {
				minIdx = j
			}
		}
		temp := reqs[i]
		reqs[i] = reqs[minIdx]
		reqs[minIdx] = temp
	}
}

// Handles the current request when the elevator reaches the start floor
func processRequest(e *Elevator, req *Request) {
	fmt.Printf("Elevator stopped at floor %d to pick up a passenger to floor %d\n", e.currentFloor, req.endFloor)
	e.standTicks = 5
	e.destination = req.endFloor
	req.startFloor = -1 // Mark the request as handled
}

// move the elevator one floor towards the target floor
func moveToFloor(e *Elevator, targetFloor int) {
	if e.currentFloor < targetFloor {
		e.currentFloor++
	} else if e.currentFloor > targetFloor {
		e.currentFloor--
	}
}

// moves the elevator towards the ground floor if idle
func returnToGroundFloor(e *Elevator) {
	if e.currentFloor != 1 {
		moveToFloor(e, 1)
	}
}

// simulates the operation of the elevator
func simulateElevator(reqs RequestList, n int) {
	elevator := Elevator{currentFloor: 1, standTicks: 0, destination: -1}
	currentTime := 0

	sortRequestsByTime(&reqs, n)

	for {
		clearScreen()
		displayElevatorStatus(elevator, reqs, currentTime, n)

		if elevator.standTicks > 0 {
			elevator.standTicks--
		} else if elevator.currentFloor == elevator.destination {
			elevator.standTicks = 5
			elevator.destination = -1
		} else if elevator.destination != -1 {
			moveToFloor(&elevator, elevator.destination)
		} else {
			closestRequestIdx := findClosestRequest(reqs, elevator.currentFloor, currentTime, n)
			if closestRequestIdx == -1 {
				returnToGroundFloor(&elevator)
			} else {
				closestRequest := &reqs[closestRequestIdx]
				if elevator.currentFloor == closestRequest.startFloor {
					processRequest(&elevator, closestRequest)
				} else {
					moveToFloor(&elevator, closestRequest.startFloor)
				}
			}
		}

		currentTime++
		waitForNextTick()
	}
}

// waits for a short period to simulate the passage of time
func waitForNextTick() {
	fmt.Println("Press Enter to proceed to the next tick...")
	fmt.Scanln()
}

// prints the current status of the elevator and requests
func displayElevatorStatus(e Elevator, reqs RequestList, currentTime, n int) {
	fmt.Printf("Time: %d\n", currentTime)
	fmt.Printf("Elevator is at floor: %d\n", e.currentFloor)
	if e.standTicks > 0 {
		fmt.Printf("Elevator is standing for %d more ticks\n", e.standTicks)
	}
	if e.destination != -1 {
		fmt.Printf("Elevator is moving to floor: %d\n", e.destination)
	}
	fmt.Println("Requests:")
	for i := 0; i < n; i++ {
		if reqs[i].startFloor > 0 {
			fmt.Printf("At time %d, start floor %d, end floor %d\n", reqs[i].time, reqs[i].startFloor, reqs[i].endFloor)
		}
	}
}

// main function
func main() {
	Introduction()
	clearScreen()
	mainMenu()
}

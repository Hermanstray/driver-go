package main

import (
	"Driver-go/elevio"
	"fmt"
)

func main() {

	numFloors := 4

	elevio.Init("localhost:15657", numFloors)

	var current_state State = idle

	var d elevio.MotorDirection = elevio.MD_Up
	//elevio.SetMotorDirection(d)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	for {

		select {
		case a := <-drv_buttons:

			switch current_state {
			case idle:

			case move_between_floors:

			case moving_passing_floor:

			case door_open:

			}

			fmt.Printf("%+v\n", a)
			elevio.SetButtonLamp(a.Button, a.Floor, true)

		case a := <-drv_floors:

			switch current_state {
			case idle:

			case move_between_floors:

			case moving_passing_floor:

			case door_open:

			}

			fmt.Printf("%+v\n", a)
			if a == numFloors-1 {
				d = elevio.MD_Down
			} else if a == 0 {
				d = elevio.MD_Up
			}
			elevio.SetMotorDirection(d)

		case a := <-drv_obstr:

			switch current_state {
			case idle:

			case move_between_floors:

			case moving_passing_floor:

			case door_open:

			}
			fmt.Printf("%+v\n", a)
			if a {
				elevio.SetMotorDirection(elevio.MD_Stop)
			} else {
				elevio.SetMotorDirection(d)
			}

		case a := <-drv_stop:

			switch current_state {
			case idle:

			case move_between_floors:

			case moving_passing_floor:

			case door_open:

			}
			fmt.Printf("%+v\n", a)
			for f := 0; f < numFloors; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}
		}
	}
}

// func checkState() {
// 	checkFloor()
// 	checkMoving()
// 	checkDoorOpen()
// }

type State int

const (
	idle                 State = 0
	move_between_floors        = 1
	door_open                  = 2
	moving_passing_floor       = 3
)

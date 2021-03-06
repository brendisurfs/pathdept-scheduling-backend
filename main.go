package main

import (
	"fmt"

	"github.com/brendisurfs/scheduler-app/scheduler"
	"github.com/brendisurfs/scheduler-app/workers"
)

const MAX_WORKERS_OFF = 2

func main() {

	dm := workers.Worker{
		Name:         "dm",
		StartingSlot: workers.Slot1,
		CurrentSlot:  workers.Slot2,
		GlobalSlot:   workers.Slot2,
		DaysOff:      []int{2, 4, 6, 9, 11, 15, 18},
	}

	j := workers.Worker{
		Name:         "j",
		StartingSlot: workers.Slot2,
		CurrentSlot:  workers.Slot3,
		GlobalSlot:   workers.Slot3,
		DaysOff:      []int{3, 4, 7, 8, 11, 12, 16, 17},
	}

	k := workers.Worker{
		Name:         "k",
		StartingSlot: workers.Slot4,
		CurrentSlot:  workers.Slot1,
		GlobalSlot:   workers.Slot1,
		DaysOff:      []int{6, 7, 8, 11, 12, 16, 17},
	}

	workers := []workers.Worker{dm, j, k}

	scheduler.AddToCalendar(&workers)

	fmt.Println(scheduler.DaySchedule.Day)
}

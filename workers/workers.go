package workers

// enum of slots to make it readable.
const (
	Slot1 = iota
	Slot2
	Slot3
	Slot4
	Slot5
)

type Scheduler interface {
	isOff() bool
	setDaysOff() []int
}

type Worker struct {
	Name         string        `json:"name"`
	StartingSlot int           `json:"StartingSlot"`
	CurrentSlot  int           `json:"CurrentSlot"`
	GlobalSlot   int           `json:"GlobalSlot"`
	WorkDays     []interface{} `json:"WorkDays"`
	DaysOff      []int         `json:"DaysOff"`
}

type NextSlotValue struct {
	Off  bool
	Slot int
}

// IsOff - checks if worker is off, returns boolean.
func (w *Worker) IsOff() bool {
	// need to check all the days, which days are labeled off.
	for i := range w.DaysOff {
		if w.CurrentSlot == int(w.DaysOff[i]) {
			return true
		}
	}
	return false
}

// IncrementWorkerDay - shifts the workers schedule day by 1, or back to 0 if greater than
func (w *Worker) IncrementWorkerDay() {
	if w.CurrentSlot+1 >= 5 {
		w.CurrentSlot = Slot1
	} else {
		w.CurrentSlot++
	}
}

// GetNextSlot - determine what the next slot will be || Off day.
func (w *Worker) GetNextSlot() int {
	ns := w.CurrentSlot + 1
	return ns
}

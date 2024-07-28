package others

type Race struct {
	Target     int
	Position   []int
	initialPos []int
	Speed      []int
	Finished   int
	Fleets     int
	fleetGroup []int
}

func carFleet(target int, position, speed []int) int {
	totalCars := len(position)
	r := Race{
		Target:     target,
		Position:   position,
		initialPos: position,
		Speed:      speed,
		fleetGroup: make([]int, totalCars),
	}

	for r.Finished < totalCars {
		r.moveCars()
		r.limitOvertaking()
		r.checkNewFleets()
	}

	return r.Fleets
}

func (r *Race) moveCars() {
	for i := 0; i < len(r.Position); i++ {
		// ignore already finished cars
		if r.Position[i] < r.Target {
			r.Position[i] += r.Speed[i]
		}
	}
	// fmt.Printf("%v\n", r.Position)
}

func (r *Race) checkNewFleets() {
	updateFleets := false

	for i := 0; i < len(r.Position); i++ {
		if r.Position[i] >= r.Target && r.fleetGroup[i] != 1 {
			r.Finished++
			r.fleetGroup[i] = 1
			updateFleets = true
		}
	}

	if updateFleets {
		r.Fleets++
	}
}

func (r *Race) limitOvertaking() {
	for i := 0; i < len(r.Position); i++ {
		tl := r.tailingList(i)
		if len(tl) < 1 {
			continue
		}
	}
}

func (r *Race) tailingList(idx int) []int {
	tailedBy := []int{}
	for i, v := range r.Position {
		if v < r.Position[idx] {
			tailedBy = append(tailedBy, i)
		}
	}
	return tailedBy
}

package birdwatcher

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (s int) {

	for _, v := range birdsPerDay {
		s += v
	}
	fmt.Println(" Birds count: ", s)
	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) (s int) {

	for i, v := range birdsPerDay {
		fmt.Println(" Week: ", week, " Sum: ", s, " Value: ", v)
		if week == 0 {
			return
		} else if week == 1 {
			s += v
		}
		if (i != 0) && (i%7) == 0 {
			week -= 1
		}
	}
	return
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}

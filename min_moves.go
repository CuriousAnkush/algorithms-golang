package algorithms

import "math"

func minMovesToSeat(seats []int, students []int) int {
	var totalMoves int
	var posSeatCountMap [101]int
	var studentPosMap [101]int

	for idx, element := range seats {
		posSeatCountMap[element]++
		studentPosMap[students[idx]]++
	}
	studentCursor, seatCursor := 0, 0

	for studentCursor < 101 && seatCursor < 101 {
		for studentCursor < 101 && studentPosMap[studentCursor] == 0 {
			studentCursor++
		}
		for seatCursor < 101 && posSeatCountMap[seatCursor] == 0 {
			seatCursor++
		}

		if seatCursor < 101 && studentCursor < 101 {
			movesToMake := studentCursor - seatCursor
			totalMoves = totalMoves + int(math.Abs(float64(movesToMake)))
			studentPosMap[studentCursor]--
			posSeatCountMap[seatCursor]--
		}
	}

	return totalMoves
}

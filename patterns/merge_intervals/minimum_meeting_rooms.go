package mergeintervals

import (
	"math"
	"sort"
)

func FindMinimumMeetingRooms(meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	rooms := map[int][]int{}
	totalRooms := 1
	rooms[totalRooms] = meetings[0]
	for i := 1; i < len(meetings); i++ {
		a := meetings[i]
		isMeetingOverlap := false
		roomId := 1
		for room, meeting := range rooms {
			if a[0] < meeting[1] {
				isMeetingOverlap = true
			} else if a[0] == meeting[1] {
				isMeetingOverlap = false
				roomId = room
				break
			}
		}
		if isMeetingOverlap {
			totalRooms++
			rooms[totalRooms] = a
		} else {
			b := rooms[totalRooms]
			rooms[roomId][0] = int(math.Min(float64(a[0]), float64(b[0])))
			rooms[roomId][1] = int(math.Max(float64(a[1]), float64(b[1])))
		}
	}
	return totalRooms
}

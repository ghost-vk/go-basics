// ASSIGNMENT
//
// We've been asked to "bucket" costs per day, in a given period.
//
// Complete the getCostsByDay() function using the append() function. It
// accepts a slice of cost structs and returns a slice of float64s, where each
// float64 represents the total cost for a day.
//
// The length of the daily costs slice should be the number of days contained
// in the costs slice, up to and including the last day. There should only be
// one "bucket" of costs per day. Be sure to include entries for intermediate
// days, even if they don't have any costs.
//
// Days are numbered and start at 0.
//
// EXAMPLE
// Given this input:
//
// []cost{
//     {0, 4.0},
//     {1, 2.1},
//     {5, 2.5},
//     {1, 3.1},
// }
// We expect this result:
//
// []float64{
//     4.0, // first day
//     5.2, // 2.1 + 3.1
//     0.0, // intermediate days with no costs
//     0.0, // ...
//     0.0, // ...
//     2.5, // last day
// }

package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	if len(costs) == 0 {
		return []float64{}
	}

	// Сначала пытался сделать через индексное обращение + make
	// Этот код отрабатывает, время такое:
	// ok  	gobasics/27-append	0.386s
	//
	// costsCount := len(costs)
	//
	// lastDayIdx := 0
	// for j := 0; j < costsCount; j++ {
	// 	if costs[j].day > lastDayIdx {
	// 		lastDayIdx = costs[j].day
	// 	}
	// }
	//
	// cGroup := make([]float64, lastDayIdx+1)
	//
	// for d := 0; d <= lastDayIdx; d++ {
	// 	for c := 0; c < len(costs); c++ {
	// 		if d == costs[c].day {
	// 			cGroup[d] += costs[c].value
	// 		}
	// 	}
	// }
	//
	// return cGroup
	// ======================================
	// Теперь через append
	// cGroup := []float64{}
	//// Можно и так:
	//// cGroup := make([]float64, 1)
	// for i := 0; i < len(costs); i++ {
	// 	curDay := costs[i].day
	// 	if len(cGroup) < curDay+1 {
	// 		toAdd := make([]float64, curDay+1-len(cGroup))
	// 		cGroup = append(cGroup, toAdd...)
	// 		cGroup[curDay] = costs[i].value
	// 	} else {
	// 		cGroup[curDay] += costs[i].value
	// 	}
	// }
	// return cGroup
	// Тест тоже прошел
	// ======================================
	// Третий способ через while (который for 😅)
	// Наверн самый читаемый, но не самый интересный)))
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for len(costsByDay) <= cost.day {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

package wc

import (
    "sort"
    "math/rand"
)

type WeightedChoice struct {
    Weights []int
}

func (wc *WeightedChoice) linear() interface{} {
        var totals []int
        running_total := 0

        for _, w := range wc.Weights {
                running_total += w
                totals = append(totals, running_total)
        }

        rnd := rand.Float64() * float64(running_total)
        for i, total := range totals {
                if rnd < float64(total) {
                        return wc.Weights[i]
                }
        }
        return nil
}

func (wc *WeightedChoice) binarySearch() interface{} {
    totals := []int{}
    running_total := 0

    for _, w := range wc.Weights {
        running_total += w
        totals = append(totals, running_total)
    }

    rnd := rand.Float64() * float64(running_total)
    i := sort.Search(len(totals), func(i int) bool { return float64(totals[i]) >= rnd } )
    return wc.Weights[i]
}

func (wc *WeightedChoice) linearNoTotals() interface {} {
    running_totals := 0
    for _, v := range wc.Weights {
        running_totals += v
    }
    rnd := rand.Float64() * float64(running_totals)
    for i, w := range wc.Weights {
        rnd -= float64(w)
        if rnd < 0 {
            return wc.Weights[i]
        }
    }
    return nil
}

func (wc *WeightedChoice) kingOftheHill() interface{} {
    total, rnd := 0, 0
    for i, w := range wc.Weights {
        total += w
        if rand.Float64() * float64(total) < float64(w) {
            rnd = i
        }
    }
    return wc.Weights[rnd]
}

package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	rate       = flag.Float64("rate", 0.0325, "the lending rate (one year)")
	months     = flag.Int("months", 15*12, "the months of your loan")
	limit      = flag.Float64("limit", 1200000, "you loan limit")
	repayments = flag.Float64("repayments", 5568, "how much pay back money per month")
	startDate  = flag.String("start-date", "2019-10-01T00:00:00Z", "start date time")
)

type info struct {
	Date          time.Time
	Left          float64
	Profit        float64
	Principal     float64
	AlreadyProfit float64
}

func print(infos []*info) {
	for i, info := range infos {
		fmt.Printf("第%5d个月 profit: %5.2f, principal: %5.2f: left: %5.2f , already profit: %5.2f \n ",
			i+1, info.Profit, info.Principal, info.Left, info.AlreadyProfit)
	}
}

func calculate(infos []*info, rate, limit, repayments float64, months int, start time.Time) ([]*info, float64, float64) {
	allProfit := float64(0)
	for ; months > 1; months-- {
		rProfit := limit * rate / 12
		allProfit += rProfit
		rPrincipal := repayments - rProfit
		rLeft := limit - rPrincipal
		limit = rLeft
		infos = append(infos, &info{
			Left:          rLeft,
			Profit:        rProfit,
			Principal:     rPrincipal,
			AlreadyProfit: allProfit,
		})
	}
	return infos, allProfit, limit
}

func do(rate, limit, repayments float64, months int, start time.Time) {
	infos := make([]*info, 0)
	infos, pro, limit := calculate(infos, rate, limit, repayments, months, start)
	print(infos)
	fmt.Printf("the last month you should pay: %5.2f \n", limit)
	fmt.Printf("you should pay profit: %5.2f \n", pro)
}

func main() {
	flag.Parse()
	fmt.Println("processing ...")
	start, err := time.Parse(time.RFC3339, *startDate)
	if err != nil {
		panic(err)
	}
	do(*rate, *limit, *repayments, *months, start)

	fmt.Println("-------------------- now ------------------")

	haveDone := 0
	// all: 12 * 15 = 180
	// 02: 2018.11
	haveDone += 2
	// 14: 2019.12
	haveDone += 12
	// 26: 2020.12
	haveDone += 12
	// 38: 2021.12
	haveDone += 12
	// 50: 2022.12
	haveDone += 12
	// 56: 2023.06
	haveDone += 6
	fmt.Printf("have done: %v", haveDone)
	//        ---- 1046748.33
	_months := 15*12 - haveDone
	_limit := 749452.93 - 000000
	//_repayments := 6500.0
	_repayments := 5445.0
	_rate := 0.0310
	_start, _ := time.Parse(time.RFC3339, "2023-03-01T00:00:00Z")
	do(_rate,
		_limit,
		_repayments,
		_months,
		_start)
}

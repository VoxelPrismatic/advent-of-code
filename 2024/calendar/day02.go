package calendar

import (
	"fmt"
	"strconv"
	"strings"
)

type day02_report []int

func (report day02_report) Increasing() bool {
	if len(report) == 0 {
		return false
	}

	prev := report[0]
	for i := 1; i < len(report); i++ {
		if report[i] >= prev {
			return false
		}
		prev = report[i]
	}
	return true
}

func (report day02_report) Decreasing() bool {
	if len(report) == 0 {
		return false
	}

	prev := report[0]
	for i := 1; i < len(report); i++ {
		if report[i] <= prev {
			return false
		}
		prev = report[i]
	}
	return true
}

func (report day02_report) Directional() bool {
	return report.Increasing() || report.Decreasing()
}

func (report day02_report) LargestDiff() int {
	if len(report) == 0 {
		return 0
	}

	prev := report[0]
	ret := 0
	diff := 0
	for _, v := range report {
		if v >= prev {
			diff = v - prev
		} else {
			diff = prev - v
		}

		if diff > ret {
			ret = diff
		}
		prev = v
	}
	return ret
}

func (report day02_report) Valid() bool {
	return report.LargestDiff() <= 3 && report.Directional()
}

func day02_reports(lines []string) []day02_report {
	reports := []day02_report{}
	for _, line := range lines {
		report := day02_report{}
		for _, part := range strings.Split(line, " ") {
			if part == "" {
				continue
			}
			i, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	return reports
}

func Day02_Part1(lines []string) {
	reports := day02_reports(lines)
	count := 0
	for _, report := range reports {
		if report.Valid() {
			count++
		}
	}

	fmt.Println(count)
}

func Day02_Part2(lines []string) {
	reports := day02_reports(lines)
	count := 0
	for _, report := range reports {
		new_report := day02_report{}
		_ = copy(new_report, report)

		for i := 0; i < len(report) && !new_report.Valid(); i++ {
			new_report = day02_report{}
			for j := 0; j < len(report); j++ {
				if j == i {
					continue
				}
				new_report = append(new_report, report[j])
			}
		}

		if new_report.Valid() {
			count++
		}
	}

	fmt.Println(count)
}

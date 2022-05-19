package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0 
	for _, integer := range b.amount {
		sum += int(integer)
	}
	return sum / len(b.amount)
}


func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0 
	for i := 0 ; i < len(c.temp) ; i++ {
		sum += int(c.temp[i])
	}
	return sum / len(c.temp)
}


func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0 
	for _, integer := range m.amount {
		sum += int(integer)
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}


func main() {
	bw := BandwidthUsage{[]Bytes{5000, 1000, 13000, 8000, 9000}}
	tmp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{8000, 80000, 81000, 85000, 90000}}

	dash := Dashboard{
		BandwidthUsage: bw,
		CpuTemp: tmp,
		MemoryUsage: memory,
	}

	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageBandwidth())
	fmt.Printf("Average cpu temp: %v\n", dash.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dash.AverageMemoryUsage())
}
package main

import (
	"fmt"
	"sort"
)

type Mass float64
type Miles float64

/*// class Employee
type Employee struct {
	Name string
	ID   string
	SSN  int
	Age  int
}

// ToString method
func (employee Employee) ToString() string {
	return fmt.Sprintf("%s: %d,%s,%d", employee.Name, employee.Age, employee.ID, employee.SSN)
}

// SortByAge type
type SortByAge []Employee

func (sortIntf SortByAge) Len() int               { return len(sortIntf) }
func (sortIntf SortByAge) Swap(i int, j int)      { sortIntf[i], sortIntf[j] = sortIntf[j], sortIntf[i] }
func (sortIntf SortByAge) Less(i int, j int) bool { return sortIntf[i].Age < sortIntf[j].Age }*/
type Thing struct {
	name string
	mass Mass
	distance Miles
	meltingpoint int
	freezingpoint int
}


type ThingSorter struct {
	Things []Thing
	byFactor func(Thing1 *Thing, Thing2 *Thing) bool
}

type ByFactor func(Thing1 *Thing, Thing2 *Thing) bool

func (byFactor ByFactor) Sort(Things []Thing)  {
	var sortedThings *ThingSorter
	sortedThings = &ThingSorter{
		Things: Things,
		byFactor: byFactor,
	}
	sort.Sort(sortedThings)
}

func (ThingSorter *ThingSorter) Len() int {
	return len(ThingSorter.Things)
}

func (ThingSorter *ThingSorter) Swap(i, j int)  {
	ThingSorter.Things[i], ThingSorter.Things[j] = ThingSorter.Things[j], ThingSorter.Things[i]
}

func (ThingSorter *ThingSorter) Less(i, j int) bool  {
	return  ThingSorter.byFactor(&ThingSorter.Things[i], &ThingSorter.Things[j])
}
func main()  {

/*	sort.Slice(employees,func(i, j int) bool {
		return employees[i].Age < employees[j].Age
	})
*/
	var Things = []Thing{
		{"IronRod", 0.055, 0.4, 3000, -180},
		{"SteelChair", 0.815, 0.7, 4000, -209},
		{"CopperBowl", 1.0, 1.0, 60, -30},
		{"BrassPot", 0.107, 1.5, 10000, -456},
	}

	var name func(*Thing, *Thing ) bool
	name = func(Thing *Thing, Thing2 *Thing) bool {
		return Thing.name < Thing2.name
	}

	var mass func(*Thing, *Thing ) bool
	mass = func(Thing *Thing, Thing2 *Thing) bool {
		return Thing.mass < Thing2.mass
	}

	ByFactor(name).Sort(Things)
	fmt.Println("By name:", Things)

	ByFactor(mass).Sort(Things)
	fmt.Println("By name:", Things)
}
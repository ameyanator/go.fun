package main

type Vehicle struct {
	strategy DriveStrategy
}

func New(d DriveStrategy) *Vehicle {
	return &Vehicle{
		strategy: d,
	}
}

func (v *Vehicle) drive() {
	v.strategy.drive()
}

func (v *Vehicle) changeStrategy(d DriveStrategy) {
	v.strategy = d
}

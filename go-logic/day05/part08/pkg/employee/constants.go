package employee

const (
	minimunWage = 100.00       // Gaji Minimum
	maximumWage = 1_000_000.00 // Gaji Maximum
)

// declare enum
type Placement string

const (
	INTERNAL  Placement = "INTERNAL"
	OUTSOURCE Placement = "OS"
)
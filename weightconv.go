package weight

import "fmt"

type Kg float64
type Lbs float64

func (lbs Lbs) String() string {
	return fmt.Sprintf("%g lbs\n", lbs)
}

func (kg Kg) String() string {
	return fmt.Sprintf("%g kg\n", kg)
}

// умножим массу на 2.205
func KgToLbs(kg Kg) Lbs {
	return Lbs(kg * 2.205)
}

// разделиим массу на 2.205
func LbsToKg(lbs Lbs) Kg {
	return Kg(lbs / 2.205)
}

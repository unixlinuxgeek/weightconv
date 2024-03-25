package weight

import (
	"math/rand/v2"
	"testing"
)

func TestLbsToKg(t *testing.T) {
	r := genRan(1, 9)
	t.Log(r)
	kg1 := LbsToKg(Lbs(r))
	kg2 := Kg(r / 2.205)

	if kg1 == kg2 {
		t.Logf("Kg1:%g  Kg2:%g", kg1, kg2)
	} else {
		t.Errorf("Test TestLbsToKg failed: %g and %g", kg1, kg2)
	}
}

func TestKgToLbs(t *testing.T) {
	r := genRan(1, 9)
	t.Log(r)
	lbs1 := KgToLbs(Kg(r))
	lbs2 := Lbs(r * 2.205)

	if lbs1 == lbs2 {
		t.Log("Test TestKgToLbs passed!!!")
	} else {
		t.Errorf("Test TestKgToLbs failed: %g and %g", lbs1, lbs2)
	}
}

func genRan(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

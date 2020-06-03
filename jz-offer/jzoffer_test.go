package jz_offer

import "testing"

func TestCuttingRope(t *testing.T) {
	t.Log(cuttingRope(120))
	t.Log(cuttingRope(2))
	t.Log(cuttingRope(3))
	t.Log(cuttingRope(4))
	t.Log(cuttingRope(10))
}

func TestHammingWeight(t *testing.T) {
	t.Log(hammingWeight(-15))
	t.Log(hammingWeight(8))
}

func TestMyPow(t *testing.T) {
	t.Log(myPow(2.0, 3))
}

func TestPrintNumbers(t *testing.T) {
	printNumbers(2)
}

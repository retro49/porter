package main

import (
	"testing"

	"github.com/retro49/porter/plogger"
)

func TestRange(t *testing.T) {
	logger := plogger.NewPlogger()
	t1 := "1..10"
	t11, t12, err := fromRange(t1)
	if err != nil {
		logger.Error("1..10", "error has occured")
		t.Fail()
	}
	if t11 != 1 && t12 != 9 {
		logger.Error("1..10", "value is not correct")
		t.Fail()
	}

	// where am i going with this variable name???
	t2 := "1..=10"
	s3, s4, err := fromRange(t2)
	if err != nil {
		logger.Error("1..=10", "error has occured")
		t.Fail()
	}
	if s3 != 1 && s4 != 10 {
		logger.Error("1..=10", "value is not correct")
		t.Fail()
	}

	t3 := "3..40"
	t31, t32, err := fromRange(t3)
	if err != nil {
		logger.Error("3..40", "error has occured")
		t.Fail()
	}
	if t31 != 3 && t32 != 39 {
		logger.Error("3..40", "value is not correct")
		t.Fail()
	}

	t4 := "3..=40"
	t41, t42, err := fromRange(t4)
	if err != nil {
		logger.Error("3..=40", "error has occured")
		t.Fail()
	}
	if t41 != 3 && t42 != 40 {
		logger.Error("3..=40", "value is not correct")
		t.Fail()
	}

	// the range wont matter here.
	t5 := "50..=10"
	t51, t52, err := fromRange(t5)
	if err != nil {
		logger.Error("50..=10", "error has occured")
		t.Fail()
	}
	if t51 != 50 && t52 != 10 {
		logger.Error("50..=10", "value is not correct")
		t.Fail()
	}

	// must fail
	t6 := "90=..10"
	t61, t62, err := fromRange(t6)
	if err == nil {
                logger.Error("90=..10", "error has occured")
		t.Fail()
	}
	if t61 != -1 && t62 != -1 {
                logger.Error("90=..10", "value is not correct")
		t.Fail()
	}

	t7 := "100..100"
	t71, t72, err := fromRange(t7)
	if err != nil {
                logger.Error("100..100", "error has occured")
		t.Fail()
	}
	if t71 != 100 && t72 != 99 {
                logger.Error("100..100", "value is not correct")
		t.Fail()
	}

	t8 := "100..=100"
	t81, t82, err := fromRange(t8)
	if err != nil {
                logger.Error("100..=100", "error has occured")
		t.Fail()
	}
	if t81 != 100 && t82 != 100 {
                logger.Error("100..=100", "value is not correct")
		t.Fail()
	}
}

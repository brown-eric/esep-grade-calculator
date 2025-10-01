package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestBoundries(t *testing.T) {
	gcA := NewGradeCalculator() // exactly 90%
    gcA.AddGrade("assignment", 90, Assignment)
    gcA.AddGrade("exam", 90, Exam)
    gcA.AddGrade("essay", 90, Essay)
    if gcA.GetFinalGrade() != "A" {
        t.Errorf("Expected A for 90, got %s", gcA.GetFinalGrade())
    }

	gcB := NewGradeCalculator() // exactly 80%
    gcB.AddGrade("assignment", 80, Assignment)
    gcB.AddGrade("exam", 80, Exam)
    gcB.AddGrade("essay", 80, Essay)
    if gcB.GetFinalGrade() != "B" {
        t.Errorf("Expected B for 80, got %s", gcB.GetFinalGrade())
    }
	
	gcC := NewGradeCalculator() // exactly 70%
    gcC.AddGrade("assignment", 70, Assignment)
    gcC.AddGrade("exam", 70, Exam)
    gcC.AddGrade("essay", 70, Essay)
    if gcC.GetFinalGrade() != "C" {
        t.Errorf("Expected C for 70, got %s", gcC.GetFinalGrade())
    }

	gcD := NewGradeCalculator() // exactly 60%
    gcD.AddGrade("assignment", 60, Assignment)
    gcD.AddGrade("exam", 60, Exam)
    gcD.AddGrade("essay", 60, Essay)
    if gcD.GetFinalGrade() != "D" {
        t.Errorf("Expected D for 60, got %s", gcD.GetFinalGrade())
    }

	gcF := NewGradeCalculator() // below 60%
    gcF.AddGrade("assignment", 50, Assignment)
    gcF.AddGrade("exam", 59, Exam)
    gcF.AddGrade("essay", 40, Essay)
    if gcF.GetFinalGrade() != "F" {
        t.Errorf("Expected F for <60, got %s", gcF.GetFinalGrade())
    }
}

func TestEmpty(t *testing.T) {
	gc := NewGradeCalculator()

	// no grades = assumed to be fail
    if gc.GetFinalGrade() != "F" {
        t.Errorf("Expected F with no grades, got %s", gc.GetFinalGrade())
    }

    // assignment only, assumed to be A
    gc1 := NewGradeCalculator()
    gc1.AddGrade("assignment1", 100, Assignment)
    if gc1.GetFinalGrade() != "A" {
        t.Errorf("Expected A with only assignment=100, got %s", gc1.GetFinalGrade())
    }

	// exam only, assumed to be A
    gc2 := NewGradeCalculator()
    gc2.AddGrade("exam1", 90, Exam)
    if gc2.GetFinalGrade() != "A" {
        t.Errorf("Expected A with only exam=90, got %s", gc2.GetFinalGrade())
    }

    // essay only, assumed to be B
    gc3 := NewGradeCalculator()
    gc3.AddGrade("essay1", 80, Essay)
    if gc3.GetFinalGrade() != "B" {
        t.Errorf("Expected B with only essay=80, got %s", gc3.GetFinalGrade())
    }
}

func TestMultipleGradesAverage(t *testing.T) {
    gc := NewGradeCalculator()
    // Assignments average = 90
    gc.AddGrade("assignment1", 100, Assignment)
    gc.AddGrade("assignment2", 80, Assignment)

    // Exams average = 80
    gc.AddGrade("exam1", 70, Exam)
    gc.AddGrade("exam2", 90, Exam)

    // Essays average = 80
    gc.AddGrade("essay1", 60, Essay)
    gc.AddGrade("essay2", 100, Essay)

    grade := gc.GetFinalGrade()
    if grade != "B" {
        t.Errorf("Expected B with mixed averages, got %s", grade)
    }
}

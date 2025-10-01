package esepunittests

type GradeCalculator struct {
	// assignments []Grade
	// exams       []Grade
	// essays      []Grade
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		// assignments: make([]Grade, 0),
		// exams:       make([]Grade, 0),
		// essays:      make([]Grade, 0),
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	// switch gradeType {
	// case Assignment:
	// 	gc.assignments = append(gc.assignments, Grade{
	// 		Name:  name,
	// 		Grade: grade,
	// 		Type:  Assignment,
	// 	})
	// case Exam:
	// 	gc.exams = append(gc.exams, Grade{
	// 		Name:  name,
	// 		Grade: grade,
	// 		Type:  Exam,
	// 	})
	// case Essay:
	// 	gc.essays = append(gc.essays, Grade{
	// 		Name:  name,
	// 		Grade: grade,
	// 		Type:  Essay,
	// 	})
	// }
	gc.grades = append(gc.grades, Grade{
        Name:  name,
        Grade: grade,
        Type:  gradeType,
    })
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	// assignment_average := computeAverage(gc.assignments)
	// exam_average := computeAverage(gc.exams)
	// essay_average := computeAverage(gc.exams)
	assignmentAvg := computeAverageByType(gc.grades, Assignment)
    examAvg := computeAverageByType(gc.grades, Exam)
    essayAvg := computeAverageByType(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

// func computeAverage(grades []Grade) int {
// 	sum := 0

// 	for grade, _ := range grades {
// 		sum += grade
// 	}

// 	return sum / len(grades)
// }
func computeAverageByType(grades []Grade, targetType GradeType) int {
    sum := 0
    count := 0
    for _, g := range grades {
        if g.Type == targetType {
            sum += g.Grade
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return sum / count
}


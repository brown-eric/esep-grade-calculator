package esepunittests

type GradeCalculator struct {
	grades []Grade // ✅ single slice for all grades
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
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAvg := computeAverageByType(gc.grades, Assignment)
	examAvg := computeAverageByType(gc.grades, Exam)
	essayAvg := computeAverageByType(gc.grades, Essay)

	weightedGrade := float64(assignmentAvg)*0.50 +
		float64(examAvg)*0.35 +
		float64(essayAvg)*0.15

	return int(weightedGrade)
}

// helper to compute averages for a specific grade type
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

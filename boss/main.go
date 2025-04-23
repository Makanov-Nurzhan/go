package main

import "fmt"

type Person interface {
	Greet() string
}
type Student struct {
	Name   string
	Age    int
	Grades map[string][]int
}

func (st *Student) AddGrade(subject string, grade int) {
	if st.Grades[subject] == nil {
		st.Grades = make(map[string][]int)
	}
	st.Grades[subject] = append(st.Grades[subject], grade)
}

func (st *Student) Average(subject string) float64 {
	var sum int
	if st.Grades[subject] != nil {
		for _, grade := range st.Grades[subject] {
			sum += grade
		}
		return float64(sum) / float64(len(st.Grades[subject]))
	}
	return -1
}

func (st *Student) Greet() string {
	msg := fmt.Sprintf("Привет, я %s, мне %d лет", st.Name, st.Age)
	return msg
}

func main() {
	student1 := Student{
		Name: "Jack",
		Age:  19,
		Grades: map[string][]int{
			"Math": {40, 50, 56},
			"CS":   {30, 20, 66},
		},
	}
	student2 := Student{
		Name: "Bob",
		Age:  20,
		Grades: map[string][]int{
			"Math": {70, 50, 76},
			"CS":   {10, 80, 36},
		},
	}
	student3 := Student{
		Name: "Lia",
		Age:  25,
		Grades: map[string][]int{
			"Math": {70, 20, 96},
			"CS":   {20, 59, 48},
		},
	}
	//student1.AddGrade("Chemistry", 29)
	var subject = "Math"
	students := []Student{student1, student2, student3}
	for _, student := range students {
		fmt.Println(student.Greet())
	}
	for _, student := range students {
		fmt.Printf("Средний бал %s по %s: %.2f\n", student.Name, subject, student.Average(subject))
	}

}

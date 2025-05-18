package main

import "fmt"

type School struct {
	grades   []Grade
	students []Student
	classes  []Class
	teachers []Teacher
}

type Student struct {
	id   int
	name string
}

type Grade struct {
	id         int
	id_student int
	id_subject int
	value      int
}

type Teacher struct {
	id         int
	id_subject int
	name       string
}

type Class struct {
	students []Student
}

type Entity interface {
	addToSchool(school *School)
	printEntity(school *School)
}

func (curSt Student) printEntity(school *School) {
	fmt.Printf("Ученики:\n")

	for _, student := range school.students {
		fmt.Printf("=ID:%d, Имя:%s\n", student.id, student.name)
	}
}

func (st Student) addToSchool(school *School) {
	school.students = append(school.students, st)
}

func (sch *School) add(entities []Entity) {
	for _, el := range entities {
		el.addToSchool(sch)
	}
}

func (sch *School) init() {
	var student []Entity = []Entity{
		Student{
			id:   1,
			name: "Иван",
		}, Student{
			id:   2,
			name: "Анна",
		},
	}

	sch.add(student)
	student[0].printEntity(sch)
}

func (sch *School) print() {
	fmt.Println(sch)
}

func main() {
	var school School = School{}

	school.init()
	school.print()
}

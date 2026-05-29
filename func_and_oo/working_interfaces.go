package main

import "fmt"

// with interface we define methods
type Food interface {
	Nutrition() string
	FoodType() string
}

type Apple struct {
}

func (a Apple) Nutrition() string {
	return "Apples are carn heavy!"
}

func (a Apple) FoodType() string {
	return "Apples are fruit"
}

type Celery struct {
}

func (c Celery) Nutrition() string {
	return "Celery has zero everything!"
}

func (c Celery) FoodType() string {
	return "Celery is a vergtable."
}

func WorkingWithInterfaces() {
	foods := []Food{Apple{}, Celery{}}
	for _, f := range foods {
		fmt.Println(f.Nutrition(), f.FoodType())
	}
}

package models

import (
	"effective-test/internal/api/schemas"
	"effective-test/internal/config"
	"errors"
	country "github.com/mikekonan/go-countries"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

func (p *Person) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	p.Name = name
	return config.Get().Database.Save(&p).Error
}

func (p *Person) SetSurname(surname string) error {
	if surname == "" {
		return errors.New("surname cannot be empty")
	}
	p.Surname = surname
	return config.Get().Database.Save(&p).Error
}

func (p *Person) SetPatronymic(patronymic string) error {
	p.Patronymic = patronymic
	return config.Get().Database.Save(&p).Error
}

func (p *Person) SetAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	p.Age = age
	return config.Get().Database.Save(&p).Error
}

func (p *Person) SetGender(gender string) error {
	if gender != "male" && gender != "female" {
		return errors.New("gender can be male or female")
	}
	p.Gender = gender
	return config.Get().Database.Save(&p).Error
}

func (p *Person) SetNation(countryCode string) error {
	_, ok := country.ByAlpha2Code(country.Alpha2Code(countryCode))
	if !ok {
		return errors.New("invalid country code")
	}
	p.Nation = countryCode
	return config.Get().Database.Save(&p).Error
}

func (p *Person) Find(personID string) error {
	return config.Get().Database.First(&p, personID).Error
}

func (p *Person) AsResponse() schemas.ResponsePerson {
	return schemas.ResponsePerson{
		ID:         p.ID,
		Name:       p.Name,
		Surname:    p.Surname,
		Patronymic: p.Patronymic,
		Age:        p.Age,
		Gender:     p.Gender,
		Nation:     p.Nation,
	}
}

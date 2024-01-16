package service

import (
	"effective-test/internal/api/schemas"
	"effective-test/internal/config"
	"effective-test/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

const ageAPI = "https://api.agify.io"
const genderAPI = "https://api.genderize.io"
const nationAPI = "https://api.nationalize.io"

const itemsPerPage = 10

func getAgeByName(name string) (int, error) {
	requestLink := fmt.Sprintf("%s/?name=%s", ageAPI, name)
	request, err := http.Get(requestLink)
	if err != nil {
		return -1, err
	}
	defer request.Body.Close()

	type ageContract struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	var response ageContract
	err = json.NewDecoder(request.Body).Decode(&response)
	if err != nil {
		return -1, err
	}

	return response.Age, nil
}

func getGenderByName(name string) (string, error) {
	requestLink := fmt.Sprintf("%s/?name=%s", genderAPI, name)
	request, err := http.Get(requestLink)
	if err != nil {
		return "", err
	}
	defer request.Body.Close()

	type genderContract struct {
		Count       int     `json:"count"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float32 `json:"probability"`
	}

	var response genderContract
	err = json.NewDecoder(request.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Gender, nil
}

func getNationByName(name string) (string, error) {
	requestLink := fmt.Sprintf("%s/?name=%s", nationAPI, name)
	request, err := http.Get(requestLink)
	if err != nil {
		return "", err
	}
	defer request.Body.Close()

	type countryNation struct {
		CountryID   string  `json:"country_id"`
		Probability float32 `json:"probability"`
	}

	type nationContract struct {
		Count   int             `json:"count"`
		Name    string          `json:"name"`
		Country []countryNation `json:"country"`
	}

	var response nationContract
	err = json.NewDecoder(request.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	result := countryNation{
		CountryID:   "",
		Probability: -1,
	}
	for _, code := range response.Country {
		if code.Probability > result.Probability {
			result = code
		}
	}
	return result.CountryID, nil
}

func fillPredictions(person *models.Person) error {
	var wg sync.WaitGroup
	wg.Add(3)

	var ageResult int
	var ageErr error
	go func() {
		defer wg.Done()
		ageResult, ageErr = getAgeByName(person.Name)
	}()

	var genderResult string
	var genderErr error
	go func() {
		defer wg.Done()
		genderResult, genderErr = getGenderByName(person.Name)
	}()

	var nationResult string
	var nationErr error
	go func() {
		defer wg.Done()
		nationResult, nationErr = getNationByName(person.Name)
	}()

	wg.Wait()

	if ageErr != nil {
		return ageErr
	}
	person.Age = ageResult

	if genderErr != nil {
		return genderErr
	}
	person.Gender = genderResult

	if nationErr != nil {
		return nationErr
	}
	person.Nation = nationResult

	return nil
}

func CreatePerson(personSchema schemas.CreatePerson) (models.Person, error) {
	if personSchema.Name == "" || personSchema.Surname == "" {
		return models.Person{}, errors.New("name or surname cannot be empty")
	}
	person := models.Person{Name: personSchema.Name, Surname: personSchema.Surname, Patronymic: personSchema.Patronymic}

	err := fillPredictions(&person)
	if err != nil {
		return person, err
	}
	err = config.Get().Database.Save(&person).Error
	return person, err
}

func DeletePersonByID(personID string) error {
	var person models.Person
	err := person.Find(personID)
	if err != nil {
		return err
	}
	return config.Get().Database.Delete(&person).Error
}

func UpdatePerson(personID string, updateSchema schemas.UpdatePerson) error {
	var person models.Person
	err := person.Find(personID)
	if err != nil {
		return err
	}

	err = person.SetName(updateSchema.Name)
	if err != nil {
		return err
	}

	err = person.SetSurname(updateSchema.Surname)
	if err != nil {
		return err
	}

	err = person.SetPatronymic(updateSchema.Patronymic)
	if err != nil {
		return err
	}

	err = person.SetAge(updateSchema.Age)
	if err != nil {
		return err
	}

	err = person.SetGender(updateSchema.Gender)
	if err != nil {
		return err
	}

	err = person.SetNation(updateSchema.CountryCode)
	if err != nil {
		return err
	}

	return nil
}

func GetPersonsWithPaging(pageNum int) ([]models.Person, error) {
	var results []models.Person
	err := config.Get().Database.Offset((pageNum - 1) * itemsPerPage).Limit(itemsPerPage).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

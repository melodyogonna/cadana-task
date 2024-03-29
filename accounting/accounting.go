package accounting

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type salary struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Person struct {
	ID     string `json:"id"`
	Name   string `json:"personName"`
	Salary salary `json:"salary"`
}

type Persons struct {
	Data []Person `json:"data"`
}

// Sort our persons data by salary in place by either ascending or descending order.
// You're expected to pass either "asc" or "desc" to indicate what order the
// persons should be sorted, any other thing returns an error.
func (persons *Persons) SortBySalary(order string) error {
	switch order {
	case "asc":
		persons.sortAsc()

	case "desc":
		persons.sortDesc()

	default:
		return errors.New("Unknown sorting order. please pass either 'asc' or 'desc'")
	}
	return nil
}

func (persons *Persons) sortAsc() {
	sort.Slice(persons.Data, func(i, j int) bool {
		cur := persons.Data[i]
		next := persons.Data[j]

		// convert numerical money values to float
		cv, _ := strconv.ParseFloat(cur.Salary.Value, 32)
		nv, _ := strconv.ParseFloat(next.Salary.Value, 32)

		return cv < nv
	})

}

func (persons *Persons) sortDesc() {
	sort.Slice(persons.Data, func(i, j int) bool {
		cur := persons.Data[i]
		next := persons.Data[j]

		// convert numerical money values to float
		cv, _ := strconv.ParseFloat(cur.Salary.Value, 32)
		nv, _ := strconv.ParseFloat(next.Salary.Value, 32)

		return cv > nv
	})

}

// Group the persons in our data by currency.
func (persons *Persons) GroupByCurrency() map[string][]Person {
	var group map[string][]Person = make(map[string][]Person)

	for _, person := range persons.Data {
		group[person.Salary.Currency] = append(group[person.Salary.Currency], person)
	}

	return group
}

// Filter the persons in our object by the salary they earn.
// The criteria takes an unsigned integer and filters persons with a salary greater than the amount.
// If the salary currency is not in USD it'll attempt to make a conversion by currency an exchangerate API
// to retrieve the exchange rate between salary currency and USD.
// Exchange rates are saved in memory and reused if the same pair is encountered again.
// This implementation is not generalised and will always filter based on the salaries that are greater than the criteria
func (persons *Persons) FilterByCurrency(criteria uint32) ([]Person, error) {
	var filteredPersons []Person
	var currentRates map[string]float64 = make(map[string]float64)

	for _, person := range persons.Data {
		if person.Salary.Currency == "USD" {
			cv, err := strconv.ParseFloat(person.Salary.Value, 64)
			if err != nil {
				return []Person{}, fmt.Errorf("Not a valid number %s for user %s", person.Salary.Value, person.ID)
			}

			if cv >= float64(criteria) {
				filteredPersons = append(filteredPersons, person)
			}
		} else {
			pair := "USD-" + person.Salary.Currency
			rate, ok := currentRates[pair]
			if ok {
				cv, err := strconv.ParseFloat(person.Salary.Value, 64)
				if err != nil {
					return []Person{}, fmt.Errorf("Not a valid number %s for user %s", person.Salary.Value, person.ID)
				}

				usdEquiv := cv * rate
				if usdEquiv >= float64(criteria) {
					filteredPersons = append(filteredPersons, person)
				}
			} else {
				exchangeRate, err := retrieveRate(pair)

				if err != nil {
					log.Print(err)
					return []Person{}, errors.New("Unable to apply filter")
				}

				if exchangeRate <= 0 {
					return []Person{}, fmt.Errorf("Did not return valid exchange rate for - %s", pair)
				}

				cv, err := strconv.ParseFloat(person.Salary.Value, 64)
				if err != nil {
					return []Person{}, fmt.Errorf("Not a valid number %s for user %s", person.Salary.Value, person.ID)
				}

				currentRates[pair] = exchangeRate
				usdEquiv := cv * exchangeRate
				if usdEquiv >= float64(criteria) {
					filteredPersons = append(filteredPersons, person)
				}

			}

		}
	}

	return filteredPersons, nil
}

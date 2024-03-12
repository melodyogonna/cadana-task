package accounting

import (
	"testing"
)

func TestSortBySalaryDesc(t *testing.T) {
	persons := []Person{
		{ID: "id", Name: "Mid", Salary: salary{Currency: "USD", Value: "10.99"}},
		{ID: "id", Name: "High", Salary: salary{Currency: "USD", Value: "20.99"}},
		{ID: "id", Name: "low", Salary: salary{Currency: "USD", Value: "5.00"}},
	}

	data := Persons{Data: persons}

	err := data.SortBySalary("desc")

	highEmployee := data.Data[0]

	if highEmployee.Name != "High" {
		t.Errorf("%s != %s", "High", highEmployee.Name)
	}

	if err != nil {
		t.Errorf("Error returned for valid sort order")
	}

}

func TestSortBySalaryAsc(t *testing.T) {
	persons := []Person{
		{ID: "id", Name: "Mid", Salary: salary{Currency: "USD", Value: "10.99"}},
		{ID: "id", Name: "High", Salary: salary{Currency: "USD", Value: "20.99"}},
		{ID: "id", Name: "Low", Salary: salary{Currency: "USD", Value: "5.00"}},
	}

	data := Persons{Data: persons}

	err := data.SortBySalary("asc")

	highEmployee := data.Data[0]

	if highEmployee.Name != "Low" {
		t.Errorf("%s != %s", "Low", highEmployee.Name)
	}

	if err != nil {
		t.Errorf("Error returned for valid sort order")
	}

}

func TestErrorOnInvalidOrder(t *testing.T) {

	persons := []Person{
		{ID: "id", Name: "Mid", Salary: salary{Currency: "USD", Value: "10.99"}},
		{ID: "id", Name: "High", Salary: salary{Currency: "USD", Value: "20.99"}},
		{ID: "id", Name: "Low", Salary: salary{Currency: "USD", Value: "5.00"}},
	}

	data := Persons{Data: persons}

	err := data.SortBySalary("invalid")

	if err == nil {
		t.Error("Error should be returned for invalid sort order")
	}
}

func TestCurrencyGrouping(t *testing.T) {
	persons := []Person{
		{ID: "id", Name: "Mid", Salary: salary{Currency: "USD", Value: "10.99"}},
		{ID: "id", Name: "High", Salary: salary{Currency: "AUD", Value: "20.99"}},
		{ID: "id", Name: "Low", Salary: salary{Currency: "USD", Value: "5.00"}},
		{ID: "id", Name: "High", Salary: salary{Currency: "GBP", Value: "15.99"}},
	}

	data := Persons{Data: persons}

	grouping := data.GroupByCurrency()

	if len(grouping["USD"]) != 2 {
		t.Error("Items not grouped properly for USD")
	}

	if len(grouping["GBP"]) != 1 {
		t.Error("Items not grouped properly for GBP")
	}

	if len(grouping["AUD"]) != 1 {
		t.Error("Items not grouped properly for AUD")
	}
}

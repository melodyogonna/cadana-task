## Accounting

A simple package that provides some aggregation of an assumed json data in the format:

```json
[{"id": string, "personName": string, "salary": {"value": string (numerical), "currency": string}}]
```

This data can be sorted by salary, grouped by currency, and filtered by some criteria. "personName" is shortened to "Name" when the json is parsed.

### Example

\# _example.go_

```go
package main

import (
    "log"
    "accounting"
)

func main(){
    jsonString := []bytes(`[
    {"id":"12", "personName": "High", "salary": {"value": "100.00", "currency": "USD"}},
    {"id":"12", "personName": "Mid", "salary": {"value": "70.00", "currency": "USD"}},
    {"id":"12", "personName": "Low", "salary": {"value": "50.00", "currency": "USD"}}
    ]`)

    persons, err := accounting.InitPersons(jsonString)
    if err != nil {
        log.Print(err)
        return
    }

    persons.SortBySalary("asc")
    log.Print(persons.Data[0].Name)
}
```

```sh
$ go run .

Low

```

This contains a lot of bugs as all the required validations and checks were not performed.

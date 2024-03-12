# Exchange rate

A simple server that simulates retrieving exchanges for different currencies.

API:
url - http://localhost:8000/exchange-rate POST
payload - `{"currency-pair": "CURRENCY-PAIR"}` - e.g `{"currency-pair": "USD-AUD"}` sample.
content-type - "application/json"

Very naive implementation that doesn't do enough reflective analysis around floats and ints.

Basically calls a test http client that doesn't really simulate timeouts and errors, so channels can potentially block forever.

Running

```sh
$ cd ./exchangerate
$ go build .
$ ./exchangerate

```

A valid request returns

```
{"USD-AUD": 0.66}
```

Which is not actually a valid exchange rate as no calls are made to any external service, instead I implemented a skeleton of how such a service would
be implemented if neccesary.

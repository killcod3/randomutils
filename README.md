# randomutils

[![Go Report Card](https://goreportcard.com/badge/github.com/killcod3/randomutils)](https://goreportcard.com/report/github.com/killcod3/randomutils)
[![GoDoc](https://godoc.org/github.com/killcod3/randomutils?status.svg)](https://godoc.org/github.com/killcod3/randomutils)

A Go package that provides a collection of random number generation functions and UUID (Universally Unique Identifier) generation functions.

## Installation

To install the `randomutils` package, run:

```bash
go get github.com/killcod3/randomutils
```


## Usage

Here is an example of how to use the `randomutils` package:

```go
package main

import (
	"fmt"

	"github.com/<your-username>/randomutils"
)

func main() {
	// Generate a random string with a specified pattern
	randomString, err := randomutils.GetRandomStr("????d?l")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(randomString)

	// Generate a random number within a given range
	randomNumber, err := randomutils.RandNumInRange(10, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(randomNumber)

	// Generate a random hexadecimal string
	randomHex := randomutils.GetRandHex(8)
	fmt.Println(randomHex)

	// Generate random bytes
	randomBytes := randomutils.GetRandBytes(16)
	fmt.Println(randomBytes)

	// Generate a random integer of a specified length
	randomInt, err := randomutils.GetRandInt(8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(randomInt)

	// Generate a UUID of version 1
	uuidv1, err := randomutils.GetUUIDv1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uuidv1)

	// Generate a UUID of version 4
	uuidv4, err := randomutils.GetUUIDv4()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uuidv4)

	// Generate a UUID of version 5
	uuidv5, err := randomutils.GetUUIDv5("6ba7b810-9dad-11d1-80b4-00c04fd430c8", "example")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uuidv5)
}
```

## Contributing

We welcome contributions to the `randomutils` package. If you find any bugs or have any suggestions for new features, please open an issue. If you would like to contribute code, please follow these steps:

1. Fork the repository
2. Create a new branch for your changes
3. Make your changes
4. Run the tests to make sure everything still works
5. Commit your changes with a descriptive commit message
6. Push your changes to your fork
7. Open a pull request to the `master` branch of the `randomutils` repository

## License

The `randomutils` package is released under the MIT License. See the LICENSE file for more details.

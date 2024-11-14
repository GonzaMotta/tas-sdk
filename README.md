# Go Project with TAS SDK

This project uses the `tas-sdk` package to perform audits and save traces in a service. Below are the steps to download, rename, and implement the basic audit and trace functions.

## Prerequisites

- Go must be installed on your system.
- Have access to GitHub and the ability to download packages from it.

## Installation

1. Download the `tas-sdk` package by running the following command:

    ```bash
    go get github.com/GonzaMotta/tas-sdk
    ```

2. Rename the package in your project to make it easier to use. This is done by adding an alias in your Go code file:

    ```go
    import sdkTas "github.com/GonzaMotta/tas-sdk"
    ```

## Usage

### Create an Audit

To create a new audit, use the `CreateAudit` function provided by the SDK. This method returns a unique `auditId` that you can use to associate with other tracking operations.

```go
auditId, err := sdkTas.CreateAudit() // Example auditId: 673228f4f0579b84116e0b6f

if err != nil {
    log.Println("Error creating audit", err)
}
```

### Save a String Trace

Once you start the CreateAudit() function, you can save the information you require in our system, using SaveTrace() review the example below. 

###  Important, the auditId parameter is optional. If you do not pass it any value, by default it will take the id created in the CreateAudit() function

```go
traceId, err := sdkTas.SaveTrace("Your service name here", "Here your data", auditId ...string ) // returns trace id

if err != nil {
    log.Println("Error creating audit", err)
}
```

### complete example
```go

package main

import (
	"log"

	sdkTas "github.com/GonzaMotta/tas-sdk"
)


func main() {

	_, err := sdkTas.CreateAudit()

	if err != nil {
		log.Println(err)
	}

	trace, err := sdkTas.SaveTrace("My service or process name", "Hi world!")

	if err != nil {
		log.Println(err)
	}

	log.Println("trade id: ", trace)

}

```
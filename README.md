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

Once you have an auditId, you can save a trace in your system using SaveStringTrace. This function takes three parameters: a description of the trace, the service name, and the auditId.

```go
traceId, err := sdkTas.SaveStringTrace("Your service name here", "Here your data", auditId) // returns trace id

if err != nil {
    log.Println("Error creating audit", err)
}
```


### Save an Object Trace

Once you have an auditId, you can save a trace in your system using SaveObjectTrace. This function takes three parameters: a description of the trace, the service name, and the auditId

```go

object := map[string]interface{}{
    "field_1": "value-1",
    "field_2": 1,
    "field_3": "value_3",
}

objectTrace, err := sdkTas.SaveObjectTrace("Your service name here", object, auditId) // returns trace id

if err != nil {
    log.Println("Error saving trace", err)
}
```


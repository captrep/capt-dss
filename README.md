# DSS RESTful API
The DSS is usually used by information systems students to create their final project, if you are a frontend dev and dont want to deal with the 'formula' thing, feel free to use this API. Currently the only DSS method available is MOORA.

## Usage
This API only have one endpoint and receive POST request with payload:
```
{
    "alternative": [
        "A1",
        "A2",
        "A3",
        "A4",
        "A5"
    ],
    "assesment": [
        [30, 50, 40, 50, 30],
        [40, 50, 30, 40, 50],
        [40, 50, 50, 30, 40],
        [20, 20, 30, 50, 40],
        [40, 50, 50, 40, 30]
    ],
    "criteria_type": [
        "benefit", "benefit", "benefit", "cost", "cost"
    ],
    "weight": [
        2.2,
        2.1,
        2.1,
        1.8,
        1.8
    ]
    
}
```

> **Note**
> This API doesnt store the results and only performs calculations, so make sure you have data store for the criteria, alternative, results, etc.
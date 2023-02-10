# Usage
- clone this repository and follow the instructions below

## Run the server first
```zsh
cd server
go run main.go
```
* output 
```zsh
server listenign at [::]:9090
```
## Run the client
* `note` : from another terminal, run the client
```zsh
cd client
go run main.go
```
in the server, you will get the output, 
```zsh
received service: wahidin
```

## Process the service with Rest API on the client side
* `note` : need to provide a curl package
```zsh
curl "http://localhost:3000"
```
* response body
```json
[
  {
    "id": 1,
    "name": "Aji 1",
    "age": 21,
    "address": {
      "line": "Line 1",
      "another_line": "Another Line 1"
    }
  },
  {
    "id": 2,
    "name": "Aji 2",
    "age": 22,
    "address": {
      "line": "Line 2",
      "another_line": "Another Line 2"
    }
  },
  {
    "id": 3,
    "name": "Aji 3",
    "age": 23,
    "address": {
      "line": "Line 3",
      "another_line": "Another Line 3"
    }
  },
  {
    "id": 4,
    "name": "Aji 4",
    "age": 24,
    "address": {
      "line": "Line 4",
      "another_line": "Another Line 4"
    }
  },
  {
    "id": 5,
    "name": "Aji 5",
    "age": 25,
    "address": {
      "line": "Line 5",
      "another_line": "Another Line 5"
    }
  },
  {
    "id": 6,
    "name": "Aji 6",
    "age": 26,
    "address": {
      "line": "Line 6",
      "another_line": "Another Line 6"
    }
  },
  {
    "id": 7,
    "name": "Aji 7",
    "age": 27,
    "address": {
      "line": "Line 7",
      "another_line": "Another Line 7"
    }
  },
  {
    "id": 8,
    "name": "Aji 8",
    "age": 28,
    "address": {
      "line": "Line 8",
      "another_line": "Another Line 8"
    }
  },
  {
    "id": 9,
    "name": "Aji 9",
    "age": 29,
    "address": {
      "line": "Line 9",
      "another_line": "Another Line 9"
    }
  },
  {
    "id": 10,
    "name": "Aji 10",
    "age": 30,
    "address": {
      "line": "Line 10",
      "another_line": "Another Line 10"
    }
  }
]
```


### notes
- [x] the sample contract or protofile goes <a href="https://github.com/WahidinAji/contract">here</a>
- [x] don't forget to eat

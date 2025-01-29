# My Golang App

This is a simple console application written in Go. It demonstrates basic arithmetic operations using utility functions.

## Project Structure

```
my-golang-app
├── cmd
│   └── main.go        # Entry point of the application
├── pkg
│   └── utils
│       └── utils.go   # Utility functions for arithmetic operations
├── go.mod             # Module definition file
└── README.md          # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Building the Application

To build the application, navigate to the project directory and run:

```
go build -o my-golang-app ./cmd
```

### Running the Application

After building, you can run the application with:

```
./my-golang-app
```

### Usage

The application currently supports basic arithmetic operations. You can modify the `main.go` file to implement specific functionalities using the utility functions defined in `pkg/utils/utils.go`. 

## Contributing

Feel free to submit issues or pull requests for improvements or additional features.
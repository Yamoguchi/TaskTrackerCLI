# TaskTrackerCLI

TaskTrackerCLI is a command-line application written in Go that helps you manage and track your tasks efficiently.

## Features

- Add new tasks with descriptions and due dates
- List all tasks with their statuses
- Mark tasks as completed
- Delete tasks
- Filter tasks by status or due date

## Installation

To install TaskTrackerCLI, you need to have Go installed on your machine. Then, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/TaskTrackerCLI.git
    ```
2. Navigate to the project directory:
    ```sh
    cd TaskTrackerCLI
    ```
3. Build the application:
    ```sh
    go build -o tasktracker
    ```

## Usage

After building the application, you can use the `tasktracker` command to manage your tasks. Here are some examples:

- Add a new task:
    ```sh
    ./tasktracker add "Task description" --due "2023-10-15"
    ```
- List all tasks:
    ```sh
    ./tasktracker list
    ```
- Mark a task as completed:
    ```sh
    ./tasktracker complete 1
    ```
- Delete a task:
    ```sh
    ./tasktracker delete 1
    ```
- Filter tasks by status:
    ```sh
    ./tasktracker list --status completed
    ```
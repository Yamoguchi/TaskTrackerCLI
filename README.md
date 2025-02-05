# TaskTrackerCLI
project-url: https://roadmap.sh/projects/task-tracker
TaskTrackerCLI is a command-line application written in Go that helps you manage and track your tasks efficiently.

## Features

- Add new tasks with descriptions
- List all tasks with their statuses
- Mark tasks as completed
- Delete tasks
- Filter tasks by status

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
    ./tasktracker add "Some new task"
    ```
- Update task:
    ```sh
    ./tasktracker update 1 "New text"
    ```
- Marking a task as in progress or done
    ```sh
    ./task-cli mark-in-progress 1
    ./task-cli mark-done 1
    ```
- Delete a task:
    ```sh
    ./tasktracker delete 1
    ```
- List all tasks:
    ```sh
    ./tasktracker list
    ```
- Filter tasks by status:
    ```sh
    ./tasktracker list done
    ./tasktracker list todo
    ./tasktracker list in-progress
    ```
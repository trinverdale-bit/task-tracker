# Task Tracker

A command-line tool written in Go to manage tasks. The tool supports adding, deleting, updating, and listing tasks, as well as marking them as "in progress" or "done". Task data is stored in a local `tasks.json` file.

https://roadmap.sh/projects/task-tracker

## Features
- Add a new task with a description
- Delete a task by its ID
- Update a task's description by its ID
- List all tasks
- List tasks that are marked as "done"
- List tasks that are marked as "todo"
- Mark tasks as "in progress" or "done"

## Installation

1. Clone the repository to your local machine:
   ```
   bash
   git clone https://github.com/trinverdale-bit/task-tracker.git
   ```
   
2. Navigate to the project directory:
   ```
   cd task-tracker
   ```

3. Build the project:
   ```
   go build -o task-tracker
   ```
   This will generate an executable file named `task-tracker`.

## Usage
After building the project, you can run the program with the following commands:

## Add a task 
`./task-tracker add "<description>"`
Adds a new task with the given description.

## Delete a task
`./task-tracker delete <id>`
Deletes the task with the specified ID.

## Update a task
`./task-tracker update <id> "<new description>"`
Updates the description of the task with the specified ID.

## List all tasks
`./task-tracker list`
Lists all tasks.

## Mark a task as "In Progress"
`./task-tracker mark-in-progress <id>`
Marks the task with the specified ID as "todo".

## Mark a task as "Done"
`./task-tracker mark-done <id>`
Marks the task with the specified ID as "done".

## List done tasks
`./task-tracker list-done`
Lists all tasks that are marked as "done".

## List todo tasks
`./task-tracker list-todo`
Lists all tasks that are marked as "todo".

## Show help
`./task-tracker --help`
Displays the available commands and usage instructions.

## Examples
1. Add a task:
`./task-tracker add "Complete Go project"`

2. List tasks:
`./task-tracker list`

3. Mark the task as done:
`./task-tracker mark-done 1`

4. List done tasks:
`./task-tracker list-done`


## Contributions
Feel free to fork the repository, make changes, and submit pull requests. Contributions are welcome!

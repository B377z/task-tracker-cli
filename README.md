# Task Tracker CLI

A simple command-line application to track and manage your tasks.

## Features

1. Add tasks.
2. Update task descriptions.
3. Delete tasks.
4. Mark tasks as `in-progress` or `done`.
5. List tasks (`all`, `todo`, `in-progress`, `done`).

## Installation

1. Build the CLI:
   go build -o task-cli

2. Move the binary to a directory in your PATH (optional)
   mv task-cli /usr/local/bin

## Usage
Add a Task:
    task-cli add "Task description"

Update a Task:
    task-cli update 1 "New task description"

Delete a Task:
    task-cli delete 1

List Tasks:
    task-cli list

## Project URL
[GitHub Repository](https://github.com/B377z/task-tracker-cli)

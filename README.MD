# Task Tracker CLI

Task Tracker is a simple Command Line Interface (CLI) application used to track and manage your tasks. It helps you keep track of what you need to do, what you are currently working on, and what you have already completed.

This project is solution for roadmap.sh backend project and designed to help practice:
- Working with the filesystem
- Handling command-line arguments
- Building a basic CLI application

---

## 🚀 Features

The Task Tracker CLI allows you to:

- Add tasks  
- Update tasks  
- Delete tasks  
- Mark tasks as **in-progress** or **done**  
- List all tasks  
- List tasks by status:
  - `todo`
  - `in-progress`
  - `done`

---

## 📁 Data Storage

- Tasks are stored in a `tasks.json` file in the current directory.
- The file is automatically created if it does not exist.
- No external libraries or frameworks are used — only the native filesystem and standard language features.

---

## 📌 Task Properties

Each task has the following fields:

- `id` – Unique identifier  
- `description` – Short task description  
- `status` – `todo`, `in-progress`, or `done`  
- `createdAt` – Date and time when the task was created  
- `updatedAt` – Date and time when the task was last updated  

---

## ⚙️ Usage

### ➕ Add a Task
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)`
```

### ✏️ Update a Task
```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### 🗑 Delete a Task
```bash
task-cli delete 1
```

### 🔄 Mark Task Status
```bash
task-cli mark-in-progress 1
task-cli mark-done 1
```

### 📋 List Tasks
```bash
task-cli list
```

### 📋 List Tasks by Status
```bash
task-cli list done
task-cli list todo
task-cli list in-progress
```

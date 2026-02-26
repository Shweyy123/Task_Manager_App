# 🗂️ Task Manager App

A full-featured **Task Management Web Application** built using **Go (Golang)** and **MySQL**.  
This application allows users to register, log in securely, and manage their daily tasks efficiently through a clean dashboard interface.

---

## 🚀 Tech Stack

- **Backend:** Go (Golang)
- **Database:** MySQL
- **Frontend:** HTML, CSS, Bootstrap, JS
- **Authentication:** Session-based authentication
- **Architecture:** MVC pattern

---

## ✨ Features

### 🔐 User Authentication

#### 📝 Register Page
- New users can create an account
- Passwords are securely hashed before storing
- Validates user inputs
- Prevents duplicate registrations

#### 🔑 Login Page
- Secure login with email and password
- Session management for authentication
- Error handling for invalid credentials
- Redirects user to dashboard after login

---

## 📊 Dashboard

After successful login, users are redirected to their personal dashboard.

### ✅ Create Task
- Add new tasks with title and description
- Tasks are saved in MySQL database
- Tasks are linked to the logged-in user

### ✏️ Edit Task
- Update task title or description
- Changes are instantly reflected in the database

### ❌ Delete Task
- Remove tasks permanently
- Only the task owner can delete their tasks

### 📋 View Tasks
- Displays all tasks created by the logged-in user
- Organized layout for better task management

---

## 🗄️ Database Structure

### Users Table

| Column   | Type      | Description |
|----------|----------|-------------|
| id       | INT (PK) | Unique user ID |
| username | VARCHAR  | Username |
| email    | VARCHAR  | User email |
| password | VARCHAR  | Hashed password |

### Tasks Table

| Column      | Type      | Description |
|-------------|----------|-------------|
| id          | INT (PK) | Task ID |
| user_id     | INT (FK) | References Users table |
| title       | VARCHAR  | Task title |
| description | TEXT     | Task details |
| created_at  | TIMESTAMP| Creation time |

---

## 🔐 Security Features

- Password hashing
- Session-based authentication
- Route protection for dashboard
- SQL injection prevention using parameterized queries

---

## 🖥️ Application Flow

1. User registers
2. User logs in
3. User accesses dashboard
4. User creates, edits, or deletes tasks
5. User logs out securely

---

## ⚙️ Installation & Setup

### 1️⃣ Clone the Repository

    git clone https://github.com/Shweyy123/Task_Manager_App.git
    cd Task_Manager_App

2️⃣ Install Dependencies

    go mod tidy

3️⃣ Setup MySQL Database

    CREATE DATABASE task_manager;

Update your database credentials inside the configuration file.

4️⃣ Run the Application

    go run main.go

Open your browser and visit:

    http://localhost:8080
    

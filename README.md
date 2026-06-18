# Kahoot Quiz Clone

A real-time multiplayer quiz application inspired by Kahoot, built with Go (Fiber) backend and Svelte frontend.

## Features

-  Host interactive quizzes with multiple-choice questions
- Real-time multiplayer support via WebSockets
-  Timed questions with customizable durations
-  Live leaderboards and score tracking
-  Clean, responsive UI built with Svelte and TailwindCSS

## Tech Stack

### Backend
- **Go** with Fiber web framework
- **MongoDB** for data persistence
- **WebSockets** for real-time communication

### Frontend
- **Svelte** with TypeScript
- **Vite** for build tooling
- **TailwindCSS** for styling
- **svelte-spa-router** for client-side routing

## Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- MongoDB 8.0 or higher

## Installation

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/kahoot-clone.git
cd kahoot-clone
```

### 2. Set up the backend
```bash
cd backend
go mod download
```

### 3. Set up the frontend
```bash
cd ../frontend
npm install
```

### 4. Start MongoDB
Make sure MongoDB is running on `localhost:27017`

```bash
# Windows
net start MongoDB

# macOS/Linux
sudo systemctl start mongod
```

## Running the Application

### Start the Backend
```bash
cd backend
go run cmd/quiz/quiz.go
```
Backend will run on `http://localhost:3000`

### Start the Frontend
```bash
cd frontend
npm run dev
```
Frontend will run on `http://localhost:5173`

## Adding Sample Quizzes

Connect to MongoDB and insert sample data:

```bash
mongosh
```

```javascript
use quiz

db.quizzes.insertOne({
  name: "Sample Quiz",
  questions: [
    {
      id: "q1",
      name: "What is 2 + 2?",
      time: 20,
      choices: [
        { id: "c1", name: "3", correct: false },
        { id: "c2", name: "4", correct: true },
        { id: "c3", name: "5", correct: false },
        { id: "c4", name: "6", correct: false }
      ]
    }
  ]
})
```

## Usage

### As a Host
1. Navigate to `http://localhost:5173/#/host`
2. Select a quiz to host
3. Share the game code with players
4. Start the quiz when players have joined

### As a Player
1. Navigate to `http://localhost:5173/`
2. Enter the game code provided by the host
3. Enter your name
4. Answer questions and compete!

## Project Structure

```
KahootClone/
├── backend/
│   ├── cmd/quiz/          # Application entry point
│   └── internal/
│       ├── collection/    # MongoDB collections
│       ├── controller/    # HTTP handlers
│       ├── entity/        # Data models
│       └── service/       # Business logic
└── frontend/
    └── src/
        ├── lib/           # Reusable components
        ├── model/         # TypeScript types
        ├── service/       # API and network services
        └── views/         # Page components
            ├── host/      # Host views
            ├── player/    # Player views
            └── edit/      # Quiz editing
```

## API Endpoints

- `GET /api/quizzes` - Get all quizzes
- `GET /api/quizzes/:quizId` - Get quiz by ID
- `PUT /api/quizzes/:quizId` - Update quiz
- `GET /ws` - WebSocket connection for real-time gameplay

## Acknowledgments

- Inspired by [Kahoot!](https://kahoot.com/)
- Built as a learning project for Go and Svelte

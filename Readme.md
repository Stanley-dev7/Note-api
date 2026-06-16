 Day 33 Notes API (Go + JWT)

A secure Notes API built with Go, SQLite, and JWT authentication.

 Features

- User registration & login
- JWT authentication
- Password hashing (bcrypt)
- Create notes
- Read notes (user-specific)
- Update notes
- Delete notes
- Protected routes using middleware

Tech Stack

- Go (Golang)
- SQLite
- Gorilla Mux
- JWT (golang-jwt)
- bcrypt


 Authentication Flow
1. Register user
2. Login user
3. Receive JWT token
4. Use token to access protected routes

 API Routes
 Public
- POST /register
- POST /login

Protected (Requires JWT)
- POST /api/notes
- GET /api/notes
- PUT /api/notes/{id}
- DELETE /api/notes/{id}

Challenges & Learning Goals.
- Implementing JWT authentication from scratch
- Working with SQLite database integration in Go
- Securing passwords using bcrypt hashing
- Restricting data access per authenticated user
- Building full CRUD operations (Create, Read, Update, Delete)
- Designing clean folder structure for scalability
- Handling middleware-based route protection
- Managing Go modules and dependencies

Future Improvements
- Add refresh tokens for better security
- Implement input validation middleware
- Add pagination for notes

 Live API
Base url: https://note-api-jlsq.onrender.com

Authentication
- JWT-based authentication system
- Users must login to access notes

 API Endpoints
POST /register  
POST /login  
POST /api/notes  
GET /api/notes  
PUT /api/notes/:id  
DELETE /api/notes/:id
PUT /api/notes/:id  
DELETE /api/notes/:id

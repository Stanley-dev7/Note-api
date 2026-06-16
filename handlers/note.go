package handlers

import (
	"encoding/json"
	"net/http"

	"day33-notes-api/database"
	"day33-notes-api/models"

	"github.com/gorilla/mux"
)

// CREATE NOTE (SECURE)
func CreateNote(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(int)

	var note models.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO notes(title, content, user_id) VALUES(?, ?, ?)",
		note.Title,
		note.Content,
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Note created successfully")
}

// GET NOTES (ONLY LOGGED-IN USER)
func GetNotes(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(int)

	rows, err := database.DB.Query(
		"SELECT id, title, content, user_id FROM notes WHERE user_id = ?",
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note

		err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.UserID,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		notes = append(notes, note)
	}

	json.NewEncoder(w).Encode(notes)
}

// UPDATE NOTE
func UpdateNote(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var note models.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE notes SET title = ?, content = ? WHERE id = ?",
		note.Title,
		note.Content,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Note updated successfully")
}

// DELETE NOTE
func DeleteNote(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := database.DB.Exec(
		"DELETE FROM notes WHERE id = ?",
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Note deleted successfully")
}
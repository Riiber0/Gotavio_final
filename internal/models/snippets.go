package models

import (
	"database/sql"
	"errors"
	"time"
	"strconv"
)

type SnippetModelInterface interface {
	Insert(title string, author int, content string, expires int) (int, error)
	Delete(id int) (error)
	Get(id int) (*Snippet, error)
	Latest() ([]*Snippet, error)
	Search(title string, author string) ([]*Snippet, error)
	GetSaved(id int) ([]*Snippet, error)
	Save(snnipetsID int, usersID int) (error)
	Remove(snnipetsID int, usersID int) (error)
	Exists(snnipetsID int, usersID int) (bool, error)
}

type Snippet struct {
	ID        int
	Title     string
	Author    string
	AuthorID  int
	Content   string
	Created   time.Time
	Expires   time.Time
}

// SnippetModel wraps a sql.DB conn pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, author int, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, author, content, created, expires)
	VALUES(?, ?, ?, DATETIME('now'), DATETIME(DATETIME('now'), '+` + strconv.Itoa(expires) + ` day'))`

	result, err := m.DB.Exec(stmt, title, author, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Delete(id int) (error) {
	stmt := `DELETE FROM snippets WHERE id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Get a specific snippet.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT S.id, S.title, S.author, U.name, S.content, S.created, S.expires
	FROM snippets AS S INNER JOIN users AS U ON S.author = U.id
	WHERE expires > DATETIME('now') AND S.id = ?`

	s := &Snippet{}

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.AuthorID, &s.Author, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Latest return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT S.id, S.title, U.name, S.content, S.created, S.expires
	FROM snippets AS S INNER JOIN users AS U ON S.author = U.id
	WHERE expires > DATETIME('now') ORDER BY S.id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Author, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (m *SnippetModel) Search(title string, author string) ([]*Snippet, error) {
	if title == "" && author == ""  { return nil, nil }

	stmt := ""

	if title == "" {
		stmt = "SELECT S.id, S.title, U.name, S.content, S.created, S.expires FROM snippets AS S INNER JOIN users AS U ON S.author = U.id WHERE U.name LIKE '%" + author + "%'"
	} else if author == "" {
		stmt = "SELECT S.id, S.title, U.name, S.content, S.created, S.expires FROM snippets AS S INNER JOIN users AS U ON S.author = U.id WHERE S.title LIKE '%" + title + "%'"
	} else {
		stmt = "SELECT S.id, S.title, U.name, S.content, S.created, S.expires FROM snippets AS S INNER JOIN users AS U ON S.author = U.id WHERE S.title LIKE '%" + title + "%' AND U.name LIKE '%" + author + "%'"
	}

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}

	var count int = 0

	for rows.Next() {
		count++
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Author, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if count == 0 { return nil, nil }

	return snippets, nil
}

func (m *SnippetModel) GetSaved(id int) ([]*Snippet, error) {
	stmt := `SELECT S.id, S.title, U.name, S.content, S.created, S.expires
	FROM snippets AS S, users AS U, saved as SD
	WHERE ? = U.id AND ? = SD.usersID AND S.id = SD.snippetsID`

	rows, err := m.DB.Query(stmt, id, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Author, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (m *SnippetModel) Save(snippetsID int, usersID int) (error) {
	stmt := `INSERT INTO saved (snippetsID, usersID) VALUES(?, ?)`

	result, err := m.DB.Exec(stmt, snippetsID, usersID)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (m *SnippetModel) Remove(snippetsID int, usersID int) (error) {
	stmt := `DELETE FROM saved WHERE snippetsID = ? AND usersID = ?`

	_, err := m.DB.Exec(stmt, snippetsID, usersID)
	if err != nil {
		return err
	}

	return nil
}

func (m *SnippetModel) Exists(snippetsID int, usersID int) (exists bool, err error) {
	stmt := "SELECT EXISTS(SELECT true FROM saved WHERE snippetsID = ? AND usersID = ?)"

	err = m.DB.QueryRow(stmt, snippetsID, usersID).Scan(&exists)

	return exists, err
}


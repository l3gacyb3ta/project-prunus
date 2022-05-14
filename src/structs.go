package main

import "errors"

type User struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Password string `json:"password"` // NEVER NEVER NEVER DO THIS GOD NO THIS IS JUST A TEST GOD NO
}

type Author struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type Book struct {
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Rating uint8 `json:"rating"`
	Author Author `json:"author"`
}

type Library struct {
	Authors []Author `json:"authors"`
	Books   []Book `json:"book"`
	Owner   *User `json:"owner"`
	Name    string `json:"name"`
}

type DB struct {
	Librarys []Library `json:"libraries"`
	Users    []User `json:"users"`
}

func (d *DB) createUser(username string, password string) {
	newUser := User{Username: username, Bio: "", Password: password}
	d.Users = append(d.Users, newUser)
}

func (d *DB) createLibrary(username string, name string) error {
	for i, user := range(d.Users) { // NOT A GOOD LONG TERM SOLUTION
		if (user.Username == username) {
			d.Librarys = append(d.Librarys, Library{
				Authors: []Author{},
				Books:   []Book{},
				Owner:   &d.Users[i],
				Name:    name,
			})
			return nil
		}
	}
	
	return errors.New("User not found")
}

func (d *DB) createBook(user string, libraryName string, title string, isbn string, rating uint8, author Author) error {
	for _, library := range(d.Librarys) {
		if (library.Owner.Username == user && library.Name == libraryName) { // this is even worse why tf would i do this but im tired and i havent taken my meds so i cant be bothered
			library.Books = append(library.Books, Book{
				Title:  title,
				ISBN:   isbn,
				Rating: rating,
				Author: author,
			})
			return nil
		}
	}

	return errors.New("Library not found")
}
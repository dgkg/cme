package model

import (
	"errors"
	//"github.com/gorilla/sessions"
)

type Session struct {
	Name   string
	Values []string
}

// fontion permettant d'être connecté
func (s Session) Connect() (err error) {
	/*
		session.Options = &sessions.Options{
			Path:     "../cessions/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		session, err := store.Get(r, "session-name")
		// Set some session values.
		session.Values["foo"] = "bar"
		session.Values[42] = 43
		// Save it.
		session.Save(r, w)
	*/
	return errors.New("Erreur")
}

package auth

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	sessionsc "github.com/sargassum-eco/fluitans/internal/clients/sessions"
)

// Identity

func SetIdentity(s *sessions.Session, username string) {
	identity := Identity{
		Authenticated: username != "",
		User:          username,
	}
	s.Values["identity"] = identity
	gob.Register(identity)
}

func GetIdentity(s sessions.Session) (identity Identity, err error) {
	if s.IsNew {
		return
	}

	rawIdentity, ok := s.Values["identity"]
	if !ok {
		// A zero value for Identity indicates that the session has no identity associated with it
		identity = Identity{}
		return
	}
	identity, ok = rawIdentity.(Identity)
	if !ok {
		err = fmt.Errorf("unexpected type for field identity in session")
		return
	}
	return
}

// Access

func Get(c echo.Context, s sessions.Session) (a Auth, err error) {
	a.Identity, err = GetIdentity(s)
	// TODO: if CSRFInput isn't used by a templated page, it'll still change the ETag and thus prevent
	// caching; we need a way to specify whether to blank out CSRFInput based on a.Authorized()
	a.CSRFInput = csrf.TemplateField(c.Request())
	return
}

func GetFromRequest(r *http.Request, s sessions.Session) (a Auth, err error) {
	a.Identity, err = GetIdentity(s)
	a.CSRFInput = csrf.TemplateField(r)
	return
}

func GetWithSession(c echo.Context, sc *sessionsc.Client) (a Auth, s *sessions.Session, err error) {
	s, err = sc.Get(c)
	if err != nil {
		return Auth{}, nil, err
	}
	a, err = Get(c, *s)
	if err != nil {
		return Auth{}, s, err
	}

	return
}
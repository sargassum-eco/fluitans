// Package routes contains the route handlers for the Fluitans server.
package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sargassum-eco/fluitans/internal/app/fluitans/auth"
	"github.com/sargassum-eco/fluitans/internal/app/fluitans/client"
	"github.com/sargassum-eco/fluitans/pkg/framework/route"
	"github.com/sargassum-eco/fluitans/pkg/framework/session"
)

var AuthnPages = []route.Templated{
	{
		Path:         "/login",
		Method:       http.MethodGet,
		HandlerMaker: getLogin,
		Templates:    []string{"auth/login.page.tmpl"},
	},
	{
		Path:         "/sessions",
		Method:       http.MethodPost,
		HandlerMaker: postSessions,
		Templates:    []string{},
	},
}

type LoginData struct {
	NoAuth bool
	ErrorMessages []string
}

func getLogin(g route.TemplateGlobals, te route.TemplateEtagSegments) (echo.HandlerFunc, error) {
	t := "auth/login.page.tmpl"
	err := te.RequireSegments("authn.getLogin", t)
	if err != nil {
		return nil, err
	}

	app, ok := g.App.(*client.Globals)
	if !ok {
		return nil, client.NewUnexpectedGlobalsTypeError(g.App)
	}
	return func(c echo.Context) error {
		// Check authentication & authorization
		a, sess, err := auth.GetWithSession(c, app.Clients.Sessions)
		if err != nil {
			return err
		}

		// Consume & save session
		errorMessages, err := session.GetErrorMessages(sess)
		if err != nil {
			return err
		}
		loginData := LoginData{
			NoAuth: app.Clients.Authn.Config.NoAuth,
			ErrorMessages: errorMessages,
		}
		if err = sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		// Produce output
		return route.Render(c, t, loginData, a, te, g)
	}, nil
}

func postSessions(
	g route.TemplateGlobals, te route.TemplateEtagSegments,
) (echo.HandlerFunc, error) {
	app, ok := g.App.(*client.Globals)
	if !ok {
		return nil, client.NewUnexpectedGlobalsTypeError(g.App)
	}
	sc := app.Clients.Sessions
	return func(c echo.Context) error {
		// Parse params
		method := c.FormValue("method")

		// Run queries
		switch method {
		default:
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf(
				"invalid POST method %s", method,
			))
		case "AUTHENTICATE":
			username := c.FormValue("username")
			password := c.FormValue("password")
			authenticated, err := app.Clients.Authn.CheckCredentials(username, password)
			if err != nil {
				return err
			}
			if !authenticated {
				sess, err := sc.Get(c)
				if err != nil {
					return err
				}

				session.AddErrorMessage(sess, "Could not log in!")
				auth.SetIdentity(sess, "")
				if err = sess.Save(c.Request(), c.Response()); err != nil {
					return err
				}
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			sess, err := sc.Regenerate(c)
			if err != nil {
				return err
			}

			// TODO: implement idle timeout, and implement renewal timeout (if we can). Refer to the
			// "Automatic Session Expiration" section of
			// https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html
			// TODO: regenerate the session upon privilege change
			// TODO: implement idle timeout and automatic renewal timeout
			// TODO: log the session life cycle
			// TODO: add intrusion detection
			auth.SetIdentity(sess, username)
			if err = sess.Save(c.Request(), c.Response()); err != nil {
				return err
			}
			// TODO: redirect to the previous page by getting the path from a form field
		case "DELETE":
			sess, err := sc.Invalidate(c)
			if err != nil {
				return err
			}
			if err := sess.Save(c.Request(), c.Response()); err != nil {
				return err
			}
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}, nil
}

package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
	"github.com/tardog/schutzstreifen/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_schutzstreifen_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.Use(SetCurrentUser)
		app.Use(Authorize)
		app.Use(addMapboxToken)

		app.GET("/", HomeHandler)
		app.GET("/points", PointsHandler)

		app.Resource("/hazards", HazardsResource{})

		adminGroup := app.Group("/")
		adminGroup.Use(authorizeAdmin)
		adminGroup.Resource("/hazard_types", HazardTypesResource{})

		usersRes := UsersResource{}
		usersRoute := adminGroup.Resource("/users", usersRes)

		app.GET("/login", AuthNew)
		app.POST("/login", AuthCreate)
		app.GET("/logout", AuthDestroy)

		// Routes not requiring login
		app.Middleware.Skip(Authorize, HomeHandler, AuthNew, AuthCreate, PointsHandler)
		usersRoute.Middleware.Skip(Authorize, usersRes.New, usersRes.Create)
		usersRoute.Middleware.Skip(authorizeAdmin, usersRes.New, usersRes.Create)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

// addMapboxToken will attempt to find a mapbox token in the environment and set it on the context.
func addMapboxToken(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		mapboxToken, err := envy.MustGet("MAPBOX_API_TOKEN")

		if err != nil {
			return err
		}

		c.Set("mapboxToken", mapboxToken)
		return next(c)
	}
}

// authorizeAdmin will check whether the current user has admin rights
func authorizeAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return errors.WithStack(errors.New("no transaction found"))
		}

		user := &models.User{}
		if err := tx.Find(user, c.Session().Get("current_user_id").(uuid.UUID)); err != nil {
			return c.Redirect(401, "loginPath()")
		}

		if !user.Admin {
			return c.Error(404, errors.New("Nothing to see here"))
		}

		return next(c)
	}
}

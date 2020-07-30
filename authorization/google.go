package authorization

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber"
	"go.uber.org/dig"

	"github.com/markonesgava/take-care/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const JWTTokenURL = "https://oauth2.googleapis.com/token"

var (
	oAuthConfig = oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:3900/auth/oauth2",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

type AuthorizationRoutes struct{}

const oauthStateString = "state"

func ProvideServices(container *dig.Container) error {
	err := provideRoutes(container)

	if err != nil {
		return err
	}

	return container.Invoke(func(_ AuthorizationRoutes) {

	})
}

func provideRoutes(container *dig.Container) error {
	return container.Provide(NewRouter)
}

func NewRouter(app *fiber.App, configuration *config.Configuration) AuthorizationRoutes {
	auth := app.Group("/auth")

	oAuthConfig.ClientID = configuration.Google.ClientID
	oAuthConfig.ClientSecret = configuration.Google.ClientSecret

	auth.Get("/", func(c *fiber.Ctx) {
		url := oAuthConfig.AuthCodeURL(oauthStateString) // TODO change to random value
		c.Redirect(url, fiber.StatusTemporaryRedirect)
	})

	auth.Get("/oauth2", func(c *fiber.Ctx) {
		content, err := getUserInfo(c.FormValue("state"), c.FormValue("code"))
		if err != nil {
			c.Redirect("/", http.StatusTemporaryRedirect)
			return
		}
		c.Send(fmt.Sprintf("Content: %s\n", content))
	})
	return AuthorizationRoutes{}
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := oAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

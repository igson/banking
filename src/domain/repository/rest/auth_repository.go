package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestOAuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type restOAuthRepository struct {
}

//NewRemoteAuthRepository remote permission to access
func NewRemoteAuthRepository() RestOAuthRepository {
	return &restOAuthRepository{}
}

func (r restOAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)
	fmt.Println("Retorno Autorização ----> ", u)
	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false

	} else {
		m := map[string]bool{}

		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			fmt.Println(err)
			return false
		}
		return m["isAuthorized"]

	}

}

/*
  This will generate a url for token verification in the below format
  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}
  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	fmt.Println("Metodo -----> ", "buildVerifyURL")

	u := url.URL{Host: "localhost:8081", Path: "/oauth/verify", Scheme: "http"}

	q := u.Query()

	q.Add("token", token)

	q.Add("routeName", routeName)

	q.Add("username", "2000")

	for k, v := range vars {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()

	return u.String()
}

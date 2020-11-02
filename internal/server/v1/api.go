package v1

import (
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/internal/data"
	review2 "github.com/goandfootball/test-api/internal/data/review"
	user2 "github.com/goandfootball/test-api/internal/data/user"
	"github.com/goandfootball/test-api/internal/server/v1/review"
	"github.com/goandfootball/test-api/internal/server/v1/user"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &user.UsrRouter{
		Repository: &user2.DbUser{
			Data: data.New(),
		},
	}

	rr := &review.RevRouter{
		Repository: &review2.DbReview{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.UserRouter())
	r.Mount("/reviews", rr.ReviewRouter())

	return r
}

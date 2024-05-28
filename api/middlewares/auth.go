package middlewares

import (
	"context"
	"errors"
	"github.com/rikuya98/goTodoApp/apperrors"
	"google.golang.org/api/idtoken"
	"net/http"
	"os"
	"strings"
)

var (
	googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
)

type userNameKey struct{}

func GetUserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{})
	if usernameStr, ok := id.(string); ok {
		return usernameStr
	}
	return ""
}
func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userNameKey{}, name)
	req = req.WithContext(ctx)
	return req
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダから Authorization フィールドを抜き出す
		authorization := req.Header.Get("Authorization")
		// Authorizationフィールドが"Bearer [IDトークン]"の形になっているか検証
		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequireAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequireAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// IDトークン検証
		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CanNotMakeValidator.Wrap(err, "failed to make validator")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.IllegalToken.Wrap(err, "invalid id token")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// nameフィールドをpayloadから抜き出す
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.IllegalToken.Wrap(errors.New("name field not found"), "name field not found")
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		// contextにユーザー名をセット
		req = SetUserName(req, name.(string))

		// 本物のハンドラへ
		next.ServeHTTP(w, req)
	})
}

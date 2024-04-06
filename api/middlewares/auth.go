package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/KakinokiKanta/go-intermediate/apperrors"
	"github.com/KakinokiKanta/go-intermediate/common"
	"google.golang.org/api/idtoken"
)

const (
	googleClientID = "301425239768-0o5svokv0dv7e6da5m94v9446d6hakqh.apps.googleusercontent.com"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダから Authorization フィールドを抜き出す
		authorization := req.Header.Get("Authorization")

		// Authorization フィールドが"Bearer [IDトークン]"の形になっているか検証
		authHeaders := strings.Split(authorization, " ") // 空白区切りで2つに分かれるか
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1] // 空白区切りで分けた1つ目がBearer、2つ目が空ではないか
		if bearer != "Bearer" || idToken != "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// ID トークン検証
		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// name フィールドを payload から抜き出す
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// context にユーザ名をセット
		req = common.SetUserName(req, name.(string))

		// 本物のハンドラへ
		next.ServeHTTP(w, req)
	})
}

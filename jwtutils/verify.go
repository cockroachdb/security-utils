package jwtutils

import (
	"context"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func ValidateAndParseJwt(
	ctx context.Context, jwtString string, jwks string,
) (jwt.Token, error) {
	keySet, err := jwk.Parse([]byte(jwks))
	if err != nil {
		return nil, err
	}

	// Parse the token again but validate it - check for expiration and that
	// it is signed by the right key.
	parsedToken, err := jwt.Parse(
		[]byte(jwtString),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
		jwt.InferAlgorithmFromKey(true),
	)
	if err != nil {
		return nil, err
	}

	//// Validate that at least one of the token audiences matches the issuer config.
	//audienceFound := false
	//for _, audience := range parsedToken.Audience() {
	//	if audience == jwtIssuer.Audience {
	//		audienceFound = true
	//		break
	//	}
	//}
	//if !audienceFound {
	//	return nil, err
	//}

	return parsedToken, nil
}

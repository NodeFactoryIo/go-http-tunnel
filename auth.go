// Copyright (C) 2017 Michał Matczuk
// Use of this source code is governed by an AGPL-style
// license that can be found in the LICENSE file.

package tunnel

// Auth holds user and password.
//type Auth struct {
//	User     string
//	Password string
//}

type Auth struct {
	Token string
}

// NewAuth creates new auth from string representation "user:password".
func NewAuth(token string) *Auth {
	if token == "" {
		return nil
	}

	a := &Auth{
		Token: token,
	}

	return a
}

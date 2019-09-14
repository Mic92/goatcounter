// Copyright © 2019 Martin Tournoij <martin@arp242.net>
// This file is part of GoatCounter and published under the terms of the AGPLv3,
// which can be found in the LICENSE file or at gnu.org/licenses/agpl.html

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebsiteTpl(t *testing.T) {
	tests := []handlerTest{
		{
			name:     "index",
			router:   NewWebsite,
			path:     "/",
			wantCode: 200,
			wantBody: "doesn’t track users;",
		},
		{
			name:     "help",
			router:   NewWebsite,
			path:     "/help",
			wantCode: 200,
			wantBody: "I don’t see my pageviews?",
		},
		{
			name:     "privacy",
			router:   NewWebsite,
			path:     "/privacy",
			wantCode: 200,
			wantBody: "GoatCounter does not collect data which can be used to identify",
		},
		{
			name:     "terms",
			router:   NewWebsite,
			path:     "/terms",
			wantCode: 200,
			wantBody: "The “services” are any software, application, product, or service",
		},

		{
			name:     "status",
			router:   NewWebsite,
			path:     "/status",
			wantCode: 200,
			wantBody: "uptime",
		},

		{
			name:     "signup",
			router:   NewWebsite,
			path:     "/signup/personal",
			wantCode: 200,
			wantBody: `<label for="name">Site name</label>`,
		},
	}

	for _, tt := range tests {
		runTest(t, tt, nil)
	}
}

func TestWebsiteSignup(t *testing.T) {
	tests := []handlerTest{
		{
			name:         "basic",
			method:       "POST",
			router:       NewWebsite,
			path:         "/signup/personal",
			body:         signupArgs{Name: "Example", Code: "example", Email: "m@example.com", UserName: "Example user", TuringTest: "9"},
			wantCode:     303,
			wantFormCode: 303,
		},

		{
			name:         "no-code",
			method:       "POST",
			router:       NewWebsite,
			path:         "/signup/personal",
			body:         signupArgs{Name: "Example", Email: "m@example.com", UserName: "Example user", TuringTest: "9"},
			wantCode:     200,
			wantBody:     "", // TODO: should return JSON
			wantFormCode: 200,
			wantFormBody: "Error: must be set, must be longer than 1 characters",
		},
	}

	for _, tt := range tests {
		runTest(t, tt, func(t *testing.T, rr *httptest.ResponseRecorder, r *http.Request) {
			// TODO: test state
		})
	}
}

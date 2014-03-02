package facebook

import (
	"encoding/json"
	"net/http"
	// "net/mail"
	"net/url"
)

type (
	// UserInfo from /me endpoint. More info:
	// https://developers.facebook.com/docs/graph-api/reference/user
	UserInfo struct {
		ID string `json:"id"`
		// TODO age range
		Bio string `json:"bio"`
		// Birthday	Date	`json:"birthday"`
		// TODO currency
		// TODO education
		// TODO cover
		Email string `json:"email,omitempty"`
		// TODO favorite atheletes
		// TODO favorite teams
		FirstName string `json:"first_name"`
		Gender    string `json:"gender"`
		// TODO Hometown	string	`json:"hometown"`
		// TODO InspirationalPeople
		Installed bool `json:"installed,omitempty"`
		// Indicates if the popular person was manually verified by Facebook
		IsVerified bool `json:"is_verified,omitempty"`
		// TODO languages
		LastName string `json:"last_name"`
		Link     string `json:"link,omitempty"`
		Locale   string `json:"locale,omitempty"`
		// TODO Location
		MiddleName         string `json:"middle_name,omitempty"`
		Name               string `json:"name"`
		NameFormat         string `json:"name_format,omitempty"`
		Political          string `json:"political,omitempty"`
		Quotes             string `json:"quotes,omitempty"`
		RelationshipStatus string `json:"relationship_status,omitempty"`
		Religion           string `json:"religion,omitempty"`
		// TODO SignificantOther	string	`json:"significant_other,omitempty"`
		ThirdPartyID string `json:"third_party_id,omitempty"`
		UserName     string `json:"username,omitempty"`
		Verified     bool   `json:"verified"`
		Website      string `json:"website,omitempty"`
		// TODO work
	}
)

// User retrieves user information. Look here:
// https://developers.facebook.com/docs/graph-api/reference/user
func (g *Graph) User() (*UserInfo, error) {
	var resp *http.Response
	var err error
	user := UserInfo{}

	userURL, _ := url.Parse("https://graph.facebook.com/me")
	query := userURL.Query()
	query.Set("access_token", g.AccessToken)
	userURL.RawQuery = query.Encode()

	if resp, err = http.Get(userURL.String()); err != nil {
		return &user, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return &user, err
	}
	return &user, nil
}

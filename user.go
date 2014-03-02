package facebook

import (
	"encoding/json"
	"errors"
	"net/http"
	// "net/mail"
	"net/url"
	"strconv"
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

	// PictureInfo is a struct for Facebook response.
	PictureInfo struct {
		Data PictureData `json:"data"`
	}
	// PictureData contains information about current user picture
	PictureData struct {
		URL          string `json:"url"`
		IsSilhouette bool   `json:"is_silhouette"`
		Height       int    `json:"height,omitempty"`
		Width        int    `json:"width,omitempty"`
	}
)

// User retrieves user information. Look here:
// https://developers.facebook.com/docs/graph-api/reference/user
func (g *Graph) User() (*UserInfo, error) {
	var resp *http.Response
	var err error
	user := UserInfo{}
	str := "https://graph.facebook.com/me"

	if g.UserID != "" {
		str = "https://graph.facebook.com/" + g.UserID
	}

	endpoint, _ := url.Parse(str)
	query := endpoint.Query()
	query.Set("access_token", g.AccessToken)
	endpoint.RawQuery = query.Encode()

	if resp, err = http.Get(endpoint.String()); err != nil {
		return &user, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return &user, err
	}
	return &user, nil
}

// Picture retrieves PictureInfo about currently used user picture. If
// height or width are set to 0 they aren't used in request. Size can be
// square, small, normal or large. More info:
// https://developers.facebook.com/docs/graph-api/reference/user/picture/
func (g *Graph) Picture(height, width int, size string) (*PictureInfo, error) {
	var resp *http.Response
	var err error
	info := PictureInfo{}

	str := "https://graph.facebook.com/me"

	if g.UserID != "" {
		str = "https://graph.facebook.com/" + g.UserID + "/picture"
	}

	endpoint, _ := url.Parse(str)
	query := endpoint.Query()
	query.Set("redirect", "false")
	if height != 0 {
		query.Set("height", strconv.Itoa(height))
	}
	if width != 0 {
		query.Set("width", strconv.Itoa(width))
	}
	if size != "" {
		if size == "square" || size == "small" || size == "normal" || size == "large" {
			query.Set("type", size)
		} else {
			return &info, errors.New("wrong picture size parameter. square, small, normal or large are allowed")
		}
	}
	endpoint.RawQuery = query.Encode()

	if resp, err = http.Get(endpoint.String()); err != nil {
		return &info, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return &info, err
	}
	return &info, nil
}

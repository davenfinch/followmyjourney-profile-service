package model

// Profile represents a user profile
type Profile struct {
	GUID                 string `json:"guid"`
	Email                string `json:"email"`
	ScreenName           string `json:"screen_name"`
	FirstName            string `json:"first_name,omitempty"`
	ImageURL             string `json:"image_url,omitempty"`
	Locale               string `json:"locale,omitempty"`
	AccountType          string `json:"account_type,omitempty"`
	CharityLink          string `json:"charity_link,omitempty"`
	PaypalSubscriptionId string `json:"paypal_subscription_id,omitempty"`
}

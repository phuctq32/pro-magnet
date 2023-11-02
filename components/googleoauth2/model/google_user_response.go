package ggoauthmodel

type GoogleUserResponse struct {
	EmailAddresses []struct {
		Value string `json:"value"`
	} `json:"emailAddresses,omitempty"`

	Names []struct {
		DisplayName string `json:"displayName"`
	} `json:"names,omitempty"`

	Photos []struct {
		URL string `json:"url"`
	} `json:"photos,omitempty"`

	Birthdays []struct {
		Date struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"date"`
	} `json:"birthdays"`

	PhoneNumbers []struct {
		Value string `json:"value"`
	} `json:"phoneNumbers"`
}

package models

type Specifics struct {
	PhotoUrl   string `json:"photoUrl"`
	ImageId    string `json:"imageId"`
	ImageLevel string `json:"imageLevel"`
	Email      string `json:"email"`
	EmailLevel string `json:"emailLevel"`
	Phone      string `json:"phone"`
	PhoneLevel string `json:"phoneLevel"`
}

func (s Specifics) addEmail(emails string) (result string) {
	if s.Email != "" {
		if emails != "" {
			result = emails + ", "
		}
		result += s.Email
	}
	return
}

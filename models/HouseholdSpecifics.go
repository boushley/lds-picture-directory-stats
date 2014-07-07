package models

type HouseholdSpecifics struct {
	HouseholdInfo   Specifics `json:"householdInfo"`
	HeadOfHousehold Specifics `json:"headOfHousehold"`
	Spouse          Specifics `json:"spouse"`
	CoupleName      string    `json:"coupleName"`
}

func (h HouseholdSpecifics) HasHouseholdPicture() bool {
	return h.HouseholdInfo.PhotoUrl != ""
}

func (h HouseholdSpecifics) HasHeadPicture() bool {
	return h.HouseholdInfo.PhotoUrl != ""
}
func (h HouseholdSpecifics) HasSpousePicture() bool {
	return h.HouseholdInfo.PhotoUrl != ""
}

package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Marriage   string
	Age        string
	XinZuo     string
	Height     string
	Weight     string
	WorkPlace  string
	Income     string
	Occupation string
	Education  string

	HoKou      string
	House      string
	Car        string
}

func FromJsonObj(o interface{})(Profile, error){
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil{
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
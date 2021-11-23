package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const OauthGoogleUrlAPI = "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,genders,birthdays&access_token="

func GetGoogleProfile(token string) (res map[string]interface{},err error) {
	response, err := http.Get(OauthGoogleUrlAPI + token)
	if err != nil {
		fmt.Println("error token")
		return res, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200{
		return res,err
	}
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("error read body")

		return res, err
	}
	err = json.Unmarshal(responseBody,&res)
	if err != nil {
		return res,err
	}

	return res,err
}

func getUser(token string) (res []byte,err error) {
	response, err := http.Get(OauthGoogleUrlAPI + token)
	if err != nil {
		fmt.Println("error token")
		return res, err
	}
	defer response.Body.Close()
	res, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error read body")

		return nil, err
	}

	return res, nil
}
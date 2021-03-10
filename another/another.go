package another

import (
	"github.com/nyaruka/phonenumbers"
)

//IsPhoneNumber Exported Method
func IsPhoneNumber(field string, countryCode string) (bool, string) {
	num, err := phonenumbers.Parse(field, countryCode)
	if err != nil {
		return false, ""
	}

	return phonenumbers.IsValidNumber(num), phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
}

package main

import (
	"fmt"
	"strings"
)

const LoginType = "login"

const (
	EmailLabel    = "EMAIL"
	UsernameLabel = "USERNAME"
	PasswordLabel = "PASSWORD"
	UrlLabel      = "URL"
)

type Login struct {
}

func (login *Login) Generate(items []EnpassItem) [][]string {
	records := make([][]string, 0)

	records = append(records, []string{"title", "website", "username", "password", "notes"})

	for _, item := range items {

		// build the map type -> slice of values
		fieldValuesByLabel := make(map[string][]string, 0)
		for _, field := range item.Fields {
			// skip the field that contains an empty value
			if field.Value == "" {
				continue
			}

			// set the uppercase to Label because "Enpass" exported values contain different values
			label := strings.ToUpper(field.Type)

			if fieldValuesByLabel[label] == nil {
				fieldValuesByLabel[label] = []string{field.Value}
			} else {
				fieldValuesByLabel[label] = append(fieldValuesByLabel[label], field.Value)
			}
		}

		var username string
		var notes string

		// fill the username by email if the not null. In other case
		if len(fieldValuesByLabel[UsernameLabel]) > 0 {
			username = joinValue(fieldValuesByLabel[UsernameLabel])
			if len(joinValue(fieldValuesByLabel[EmailLabel])) > 0 {
				notes = fmt.Sprintf("Emails(s): %s;", joinValue(fieldValuesByLabel[EmailLabel]))
			}
		} else if len(joinValue(fieldValuesByLabel[EmailLabel])) > 0 {
			username = joinValue(fieldValuesByLabel[EmailLabel])
		}

		if len(fieldValuesByLabel[UrlLabel]) > 1 {
			notes = notes + fmt.Sprintf(" Url(s): %s;", joinValue(fieldValuesByLabel[UrlLabel]))
		}
		if len(fieldValuesByLabel[PasswordLabel]) > 1 {
			notes = notes + fmt.Sprintf(" Passwords(s): %s;", joinValue(fieldValuesByLabel[PasswordLabel]))
		}
		website := oneValue(fieldValuesByLabel[UrlLabel])
		password := oneValue(fieldValuesByLabel[PasswordLabel])

		if item.Note != "" {
			notes = notes + item.Note
		}

		records = append(records, []string{item.Title, website, username, password, notes})
	}

	return records
}

func (login *Login) Type() string {
	return LoginType
}

func joinValue(source []string) (result string) {
	for _, v := range source {
		if result == "" {
			result = v
		} else {
			result = result + ", " + v
		}
	}

	return result
}

func oneValue(source []string) (result string) {
	for _, v := range source {
		if result == "" {
			return v
		} else {
			return v
		}
	}
	return ""
}

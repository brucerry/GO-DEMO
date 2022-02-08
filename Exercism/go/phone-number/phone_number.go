package phonenumber

import (
    "fmt"
    "strings"
)

func Number(phoneNumber string) (string, error) {
	s := []string{}
    for _,v := range phoneNumber {
        if v >= '0' && v <= '9' {
            s = append(s, string(v))
        }
    }
	if len(s) == 11 {
        if s[0] != "1" {
    		return "", fmt.Errorf("")
        }
        s2 := strings.Join(s[1:], "")
        if s2[0] < '2' || s2[3] < '2' {
            return "", fmt.Errorf("")
        }
        return s2, nil
    } else if len(s) == 10 {
        if s[0] < "2" || s[3] < "2" {
            return "", fmt.Errorf("")
        }
    	return strings.Join(s, ""), nil
    } else {
    	return "", fmt.Errorf("")
    }
}

func AreaCode(phoneNumber string) (string, error) {
    s,e := Number(phoneNumber)
	if e != nil {
        return "", fmt.Errorf("")
    }
	return s[:3], nil
}

func Format(phoneNumber string) (string, error) {
    s,e := Number(phoneNumber)
	if e != nil {
        return "", fmt.Errorf("")
    }
	s2 := fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:])
    return s2, nil
}

package entity

import (
	"crypto/sha1"
	"regexp"
	"strconv"
	"time"
)

type Customer struct {
	Username string
	Password string
	ID       string
}

type DataCustomer struct {
	Data []Customer
}

func (c *DataCustomer) Getcustomer() []Customer {
	return c.Data
}

func setID() string {
	idTime := time.Now().UnixNano()
	id := strconv.Itoa(int(idTime))
	return id
}

func setPassword(pass, id string) string {
	saltPass := pass + id
	sha := sha1.New()
	sha.Write([]byte(saltPass))
	password := sha.Sum(nil)
	return string(password)
}

func NewCustomer() *DataCustomer {
	return &DataCustomer{}
}

func (c *DataCustomer) Register(username, password string) (string, error) {
	regex, err := regexp.Compile(`[a-zA-Z]{5,}`)
	if err != nil {
		return "", CustomError{MsgError: "Regex error"}
	}

	if !regex.MatchString(password) {
		return "", CustomError{MsgError: "Password must be at least 5 characters"}
	}

	for i := range c.Data {
		if c.Data[i].Username == username {
			return "", CustomError{MsgError: "Username already exist"}
		}
	}
	id := setID()
	password = setPassword(password, id)
	c.Data = append(c.Data, Customer{username, password, id})
	return "Registrasi Success", nil
}

func (c *DataCustomer) Login(username, password string) (string, string, error) {
	for i := range c.Data {
		if c.Data[i].Username == username {
			check := setPassword(password, c.Data[i].ID)
			if check != c.Data[i].Password {
				return "", "", CustomError{MsgError: "Password is wrong"}
			}
			return c.Data[i].Username, "Login Success", nil
		}
	}
	return "", "", CustomError{MsgError: "Username is wrong"}
}

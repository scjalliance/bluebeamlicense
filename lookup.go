package bluebeamlicense

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// License is the license info
type License struct {
	ProductName    string
	SerialNumber   string
	ProductKey     string
	Maintenance    string
	UsersAllowed   int
	UsersInstalled int
	CompanyName    string
	CompanyAddress string
	CompanyCity    string
	CompanyState   string
	CompanyPostal  string
	CompanyCountry string
	ContactName    string
	ContactTitle   string
	ContactEmail   string
	ContactPhone   string
	ResellerName   string
	ResellerPhone  string
	Computers      []Computer
}

// Computer is a computer that has a license attached to it
type Computer struct {
	ID                int
	Name              string
	AuthorizationCode string
	Version           string
	DateAuthorized    time.Time
}

// Lookup will lookup the License
func Lookup(SerialNumber string, ProductKey string, ContactEmail string) (*License, error) {
	/////
	resp, err := http.PostForm("http://www.bluebeam.com/us/license/reglookup.asp", url.Values{
		"sSN":       {SerialNumber},
		"sPK":       {ProductKey},
		"sEM":       {ContactEmail},
		"cmdSubmit": {"Get List"},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc := html.NewTokenizer(resp.Body)
	/////

	/////
	// f, err := os.Open("../test.html")
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()
	// doc := html.NewTokenizer(f)
	/////

	license := new(License)
	m := make(map[string]string)
	currentTag := ""
	section := ""

TokenLoop:
	for {
		tt := doc.Next()

		switch tt {
		case html.ErrorToken:
			break TokenLoop

		case html.StartTagToken:
			t := doc.Token()
			if t.Data == "tr" {
				tdI := 0
				tdKey := ""
				computer := Computer{}
			TDLoop:
				for {
					tt = doc.Next()
					switch tt {
					case html.ErrorToken:
						break TDLoop
					case html.EndTagToken:
						d := doc.Token().Data
						if d == "tr" {
							break TDLoop
						}
					case html.StartTagToken:
						t = doc.Token()
						currentTag = t.Data
						if t.Data == "td" {
							tdI++
						}
					case html.TextToken:
						t = doc.Token()
						if currentTag == "h3" && tdI == 1 {
							key := strings.TrimSpace(t.Data)
							if key != "" {
								section = key
							}
						}
						if currentTag == "td" {
							if section == "Installed Computers" {
								value := strings.TrimSpace(t.Data)
								if value != "" {
									if tdI == 1 {
										computer.ID, _ = strconv.Atoi(value)
									}
									if tdI == 2 {
										computer.Name = value
									}
									if tdI == 3 {
										computer.AuthorizationCode = value
									}
									if tdI == 4 {
										computer.Version = value
									}
									if tdI == 5 {
										computer.DateAuthorized, _ = time.Parse("1/2/2006 3:04:05 PM", value)
										license.Computers = append(license.Computers, computer)
									}
								}
							} else {
								if tdI == 2 {
									key := strings.TrimSpace(t.Data)
									if key != "" {
										tdKey = section + ":" + key
									}
								}
								if tdI == 4 {
									value := strings.TrimSpace(t.Data)
									if value != "" {
										m[tdKey] = value
									}
								}
							}
						}
					}
				}
			}
		}
	}

	license.ProductName = m["License Information:Product Name"]
	license.SerialNumber = m["License Information:Serial Number"]
	license.ProductKey = m["License Information:Product Key"]
	license.Maintenance = m["License Information:Maintenance"]
	license.UsersAllowed, _ = strconv.Atoi(m["License Information:Users Allowed"])
	license.UsersInstalled, _ = strconv.Atoi(m["License Information:Users Installed"])
	license.CompanyName = m["Customer Information:Company Name"]
	license.CompanyAddress = m["Customer Information:Address"]
	license.CompanyCity = m["Customer Information:City"]
	license.CompanyState = m["Customer Information:State/Province"]
	license.CompanyPostal = m["Customer Information:Zip/Postal Code"]
	license.CompanyCountry = m["Customer Information:Country"]
	license.ContactName = m["Customer Contact:Contact"]
	license.ContactTitle = m["Customer Contact:Title"]
	license.ContactEmail = m["Customer Contact:Email"]
	license.ContactPhone = m["Customer Contact:Phone"]
	license.ResellerName = m["Reseller Contact:Name"]
	license.ResellerPhone = m["Reseller Contact:Phone"]

	return license, nil
}

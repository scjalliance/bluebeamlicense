package main

import (
	"fmt"
	"os"

	"github.com/scjalliance/bluebeamlicense"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println(os.Args[0] + " <serial> <product-key> <contact-email>")
		os.Exit(1)
	}

	license, err := bluebeamlicense.Lookup(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("[License Information]")
	fmt.Printf("%20s:\t%s\n", "Product Name", license.ProductName)
	fmt.Printf("%20s:\t%s\n", "Serial Number", license.SerialNumber)
	fmt.Printf("%20s:\t%s\n", "Product Key", license.ProductKey)
	fmt.Printf("%20s:\t%s\n", "Maintenance", license.Maintenance)
	fmt.Printf("%20s:\t%d\n", "Users Allowed", license.UsersAllowed)
	fmt.Printf("%20s:\t%d\n", "Users Installed", license.UsersInstalled)
	fmt.Println("")
	fmt.Println("[Customer Information]")
	fmt.Printf("%20s:\t%s\n", "Company Name", license.CompanyName)
	fmt.Printf("%20s:\t%s\n", "Address", license.CompanyAddress)
	fmt.Printf("%20s:\t%s\n", "Region", license.CompanyState)
	fmt.Printf("%20s:\t%s\n", "Postal Code", license.CompanyPostal)
	fmt.Printf("%20s:\t%s\n", "Country", license.CompanyCountry)
	fmt.Println("")
	fmt.Println("[Customer Contact]")
	fmt.Printf("%20s:\t%s\n", "Contact Name", license.ContactName)
	fmt.Printf("%20s:\t%s\n", "Contact Title", license.ContactTitle)
	fmt.Printf("%20s:\t%s\n", "Contact Email", license.ContactEmail)
	fmt.Printf("%20s:\t%s\n", "Contact Phone", license.ContactPhone)
	fmt.Println("")
	fmt.Println("[Reseller Contact]")
	fmt.Printf("%20s:\t%s\n", "Reseller Name", license.ResellerName)
	fmt.Printf("%20s:\t%s\n", "Reseller Phone", license.ResellerPhone)
	fmt.Println("")
	fmt.Println("[Installed Computers]")
	fmt.Printf("%5s  %-20s  %-10s  %-20s  %-25s\n", "ID", "Name", "Version", "Date Authorized", "Authorization Code")
	fmt.Printf("%5s  %-20s  %-10s  %-20s  %-25s\n", "-----", "--------------------", "----------", "--------------------", "-------------------------")
	for _, computer := range license.Computers {
		fmt.Printf("%5d  %-20s  %-10s  %-20s  %-25s\n", computer.ID, computer.Name, computer.Version, computer.DateAuthorized.Format("2006-01-02 03:04p"), computer.AuthorizationCode)
	}
}

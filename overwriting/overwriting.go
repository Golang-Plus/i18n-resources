package overwriting

import (
	"github.com/golang-plus/errors"
	"github.com/golang-plus/i18n"
)

// Category represents the category of resource.
type Category byte

const (
	LanguageName Category = iota + 1
	CountryName
	CountryAlias
	CurrencyName
	CultureName
)

// Overwrite overwrites the resource for category.
func (r Category) Overwrite(resource *Resource) {
	switch r {
	case LanguageName:
		for code, value := range resource.Values {
			language := i18n.GetLanguage(code)
			if language != nil { // data wrong allowed
				language.Name.SetValue(resource.Language, value)
			}
		}

	case CountryName:
		for code, value := range resource.Values {
			country := i18n.GetCountry(nil, code)
			if country != nil { // data wrong allowed
				country.Name.SetValue(resource.Language, value)
			}
		}

	case CountryAlias:
		for code, value := range resource.Values {
			country := i18n.GetCountry(nil, code)
			if country != nil { // data wrong allowed
				country.Aliases.SetValue(resource.Language, value)
			}
		}

	case CurrencyName:
		for code, value := range resource.Values {
			currency := i18n.GetCurrency(code)
			if currency != nil { // data wrong allowed
				currency.Name.SetValue(resource.Language, value)
			}
		}

	case CultureName:
		for code, value := range resource.Values {
			culture := i18n.GetCulture(code)
			if culture != nil { // data wrong allowed
				culture.Name.SetValue(resource.Language, value)
			}
		}

	default:
		panic(errors.Newf("category #%d is invalid", r))
	}
}

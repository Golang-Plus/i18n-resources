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
			language, ok := i18n.LookupLanguage(code)
			if ok { // data wrong allowed
				language.Name.SetValue(resource.Language, value)
			}
		}

	case CountryName:
		for code, value := range resource.Values {
			country, ok := i18n.LookupCountry(nil, code)
			if ok { // data wrong allowed
				country.Name.SetValue(resource.Language, value)
			}
		}

	case CountryAlias:
		for code, value := range resource.Values {
			country, ok := i18n.LookupCountry(nil, code)
			if ok { // data wrong allowed
				country.Aliases.SetValue(resource.Language, value)
			}
		}

	case CurrencyName:
		for code, value := range resource.Values {
			currency, ok := i18n.LookupCurrency(code)
			if ok { // data wrong allowed
				currency.Name.SetValue(resource.Language, value)
			}
		}

	case CultureName:
		for code, value := range resource.Values {
			culture, ok := i18n.LookupCulture(code)
			if ok { // data wrong allowed
				culture.Name.SetValue(resource.Language, value)
			}
		}

	default:
		panic(errors.Newf("category #%d is invalid", r))
	}
}

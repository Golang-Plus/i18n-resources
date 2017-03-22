package en

import (
	"github.com/golang-plus/i18n"
)

// load resources.
func init() {
	lang := i18n.GetLanguage("en")

	for code, name := range languageNames {
		language := i18n.GetLanguage(code)
		if language != nil {
			language.Name.SetValue(lang, name)
		}
	}

	for code, name := range countryNames {
		country := i18n.GetCountry(nil, code)
		if country != nil {
			country.Name.SetValue(lang, name)
		}
	}

	for code, name := range currencyNames {
		currency := i18n.GetCurrency(code)
		if currency != nil {
			currency.Name.SetValue(lang, name)
		}
	}

	for code, name := range cultureNames {
		culture := i18n.GetCulture(code)
		if culture != nil {
			culture.Name.SetValue(lang, name)
		}
	}
}

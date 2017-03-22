package en

import (
	"github.com/golang-plus/i18n"
)

// load resources.
func init() {
	lang := i18n.GetLanguage("en")

	for code, name := range languageNames {
		language := i18n.GetLanguage(code)
		//if language != nil {  // we need panic if data wrong
		language.Name.SetValue(lang, name)
		//}
	}

	for code, name := range countryNames {
		country := i18n.GetCountry(nil, code)
		//if country != nil {  // we need panic if data wrong
		country.Name.SetValue(lang, name)
		//}
	}

	for code, name := range currencyNames {
		currency := i18n.GetCurrency(code)
		//if currency != nil {  // we need panic if data wrong
		currency.Name.SetValue(lang, name)
		//}
	}

	for code, name := range cultureNames {
		culture := i18n.GetCulture(code)
		//if cculture != nil {  // we need panic if data wrong
		culture.Name.SetValue(lang, name)
		//}
	}
}

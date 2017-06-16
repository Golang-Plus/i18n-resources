package zh_hans

import (
	"github.com/golang-plus/i18n"
)

// load resources.
func init() {
	lang, _ := i18n.LookupLanguage("zh-Hans")

	for code, name := range languageNames {
		language, ok := i18n.LookupLanguage(code)
		if ok {
			language.Name.SetValue(lang, name)
		}
	}

	for code, name := range countryNames {
		country, ok := i18n.LookupCountry(nil, code)
		if ok {
			country.Name.SetValue(lang, name)
		}
	}

	for code, name := range currencyNames {
		currency, ok := i18n.LookupCurrency(code)
		if ok {
			currency.Name.SetValue(lang, name)
		}
	}

	for code, name := range cultureNames {
		culture, ok := i18n.LookupCulture(code)
		if ok {
			culture.Name.SetValue(lang, name)
		}
	}
}

package deepl

// https://www.deepl.com/docs-api/translate-text
// TODO: Retrieve from DeepL API (https://www.deepl.com/docs-api/general/get-languages/)
var (
	lang = []string{
		"Bulgarian",
		"Czech",
		"Danish",
		"German",
		"Greek",
		"English",
		"Spanish",
		"Estonian",
		"Finnish",
		"French",
		"Hungarian",
		"Indonesian",
		"Italian",
		"Japanese",
		"Korean",
		"Lithuanian",
		"Latvian",
		"Norwegian",
		"Dutch",
		"Polish",
		"Portuguese",
		"Romanian",
		"Russian",
		"Slovak",
		"Slovenian",
		"Swedish",
		"Turkish",
		"Ukrainian",
		"Chinese",
	}
	langCode = map[string]string{
		"Bulgarian":  "BG",
		"Czech":      "CS",
		"Danish":     "DA",
		"German":     "DE",
		"Greek":      "EL",
		"English":    "EN",
		"Spanish":    "ES",
		"Estonian":   "ET",
		"Finnish":    "FI",
		"French":     "FR",
		"Hungarian":  "HU",
		"Indonesian": "ID",
		"Italian":    "IT",
		"Japanese":   "JA",
		"Korean":     "KO",
		"Lithuanian": "LT",
		"Latvian":    "LV",
		"Norwegian":  "NB",
		"Dutch":      "NL",
		"Polish":     "PL",
		"Portuguese": "PT",
		"Romanian":   "RO",
		"Russian":    "RU",
		"Slovak":     "SK",
		"Slovenian":  "SL",
		"Swedish":    "SV",
		"Turkish":    "TR",
		"Ukrainian":  "UK",
		"Chinese":    "ZH",
	}
)

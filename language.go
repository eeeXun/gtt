package main

// https://cloud.google.com/translate/docs/languages
var (
	Lang = []string{
		"Afrikaans",
		"Albanian",
		"Amharic",
		"Arabic",
		"Armenian",
		"Azerbaijani",
		"Basque",
		"Belarusian",
		"Bengali",
		"Bosnian",
		"Bulgarian",
		"Catalan",
		"Cebuano",
		"Chinese (Simplified)",
		"Chinese (Traditional)",
		"Corsican",
		"Croatian",
		"Czech",
		"Danish",
		"Dutch",
		"English",
		"Esperanto",
		"Estonian",
		"Finnish",
		"French",
		"Frisian",
		"Galician",
		"Georgian",
		"German",
		"Greek",
		"Gujarati",
		"Haitian Creole",
		"Hausa",
		"Hawaiian",
		"Hebrew",
		"Hindi",
		"Hmong",
		"Hungarian",
		"Icelandic",
		"Igbo",
		"Indonesian",
		"Irish",
		"Italian",
		"Japanese",
		"Javanese",
		"Kannada",
		"Kazakh",
		"Khmer",
		"Kinyarwanda",
		"Korean",
		"Kurdish",
		"Kyrgyz",
		"Lao",
		"Latin",
		"Latvian",
		"Lithuanian",
		"Luxembourgish",
		"Macedonian",
		"Malagasy",
		"Malay",
		"Malayalam",
		"Maltese",
		"Maori",
		"Marathi",
		"Mongolian",
		"Myanmar (Burmese)",
		"Nepali",
		"Norwegian",
		"Nyanja (Chichewa)",
		"Odia (Oriya)",
		"Pashto",
		"Persian",
		"Polish",
		"Portuguese (Portugal, Brazil)",
		"Punjabi",
		"Romanian",
		"Russian",
		"Samoan",
		"Scots Gaelic",
		"Serbian",
		"Sesotho",
		"Shona",
		"Sindhi",
		"Sinhala (Sinhalese)",
		"Slovak",
		"Slovenian",
		"Somali",
		"Spanish",
		"Sundanese",
		"Swahili",
		"Swedish",
		"Tagalog (Filipino)",
		"Tajik",
		"Tamil",
		"Tatar",
		"Telugu",
		"Thai",
		"Turkish",
		"Turkmen",
		"Ukrainian",
		"Urdu",
		"Uyghur",
		"Uzbek",
		"Vietnamese",
		"Welsh",
		"Xhosa",
		"Yiddish",
		"Yoruba",
		"Zulu",
	}
	LangCode = map[string]string{
		"Afrikaans":                     "af",
		"Albanian":                      "sq",
		"Amharic":                       "am",
		"Arabic":                        "ar",
		"Armenian":                      "hy",
		"Azerbaijani":                   "az",
		"Basque":                        "eu",
		"Belarusian":                    "be",
		"Bengali":                       "bn",
		"Bosnian":                       "bs",
		"Bulgarian":                     "bg",
		"Catalan":                       "ca",
		"Cebuano":                       "ceb",
		"Chinese (Simplified)":          "zh-CN",
		"Chinese (Traditional)":         "zh-TW",
		"Corsican":                      "co",
		"Croatian":                      "hr",
		"Czech":                         "cs",
		"Danish":                        "da",
		"Dutch":                         "nl",
		"English":                       "en",
		"Esperanto":                     "eo",
		"Estonian":                      "et",
		"Finnish":                       "fi",
		"French":                        "fr",
		"Frisian":                       "fy",
		"Galician":                      "gl",
		"Georgian":                      "ka",
		"German":                        "de",
		"Greek":                         "el",
		"Gujarati":                      "gu",
		"Haitian Creole":                "ht",
		"Hausa":                         "ha",
		"Hawaiian":                      "haw",
		"Hebrew":                        "he",
		"Hindi":                         "hi",
		"Hmong":                         "hmn",
		"Hungarian":                     "hu",
		"Icelandic":                     "is",
		"Igbo":                          "ig",
		"Indonesian":                    "id",
		"Irish":                         "ga",
		"Italian":                       "it",
		"Japanese":                      "ja",
		"Javanese":                      "jv",
		"Kannada":                       "kn",
		"Kazakh":                        "kk",
		"Khmer":                         "km",
		"Kinyarwanda":                   "rw",
		"Korean":                        "ko",
		"Kurdish":                       "ku",
		"Kyrgyz":                        "ky",
		"Lao":                           "lo",
		"Latin":                         "la",
		"Latvian":                       "lv",
		"Lithuanian":                    "lt",
		"Luxembourgish":                 "lb",
		"Macedonian":                    "mk",
		"Malagasy":                      "mg",
		"Malay":                         "ms",
		"Malayalam":                     "ml",
		"Maltese":                       "mt",
		"Maori":                         "mi",
		"Marathi":                       "mr",
		"Mongolian":                     "mn",
		"Myanmar (Burmese)":             "my",
		"Nepali":                        "ne",
		"Norwegian":                     "no",
		"Nyanja (Chichewa)":             "ny",
		"Odia (Oriya)":                  "or",
		"Pashto":                        "ps",
		"Persian":                       "fa",
		"Polish":                        "pl",
		"Portuguese (Portugal, Brazil)": "pt",
		"Punjabi":                       "pa",
		"Romanian":                      "ro",
		"Russian":                       "ru",
		"Samoan":                        "sm",
		"Scots Gaelic":                  "gd",
		"Serbian":                       "sr",
		"Sesotho":                       "st",
		"Shona":                         "sn",
		"Sindhi":                        "sd",
		"Sinhala (Sinhalese)":           "si",
		"Slovak":                        "sk",
		"Slovenian":                     "sl",
		"Somali":                        "so",
		"Spanish":                       "es",
		"Sundanese":                     "su",
		"Swahili":                       "sw",
		"Swedish":                       "sv",
		"Tagalog (Filipino)":            "tl",
		"Tajik":                         "tg",
		"Tamil":                         "ta",
		"Tatar":                         "tt",
		"Telugu":                        "te",
		"Thai":                          "th",
		"Turkish":                       "tr",
		"Turkmen":                       "tk",
		"Ukrainian":                     "uk",
		"Urdu":                          "ur",
		"Uyghur":                        "ug",
		"Uzbek":                         "uz",
		"Vietnamese":                    "vi",
		"Welsh":                         "cy",
		"Xhosa":                         "xh",
		"Yiddish":                       "yi",
		"Yoruba":                        "yo",
		"Zulu":                          "zu",
	}
)

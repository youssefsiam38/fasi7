package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/viper"
)

func GetConfigString(key string) string {
	return os.ExpandEnv(viper.GetString(key))
}

func GetFilesRecursive(dir string) ([]string, error) {
	var files []string
	fasi7ignore, err := parseFasi7ignore(dir)
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		if lo.Contains(fasi7ignore, relPath) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	// exlcude .fasi7ignore file
	for i, file := range files {
		if strings.HasSuffix(file, ".fasi7ignore") {
			files = append(files[:i], files[i+1:]...)
			break
		}
	}

	return files, err
}

func parseFasi7ignore(dir string) ([]string, error) {
	fasi7ignorePath := filepath.Join(dir, ".fasi7ignore")
	data, err := os.ReadFile(fasi7ignorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var fasi7ignore []string
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fasi7ignore = append(fasi7ignore, line)
	}

	return fasi7ignore, scanner.Err()
}

func WriteFile(path string, data []byte) error {
	// Create all the parent directories
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Write the file
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

func IsoToLanguage(iso string) string {
	switch iso {
	case "af":
		return "Afrikaans"
	case "af-ZA":
		return "Afrikaans (South Africa)"
	case "ar":
		return "Arabic"
	case "ar-AE":
		return "Arabic (U.A.E.)"
	case "ar-BH":
		return "Arabic (Bahrain)"
	case "ar-DZ":
		return "Arabic (Algeria)"
	case "ar-EG":
		return "Arabic (Egypt)"
	case "ar-IQ":
		return "Arabic (Iraq)"
	case "ar-JO":
		return "Arabic (Jordan)"
	case "ar-KW":
		return "Arabic (Kuwait)"
	case "ar-LB":
		return "Arabic (Lebanon)"
	case "ar-LY":
		return "Arabic (Libya)"
	case "ar-MA":
		return "Arabic (Morocco)"
	case "ar-OM":
		return "Arabic (Oman)"
	case "ar-QA":
		return "Arabic (Qatar)"
	case "ar-SA":
		return "Arabic (Saudi Arabia)"
	case "ar-SY":
		return "Arabic (Syria)"
	case "ar-TN":
		return "Arabic (Tunisia)"
	case "ar-YE":
		return "Arabic (Yemen)"
	case "az":
		return "Azeri (Latin)"
	case "az-AZ":
		return "Azeri (Latin) (Azerbaijan)"
	case "be":
		return "Belarusian"
	case "be-BY":
		return "Belarusian (Belarus)"
	case "bg":
		return "Bulgarian"
	case "bg-BG":
		return "Bulgarian (Bulgaria)"
	case "bs-BA":
		return "Bosnian (Bosnia and Herzegovina)"
	case "ca":
		return "Catalan"
	case "ca-ES":
		return "Catalan (Spain)"
	case "cs":
		return "Czech"
	case "cs-CZ":
		return "Czech (Czech Republic)"
	case "cy":
		return "Welsh"
	case "cy-GB":
		return "Welsh (United Kingdom)"
	case "da":
		return "Danish"
	case "da-DK":
		return "Danish (Denmark)"
	case "de":
		return "German"
	case "de-AT":
		return "German (Austria)"
	case "de-CH":
		return "German (Switzerland)"
	case "de-DE":
		return "German (Germany)"
	case "de-LI":
		return "German (Liechtenstein)"
	case "de-LU":
		return "German (Luxembourg)"
	case "dv":
		return "Divehi"
	case "dv-MV":
		return "Divehi (Maldives)"
	case "el":
		return "Greek"
	case "el-GR":
		return "Greek (Greece)"
	case "en":
		return "English"
	case "en-AU":
		return "English (Australia)"
	case "en-BZ":
		return "English (Belize)"
	case "en-CA":
		return "English (Canada)"
	case "en-CB":
		return "English (Caribbean)"
	case "en-GB":
		return "English (United Kingdom)"
	case "en-IE":
		return "English (Ireland)"
	case "en-JM":
		return "English (Jamaica)"
	case "en-NZ":
		return "English (New Zealand)"
	case "en-PH":
		return "English (Republic of the Philippines)"
	case "en-TT":
		return "English (Trinidad and Tobago)"
	case "en-US":
		return "English (United States)"
	case "en-ZA":
		return "English (South Africa)"
	case "en-ZW":
		return "English (Zimbabwe)"
	case "eo":
		return "Esperanto"
	case "es":
		return "Spanish"
	case "es-AR":
		return "Spanish (Argentina)"
	case "es-BO":
		return "Spanish (Bolivia)"
	case "es-CL":
		return "Spanish (Chile)"
	case "es-CO":
		return "Spanish (Colombia)"
	case "es-CR":
		return "Spanish (Costa Rica)"
	case "es-DO":
		return "Spanish (Dominican Republic)"
	case "es-EC":
		return "Spanish (Ecuador)"
	case "es-ES":
		return "Spanish (Spain)"
	case "es-GT":
		return "Spanish (Guatemala)"
	case "es-HN":
		return "Spanish (Honduras)"
	case "es-MX":
		return "Spanish (Mexico)"
	case "es-NI":
		return "Spanish (Nicaragua)"
	case "es-PA":
		return "Spanish (Panama)"
	case "es-PE":
		return "Spanish (Peru)"
	case "es-PR":
		return "Spanish (Puerto Rico)"
	case "es-PY":
		return "Spanish (Paraguay)"
	case "es-SV":
		return "Spanish (El Salvador)"
	case "es-UY":
		return "Spanish (Uruguay)"
	case "es-VE":
		return "Spanish (Venezuela)"
	case "et":
		return "Estonian"
	case "et-EE":
		return "Estonian (Estonia)"
	case "eu":
		return "Basque"
	case "eu-ES":
		return "Basque (Spain)"
	case "fa":
		return "Farsi"
	case "fa-IR":
		return "Farsi (Iran)"
	case "fi":
		return "Finnish"
	case "fi-FI":
		return "Finnish (Finland)"
	case "fo":
		return "Faroese"
	case "fo-FO":
		return "Faroese (Faroe Islands)"
	case "fr":
		return "French"
	case "fr-BE":
		return "French (Belgium)"
	case "fr-CA":
		return "French (Canada)"
	case "fr-CH":
		return "French (Switzerland)"
	case "fr-FR":
		return "French (France)"
	case "fr-LU":
		return "French (Luxembourg)"
	case "fr-MC":
		return "French (Principality of Monaco)"
	case "gl":
		return "Galician"
	case "gl-ES":
		return "Galician (Spain)"
	case "gu":
		return "Gujarati"
	case "gu-IN":
		return "Gujarati (India)"
	case "he":
		return "Hebrew"
	case "he-IL":
		return "Hebrew"
	case "hi":
		return "Hindi"
	case "hi-IN":
		return "Hindi (India)"
	case "hr":
		return "Croatian"
	case "hr-BA":
		return "Croatian (Bosnia and Herzegovina)"
	case "hr-HR":
		return "Croatian (Croatia)"
	case "hu":
		return "Hungarian"
	case "hu-HU":
		return "Hungarian (Hungary)"
	case "hy":
		return "Armenian"
	case "hy-AM":
		return "Armenian (Armenia)"
	case "id":
		return "Indonesian"
	case "id-ID":
		return "Indonesian (Indonesia)"
	case "is":
		return "Icelandic"
	case "is-IS":
		return "Icelandic (Iceland)"
	case "it":
		return "Italian"
	case "it-CH":
		return "Italian (Switzerland)"
	case "it-IT":
		return "Italian (Italy)"
	case "ja":
		return "Japanese"
	case "ja-JP":
		return "Japanese (Japan)"
	case "ka":
		return "Georgian"
	case "ka-GE":
		return "Georgian (Georgia)"
	case "kk":
		return "Kazakh"
	case "kk-KZ":
		return "Kazakh (Kazakhstan)"
	case "kn":
		return "Kannada"
	case "kn-IN":
		return "Kannada (India)"
	case "ko":
		return "Korean"
	case "ko-KR":
		return "Korean (Korea)"
	case "kok":
		return "Konkani"
	case "kok-IN":
		return "Konkani (India)"
	case "ky":
		return "Kyrgyz"
	case "ky-KG":
		return "Kyrgyz (Kyrgyzstan)"
	case "lt":
		return "Lithuanian"
	case "lt-LT":
		return "Lithuanian (Lithuania)"
	case "lv":
		return "Latvian"
	case "lv-LV":
		return "Latvian (Latvia)"
	case "mi":
		return "Maori"
	case "mi-NZ":
		return "Maori (New Zealand)"
	case "mk":
		return "FYRO Macedonian"
	case "mk-MK":
		return "FYRO Macedonian (Former Yugoslav Republic of Macedonia)"
	case "mn":
		return "Mongolian"
	case "mn-MN":
		return "Mongolian (Mongolia)"
	case "mr":
		return "Marathi"
	case "mr-IN":
		return "Marathi (India)"
	case "ms":
		return "Malay"
	case "ms-BN":
		return "Malay (Brunei Darussalam)"
	case "ms-MY":
		return "Malay (Malaysia)"

	case "mt":
		return "Maltese"
	case "mt-MT":
		return "Maltese (Malta)"
	case "nb":
		return "Norwegian (Bokmål)"
	case "nb-NO":
		return "Norwegian (Bokmål) (Norway)"
	case "nl":
		return "Dutch"
	case "nl-BE":
		return "Dutch (Belgium)"
	case "nl-NL":
		return "Dutch (Netherlands)"
	case "nn-NO":
		return "Norwegian (Nynorsk) (Norway)"
	case "ns":
		return "Northern Sotho"
	case "ns-ZA":
		return "Northern Sotho (South Africa)"
	case "pa":
		return "Punjabi"
	case "pa-IN":
		return "Punjabi (India)"
	case "pl":
		return "Polish"
	case "pl-PL":
		return "Polish (Poland)"
	case "ps":
		return "Pashto"
	case "ps-AR":
		return "Pashto (Afghanistan)"
	case "pt":
		return "Portuguese"
	case "pt-BR":
		return "Portuguese (Brazil)"
	case "pt-PT":
		return "Portuguese (Portugal)"
	case "qu":
		return "Quechua"
	case "qu-BO":
		return "Quechua (Bolivia)"
	case "qu-EC":
		return "Quechua (Ecuador)"
	case "qu-PE":
		return "Quechua (Peru)"
	case "ro":
		return "Romanian"
	case "ro-RO":
		return "Romanian (Romania)"
	case "ru":
		return "Russian"
	case "ru-RU":
		return "Russian (Russia)"
	case "sa":
		return "Sanskrit"
	case "sa-IN":
		return "Sanskrit (India)"
	case "se":
		return "Sami (Northern)"
	case "se-FI":
		return "Sami (Northern) (Finland)"
	case "se-NO":
		return "Sami (Northern) (Norway)"
	case "se-SE":
		return "Sami (Northern) (Sweden)"
	case "sk":
		return "Slovak"
	case "sk-SK":
		return "Slovak (Slovakia)"
	case "sl":
		return "Slovenian"
	case "sl-SI":
		return "Slovenian (Slovenia)"
	case "sq":
		return "Albanian"
	case "sq-AL":
		return "Albanian (Albania)"
	case "sr-BA":
		return "Serbian (Latin) (Bosnia and Herzegovina)"
	case "sr-SP":
		return "Serbian (Latin) (Serbia and Montenegro)"
	case "sv":
		return "Swedish"
	case "sv-FI":
		return "Swedish (Finland)"
	case "sv-SE":
		return "Swedish (Sweden)"
	case "sw":
		return "Swahili"
	case "sw-KE":
		return "Swahili (Kenya)"
	case "syr":
		return "Syriac"
	case "syr-SY":
		return "Syriac (Syria)"
	case "ta":
		return "Tamil"
	case "ta-IN":
		return "Tamil (India)"
	case "te":
		return "Telugu"
	case "te-IN":
		return "Telugu (India)"
	case "th":
		return "Thai"
	case "th-TH":
		return "Thai (Thailand)"
	case "tl":
		return "Tagalog"
	case "tl-PH":
		return "Tagalog (Philippines)"
	case "tn":
		return "Tswana"
	case "tn-ZA":
		return "Tswana (South Africa)"
	case "tr":
		return "Turkish"
	case "tr-TR":
		return "Turkish (Turkey)"
	case "tt":
		return "Tatar"
	case "tt-RU":
		return "Tatar (Russia)"
	case "ts":
		return "Tsonga"
	case "uk":
		return "Ukrainian"
	case "uk-UA":
		return "Ukrainian (Ukraine)"
	case "ur":
		return "Urdu"
	case "ur-PK":
		return "Urdu (Islamic Republic of Pakistan)"
	case "uz":
		return "Uzbek (Latin)"
	case "uz-UZ":
		return "Uzbek (Latin) (Uzbekistan)"
	case "vi":
		return "Vietnamese"
	case "vi-VN":
		return "Vietnamese (Viet Nam)"
	case "xh":
		return "Xhosa"
	case "xh-ZA":
		return "Xhosa (South Africa)"
	case "zh":
		return "Chinese"
	case "zh-CN":
		return "Chinese (S)"
	case "zh-HK":
		return "Chinese (Hong Kong)"
	case "zh-MO":
		return "Chinese (Macau)"
	case "zh-SG":
		return "Chinese (Singapore)"
	case "zh-TW":
		return "Chinese (T)"
	case "zu":
		return "Zulu"
	case "zu-ZA":
		return "Zulu (South Africa)"
	default:
		return iso
	}
}

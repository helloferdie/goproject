package liblocale

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
	"spun/pkg/liblogger"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Regex for detect <T syntax>
var regexTemplate = regexp.MustCompile(`^<T (.+)>$`)

// LoadBundle
func LoadBundle() (*i18n.Bundle, error) {
	// Load bundle with default language
	defaultLocale := os.Getenv("locale_default")
	bundle := i18n.NewBundle(language.English)
	if defaultLocale != "" {
		tag, err := language.Parse(defaultLocale)
		if err != nil {
			bundle = i18n.NewBundle(tag)
		}
	}

	// Load localization
	dirLocale := os.Getenv("dir_locale")
	files, err := os.ReadDir(dirLocale)
	if err != nil {
		liblogger.Errorf("Error read localization directory %v", err)
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			// Check for valid language
			dirName := file.Name()
			_, err := language.Parse(dirName)
			if err != nil {
				liblogger.Errorf("Error localization directory %s is not a valid language tag", dirName)
				continue
			}

			dirLang := dirLocale + "/" + dirName
			languageFiles, err := os.ReadDir(dirLang)
			for _, languageFile := range languageFiles {
				if !languageFile.IsDir() {
					filename := languageFile.Name()
					extension := path.Ext(filename)
					if extension == ".json" {
						// Load localization json file
						languageFilename := dirLang + "/" + filename
						b, err := os.ReadFile(languageFilename)
						if err != nil {
							liblogger.Errorf("Error localization failed to read %s", languageFilename)
							continue
						}

						data := map[string]interface{}{}
						err = json.Unmarshal(b, &data)
						if err != nil {
							liblogger.Errorf("Error localization failed to parse %s", languageFilename)
							continue
						}

						// Flatten json
						flattenData := make(map[string]string)
						flattenJSON(strings.TrimSuffix(filename, extension), data, flattenData)
						flattenBytes, err := json.Marshal(flattenData)
						if err != nil {
							liblogger.Errorf("Error localization failed to flatten %s", languageFilename)
							continue
						}

						// Parse and set to bundle
						_, err = bundle.ParseMessageFileBytes(flattenBytes, dirName+extension)
						if err != nil {
							liblogger.Errorf("Error localization bundle failed to parse %s", languageFilename)
							continue
						}
					}
				}
			}
		}
	}
	return bundle, nil
}

// Translate
func Translate(loc *i18n.Localizer, syntax string, data interface{}) string {
	s, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID:    syntax,
		TemplateData: data,
	})
	if err != nil {
		s, err = loc.Localize(&i18n.LocalizeConfig{
			MessageID: "common.error.i18n.syntax",
			TemplateData: map[string]interface{}{
				"syntax": syntax,
			},
		})
		if err != nil {
			return syntax
		}
	}

	// Do translate SYTNTAX if return string is chained syntax <<SYNTAX>>
	isChain, syntaxChain := checkSyntax(s)
	if isChain {
		s = Translate(loc, syntaxChain, data)
	}
	return s
}

// flattenJSON convert map[string]interface{}{} to map[string]string, where the key delimiter with "."
func flattenJSON(prefix string, value interface{}, out map[string]string) {
	// Check the type of the value
	switch value := value.(type) {
	case map[string]interface{}:
		// If it's a map, iterate through its keys
		for k, v := range value {
			// Construct the new key
			newKey := k
			if prefix != "" {
				newKey = prefix + "." + k
			}
			// Recursively flatten the value
			flattenJSON(newKey, v, out)
		}
	case string:
		// If it's a string, add it to the output
		out[prefix] = value
	}
}

// checkSyntax verify syntax is chain syntax
func checkSyntax(input string) (bool, string) {
	// Check if the input matches the pattern
	matches := regexTemplate.FindStringSubmatch(input)

	// If it's a match, return true and the extracted key
	if len(matches) > 1 {
		return true, matches[1]
	}

	// If it's not a match, return false and an empty string
	return false, ""
}

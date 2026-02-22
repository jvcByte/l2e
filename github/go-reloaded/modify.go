package main

import (
	"regexp"
	"strconv"
	"strings"
)

func applyModifications(text string) string {
	text = normalizeModifiers(text)
	text = applyModifiers(text)
	text = applyPunctuation(text)
	text = applyQuotes(text)
	text = applyArticle(text)
	return text
}

func applyModifiers(text string) string {
	words := strings.Fields(text)
	// matches (up, 2), (low, 3), (cap, 1) etc.
	reWithNum := regexp.MustCompile(`^\((up|low|cap),\s*(\d+)\)$`)

	for i, word := range words {
		// handle (up, n), (low, n), (cap, n)
		if matches := reWithNum.FindStringSubmatch(word); matches != nil {
			op := matches[1]
			n, _ := strconv.Atoi(matches[2])
			words[i] = ""
			// go back n words and apply
			count := 0
			for j := i - 1; j >= 0 && count < n; j-- {
				if words[j] != "" {
					switch op {
					case "up":
						words[j] = strings.ToUpper(words[j])
					case "low":
						words[j] = strings.ToLower(words[j])
					case "cap":
						words[j] = capitalize(words[j])
					}
					count++
				}
			}
			continue
		}

		// handle (up), (low), (cap), (hex), (bin)
		switch word {
			case "(up)":
				if i > 0 {
					words[i-1] = strings.ToUpper(words[i-1])
					words[i] = ""
				}
			case "(low)":
				if i > 0 {
					words[i-1] = strings.ToLower(words[i-1])
					words[i] = ""
				}
			case "(cap)":
				if i > 0 {
					words[i-1] = capitalize(words[i-1])
					words[i] = ""
				}

			case "(hex)":
				if i > 0 {
					n, err := strconv.ParseInt(words[i-1], 16, 64)
					if err == nil {
						words[i-1] = strconv.FormatInt(n, 10)
					}
					words[i] = ""
				}
			case "(bin)":
				if i > 0 {
					n, err := strconv.ParseInt(words[i-1], 2, 64)
					if err == nil {
						words[i-1] = strconv.FormatInt(n, 10)
					}
					words[i] = ""
				}
			default:
				words[i] = word
		}
	}

	var results []string
	for _, word := range words {
		if word != "" {
			results = append(results, word)
		}
	}
	return strings.Join(results, " ")
}

func applyPunctuation(text string) string {
	// group punctuation together (remove spaces between them)
	re := regexp.MustCompile(`\s*([.,!?:;]+)\s*`)
	text = re.ReplaceAllString(text, "$1 ")

	// trim trailing space
	text = strings.TrimSpace(text)
	return text
}

func applyQuotes(text string) string {
	re := regexp.MustCompile(`'\s+(.+?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}

func applyArticle(text string) string {
	words := strings.Fields(text)
	vowelsAndH := "aeiouAEIOUhH"

	for i, word := range words {
		if (word == "a" || word == "A") && i+1 < len(words) {
			nextWord := words[i+1]
			if strings.ContainsRune(vowelsAndH, rune(nextWord[0])) {
				if word == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}
	}

	return strings.Join(words, " ")
}

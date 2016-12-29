package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		helpMessage()
		os.Exit(1)
	} else if os.Args[1] == "--help" || os.Args[1] == "-h" {
		helpPage()
	} else {
		options(os.Args[1], os.Args[2])
	}
}

func helpPage() {
	helpMessage()
	fmt.Println("Options:")
	fmt.Println("-r, --reverse\tDisplay horizontally reflected characters")
	fmt.Println("-v, --vertical\tDisplay vertically reflected characters")
	fmt.Println("-u, --rotate180\tDisplay characters rotated 180°")
	fmt.Println("    --funny\tDisplay funny characters")
	fmt.Println("    --1337\tDisplay 1337 characters")
}

func helpMessage() {
	fmt.Printf("usage: %s [options...] <string>\n", filepath.Base(os.Args[0]))
}

func options(option string, text string) {
	switch option {
	case "-r":
		handleString(text, reverse)
	case "--reverse":
		handleString(text, reverse)
	case "-v":
		handleString(text, verticalReflection)
	case "--vertical":
		handleString(text, verticalReflection)
	case "-u":
		handleString(text, rotate180)
	case "--rotate180":
		handleString(text, rotate180)
	case "--1337":
		handleString(text, _1337)
	case "--funny":
		handleString(text, funny)
	default:
		fmt.Printf("There is no %s option (\"%s --help\" for help)\n", os.Args[1], filepath.Base(os.Args[0]))
		log.Fatal("invalid option")
		os.Exit(1)
	}
}

func handleString(input string, hash map[string]string) {
	var inputWithDelimiter string
	if hash["th"] != "" {
		inputWithDelimiter = strings.Replace(strings.Replace(input, "", ",", -1), "t,h", "th", -1)
	} else {
		inputWithDelimiter = strings.Replace(input, "", ",", -1)
	}
	inputEscape := handleEscapes(inputWithDelimiter)
	input_array := strings.Split(inputEscape, ",")
	setEscapeValues(hash)
	for i := range input_array {
		generateCharacter(input_array[i], hash)
	}
}

func handleEscapes(inputString string) string {
	replacer := strings.NewReplacer("\\,n,", "\\n,",
		" ,", "\\s,",
		"\\,r", "\\r",
		",,,", ",u002c,",
		",!,", ",\\!\\,")
	newlines := replacer.Replace(inputString)
	return newlines

}

func setEscapeValues(hash map[string]string) {
	for k, v := range escapeCharacters {
		if hash[k] == "" {
			hash[k] = v
		}
	}
}

func generateCharacter(character string, hash map[string]string) {
	fmt.Printf(hash[character])
}

var escapeCharacters = map[string]string{
	"\\s":   " ",
	"\\!\\": "!",
	"\\n":   "\n",
	"\\r":   "\r",
	"u002c": ",",
	"\"":    "\"",
	"'":     "'",
	"*":     "*",
	"-":     "-",
	".":     ".",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

var reverse = map[string]string{
	"A": "A",
	"B": "ᙠ",
	"C": "Ɔ",
	"D": "ᗡ",
	"E": "Ǝ",
	"F": "ꟻ",
	"G": "Ꭾ",
	"H": "H",
	"I": "I",
	"J": "Ⴑ",
	"K": "ᐴ",
	"L": "⅃",
	"M": "M",
	"N": "И",
	"O": "O",
	"P": "ꟼ",
	"Q": "Ọ",
	"R": "Я",
	"S": "Ƨ",
	"T": "T",
	"U": "U",
	"V": "V",
	"W": "W",
	"X": "X",
	"Y": "Y",
	"Z": "Ƹ",
	"a": "ɒ",
	"b": "d",
	"c": "ɔ",
	"d": "b",
	"e": "ɘ",
	"f": "ʇ",
	"g": "ǫ",
	"h": "ʜ",
	"i": "i",
	"j": "Ⴑ",
	"k": "ʞ",
	"l": "l",
	"m": "m",
	"n": "n",
	"o": "o",
	"q": "p",
	"p": "q",
	"r": "ɿ",
	"s": "ƨ",
	"t": "ƚ",
	"u": "u",
	"v": "v",
	"w": "w",
	"x": "x",
	"y": "y",
	"z": "ƹ",
	"?": "⸮",
	"!": "!",
	"(": ")",
	")": "(",
	"&": "⅋",
}

var verticalReflection = map[string]string{
	"A": "ᗄ",
	"B": "ᗷ",
	"C": "C",
	"D": "D",
	"E": "E",
	"F": "ᖶ",
	"G": "⅁",
	"H": "H",
	"I": "I",
	"J": "ᘃ",
	"K": "K",
	"L": "Ѓ",
	"M": "W",
	"N": "И",
	"O": "O",
	"P": "b",
	"Q": "ⵚ",
	"R": "ᖉ",
	"S": "Ƨ",
	"T": "⊥",
	"U": "∩",
	"V": "⋀",
	"W": "M",
	"X": "X",
	"Y": "⅄",
	"Z": "Z",
	"a": "ɐ",
	"b": "p",
	"c": "c",
	"d": "q",
	"e": "ǝ",
	"f": "ʈ",
	"g": "ɓ",
	"h": "µ",
	"i": "!",
	"j": "ɾ",
	"k": "ʞ",
	"l": "ꞁ",
	"m": "ɯ",
	"n": "u",
	"o": "o",
	"p": "b",
	"q": "d",
	"r": "ɹ",
	"s": "ƨ",
	"t": "f",
	"u": "∩",
	"v": "٨",
	"w": "ʍ",
	"x": "x",
	"y": "ʎ",
	"z": "s",
	"'": "،",
	".": "˙",
	"?": "j",
	"!": "¡",
	"(": "(",
	")": ")",
	"&": "⅋",
}

var rotate180 = map[string]string{
	"A": "∀",
	"B": "𐐒",
	"C": "Ɔ",
	"D": "◖",
	"E": "Ǝ",
	"F": "Ⅎ",
	"G": "⅁",
	"H": "H",
	"I": "I",
	"J": "ſ",
	"K": "⋊",
	"L": "˥",
	"M": "W",
	"N": "N",
	"O": "O",
	"P": "Ԁ",
	"Q": "Ό",
	"R": "ᴚ",
	"S": "S",
	"T": "⊥",
	"U": "∩",
	"V": "Λ",
	"W": "M",
	"X": "X",
	"Y": "⅄",
	"Z": "Z",
	"a": "ɐ",
	"b": "q",
	"c": "ɔ",
	"d": "p",
	"e": "ǝ",
	"f": "ɟ",
	"g": "ƃ",
	"h": "ɥ",
	"i": "ı",
	"j": "ɾ",
	"k": "ʞ",
	"l": "ן",
	"m": "ɯ",
	"n": "u",
	"o": "o",
	"p": "d",
	"q": "b",
	"r": "ɹ",
	"s": "s",
	"t": "ʇ",
	"u": "n",
	"v": "ʌ",
	"w": "ʍ",
	"x": "x",
	"y": "ʎ",
	"z": "z",
	"'": "،",
	".": "˙",
	"?": "¿",
	"!": "¡",
	"(": "(",
	")": "(",
	"&": "⅋",
}

var _1337 = map[string]string{
	"A": "4",
	"B": "8",
	"C": "[",
	"D": "|>",
	"E": "3",
	"F": "|=",
	"G": "6",
	"H": "#",
	"I": "][",
	"J": "_|",
	"K": "|<",
	"L": "1",
	"M": "(\\/)",
	"N": "|\\|",
	"O": "0",
	"P": "|D",
	"Q": "(,)",
	"R": "|2",
	"S": "5",
	"T": "7",
	"U": "|_|",
	"V": "\\/",
	"W": "\\/\\/",
	"X": "><",
	"Y": "`/",
	"Z": "2",
	"a": "4",
	"b": "8",
	"c": "[",
	"d": "|>",
	"e": "3",
	"f": "|=",
	"g": "6",
	"h": "#",
	"i": "][",
	"j": "_|",
	"k": "|<",
	"l": "1",
	"m": "(\\/)",
	"n": "|\\|",
	"o": "0",
	"p": "|D",
	"q": "(,)",
	"r": "|2",
	"s": "5",
	"t": "7",
	"u": "|_|",
	"v": "\\/",
	"w": "\\/\\/",
	"x": "><",
	"y": "`/",
	"z": "2",
	"?": "@",
	"!": "@@",
	"(": "[[",
	")": "]]",
	"&": "&&",
}

var funny = map[string]string{
	"A":  "Å",
	"B":  "Ƀ",
	"C":  "Ç",
	"D":  "Δ",
	"E":  "З",
	"F":  "ߓ",
	"G":  "Γ",
	"H":  "Ħ",
	"I":  "Ϊ",
	"J":  "ʤ",
	"K":  "Қ",
	"L":  "£",
	"M":  "ᛖ",
	"N":  "Ñ",
	"O":  "Ø",
	"P":  "Π",
	"Q":  "Ϙ",
	"R":  "Я",
	"S":  "§",
	"T":  "Ꮦ",
	"U":  "ᦢ",
	"V":  "Ѵ",
	"W":  "Ϣ",
	"X":  "Ӂ",
	"Y":  "Ӳ",
	"Z":  "Ǯ",
	"Th": "Ð",
	"a":  "å",
	"b":  "ƀ",
	"c":  "ç",
	"d":  "δ",
	"e":  "з",
	"f":  "ϝ",
	"g":  "γ",
	"h":  "ħ",
	"i":  "ϊ",
	"j":  "ʝ",
	"k":  "қ",
	"l":  "£",
	"m":  "ϻ",
	"n":  "ñ",
	"o":  "ø",
	"p":  "π",
	"q":  "ϙ",
	"r":  "я",
	"s":  "§",
	"t":  "Ꮏ",
	"u":  "ᦢ",
	"v":  "ѵ",
	"w":  "ϣ",
	"x":  "ӂ",
	"y":  "ӳ",
	"z":  "ǯ",
	"th": "ð",
	"?":  "¿",
	"!":  "!¡",
	"(":  "<",
	")":  ">",
	"&":  "⅋",
}

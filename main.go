package main

import (
	"fmt"
)

func main() {
	for k, v := range reverse {
		fmt.Printf(k)
		fmt.Printf(v)
	}
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
}

//ABCDEFGHIJKLMNOPQRSTUVWXYZ
//abcdefghijklmnopqrstuvwxyz

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
	"r": "ʁ",
	"s": "ƨ",
	"t": "f",
	"u": "∩",
	"v": "٨",
	"w": "ʍ",
	"x": "x",
	"y": "ʎ",
	"z": "s",
}

var rotate180 = map[string]string{
//ɐqɔpǝɟƃɥıɾʞןɯuodbɹsʇnʌʍxʎz
//Z⅄XMΛ∩⊥SᴚΌԀONW˥⋊ſIH⅁ℲƎ◖Ɔ𐐒∀
//ɐqɔpǝɟƃɥıɾʞןɯuodbɹsʇnʌʍxʎz
//Z⅄XMΛ∩⊥SᴚΌԀONW˥⋊ſIH⅁ℲƎ◖Ɔ𐐒∀
}

var _1337 = map[string]string{
//1337
//"4","8", "[","|>", "3", "|=", "6", "#", "][", "_|", "|<", "1", "(\/)", "|\|", "0", "|D", "(,)", "|2", "5", "7", "|_|", "\/", "\/\/", "><", "`/", "2"
}

var funny = map[string]string{
//Foreign
//ÅɃÇΔЗߓΓĦΪʤҚ£ᛖÑØΠϘЯ§ᏖᦢѴϢӁӲǮ
//åƀçδзϝγħϊʝқ£ϻñøπϙя§Ꮏᦢѵϣӂӳǯ
//th = ðÐ
}

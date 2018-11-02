// this file was generated by gomacro command: import _b "unicode"
// DO NOT EDIT! Any change will be lost when the file is re-generated

// +build go1.11

package go1_11

import (
	. "reflect"
	unicode "unicode"
)

// reflection: allow interpreted code to import "unicode"
func init() {
	Packages["unicode"] = Package{
	Binds: map[string]Value{
		"ASCII_Hex_Digit":	ValueOf(&unicode.ASCII_Hex_Digit).Elem(),
		"Adlam":	ValueOf(&unicode.Adlam).Elem(),
		"Ahom":	ValueOf(&unicode.Ahom).Elem(),
		"Anatolian_Hieroglyphs":	ValueOf(&unicode.Anatolian_Hieroglyphs).Elem(),
		"Arabic":	ValueOf(&unicode.Arabic).Elem(),
		"Armenian":	ValueOf(&unicode.Armenian).Elem(),
		"Avestan":	ValueOf(&unicode.Avestan).Elem(),
		"AzeriCase":	ValueOf(&unicode.AzeriCase).Elem(),
		"Balinese":	ValueOf(&unicode.Balinese).Elem(),
		"Bamum":	ValueOf(&unicode.Bamum).Elem(),
		"Bassa_Vah":	ValueOf(&unicode.Bassa_Vah).Elem(),
		"Batak":	ValueOf(&unicode.Batak).Elem(),
		"Bengali":	ValueOf(&unicode.Bengali).Elem(),
		"Bhaiksuki":	ValueOf(&unicode.Bhaiksuki).Elem(),
		"Bidi_Control":	ValueOf(&unicode.Bidi_Control).Elem(),
		"Bopomofo":	ValueOf(&unicode.Bopomofo).Elem(),
		"Brahmi":	ValueOf(&unicode.Brahmi).Elem(),
		"Braille":	ValueOf(&unicode.Braille).Elem(),
		"Buginese":	ValueOf(&unicode.Buginese).Elem(),
		"Buhid":	ValueOf(&unicode.Buhid).Elem(),
		"C":	ValueOf(&unicode.C).Elem(),
		"Canadian_Aboriginal":	ValueOf(&unicode.Canadian_Aboriginal).Elem(),
		"Carian":	ValueOf(&unicode.Carian).Elem(),
		"CaseRanges":	ValueOf(&unicode.CaseRanges).Elem(),
		"Categories":	ValueOf(&unicode.Categories).Elem(),
		"Caucasian_Albanian":	ValueOf(&unicode.Caucasian_Albanian).Elem(),
		"Cc":	ValueOf(&unicode.Cc).Elem(),
		"Cf":	ValueOf(&unicode.Cf).Elem(),
		"Chakma":	ValueOf(&unicode.Chakma).Elem(),
		"Cham":	ValueOf(&unicode.Cham).Elem(),
		"Cherokee":	ValueOf(&unicode.Cherokee).Elem(),
		"Co":	ValueOf(&unicode.Co).Elem(),
		"Common":	ValueOf(&unicode.Common).Elem(),
		"Coptic":	ValueOf(&unicode.Coptic).Elem(),
		"Cs":	ValueOf(&unicode.Cs).Elem(),
		"Cuneiform":	ValueOf(&unicode.Cuneiform).Elem(),
		"Cypriot":	ValueOf(&unicode.Cypriot).Elem(),
		"Cyrillic":	ValueOf(&unicode.Cyrillic).Elem(),
		"Dash":	ValueOf(&unicode.Dash).Elem(),
		"Deprecated":	ValueOf(&unicode.Deprecated).Elem(),
		"Deseret":	ValueOf(&unicode.Deseret).Elem(),
		"Devanagari":	ValueOf(&unicode.Devanagari).Elem(),
		"Diacritic":	ValueOf(&unicode.Diacritic).Elem(),
		"Digit":	ValueOf(&unicode.Digit).Elem(),
		"Duployan":	ValueOf(&unicode.Duployan).Elem(),
		"Egyptian_Hieroglyphs":	ValueOf(&unicode.Egyptian_Hieroglyphs).Elem(),
		"Elbasan":	ValueOf(&unicode.Elbasan).Elem(),
		"Ethiopic":	ValueOf(&unicode.Ethiopic).Elem(),
		"Extender":	ValueOf(&unicode.Extender).Elem(),
		"FoldCategory":	ValueOf(&unicode.FoldCategory).Elem(),
		"FoldScript":	ValueOf(&unicode.FoldScript).Elem(),
		"Georgian":	ValueOf(&unicode.Georgian).Elem(),
		"Glagolitic":	ValueOf(&unicode.Glagolitic).Elem(),
		"Gothic":	ValueOf(&unicode.Gothic).Elem(),
		"Grantha":	ValueOf(&unicode.Grantha).Elem(),
		"GraphicRanges":	ValueOf(&unicode.GraphicRanges).Elem(),
		"Greek":	ValueOf(&unicode.Greek).Elem(),
		"Gujarati":	ValueOf(&unicode.Gujarati).Elem(),
		"Gurmukhi":	ValueOf(&unicode.Gurmukhi).Elem(),
		"Han":	ValueOf(&unicode.Han).Elem(),
		"Hangul":	ValueOf(&unicode.Hangul).Elem(),
		"Hanunoo":	ValueOf(&unicode.Hanunoo).Elem(),
		"Hatran":	ValueOf(&unicode.Hatran).Elem(),
		"Hebrew":	ValueOf(&unicode.Hebrew).Elem(),
		"Hex_Digit":	ValueOf(&unicode.Hex_Digit).Elem(),
		"Hiragana":	ValueOf(&unicode.Hiragana).Elem(),
		"Hyphen":	ValueOf(&unicode.Hyphen).Elem(),
		"IDS_Binary_Operator":	ValueOf(&unicode.IDS_Binary_Operator).Elem(),
		"IDS_Trinary_Operator":	ValueOf(&unicode.IDS_Trinary_Operator).Elem(),
		"Ideographic":	ValueOf(&unicode.Ideographic).Elem(),
		"Imperial_Aramaic":	ValueOf(&unicode.Imperial_Aramaic).Elem(),
		"In":	ValueOf(unicode.In),
		"Inherited":	ValueOf(&unicode.Inherited).Elem(),
		"Inscriptional_Pahlavi":	ValueOf(&unicode.Inscriptional_Pahlavi).Elem(),
		"Inscriptional_Parthian":	ValueOf(&unicode.Inscriptional_Parthian).Elem(),
		"Is":	ValueOf(unicode.Is),
		"IsControl":	ValueOf(unicode.IsControl),
		"IsDigit":	ValueOf(unicode.IsDigit),
		"IsGraphic":	ValueOf(unicode.IsGraphic),
		"IsLetter":	ValueOf(unicode.IsLetter),
		"IsLower":	ValueOf(unicode.IsLower),
		"IsMark":	ValueOf(unicode.IsMark),
		"IsNumber":	ValueOf(unicode.IsNumber),
		"IsOneOf":	ValueOf(unicode.IsOneOf),
		"IsPrint":	ValueOf(unicode.IsPrint),
		"IsPunct":	ValueOf(unicode.IsPunct),
		"IsSpace":	ValueOf(unicode.IsSpace),
		"IsSymbol":	ValueOf(unicode.IsSymbol),
		"IsTitle":	ValueOf(unicode.IsTitle),
		"IsUpper":	ValueOf(unicode.IsUpper),
		"Javanese":	ValueOf(&unicode.Javanese).Elem(),
		"Join_Control":	ValueOf(&unicode.Join_Control).Elem(),
		"Kaithi":	ValueOf(&unicode.Kaithi).Elem(),
		"Kannada":	ValueOf(&unicode.Kannada).Elem(),
		"Katakana":	ValueOf(&unicode.Katakana).Elem(),
		"Kayah_Li":	ValueOf(&unicode.Kayah_Li).Elem(),
		"Kharoshthi":	ValueOf(&unicode.Kharoshthi).Elem(),
		"Khmer":	ValueOf(&unicode.Khmer).Elem(),
		"Khojki":	ValueOf(&unicode.Khojki).Elem(),
		"Khudawadi":	ValueOf(&unicode.Khudawadi).Elem(),
		"L":	ValueOf(&unicode.L).Elem(),
		"Lao":	ValueOf(&unicode.Lao).Elem(),
		"Latin":	ValueOf(&unicode.Latin).Elem(),
		"Lepcha":	ValueOf(&unicode.Lepcha).Elem(),
		"Letter":	ValueOf(&unicode.Letter).Elem(),
		"Limbu":	ValueOf(&unicode.Limbu).Elem(),
		"Linear_A":	ValueOf(&unicode.Linear_A).Elem(),
		"Linear_B":	ValueOf(&unicode.Linear_B).Elem(),
		"Lisu":	ValueOf(&unicode.Lisu).Elem(),
		"Ll":	ValueOf(&unicode.Ll).Elem(),
		"Lm":	ValueOf(&unicode.Lm).Elem(),
		"Lo":	ValueOf(&unicode.Lo).Elem(),
		"Logical_Order_Exception":	ValueOf(&unicode.Logical_Order_Exception).Elem(),
		"Lower":	ValueOf(&unicode.Lower).Elem(),
		"LowerCase":	ValueOf(unicode.LowerCase),
		"Lt":	ValueOf(&unicode.Lt).Elem(),
		"Lu":	ValueOf(&unicode.Lu).Elem(),
		"Lycian":	ValueOf(&unicode.Lycian).Elem(),
		"Lydian":	ValueOf(&unicode.Lydian).Elem(),
		"M":	ValueOf(&unicode.M).Elem(),
		"Mahajani":	ValueOf(&unicode.Mahajani).Elem(),
		"Malayalam":	ValueOf(&unicode.Malayalam).Elem(),
		"Mandaic":	ValueOf(&unicode.Mandaic).Elem(),
		"Manichaean":	ValueOf(&unicode.Manichaean).Elem(),
		"Marchen":	ValueOf(&unicode.Marchen).Elem(),
		"Mark":	ValueOf(&unicode.Mark).Elem(),
		"Masaram_Gondi":	ValueOf(&unicode.Masaram_Gondi).Elem(),
		"MaxASCII":	ValueOf(unicode.MaxASCII),
		"MaxCase":	ValueOf(unicode.MaxCase),
		"MaxLatin1":	ValueOf(unicode.MaxLatin1),
		"MaxRune":	ValueOf(unicode.MaxRune),
		"Mc":	ValueOf(&unicode.Mc).Elem(),
		"Me":	ValueOf(&unicode.Me).Elem(),
		"Meetei_Mayek":	ValueOf(&unicode.Meetei_Mayek).Elem(),
		"Mende_Kikakui":	ValueOf(&unicode.Mende_Kikakui).Elem(),
		"Meroitic_Cursive":	ValueOf(&unicode.Meroitic_Cursive).Elem(),
		"Meroitic_Hieroglyphs":	ValueOf(&unicode.Meroitic_Hieroglyphs).Elem(),
		"Miao":	ValueOf(&unicode.Miao).Elem(),
		"Mn":	ValueOf(&unicode.Mn).Elem(),
		"Modi":	ValueOf(&unicode.Modi).Elem(),
		"Mongolian":	ValueOf(&unicode.Mongolian).Elem(),
		"Mro":	ValueOf(&unicode.Mro).Elem(),
		"Multani":	ValueOf(&unicode.Multani).Elem(),
		"Myanmar":	ValueOf(&unicode.Myanmar).Elem(),
		"N":	ValueOf(&unicode.N).Elem(),
		"Nabataean":	ValueOf(&unicode.Nabataean).Elem(),
		"Nd":	ValueOf(&unicode.Nd).Elem(),
		"New_Tai_Lue":	ValueOf(&unicode.New_Tai_Lue).Elem(),
		"Newa":	ValueOf(&unicode.Newa).Elem(),
		"Nko":	ValueOf(&unicode.Nko).Elem(),
		"Nl":	ValueOf(&unicode.Nl).Elem(),
		"No":	ValueOf(&unicode.No).Elem(),
		"Noncharacter_Code_Point":	ValueOf(&unicode.Noncharacter_Code_Point).Elem(),
		"Number":	ValueOf(&unicode.Number).Elem(),
		"Nushu":	ValueOf(&unicode.Nushu).Elem(),
		"Ogham":	ValueOf(&unicode.Ogham).Elem(),
		"Ol_Chiki":	ValueOf(&unicode.Ol_Chiki).Elem(),
		"Old_Hungarian":	ValueOf(&unicode.Old_Hungarian).Elem(),
		"Old_Italic":	ValueOf(&unicode.Old_Italic).Elem(),
		"Old_North_Arabian":	ValueOf(&unicode.Old_North_Arabian).Elem(),
		"Old_Permic":	ValueOf(&unicode.Old_Permic).Elem(),
		"Old_Persian":	ValueOf(&unicode.Old_Persian).Elem(),
		"Old_South_Arabian":	ValueOf(&unicode.Old_South_Arabian).Elem(),
		"Old_Turkic":	ValueOf(&unicode.Old_Turkic).Elem(),
		"Oriya":	ValueOf(&unicode.Oriya).Elem(),
		"Osage":	ValueOf(&unicode.Osage).Elem(),
		"Osmanya":	ValueOf(&unicode.Osmanya).Elem(),
		"Other":	ValueOf(&unicode.Other).Elem(),
		"Other_Alphabetic":	ValueOf(&unicode.Other_Alphabetic).Elem(),
		"Other_Default_Ignorable_Code_Point":	ValueOf(&unicode.Other_Default_Ignorable_Code_Point).Elem(),
		"Other_Grapheme_Extend":	ValueOf(&unicode.Other_Grapheme_Extend).Elem(),
		"Other_ID_Continue":	ValueOf(&unicode.Other_ID_Continue).Elem(),
		"Other_ID_Start":	ValueOf(&unicode.Other_ID_Start).Elem(),
		"Other_Lowercase":	ValueOf(&unicode.Other_Lowercase).Elem(),
		"Other_Math":	ValueOf(&unicode.Other_Math).Elem(),
		"Other_Uppercase":	ValueOf(&unicode.Other_Uppercase).Elem(),
		"P":	ValueOf(&unicode.P).Elem(),
		"Pahawh_Hmong":	ValueOf(&unicode.Pahawh_Hmong).Elem(),
		"Palmyrene":	ValueOf(&unicode.Palmyrene).Elem(),
		"Pattern_Syntax":	ValueOf(&unicode.Pattern_Syntax).Elem(),
		"Pattern_White_Space":	ValueOf(&unicode.Pattern_White_Space).Elem(),
		"Pau_Cin_Hau":	ValueOf(&unicode.Pau_Cin_Hau).Elem(),
		"Pc":	ValueOf(&unicode.Pc).Elem(),
		"Pd":	ValueOf(&unicode.Pd).Elem(),
		"Pe":	ValueOf(&unicode.Pe).Elem(),
		"Pf":	ValueOf(&unicode.Pf).Elem(),
		"Phags_Pa":	ValueOf(&unicode.Phags_Pa).Elem(),
		"Phoenician":	ValueOf(&unicode.Phoenician).Elem(),
		"Pi":	ValueOf(&unicode.Pi).Elem(),
		"Po":	ValueOf(&unicode.Po).Elem(),
		"Prepended_Concatenation_Mark":	ValueOf(&unicode.Prepended_Concatenation_Mark).Elem(),
		"PrintRanges":	ValueOf(&unicode.PrintRanges).Elem(),
		"Properties":	ValueOf(&unicode.Properties).Elem(),
		"Ps":	ValueOf(&unicode.Ps).Elem(),
		"Psalter_Pahlavi":	ValueOf(&unicode.Psalter_Pahlavi).Elem(),
		"Punct":	ValueOf(&unicode.Punct).Elem(),
		"Quotation_Mark":	ValueOf(&unicode.Quotation_Mark).Elem(),
		"Radical":	ValueOf(&unicode.Radical).Elem(),
		"Regional_Indicator":	ValueOf(&unicode.Regional_Indicator).Elem(),
		"Rejang":	ValueOf(&unicode.Rejang).Elem(),
		"ReplacementChar":	ValueOf(unicode.ReplacementChar),
		"Runic":	ValueOf(&unicode.Runic).Elem(),
		"S":	ValueOf(&unicode.S).Elem(),
		"STerm":	ValueOf(&unicode.STerm).Elem(),
		"Samaritan":	ValueOf(&unicode.Samaritan).Elem(),
		"Saurashtra":	ValueOf(&unicode.Saurashtra).Elem(),
		"Sc":	ValueOf(&unicode.Sc).Elem(),
		"Scripts":	ValueOf(&unicode.Scripts).Elem(),
		"Sentence_Terminal":	ValueOf(&unicode.Sentence_Terminal).Elem(),
		"Sharada":	ValueOf(&unicode.Sharada).Elem(),
		"Shavian":	ValueOf(&unicode.Shavian).Elem(),
		"Siddham":	ValueOf(&unicode.Siddham).Elem(),
		"SignWriting":	ValueOf(&unicode.SignWriting).Elem(),
		"SimpleFold":	ValueOf(unicode.SimpleFold),
		"Sinhala":	ValueOf(&unicode.Sinhala).Elem(),
		"Sk":	ValueOf(&unicode.Sk).Elem(),
		"Sm":	ValueOf(&unicode.Sm).Elem(),
		"So":	ValueOf(&unicode.So).Elem(),
		"Soft_Dotted":	ValueOf(&unicode.Soft_Dotted).Elem(),
		"Sora_Sompeng":	ValueOf(&unicode.Sora_Sompeng).Elem(),
		"Soyombo":	ValueOf(&unicode.Soyombo).Elem(),
		"Space":	ValueOf(&unicode.Space).Elem(),
		"Sundanese":	ValueOf(&unicode.Sundanese).Elem(),
		"Syloti_Nagri":	ValueOf(&unicode.Syloti_Nagri).Elem(),
		"Symbol":	ValueOf(&unicode.Symbol).Elem(),
		"Syriac":	ValueOf(&unicode.Syriac).Elem(),
		"Tagalog":	ValueOf(&unicode.Tagalog).Elem(),
		"Tagbanwa":	ValueOf(&unicode.Tagbanwa).Elem(),
		"Tai_Le":	ValueOf(&unicode.Tai_Le).Elem(),
		"Tai_Tham":	ValueOf(&unicode.Tai_Tham).Elem(),
		"Tai_Viet":	ValueOf(&unicode.Tai_Viet).Elem(),
		"Takri":	ValueOf(&unicode.Takri).Elem(),
		"Tamil":	ValueOf(&unicode.Tamil).Elem(),
		"Tangut":	ValueOf(&unicode.Tangut).Elem(),
		"Telugu":	ValueOf(&unicode.Telugu).Elem(),
		"Terminal_Punctuation":	ValueOf(&unicode.Terminal_Punctuation).Elem(),
		"Thaana":	ValueOf(&unicode.Thaana).Elem(),
		"Thai":	ValueOf(&unicode.Thai).Elem(),
		"Tibetan":	ValueOf(&unicode.Tibetan).Elem(),
		"Tifinagh":	ValueOf(&unicode.Tifinagh).Elem(),
		"Tirhuta":	ValueOf(&unicode.Tirhuta).Elem(),
		"Title":	ValueOf(&unicode.Title).Elem(),
		"TitleCase":	ValueOf(unicode.TitleCase),
		"To":	ValueOf(unicode.To),
		"ToLower":	ValueOf(unicode.ToLower),
		"ToTitle":	ValueOf(unicode.ToTitle),
		"ToUpper":	ValueOf(unicode.ToUpper),
		"TurkishCase":	ValueOf(&unicode.TurkishCase).Elem(),
		"Ugaritic":	ValueOf(&unicode.Ugaritic).Elem(),
		"Unified_Ideograph":	ValueOf(&unicode.Unified_Ideograph).Elem(),
		"Upper":	ValueOf(&unicode.Upper).Elem(),
		"UpperCase":	ValueOf(unicode.UpperCase),
		"UpperLower":	ValueOf(unicode.UpperLower),
		"Vai":	ValueOf(&unicode.Vai).Elem(),
		"Variation_Selector":	ValueOf(&unicode.Variation_Selector).Elem(),
		"Version":	ValueOf(unicode.Version),
		"Warang_Citi":	ValueOf(&unicode.Warang_Citi).Elem(),
		"White_Space":	ValueOf(&unicode.White_Space).Elem(),
		"Yi":	ValueOf(&unicode.Yi).Elem(),
		"Z":	ValueOf(&unicode.Z).Elem(),
		"Zanabazar_Square":	ValueOf(&unicode.Zanabazar_Square).Elem(),
		"Zl":	ValueOf(&unicode.Zl).Elem(),
		"Zp":	ValueOf(&unicode.Zp).Elem(),
		"Zs":	ValueOf(&unicode.Zs).Elem(),
	}, Types: map[string]Type{
		"CaseRange":	TypeOf((*unicode.CaseRange)(nil)).Elem(),
		"Range16":	TypeOf((*unicode.Range16)(nil)).Elem(),
		"Range32":	TypeOf((*unicode.Range32)(nil)).Elem(),
		"RangeTable":	TypeOf((*unicode.RangeTable)(nil)).Elem(),
		"SpecialCase":	TypeOf((*unicode.SpecialCase)(nil)).Elem(),
	}, Untypeds: map[string]string{
		"LowerCase":	"int:1",
		"MaxASCII":	"rune:127",
		"MaxCase":	"int:3",
		"MaxLatin1":	"rune:255",
		"MaxRune":	"rune:1114111",
		"ReplacementChar":	"rune:65533",
		"TitleCase":	"int:2",
		"UpperCase":	"int:0",
		"UpperLower":	"rune:1114112",
		"Version":	"string:10.0.0",
	}, 
	}
}

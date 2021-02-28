package static

type RunicRune struct {
	Rune rune
	Tags []string
	Nemonics []string
}

const (
	EF = "ELDER-FUTHARK"
	AS = "ANGLO-SAXON"
	YFL = "YOUNGER-FUTHARK-LONG"
	YFS = "YOUNGER-FUTHARK-SHORT"
	M = "MEDIEVAL"
	P = "PUNCTUATION"
	T = "TOLKEIN"
	FC = "FRANKS-CASKET"
	GN = "GOLDEN-NUMBER"
)

var RunicAlphabet = []RunicRune{
	{'ᚠ', []string{EF, AS, YFL, YFS, M}, []string{"FEHU", "FEOH", "FE", "F"}},
	{'ᚡ', []string{M}, []string{"V"}},
	{'ᚢ', []string{EF, AS, YFL, YFS, M}, []string{"URUZ", "UR", "U"}},
	{'ᚣ', []string{AS}, []string{"YR"}},
	{'ᚤ', []string{M},  []string{"Y"}},
	{'ᚥ', []string{M}, []string{"W"}},
	{'ᚦ', []string{EF, AS, YFL, YFS, M}, []string{"THURISAZ", "THURS", "THORN"}},
	{'ᚧ', []string{M}, []string{"ETH"}},
	{'ᚨ', []string{EF, AS}, []string{"ANSUZ", "A"}},
	{'ᚩ', []string{AS},  []string{"OS", "O"}},
	{'ᚪ', []string{AS},  []string{"AC", "A"}},
	{'ᚫ', []string{AS}, []string{"AESC"}},
	{'ᚬ', []string{YFL}, []string{"OSS", "O"}},
	{'ᚭ', []string{YFS}, []string{"OSS", "O"}},
	{'ᚮ', []string{M}, []string{"O"}},
	{'ᚯ', []string{M}, []string{"OE"}},
	{'ᚰ', []string{M}, []string{"ON"}},
	{'ᚱ', []string{EF, AS, YFL, YFS, M}, []string{"RAIDO", "RAD", "REID", "R"}},
	{'ᚲ', []string{EF}, []string{"KAUNA"}},
	{'ᚳ', []string{AS}, []string{"CEN"}},
	{'ᚴ', []string{YFL, YFS, M}, []string{"KAUN", "K"}},
	{'ᚵ', []string{M}, []string{"G"}},
	{'ᚶ', []string{M}, []string{"ENG"}},
	{'ᚷ', []string{EF, AS}, []string{"GEBO", "GYFU", "G"}},
	{'ᚸ', []string{AS}, []string{"GAR"}},
	{'ᚹ', []string{EF, AS}, []string{"WUNJO", "WYNN", "W"}},
	{'ᚺ', []string{EF}, []string{"HAGLAZ", "H"}},
	{'ᚻ', []string{AS}, []string{"HAEGL", "H"}},
	{'ᚼ', []string{YFL}, []string{"HAGALL", "H"}},
	{'ᚽ', []string{YFS}, []string{"HAGALL", "H"}},
	{'ᚾ', []string{EF, AS, YFL}, []string{"NAUDIZ", "NYD", "NAUD", "N"}},
	{'ᚿ', []string{YFS, M}, []string{"NAUD", "N"}},
	{'ᛀ', []string{M, "DOTTED"}, []string{"N"}},
	{'ᛁ', []string{EF, AS, YFL, YFS, M}, []string{"ISAZ", "IS", "ISS", "I"}},
	{'ᛂ', []string{M}, []string{"E"}},
	{'ᛃ', []string{EF}, []string{"JERAN", "J"}},
	{'ᛄ', []string{AS}, []string{"GER"}},
	{'ᛅ', []string{YFL, M}, []string{"AR", "AE"}},
	{'ᛆ', []string{YFS, M}, []string{"AR", "A"}},
	{'ᛇ', []string{EF, AS}, []string{"IWAZ", "EOH"}},
	{'ᛈ', []string{EF, AS}, []string{"PERTHO", "PEORTH", "P"}},
	{'ᛉ', []string{EF, AS}, []string{"ALGIZ", "EOLHX"}},
	{'ᛊ', []string{EF}, []string{"SOWILO", "S"}},
	{'ᛋ', []string{AS, YFL, M}, []string{"SIGEL", "SOL", "S"}},
	{'ᛌ', []string{YFS, M}, []string{"SOL", "S"}},
	{'ᛍ', []string{M}, []string{"C"}},
	{'ᛎ', []string{M}, []string{"Z"}},
	{'ᛏ', []string{EF, AS, YFL}, []string{"TIWAZ", "TIR", "TYR", "T"}},
	{'ᛐ', []string{YFS, M}, []string{"TYR", "T"}},
	{'ᛑ', []string{M}, []string{"D"}},
	{'ᛒ', []string{EF, AS, YFL, M}, []string{"BERKANAN", "BEORC", "BJARKAN", "B"}},
	{'ᛓ', []string{YFS}, []string{"BJARKAN", "B"}},
	{'ᛔ', []string{M, "DOTTED"}, []string{"P"}},
	{'ᛕ', []string{M, "OPEN"}, []string{"P"}},
	{'ᛖ', []string{EF, AS}, []string{"EHWAZ", "EH", "E"}},
	{'ᛗ', []string{EF, AS}, []string{"MANNAZ", "MAN", "M"}},
	{'ᛘ', []string{YFL, M}, []string{"MADR", "M"}},
	{'ᛙ', []string{YFS, M}, []string{"MADR", "M"}},
	{'ᛚ', []string{EF, AS, YFL, YFS, M}, []string{"LAUKAZ", "LAGU", "LOGR", "L"}},
	{'ᛛ', []string{M, "DOTTED"}, []string{"L"}},
	{'ᛜ', []string{EF}, []string{"INGWAZ"}},
	{'ᛝ', []string{AS},  []string{"ING"}},
	{'ᛞ', []string{EF, AS}, []string{"DAGAZ", "DAEG", "D"}},
	{'ᛟ', []string{EF, AS}, []string{"OTHALAN", "ETHEL", "O"}},
	{'ᛠ', []string{AS}, []string{"EAR"}},
	{'ᛡ', []string{AS}, []string{"IOR"}},
	{'ᛢ', []string{AS}, []string{"CWEORTH"}},
	{'ᛣ', []string{AS}, []string{"CALC"}},
	{'ᛤ', []string{AS}, []string{"CEALC"}},
	{'ᛥ', []string{AS}, []string{"STAN"}},
	{'ᛦ', []string{YFL, M}, []string{"YR"}},
	{'ᛧ', []string{YFS}, []string{"YR"}},
	{'ᛨ', []string{M, "ICELANDIC"}, []string{"YR"}},
	{'ᛩ', []string{M}, []string{"Q"}},
	{'ᛪ', []string{M}, []string{"X"}},
	{'᛫', []string{"SINGLE", P}, nil},
	{'᛬', []string{"MULTIPLE", P}, nil},
	{'᛭', []string{"CROSS", P}, nil},
	{'ᛮ', []string{GN, "17", "ARLAUG"}, nil},
	{'ᛯ', []string{GN, "18", "TVIMADUR"}, nil},
	{'ᛰ', []string{GN, "19", "BELGTHOR"}, nil},
	{'ᛱ', []string{T}, []string{"K"}},
	{'ᛲ', []string{T}, []string{"SH"}},
	{'ᛳ', []string{T}, []string{"OO"}},
	{'ᛴ', []string{FC}, []string{"OS"}},
	{'ᛵ', []string{FC}, []string{"IS"}},
	{'ᛶ', []string{FC}, []string{"EH"}},
	{'ᛷ', []string{FC}, []string{"AC"}},
	{'ᛸ', []string{FC}, []string{"AESC"}},
}

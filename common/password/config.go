package password

const (
	// LengthWeak weak length password
	// 弱密码长度
	LengthWeak = 6

	// LengthOK ok length password
	// 标准密码长度
	LengthOK = 12

	// LengthStrong strong length password
	// 强密码长度
	LengthStrong = 24

	// LengthVeryStrong very strong length password
	// 非常强密码长度
	LengthVeryStrong = 36

	// DefaultLetterSet is the letter set that is defaulted to - just the
	// alphabet
	// 默认字母集，小写26个字母
	DefaultLetterSet = "abcdefghijklmnopqrstuvwxyz"

	// DefaultLetterAmbiguousSet are letters which are removed from the
	// chosen character set if removing similar characters
	// 默认删除指定字母
	DefaultLetterAmbiguousSet = "ijlo"

	// DefaultNumberSet the default symbol set if character set hasn't been
	// selected
	// 默认数字集，0-9
	DefaultNumberSet = "0123456789"

	// DefaultNumberAmbiguousSet are the numbers which are removed from the
	// chosen character set if removing similar characters
	// 默认删除指定数字
	DefaultNumberAmbiguousSet = "01"

	// DefaultSymbolSet the default symbol set if character set hasn't been
	// selected
	// 默认特殊字符集
	DefaultSymbolSet = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"

	// DefaultSymbolAmbiguousSet are the symbols which are removed from the
	// chosen character set if removing ambiguous characters
	// 默认删除指定特殊字符
	DefaultSymbolAmbiguousSet = "<>[](){}:;'/|\\,"
)

// DefaultConfig is the default configuration, defaults to:
//   - length = 24
//   - Includes symbols, numbers, lowercase and uppercase letters.
//   - Excludes similar and ambiguous characters
var DefaultConfig = Config{
	Length:                     LengthStrong,
	IncludeSymbols:             true,
	IncludeNumbers:             true,
	IncludeLowercaseLetters:    true,
	IncludeUppercaseLetters:    true,
	ExcludeSimilarCharacters:   true,
	ExcludeAmbiguousCharacters: true,
}

// Config is the config struct to hold the settings about
// what type of password to generate
type Config struct {
	// Length is the length of password to generate
	// 生成密码的长度
	Length int

	// CharacterSet is the setting to manually set the
	// character set
	// 手动设置的字符集
	CharacterSet string

	// IncludeSymbols is the setting to include symbols in
	// the character set
	// i.e. !"£*
	IncludeSymbols bool

	// IncludeNumbers is the setting to include number in
	// the character set
	// i.e. 1234
	IncludeNumbers bool

	// IncludeLowercaseLetters is the setting to include
	// lowercase letters in the character set
	// i.e. abcde
	IncludeLowercaseLetters bool

	// IncludeUppercaseLetters is the setting to include
	// uppercase letters in the character set
	// i.e. ABCD
	IncludeUppercaseLetters bool

	// ExcludeSimilarCharacters is the setting to exclude
	// characters that look the same in the character set
	// i.e. i1jIo0
	// 是否删除指定字母或数字
	ExcludeSimilarCharacters bool

	// ExcludeAmbiguousCharacters is the setting to exclude
	// characters that can be hard to remember or symbols
	// that are rarely used
	// i.e. <>{}[]()/|\`
	// 是否删除指定特殊字符
	ExcludeAmbiguousCharacters bool
}

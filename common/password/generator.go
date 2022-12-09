package password

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Generator is what generates the password
type Generator struct {
	*Config
}

// NewWithDefault returns a new generator with the default
// config
// 创建一个默认配置的密码生成器
func NewWithDefault() *Generator {
	return New(&DefaultConfig)
}

// New returns a new generator
// 创建一个自定义配置的密码生成器
func New(config *Config) *Generator {
	if config == nil {
		config = &DefaultConfig
	}

	if !config.IncludeSymbols &&
		!config.IncludeLowercaseLetters &&
		!config.IncludeUppercaseLetters &&
		!config.IncludeNumbers &&
		config.CharacterSet == "" {
		config = &DefaultConfig
	}

	if config.Length == 0 {
		config.Length = LengthStrong
	}

	if config.CharacterSet == "" {
		config.CharacterSet = buildCharacterSet(config)
	}

	return &Generator{Config: config}
}

// GenerateMany generates multiple passwords with length set
// in the config
// 根据配置中的密码长度，生成多个密码
func (g Generator) GenerateMany(amount int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.Generate()
		if err != nil {
			return nil, err
		}

		generated = append(generated, *str)
	}

	return generated, nil
}

// Generate generates one password with length set in the
// config
// 根据配置中的密码长度，生成一个密码
func (g Generator) Generate() (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := 0; i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}

	return &generated, nil
}

// GenerateManyWithLength generates multiple passwords with set length
// 根据自定义密码长度，生成多个密码

func (g Generator) GenerateManyWithLength(length, amount int) ([]string, error) {
	var generated []string

	for i := 0; i < amount; i++ {
		str, err := g.GenerateWithLength(length)
		if err != nil {
			return nil, err
		}

		generated = append(generated, *str)
	}

	return generated, nil
}

// GenerateWithLength generate one password with set length
// 根据自定义密码长度，生成一个密码
func (g Generator) GenerateWithLength(length int) (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := 0; i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}

		generated += characterSet[val.Int64()]
	}

	return &generated, nil
}

// buildCharacterSet 生成密码的原始数据集
func buildCharacterSet(config *Config) string {
	var characterSet string

	if config.IncludeLowercaseLetters {
		characterSet += DefaultLetterSet
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, DefaultLetterAmbiguousSet)
		}
	}

	if config.IncludeUppercaseLetters {
		characterSet += strings.ToUpper(DefaultLetterSet)
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, strings.ToUpper(DefaultLetterAmbiguousSet))
		}
	}

	if config.IncludeNumbers {
		characterSet += DefaultNumberSet
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, DefaultNumberAmbiguousSet)
		}
	}

	if config.IncludeSymbols {
		characterSet += DefaultSymbolSet
		if config.ExcludeAmbiguousCharacters {
			characterSet = removeCharacters(characterSet, DefaultSymbolAmbiguousSet)
		}
	}

	return characterSet
}

// removeCharacters 删除字符，characters 在 str 存在则删除对应字符，不存在则保留
func removeCharacters(str, characters string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}, str)
}

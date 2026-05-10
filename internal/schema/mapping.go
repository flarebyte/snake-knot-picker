// purpose: Map compiled schema flags into concrete validator option structs.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package schema

import (
	"strconv"

	"github.com/flarebyte/snake-knot-picker/internal/validators"
)

func StringOptionsFromSpec(spec *CompiledSpec) (validators.StringOptions, error) {
	var out validators.StringOptions
	if spec == nil {
		return out, nil
	}
	out.Alphabetic = hasFlag(spec.Flags, "--alphabetic")
	out.Whitespace = hasFlag(spec.Flags, "--whitespace")
	out.Lowercase = hasFlag(spec.Flags, "--lowercase")
	out.Uppercase = hasFlag(spec.Flags, "--uppercase")
	out.Punctuation = hasFlag(spec.Flags, "--punctuation")
	out.Hexa = hasFlag(spec.Flags, "--hexa")
	out.Blank = hasFlag(spec.Flags, "--blank")
	out.UnicodeLetter = hasFlag(spec.Flags, "--unicode-letter")
	out.UnicodeNumber = hasFlag(spec.Flags, "--unicode-number")
	out.UnicodePunctuation = hasFlag(spec.Flags, "--unicode-punctuation")
	out.UnicodeSymbol = hasFlag(spec.Flags, "--unicode-symbol")
	out.UnicodeSeparator = hasFlag(spec.Flags, "--unicode-separator")
	out.Latin = hasFlag(spec.Flags, "--latin")
	out.Han = hasFlag(spec.Flags, "--han")
	out.Devanagari = hasFlag(spec.Flags, "--devanagari")
	out.Arabic = hasFlag(spec.Flags, "--arabic")
	out.Hiragana = hasFlag(spec.Flags, "--hiragana")
	out.Katakana = hasFlag(spec.Flags, "--katakana")
	out.Hangul = hasFlag(spec.Flags, "--hangul")
	out.Tamil = hasFlag(spec.Flags, "--tamil")
	out.Gujarati = hasFlag(spec.Flags, "--gujarati")
	out.Ethiopic = hasFlag(spec.Flags, "--ethiopic")
	out.Base64 = hasFlag(spec.Flags, "--base64")
	out.BooleanString = hasFlag(spec.Flags, "--boolean")
	out.StartsWith = firstFlagValue(spec.Flags, "--starts-with")

	enumRaw := firstFlagValue(spec.Flags, "--enum")
	if enumRaw != "" {
		sep := firstFlagValueDefault(spec.Flags, "--enum-separator", ",")
		enum, err := validators.ParseEnumCandidates(enumRaw, sep)
		if err != nil {
			return out, err
		}
		out.Enum = enum
	}
	return out, nil
}

func NumberOptionsFromSpec(spec *CompiledSpec) (validators.NumberOptions, error) {
	var out validators.NumberOptions
	if spec == nil {
		return out, nil
	}
	out.Int = hasFlag(spec.Flags, "--int")
	if v := firstFlagValue(spec.Flags, "--min"); v != "" {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return out, err
		}
		out.Min = &n
	}
	if v := firstFlagValue(spec.Flags, "--max"); v != "" {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return out, err
		}
		out.Max = &n
	}
	if v := firstFlagValue(spec.Flags, "--multiple-of"); v != "" {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return out, err
		}
		out.MultipleOf = &n
	}
	return out, nil
}

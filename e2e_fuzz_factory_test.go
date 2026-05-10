// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import (
	"fmt"
)

var (
	interestingModeValues = []string{
		"normal", "delicate", "whites",
		"", " normal", "normal ", "NORMAL", "nørmal",
	}
	interestingSpinValues = []string{
		"-1", "0", "1", "1200", "999999999", "9223372036854775807",
		"abc", "1e309", "NaN", "+12", "12x",
	}
	interestingTupleValues = []string{
		"0,1", "-1,0", "1,2,3", "1", "", ",", " , ",
		"10,20", "000,001", "1e2,2e2", "x,y",
	}
	interestingAddValues = []string{
		"soap", "bleach", "softener", "",
		"soap,bleach", "soap,,bleach", ",soap", "soap,",
		" soap ", "x", "x,y,z",
	}
	interestingUnknownFlags = []string{
		"--not-a-flag", "--unknown", "--__proto__", "--modee",
	}
)

type fuzzRand struct {
	data []byte
	pos  int
}

func newFuzzRand(data []byte) *fuzzRand {
	return &fuzzRand{data: data}
}

func (r *fuzzRand) next() byte {
	if len(r.data) == 0 {
		return 0
	}
	b := r.data[r.pos%len(r.data)]
	r.pos++
	return b
}

func (r *fuzzRand) bool() bool {
	return r.next()%2 == 0
}

func (r *fuzzRand) intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.next()) % n
}

func (r *fuzzRand) pick(values []string) string {
	if len(values) == 0 {
		return ""
	}
	return values[r.intn(len(values))]
}

func (r *fuzzRand) shuffle(values []string) {
	for i := len(values) - 1; i > 0; i-- {
		j := r.intn(i + 1)
		values[i], values[j] = values[j], values[i]
	}
}

func buildRandomWashStartArgv(data []byte) []string {
	r := newFuzzRand(data)
	argv := []string{"wash", "start"}

	scenario := r.intn(5)
	switch scenario {
	case 0:
		// Valid-ish scenario with randomized values and order.
		flags := []string{
			fmt.Sprintf("--mode=%s", r.pick(interestingModeValues)),
			fmt.Sprintf("--spin=%s", r.pick(interestingSpinValues)),
			fmt.Sprintf("--range=%s", r.pick(interestingTupleValues)),
		}
		if r.bool() {
			flags = append(flags, "--extra-rinse")
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	case 1:
		// Unknown flag scenario.
		argv = append(argv, r.pick(interestingUnknownFlags), r.pick([]string{"x", "y", "z", "", "123"}))
	case 2:
		// Schema command injection attempt.
		argv = append(argv, "schema", "string", "--required")
	case 3:
		// Invalid type for number.
		argv = append(argv, "--spin", r.pick(interestingSpinValues))
	default:
		// Mixed flag forms and shuffled order.
		flags := []string{
			"--mode", r.pick(interestingModeValues),
			"--range", r.pick(interestingTupleValues),
		}
		if r.bool() {
			flags = append(flags, "--spin", r.pick(interestingSpinValues))
		} else {
			flags = append(flags, fmt.Sprintf("--spin=%s", r.pick(interestingSpinValues)))
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	}
	return argv
}

func buildRandomTupleRepeatableArgv(data []byte) []string {
	r := newFuzzRand(data)
	argv := []string{"wash", "start"}

	scenario := r.intn(7)
	switch scenario {
	case 0:
		// Valid tuple + repeatable list with randomized order.
		flags := []string{
			fmt.Sprintf("--range=%s", r.pick(interestingTupleValues)),
			"--add", r.pick(interestingAddValues),
			"--add", r.pick(interestingAddValues),
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	case 1:
		// Tuple arity too small.
		argv = append(argv, "--range", r.pick([]string{"", "1", "-1", "x"}))
	case 2:
		// Tuple arity too large.
		argv = append(argv, "--range", r.pick([]string{"1,2,3", "0,0,0,0", "-1,0,1"}))
	case 3:
		// Empty tuple payload.
		argv = append(argv, "--range", "")
	case 4:
		// Repeatable only.
		argv = append(argv,
			"--add", r.pick(interestingAddValues),
			"--add", r.pick(interestingAddValues),
		)
	case 5:
		// Mixed malformed + unknown.
		argv = append(argv, "--add", r.pick(interestingUnknownFlags), "x")
	default:
		// Valid-ish tuple, mode and spin shuffled around.
		flags := []string{
			fmt.Sprintf("--range=%s", r.pick(interestingTupleValues)),
			fmt.Sprintf("--mode=%s", r.pick(interestingModeValues)),
			fmt.Sprintf("--spin=%s", r.pick(interestingSpinValues)),
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	}
	return argv
}

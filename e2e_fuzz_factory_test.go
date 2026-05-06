package picker

import (
	"fmt"
	"strconv"
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
			fmt.Sprintf("--mode=%s", r.pick([]string{"normal", "delicate", "whites"})),
			fmt.Sprintf("--spin=%s", strconv.Itoa(200+r.intn(2600))),
			fmt.Sprintf("--range=%d,%d", r.intn(100), r.intn(100)),
		}
		if r.bool() {
			flags = append(flags, "--extra-rinse")
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	case 1:
		// Unknown flag scenario.
		argv = append(argv, "--not-a-flag", r.pick([]string{"x", "y", "z", "123"}))
	case 2:
		// Schema command injection attempt.
		argv = append(argv, "schema", "string", "--required")
	case 3:
		// Invalid type for number.
		argv = append(argv, "--spin", r.pick([]string{"abc", "12x", "∞"}))
	default:
		// Mixed flag forms and shuffled order.
		flags := []string{
			"--mode", r.pick([]string{"normal", "delicate"}),
			"--range", fmt.Sprintf("%d,%d", r.intn(100), r.intn(100)),
		}
		if r.bool() {
			flags = append(flags, "--spin", strconv.Itoa(r.intn(3000)))
		} else {
			flags = append(flags, fmt.Sprintf("--spin=%d", r.intn(3000)))
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
			fmt.Sprintf("--range=%d,%d", r.intn(100), r.intn(100)),
			"--add", r.pick([]string{"soap", "bleach", "softener"}),
			fmt.Sprintf("--add=%s,%s", r.pick([]string{"x", "y", "z"}), r.pick([]string{"a", "b", "c"})),
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	case 1:
		// Tuple arity too small.
		argv = append(argv, "--range", strconv.Itoa(r.intn(100)))
	case 2:
		// Tuple arity too large.
		argv = append(argv, "--range", fmt.Sprintf("%d,%d,%d", r.intn(100), r.intn(100), r.intn(100)))
	case 3:
		// Empty tuple payload.
		argv = append(argv, "--range", "")
	case 4:
		// Repeatable only.
		argv = append(argv,
			"--add", r.pick([]string{"soap", "bleach"}),
			"--add", r.pick([]string{"rinse", "freshener"}),
		)
	case 5:
		// Mixed malformed + unknown.
		argv = append(argv, "--add", "--unknown", "x")
	default:
		// Valid-ish tuple, mode and spin shuffled around.
		flags := []string{
			fmt.Sprintf("--range=%d,%d", r.intn(100), r.intn(100)),
			fmt.Sprintf("--mode=%s", r.pick([]string{"normal", "delicate", "whites"})),
			fmt.Sprintf("--spin=%d", 200+r.intn(2600)),
		}
		r.shuffle(flags)
		argv = append(argv, flags...)
	}
	return argv
}

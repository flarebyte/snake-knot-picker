package picker

import "testing"

func FuzzE2ESecurityRejectsFlagLikeValues(f *testing.F) {
	raw := mustLoadArgsCommandFixture(f)
	f.Add([]byte{0})
	f.Add([]byte{1})
	f.Add([]byte{2})
	f.Add([]byte{3})
	f.Add([]byte{4})

	f.Fuzz(func(t *testing.T, noise []byte) {
		argv := buildSecurityArgv(noise)
		_, err := ValidateWithDocumentJSON(raw, argv)
		for i := 0; i+1 < len(argv); i++ {
			if argv[i] == "--mode" || argv[i] == "--add" || argv[i] == "--range" {
				if isFlagLikeValueToken(argv[i+1]) {
					assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)
					return
				}
			}
		}
	})
}

func FuzzE2ESecurityRejectsFlagLikeCSVSegments(f *testing.F) {
	raw := mustLoadArgsCommandFixture(f)
	f.Add("--x,soap")
	f.Add("soap, --x")
	f.Add("soap,\t--x")
	f.Add(" --x ")
	f.Add("a,b,c")

	f.Fuzz(func(t *testing.T, csv string) {
		argv := []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20", "--add", csv}
		_, err := ValidateWithDocumentJSON(raw, argv)
		if isFlagLikeValueToken(csv) || hasFlagLikeCSVSegment(csv) {
			assertErrorDetail(t, err, ErrorIDValidationInvalidType, ErrorKindValidation)
		}
	})
}

func isFlagLikeValueToken(v string) bool {
	return hasFlagLikeValue([]string{v})
}

func hasFlagLikeCSVSegment(csv string) bool {
	for _, part := range splitCSV(csv) {
		if isFlagLikeValueToken(part) {
			return true
		}
	}
	return false
}

func buildSecurityArgv(noise []byte) []string {
	r := newFuzzRand(noise)
	switch r.intn(5) {
	case 0:
		return []string{"wash", "start", "--mode", "--x"}
	case 1:
		return []string{"wash", "start", "--mode", " --x"}
	case 2:
		return []string{"wash", "start", "--mode", "\t--x"}
	case 3:
		return []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", "10,20", "--add", " --x"}
	default:
		return []string{"wash", "start", "--mode", "normal", "--spin", "1200", "--range", " --x,2"}
	}
}

package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if a == "" && b == "" {
		return 0, nil
	}
	if (a == "" && b != "") || (a != "" && b == "") {
		return 0, errors.New("empty strand")
	}
	if len(a) == 1 && len(b) == 1 {
		if a == b {
			return 0, nil
		} else {
			return 1, nil
		}
	}
	if (len(a) > len(b)) || (len(b) > len(a)) {
		return 0, errors.New("strand longer")
	}
	var distance int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

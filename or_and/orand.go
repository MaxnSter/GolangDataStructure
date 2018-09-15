package or_and

func OR(chs ...<-chan interface{}) <-chan interface{} {
	switch len(chs) {
	case 0:
		return nil
	case 1:
		return chs[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(chs) {
		case 2:
			select {
			case <-chs[0]:
			case <-chs[1]:
			}
		default:
			select {
			case <-chs[0]:
			case <-chs[1]:
			case <-chs[2]:
			case <-OR(append(chs[3:], orDone)...):
			}
		}

	}()
	return orDone
}

func AND(chs ...<-chan interface{}) <-chan interface{} {
	switch len(chs) {
	case 0:
		return nil
	case 1:
		return chs[1]
	}

	andDone := make(chan interface{})
	collector := make(chan interface{}, len(chs))
	go func() {
		defer close(andDone)

		switch len(chs) {
		case 2:
			select {
			case collector <- <-chs[0]:
			case collector <- <-chs[1]:
			}
		default:
			select {
			case collector <- <-chs[0]:
			case collector <- <-chs[1]:
			case collector <- <-chs[2]:
			case collector <- <-AND(append(chs[3:], andDone)...):
			}
		}
	}()
	return andDone
}

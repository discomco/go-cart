package status

// HasFlag accepts an input status and will check if the flag is set using bitwise OR and return true if the flag is set.
func HasFlag[T ~int](status T, flag T) bool {
	return status|flag == status
}

// NotHasFlag accepts an input status and will check if the flag is not set.
func NotHasFlag[T ~int](status T, flag T) bool {
	return !HasFlag(status, flag)
}

// UnsetFlag accepts an input status reference and will unset the flag using bitwise AND NOT.
func UnsetFlag[T ~int](status *T, flag T) {
	*status = *status &^ flag
}

// ForceFlag accepts an input status reference and will simply overwrite the status with the flag.
func ForceFlag[T ~int](status *T, flag T) {
	*status = flag
}

// SetFlag accepts an input status reference and will set the flag to the status using bitwise OR.
func SetFlag[T ~int](status *T, flag T) {
	*status = *status | flag
}

// SetFlags allows you to set multiple flags at once
func SetFlags[T ~int](status *T, flags ...T) {
	if len(flags) == 0 {
		return
	}
	for _, flag := range flags {
		SetFlag[T](status, flag)
	}
}

// HasFlags allows you to check for many flags at once.
func HasFlags[T ~int](status T, flags ...T) bool {
	if len(flags) == 0 {
		return false
	}
	for _, flag := range flags {
		if NotHasFlag[T](status, flag) {
			return false
		}
	}
	return true
}

// UnsetFlags allows you to clear multiple flags at once.
func UnsetFlags[T ~int](status *T, flags ...T) {
	if len(flags) == 0 {
		return
	}
	for _, flag := range flags {
		UnsetFlag[T](status, flag)
	}
}

func majorityFrequencyGroup(s string) string {
    charMap := make(map[byte]int, len(s))
    for i := 0; i < len(s); i++ {
        charMap[s[i]]++
    }

    groupMap := make(map[int][]byte) // freq -> list of chars
	for ch, freq := range charMap {
		groupMap[freq] = append(groupMap[freq], ch)
	}

	maxSize := 0
	bestFreq := 0
	for freq, chars := range groupMap {
		if len(chars) > maxSize || (len(chars) == maxSize && freq > bestFreq) {
			maxSize = len(chars)
			bestFreq = freq
		}
	}

    result := ""
	for _, ch := range groupMap[bestFreq] {
		result += string(ch)
	}

	return result
}

func reorderSpaces(text string) string {
    totalSpaces := strings.Count(text, " ")
	words := strings.Fields(text)
    if len(words) == 1 {
        return (words[0] + strings.Repeat(" ", totalSpaces))
    }
    inBetweenSpace := strings.Repeat(" ", totalSpaces/(len(words)-1))
	endSpace := strings.Repeat(" ", totalSpaces%(len(words)-1))
    return strings.Join(words, inBetweenSpace) + endSpace
}

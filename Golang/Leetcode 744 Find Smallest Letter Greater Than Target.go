func nextGreatestLetter(letters []byte, target byte) byte {
    for _, char := range letters {
      if char > target {
          return char
      }
    }
    return letters[0]
}

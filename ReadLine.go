// ReadLine - Reads exactly one line of input text from standard-input and works in Mac, Linux and Windows systems.
func ReadLine() string  {
  inputLine := ""
  for {
    var character byte
    inputCount, err := fmt.Scanf("%c", &character)  // Read just one character
    if err != nil { panic(err) }  // If there is IO error then panic
    if inputCount <= 0 { break }  // Exit if there is not any new input character
    // Guarantee it Works on Linux, Windows and Mac
    if character == '\n' || character == '\r' { break }  // Reach to end of the line
    inputLine += string(character)
  }
  return inputLine
}

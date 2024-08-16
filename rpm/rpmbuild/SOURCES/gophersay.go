package main   // This needs to be a binary

  import (
    "fmt"      // For fmt.Println - to produce output
    "os"       // For os.Args - to import and pass CLI arguments from the terminal
    "strings"  // For srtings. - many times we arrange string text
  )

  // Handy word-wrap function from https://rosettacode.org/wiki/Word_wrap#Go (GNU Doc 1.3)
  func wrap(text string, lineWidth int) (wrapped string) {
    words := strings.Fields(text)
    if len(words) == 0 {
      return
    }
    wrapped = words[0]
    spaceLeft := lineWidth - len(wrapped)
    for _, word := range words[1:] {
      if len(word)+1 > spaceLeft {
        wrapped += "\n" + word
        spaceLeft = lineWidth - len(word)
      } else {
        wrapped += " " + word
        spaceLeft -= 1 + len(word)
      }
    }
    return
  }
  
  // The actual execution
  func main() {
    stringargs := strings.Join(os.Args, " ")  // Convert CLI argued line to string
    
   // Re-assign speech as slogan if no arguments entered
    if stringargs == os.Args[0] {
      stringargs = os.Args[0] + " Gopher talkback written in Go for Linux"
    }

    shiftedinput := strings.SplitN(stringargs, " ", 2)  // Cut first word ($0) from string to only what the gopher says, converting back into single-item array
    input := shiftedinput[1]  // Assign first item in single-item array to $input variable

    // Break speech into lines of 50 characters or less
    speech := wrap(input, 50)  // Use our function to wrap-by-word to 50 characters per line
    
    // See which line is longest
    length := 0
    for _, line := range strings.Split(strings.TrimSuffix(speech, "\n"), "\n") {
      if len(line) > length {
        length = len(line)
      }
    }

    // Formatting length
    blength := length + 2

    // Create our filler lines
    bubble := strings.Repeat("-", blength)
    buflen := blength - 3
    buffer := strings.Repeat(" ", buflen)

    fmt.Println(" " + bubble + " ")

    // Println each quoted line
    fillnum := 0
    fill := ""
    for _, line := range strings.Split(strings.TrimSuffix(speech, "\n"), "\n") {
      fillnum = length - len(line)
      fill = strings.Repeat(" ", fillnum)
      fmt.Println("< " + line + fill + " >")

    }

    fmt.Println(" " + bubble + " ")
    fmt.Println(buffer + `  \ \`)
    fmt.Println(buffer + `   \\`)
    fmt.Println(buffer + `    \ (o)—————(o)`)
    fmt.Println(buffer + `       / Q   Q \`)
    fmt.Println(buffer + `      |   o@o   |`)
    fmt.Println(buffer + `      |    W    |`)
    fmt.Println(buffer + `     c|         |ɔ`)
    fmt.Println(buffer + `      |         |`)
    fmt.Println(buffer + `      |         |`)
    fmt.Println(buffer + `       \_______/`)
    fmt.Println(buffer + `       C       Ɔ`)

  }
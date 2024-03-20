package translation

import "strings"

func Translate(word string, language string) string {
  word = sanitise(word)
  language = sanitise((language))
  if word != "hello" {
    return ""
  }
  switch language {
    case "english":
      return "hello"
    case "german":
      return "hallo"
    case "finnish":
      return "hei"
    default:
      return ""
  }
}

func sanitise(w string) string {
  w = strings.TrimSpace(w)
  return strings.ToLower(w) 
}
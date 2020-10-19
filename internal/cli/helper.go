package cli

import (
	"fmt"
	"strings"
)

func WordWrap(text string, limit int) string {
	words := strings.Fields(text)

	if len(words) == 0 {
		return ""
	}

	wrapped := words[0]

	//spaceLeft := limit - len(wrapped)
	//
	//for _, word := range words[1:] {
	//
	//}
	fmt.Println(`hello world`)
	fmt.Println(limit)
	fmt.Println(words[2:])
	return wrapped
}

//const Logo = `
//                      &&&&&&&&&&&& %&&&&&&&&&/
//                        &&&&&&&&&%  %&&&&&&&&
//                          &&&&&&&   %&&&&&&,
//                                &% %&&&&%
//                              &&&% %&&&
//                              &&&&% %/
//                            &&&&&&%    &&&&.
//                          &&&&&&&&%  &&&&&&&&
//                        &&&&&&&&&&% &&&&&&&&&&
//`

/*repeatalpha
Instructions

Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
Expected Function*/

func RepeatAlpha(s string) string {
	result := []rune{}
	for _, ch := range s{

		if ch >= 'a' && ch <= 'z'{
			index := ch - 'a' + 1

			for i := 0, i <int(len(index)); i++{

				result = append(result, ch)
			}
		}
	}else if ch >= 'A' && ch <= 'Z'{
		index := ch - 'A' + 1
		for i := 0; i < int(len(index)); i++{

			result = append(result, ch)
		}
	}else{
		result = (append(result, ch))
	}
	return result
}

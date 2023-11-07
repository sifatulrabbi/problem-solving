package main

func runAltWordList() {

}

func mergeAlternately(word1 string, word2 string) string {
	r := ""
	l := len(word1)
	l2 := len(word2)
	i := 0
	for i < l || i < l2 {
		if i < l {
			r += string(word1[i])
		}
		if i < l2 {
			r += string(word2[i])
		}
		i++
	}
	return r
}

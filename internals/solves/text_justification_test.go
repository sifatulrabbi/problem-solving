package solves

import "testing"

func TestTextJustification(t *testing.T) {
	result := TextJustification([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"})
	t.Log(result)
}

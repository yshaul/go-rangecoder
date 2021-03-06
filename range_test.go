package compress

import (
	"testing"
	"log"
	"time"
)

func TestEncode(t *testing.T) {
	str := "cbababaaaaacaacaacabaaababaaacaacaaabaaaaaabecgacccgcebddhaacacccccaaaacabcbccaaacaaaaabbbaaaaacaacacaaaaap|gganaaaqbfdaccaaabaaaatxbabraaananakohooaababaabaabcbabcabaabcaacaabaaacbaabcababbacccaecaabcccbbawdabaqaabababaabcbacbcacabaabcbaacaacaaaapabqaaccacaaadaaccbabaaaaaaaaaaabcbaababaaaaaaaaacbcbabaaabaacaacaaacabababaaacaaacbaaaaaaaaaababcaaacacabbaabaaaacaccaaaabaacaababaaaabaaacacacbabaabbabacacccccaababbbabaaaaaaaccccacabbabaaaabaaaaccccaccabbbdabbaaaaaacacacaaaaaaaaabaaaaacaaaaaaabcaaaaabaacaaaabaaaaacaaaabaacaaacaabaacaaaabbaacacabaaaaacababaacbaabaaaaacaaaabaaaaaaacaaaacaaaaaacacaaaaabaaaaaaaaaaaabaaaaaacaaababaacaccabbaabaaaaaaaaaaabaaaaaacbaaaacceeaoaaeoieddaaaaatacyaaaabxbubtabcqrrlboocaubabbbaabaaababadarbasabaaaaaabbcbcbabbabacaacbacaaacabcbabaabbabcabcbaaaaaaabcabcabcdcabcbaaabcaaaaaaacaaabbcbcaaabcaacacbaaaccaaaadabbaaaaaeaaaabbaabcaccacbacaaabbbaaaaaaaaacaacbcaabababaacaeacabababaaaacaacccaabbabghacfgabacaaaccaabbaabaaaaacacaeccbababbbabaaaaaaacbcaaaaaaaaeacbaababaaabaaaaacacccccbbbbbbabaaacaacccccbbababbaaaaaaceeccabdbbbababaaacceegceaccabbbbdabaaac|aacaa}baaaaaaaaaaeaaac~abaaa{acaaaaaaabaaaaaaaaaaaacbbbaaaacacaabbaaaaaabaababcbaabaihccfgaaaaaabaaaaaaabbaaaaifeeaaaaabbabbabaaaaaaaeaaabcbcbaabcacabaaabaaaaaaaaececacaeacbddaaaaaanaaacqcaaabacacacaaaaaacb~bjtaqaaaabaabaaabaaaaaaaabaabcbcbaaaaacbabcaaaaaaaaaaaaabaaacbccecababaaabbaabaaacaaaababaaacbccacaaaaabaabaaaaacceeaabbbbbabaabaaacaacaaabaacccaaaaaababbaacbcbaabaaaccgecbababbbbbaaaaabaaaacacbcccccaababbbbaaaaccaccaabbabbaaaaaaaaaccccacbbabaabaccccdaaacbabaabaabcaaaaacaeacacagecaaaagagacabbacaaaecaaaaaacaaafdaraaaasdaaaadadaaaaabbbabaacaaacacca|aaaaaaaaafbbabaacaabbabaadaaaaaaaaaaaaaaaaaaaaaaabaihccaaabaabacaaaaaaabaabcaaacbaacaaaaacacaababababaaaaacbcbcaabaaaaabaaacaaecaabbbaaabacacceccabbc"

	t0 := time.Now()
    
	a := Compress(str)

	t1 := time.Now()
    log.Println("Compress time \t", t1.Sub(t0))

	log.Println(len(a))

	t0 = time.Now()
	
	b := Decompress(a)

	t1 = time.Now()
	log.Println("Decompress time \t", t1.Sub(t0))
	log.Println(len(b))

	s := string(b[:])
log.Println(s)
}

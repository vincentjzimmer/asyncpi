//line asyncpi.y:2
package asyncpi

import __yyfmt__ "fmt"

//line asyncpi.y:2
import (
	"io"
)

var proc Process

//line asyncpi.y:11
type asyncpiSymType struct {
	yys    int
	strval string
	proc   Process
	name   Name
	names  []Name
}

const LANGLE = 57346
const RANGLE = 57347
const LPAREN = 57348
const RPAREN = 57349
const PREFIX = 57350
const SEMICOLON = 57351
const COLON = 57352
const NIL = 57353
const NAME = 57354
const REPEAT = 57355
const NEW = 57356
const COMMA = 57357
const PAR = 57358
const REP = 57359

var asyncpiToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LANGLE",
	"RANGLE",
	"LPAREN",
	"RPAREN",
	"PREFIX",
	"SEMICOLON",
	"COLON",
	"NIL",
	"NAME",
	"REPEAT",
	"NEW",
	"COMMA",
	"PAR",
	"REP",
}
var asyncpiStatenames = [...]string{}

const asyncpiEofCode = 1
const asyncpiErrCode = 2
const asyncpiInitialStackSize = 16

//line asyncpi.y:65

// Parse is the entry point to the asyncpi calculus parser.
func Parse(r io.Reader) (Process, error) {
	l := NewLexer(r)
	asyncpiParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return proc, nil
	}
}

//line yacctab:1
var asyncpiExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const asyncpiNprod = 20
const asyncpiPrivate = 57344

var asyncpiTokenNames []string
var asyncpiStates []string

const asyncpiLast = 50

var asyncpiAct = [...]int{

	3, 31, 2, 16, 17, 8, 6, 11, 38, 13,
	12, 4, 5, 7, 11, 33, 19, 8, 18, 37,
	4, 5, 7, 30, 25, 22, 32, 23, 29, 24,
	34, 35, 26, 23, 27, 6, 36, 20, 32, 39,
	4, 5, 7, 15, 28, 1, 9, 21, 10, 14,
}
var asyncpiPact = [...]int{

	29, -1000, -11, -1000, -1000, 42, -7, 29, 29, 31,
	6, 6, -1000, -1000, 32, -1000, 18, -1000, 19, 17,
	-1000, 22, 36, 6, 11, 9, 6, -1000, 29, -1000,
	-1000, -1000, -1000, 0, 12, -1000, 1, 9, -1000, -1000,
}
var asyncpiPgo = [...]int{

	0, 2, 0, 1, 4, 3, 49, 45,
}
var asyncpiR1 = [...]int{

	0, 7, 1, 1, 2, 2, 2, 2, 2, 2,
	4, 4, 3, 3, 5, 5, 5, 6, 6, 6,
}
var asyncpiR2 = [...]int{

	0, 1, 1, 3, 1, 4, 6, 5, 7, 2,
	1, 3, 1, 3, 0, 1, 3, 0, 1, 3,
}
var asyncpiChk = [...]int{

	-1000, -7, -1, -2, 11, 12, 6, 13, 16, 4,
	6, 14, -1, -2, -6, 12, -5, -4, 12, -4,
	5, 15, 7, 15, 10, 7, 15, 12, 8, -4,
	12, -3, -2, 6, -5, -1, -1, 7, 7, -3,
}
var asyncpiDef = [...]int{

	0, -2, 1, 2, 4, 0, 0, 0, 0, 17,
	14, 0, 9, 3, 0, 18, 0, 15, 10, 0,
	5, 0, 0, 0, 0, 0, 14, 19, 0, 16,
	11, 7, 12, 0, 0, 6, 0, 0, 13, 8,
}
var asyncpiTok1 = [...]int{

	1,
}
var asyncpiTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
}
var asyncpiTok3 = [...]int{
	0,
}

var asyncpiErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	asyncpiDebug        = 0
	asyncpiErrorVerbose = false
)

type asyncpiLexer interface {
	Lex(lval *asyncpiSymType) int
	Error(s string)
}

type asyncpiParser interface {
	Parse(asyncpiLexer) int
	Lookahead() int
}

type asyncpiParserImpl struct {
	lval  asyncpiSymType
	stack [asyncpiInitialStackSize]asyncpiSymType
	char  int
}

func (p *asyncpiParserImpl) Lookahead() int {
	return p.char
}

func asyncpiNewParser() asyncpiParser {
	return &asyncpiParserImpl{}
}

const asyncpiFlag = -1000

func asyncpiTokname(c int) string {
	if c >= 1 && c-1 < len(asyncpiToknames) {
		if asyncpiToknames[c-1] != "" {
			return asyncpiToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func asyncpiStatname(s int) string {
	if s >= 0 && s < len(asyncpiStatenames) {
		if asyncpiStatenames[s] != "" {
			return asyncpiStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func asyncpiErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !asyncpiErrorVerbose {
		return "syntax error"
	}

	for _, e := range asyncpiErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + asyncpiTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := asyncpiPact[state]
	for tok := TOKSTART; tok-1 < len(asyncpiToknames); tok++ {
		if n := base + tok; n >= 0 && n < asyncpiLast && asyncpiChk[asyncpiAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if asyncpiDef[state] == -2 {
		i := 0
		for asyncpiExca[i] != -1 || asyncpiExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; asyncpiExca[i] >= 0; i += 2 {
			tok := asyncpiExca[i]
			if tok < TOKSTART || asyncpiExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if asyncpiExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += asyncpiTokname(tok)
	}
	return res
}

func asyncpilex1(lex asyncpiLexer, lval *asyncpiSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = asyncpiTok1[0]
		goto out
	}
	if char < len(asyncpiTok1) {
		token = asyncpiTok1[char]
		goto out
	}
	if char >= asyncpiPrivate {
		if char < asyncpiPrivate+len(asyncpiTok2) {
			token = asyncpiTok2[char-asyncpiPrivate]
			goto out
		}
	}
	for i := 0; i < len(asyncpiTok3); i += 2 {
		token = asyncpiTok3[i+0]
		if token == char {
			token = asyncpiTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = asyncpiTok2[1] /* unknown char */
	}
	if asyncpiDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", asyncpiTokname(token), uint(char))
	}
	return char, token
}

func asyncpiParse(asyncpilex asyncpiLexer) int {
	return asyncpiNewParser().Parse(asyncpilex)
}

func (asyncpircvr *asyncpiParserImpl) Parse(asyncpilex asyncpiLexer) int {
	var asyncpin int
	var asyncpiVAL asyncpiSymType
	var asyncpiDollar []asyncpiSymType
	_ = asyncpiDollar // silence set and not used
	asyncpiS := asyncpircvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	asyncpistate := 0
	asyncpircvr.char = -1
	asyncpitoken := -1 // asyncpircvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		asyncpistate = -1
		asyncpircvr.char = -1
		asyncpitoken = -1
	}()
	asyncpip := -1
	goto asyncpistack

ret0:
	return 0

ret1:
	return 1

asyncpistack:
	/* put a state and value onto the stack */
	if asyncpiDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", asyncpiTokname(asyncpitoken), asyncpiStatname(asyncpistate))
	}

	asyncpip++
	if asyncpip >= len(asyncpiS) {
		nyys := make([]asyncpiSymType, len(asyncpiS)*2)
		copy(nyys, asyncpiS)
		asyncpiS = nyys
	}
	asyncpiS[asyncpip] = asyncpiVAL
	asyncpiS[asyncpip].yys = asyncpistate

asyncpinewstate:
	asyncpin = asyncpiPact[asyncpistate]
	if asyncpin <= asyncpiFlag {
		goto asyncpidefault /* simple state */
	}
	if asyncpircvr.char < 0 {
		asyncpircvr.char, asyncpitoken = asyncpilex1(asyncpilex, &asyncpircvr.lval)
	}
	asyncpin += asyncpitoken
	if asyncpin < 0 || asyncpin >= asyncpiLast {
		goto asyncpidefault
	}
	asyncpin = asyncpiAct[asyncpin]
	if asyncpiChk[asyncpin] == asyncpitoken { /* valid shift */
		asyncpircvr.char = -1
		asyncpitoken = -1
		asyncpiVAL = asyncpircvr.lval
		asyncpistate = asyncpin
		if Errflag > 0 {
			Errflag--
		}
		goto asyncpistack
	}

asyncpidefault:
	/* default state action */
	asyncpin = asyncpiDef[asyncpistate]
	if asyncpin == -2 {
		if asyncpircvr.char < 0 {
			asyncpircvr.char, asyncpitoken = asyncpilex1(asyncpilex, &asyncpircvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if asyncpiExca[xi+0] == -1 && asyncpiExca[xi+1] == asyncpistate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			asyncpin = asyncpiExca[xi+0]
			if asyncpin < 0 || asyncpin == asyncpitoken {
				break
			}
		}
		asyncpin = asyncpiExca[xi+1]
		if asyncpin < 0 {
			goto ret0
		}
	}
	if asyncpin == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			asyncpilex.Error(asyncpiErrorMessage(asyncpistate, asyncpitoken))
			Nerrs++
			if asyncpiDebug >= 1 {
				__yyfmt__.Printf("%s", asyncpiStatname(asyncpistate))
				__yyfmt__.Printf(" saw %s\n", asyncpiTokname(asyncpitoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for asyncpip >= 0 {
				asyncpin = asyncpiPact[asyncpiS[asyncpip].yys] + asyncpiErrCode
				if asyncpin >= 0 && asyncpin < asyncpiLast {
					asyncpistate = asyncpiAct[asyncpin] /* simulate a shift of "error" */
					if asyncpiChk[asyncpistate] == asyncpiErrCode {
						goto asyncpistack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if asyncpiDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", asyncpiS[asyncpip].yys)
				}
				asyncpip--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if asyncpiDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", asyncpiTokname(asyncpitoken))
			}
			if asyncpitoken == asyncpiEofCode {
				goto ret1
			}
			asyncpircvr.char = -1
			asyncpitoken = -1
			goto asyncpinewstate /* try again in the same state */
		}
	}

	/* reduction by production asyncpin */
	if asyncpiDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", asyncpin, asyncpiStatname(asyncpistate))
	}

	asyncpint := asyncpin
	asyncpipt := asyncpip
	_ = asyncpipt // guard against "declared and not used"

	asyncpip -= asyncpiR2[asyncpin]
	// asyncpip is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if asyncpip+1 >= len(asyncpiS) {
		nyys := make([]asyncpiSymType, len(asyncpiS)*2)
		copy(nyys, asyncpiS)
		asyncpiS = nyys
	}
	asyncpiVAL = asyncpiS[asyncpip+1]

	/* consult goto table to find next state */
	asyncpin = asyncpiR1[asyncpin]
	asyncpig := asyncpiPgo[asyncpin]
	asyncpij := asyncpig + asyncpiS[asyncpip].yys + 1

	if asyncpij >= asyncpiLast {
		asyncpistate = asyncpiAct[asyncpig]
	} else {
		asyncpistate = asyncpiAct[asyncpij]
		if asyncpiChk[asyncpistate] != -asyncpin {
			asyncpistate = asyncpiAct[asyncpig]
		}
	}
	// dummy call; replaced with literal code
	switch asyncpint {

	case 1:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:33
		{
			proc = asyncpiDollar[1].proc
		}
	case 2:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:36
		{
			asyncpiVAL.proc = asyncpiDollar[1].proc
		}
	case 3:
		asyncpiDollar = asyncpiS[asyncpipt-3 : asyncpipt+1]
		//line asyncpi.y:37
		{
			asyncpiVAL.proc = NewPar(asyncpiDollar[1].proc, asyncpiDollar[3].proc)
		}
	case 4:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:40
		{
			asyncpiVAL.proc = NewNilProcess()
		}
	case 5:
		asyncpiDollar = asyncpiS[asyncpipt-4 : asyncpipt+1]
		//line asyncpi.y:41
		{
			asyncpiVAL.proc = NewSend(newPiName(asyncpiDollar[1].strval))
			asyncpiVAL.proc.(*Send).SetVals(asyncpiDollar[3].names)
		}
	case 6:
		asyncpiDollar = asyncpiS[asyncpipt-6 : asyncpipt+1]
		//line asyncpi.y:42
		{
			asyncpiVAL.proc = NewRecv(newPiName(asyncpiDollar[1].strval), asyncpiDollar[6].proc)
			asyncpiVAL.proc.(*Recv).SetVars(asyncpiDollar[3].names)
		}
	case 7:
		asyncpiDollar = asyncpiS[asyncpipt-5 : asyncpipt+1]
		//line asyncpi.y:43
		{
			asyncpiVAL.proc = NewRestrict(asyncpiDollar[3].name, asyncpiDollar[5].proc)
		}
	case 8:
		asyncpiDollar = asyncpiS[asyncpipt-7 : asyncpipt+1]
		//line asyncpi.y:44
		{
			asyncpiVAL.proc = NewRestricts(append([]Name{asyncpiDollar[3].name}, asyncpiDollar[5].names...), asyncpiDollar[7].proc)
		}
	case 9:
		asyncpiDollar = asyncpiS[asyncpipt-2 : asyncpipt+1]
		//line asyncpi.y:45
		{
			asyncpiVAL.proc = NewRepeat(asyncpiDollar[2].proc)
		}
	case 10:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:48
		{
			asyncpiVAL.name = newPiName(asyncpiDollar[1].strval)
		}
	case 11:
		asyncpiDollar = asyncpiS[asyncpipt-3 : asyncpipt+1]
		//line asyncpi.y:49
		{
			asyncpiVAL.name = newTypedPiName(asyncpiDollar[1].strval, asyncpiDollar[3].strval)
		}
	case 12:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:51
		{
			asyncpiVAL.proc = asyncpiDollar[1].proc
		}
	case 13:
		asyncpiDollar = asyncpiS[asyncpipt-3 : asyncpipt+1]
		//line asyncpi.y:52
		{
			asyncpiVAL.proc = asyncpiDollar[2].proc
		}
	case 14:
		asyncpiDollar = asyncpiS[asyncpipt-0 : asyncpipt+1]
		//line asyncpi.y:55
		{
			asyncpiVAL.names = []Name{}
		}
	case 15:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:56
		{
			asyncpiVAL.names = []Name{asyncpiDollar[1].name}
		}
	case 16:
		asyncpiDollar = asyncpiS[asyncpipt-3 : asyncpipt+1]
		//line asyncpi.y:57
		{
			asyncpiVAL.names = append(asyncpiDollar[1].names, asyncpiDollar[3].name)
		}
	case 17:
		asyncpiDollar = asyncpiS[asyncpipt-0 : asyncpipt+1]
		//line asyncpi.y:60
		{
			asyncpiVAL.names = []Name{}
		}
	case 18:
		asyncpiDollar = asyncpiS[asyncpipt-1 : asyncpipt+1]
		//line asyncpi.y:61
		{
			asyncpiVAL.names = []Name{newPiName(asyncpiDollar[1].strval)}
		}
	case 19:
		asyncpiDollar = asyncpiS[asyncpipt-3 : asyncpipt+1]
		//line asyncpi.y:62
		{
			asyncpiVAL.names = append(asyncpiDollar[1].names, newPiName(asyncpiDollar[3].strval))
		}
	}
	goto asyncpistack /* stack new state and value */
}

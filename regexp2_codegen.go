package tiktoken

import (
	"github.com/dlclark/regexp2"
	"github.com/dlclark/regexp2/helpers"
	"github.com/dlclark/regexp2/syntax"
	"unicode"
)

/*
Capture(index = 0, unindex = -1)
 Atomic
  Alternate
   Concatenate
    Setloop(Set = [^\n\r\p{L}\p{N}])(Min = 0, Max = 1)
    Setloop(Set = [\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}])(Min = 0, Max = inf)
    Setloop(Set = [\p{Ll}\p{Lm}\p{Lo}\p{M}])(Min = 1, Max = inf)
    Atomic
     Loop(Min = 0, Max = 1)
      Concatenate
       One(Ch = ')
       Atomic
        Alternate
         Set(Set = [STstſ])
         Concatenate
          Set(Set = [Rr])
          Set(Set = [Ee])
         Concatenate
          Set(Set = [Vv])
          Set(Set = [Ee])
         Set(Set = [Mm])
         SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
         Set(Set = [Dd])
   Concatenate
    Setloop(Set = [^\n\r\p{L}\p{N}])(Min = 0, Max = 1)
    Setloop(Set = [\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}])(Min = 1, Max = inf)
    Setloop(Set = [\p{Ll}\p{Lm}\p{Lo}\p{M}])(Min = 0, Max = inf)
    Atomic
     Loop(Min = 0, Max = 1)
      Concatenate
       One(Ch = ')
       Atomic
        Alternate
         Set(Set = [STstſ])
         Concatenate
          Set(Set = [Rr])
          Set(Set = [Ee])
         Concatenate
          Set(Set = [Vv])
          Set(Set = [Ee])
         Set(Set = [Mm])
         SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
         Set(Set = [Dd])
   SetloopAtomic(Set = [\p{N}])(Min = 1, Max = 3)
   Concatenate
    OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
    SetloopAtomic(Set = [^\s\p{L}\p{N}])(Min = 1, Max = inf)
    SetloopAtomic(Set = [\n\r/])(Min = 0, Max = inf)
   Concatenate
    Setloop(Set = [\s])(Min = 0, Max = inf)
    SetloopAtomic(Set = [\n\r])(Min = 1, Max = inf)
   Concatenate
    Setloop(Set = [\s])(Min = 1, Max = inf)
    NegLook
     Set(Set = [^\s])
   SetloopAtomic(Set = [\s])(Min = 1, Max = inf)
*/
// From encoding.go:132:29
// Pattern: "[^\\r\\n\\p{L}\\p{N}]?[\\p{Lu}\\p{Lt}\\p{Lm}\\p{Lo}\\p{M}]*[\\p{Ll}\\p{Lm}\\p{Lo}\\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?|[^\\r\\n\\p{L}\\p{N}]?[\\p{Lu}\\p{Lt}\\p{Lm}\\p{Lo}\\p{M}]+[\\p{Ll}\\p{Lm}\\p{Lo}\\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?|\\p{N}{1,3}| ?[^\\s\\p{L}\\p{N}]+[\\r\\n/]*|\\s*[\\r\\n]+|\\s+(?!\\S)|\\s+"
// Options: regexp2.None
type reg_Engine struct{}

func (reg_Engine) Caps() map[int]int        { return nil }
func (reg_Engine) CapNames() map[string]int { return nil }
func (reg_Engine) CapsList() []string       { return nil }
func (reg_Engine) CapSize() int             { return 1 }

func (reg_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Empty matches aren't possible
	if pos < len(r.Runtext) {
		return true
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (reg_Engine) Execute(r *regexp2.Runner) error {
	atomic_stackpos := 0
	alternation_starting_pos := 0
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	var charloop_starting_pos1, charloop_ending_pos1 = 0, 0
	iteration := 0
	var charloop_starting_pos2, charloop_ending_pos2 = 0, 0
	iteration1 := 0
	loop_iteration := 0
	startingStackpos := 0
	var charloop_starting_pos3, charloop_ending_pos3 = 0, 0
	var charloop_starting_pos4, charloop_ending_pos4 = 0, 0
	iteration2 := 0
	var charloop_starting_pos5, charloop_ending_pos5 = 0, 0
	iteration3 := 0
	loop_iteration1 := 0
	startingStackpos1 := 0
	iteration4 := 0
	iteration5 := 0
	iteration6 := 0
	var charloop_starting_pos6, charloop_ending_pos6 = 0, 0
	iteration7 := 0
	iteration8 := 0
	var charloop_starting_pos7, charloop_ending_pos7 = 0, 0
	iteration9 := 0
	negativelookahead_starting_pos := 0
	iteration10 := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Atomic
	// Atomic group.
	atomic_stackpos = r.Runstackpos

	// Node: Alternate
	// Match with 7 alternative expressions, atomically.
	alternation_starting_pos = pos

	// Branch 0
	// Node: Concatenate
	// Node: Setloop(Set = [^\n\r\p{L}\p{N}])(Min = 0, Max = 1)
	// Match [^\n\r\p{L}\p{N}] greedily, optionally.
	charloop_starting_pos = pos

	if len(slice) > 0 && set_4a7765fc40e8e8c7561121585d1545f9a51b44bf010a50f8b1b086a02ec1f07f.CharIn(slice[0]) {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos = pos
	goto CharLoopEnd

CharLoopBacktrack:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		goto AlternationBranch
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:

	// Node: Setloop(Set = [\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}])(Min = 0, Max = inf)
	// Match [\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}] greedily any number of times.
	charloop_starting_pos1 = pos

	iteration = 0
	for iteration < len(slice) && unicode.In(slice[iteration], unicode.Lu, unicode.Lt, unicode.Lm, unicode.Lo, unicode.M) {
		iteration++
	}

	slice = slice[iteration:]
	pos += iteration

	charloop_ending_pos1 = pos
	goto CharLoopEnd1

CharLoopBacktrack1:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos1 >= charloop_ending_pos1 {
		goto CharLoopBacktrack
	}
	charloop_ending_pos1--
	pos = charloop_ending_pos1
	slice = r.Runtext[pos:]

CharLoopEnd1:

	// Node: Setloop(Set = [\p{Ll}\p{Lm}\p{Lo}\p{M}])(Min = 1, Max = inf)
	// Match [\p{Ll}\p{Lm}\p{Lo}\p{M}] greedily at least once.
	charloop_starting_pos2 = pos

	iteration1 = 0
	for iteration1 < len(slice) && unicode.In(slice[iteration1], unicode.Ll, unicode.Lm, unicode.Lo, unicode.M) {
		iteration1++
	}

	if iteration1 == 0 {
		goto CharLoopBacktrack1
	}

	slice = slice[iteration1:]
	pos += iteration1

	charloop_ending_pos2 = pos
	charloop_starting_pos2++
	goto CharLoopEnd2

CharLoopBacktrack2:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos2 >= charloop_ending_pos2 {
		goto CharLoopBacktrack1
	}
	charloop_ending_pos2 = helpers.LastIndexOfAny1(r.Runtext[charloop_starting_pos2:charloop_ending_pos2], '\'')
	if charloop_ending_pos2 < 0 { // miss
		goto CharLoopBacktrack1
	}
	charloop_ending_pos2 += charloop_starting_pos2
	pos = charloop_ending_pos2
	slice = r.Runtext[pos:]

CharLoopEnd2:

	// Node: Atomic
	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	startingStackpos = r.Runstackpos
	loop_iteration = 0

LoopBody:
	;
	r.StackPush(pos)

	loop_iteration++

	// Node: Concatenate
	// Node: One(Ch = ')
	// Match '\”.
	if len(slice) == 0 || slice[0] != '\'' {
		goto LoopIterationNoMatch
	}

	// Node: Atomic
	// Node: Alternate
	// Match with 6 alternative expressions, atomically.
	if len(slice) < 2 {
		goto LoopIterationNoMatch
	}

	switch slice[1] {
	case 'S', 'T', 's', 't', 'ſ':
		pos += 2
		slice = r.Runtext[pos:]

	case 'R', 'r':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ee])
		// Match [Ee].
		if len(slice) < 3 || (slice[2]|0x20 != 'e') {
			goto LoopIterationNoMatch
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'V', 'v':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ee])
		// Match [Ee].
		if len(slice) < 3 || (slice[2]|0x20 != 'e') {
			goto LoopIterationNoMatch
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'M', 'm':
		pos += 2
		slice = r.Runtext[pos:]

	case 'L', 'l':
		// Node: SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
		// Match [Ll] exactly 2 times.
		if len(slice) < 3 ||
			(slice[1]|0x20 != 'l') ||
			(slice[2]|0x20 != 'l') {
			goto LoopIterationNoMatch
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'D', 'd':
		pos += 2
		slice = r.Runtext[pos:]

	default:
		goto LoopIterationNoMatch
	}

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration == 0 {
		goto LoopBody
	}
	goto LoopEnd

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch:
	loop_iteration--
	if loop_iteration < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto CharLoopBacktrack2
	}
	pos = r.StackPop()
	slice = r.Runtext[pos:]
LoopEnd:
	r.Runstackpos = startingStackpos // Ensure any remaining backtracking state is removed.

	goto AlternationMatch

AlternationBranch:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Concatenate
	// Node: Setloop(Set = [^\n\r\p{L}\p{N}])(Min = 0, Max = 1)
	// Match [^\n\r\p{L}\p{N}] greedily, optionally.
	charloop_starting_pos3 = pos

	if len(slice) > 0 && set_4a7765fc40e8e8c7561121585d1545f9a51b44bf010a50f8b1b086a02ec1f07f.CharIn(slice[0]) {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos3 = pos
	goto CharLoopEnd3

CharLoopBacktrack3:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos3 >= charloop_ending_pos3 {
		goto AlternationBranch1
	}
	charloop_ending_pos3--
	pos = charloop_ending_pos3
	slice = r.Runtext[pos:]

CharLoopEnd3:

	// Node: Setloop(Set = [\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}])(Min = 1, Max = inf)
	// Match [\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}] greedily at least once.
	charloop_starting_pos4 = pos

	iteration2 = 0
	for iteration2 < len(slice) && unicode.In(slice[iteration2], unicode.Lu, unicode.Lt, unicode.Lm, unicode.Lo, unicode.M) {
		iteration2++
	}

	if iteration2 == 0 {
		goto CharLoopBacktrack3
	}

	slice = slice[iteration2:]
	pos += iteration2

	charloop_ending_pos4 = pos
	charloop_starting_pos4++
	goto CharLoopEnd4

CharLoopBacktrack4:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos4 >= charloop_ending_pos4 {
		goto CharLoopBacktrack3
	}
	charloop_ending_pos4--
	pos = charloop_ending_pos4
	slice = r.Runtext[pos:]

CharLoopEnd4:

	// Node: Setloop(Set = [\p{Ll}\p{Lm}\p{Lo}\p{M}])(Min = 0, Max = inf)
	// Match [\p{Ll}\p{Lm}\p{Lo}\p{M}] greedily any number of times.
	charloop_starting_pos5 = pos

	iteration3 = 0
	for iteration3 < len(slice) && unicode.In(slice[iteration3], unicode.Ll, unicode.Lm, unicode.Lo, unicode.M) {
		iteration3++
	}

	slice = slice[iteration3:]
	pos += iteration3

	charloop_ending_pos5 = pos
	goto CharLoopEnd5

CharLoopBacktrack5:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos5 >= charloop_ending_pos5 {
		goto CharLoopBacktrack4
	}
	charloop_ending_pos5 = helpers.LastIndexOfAny1(r.Runtext[charloop_starting_pos5:charloop_ending_pos5], '\'')
	if charloop_ending_pos5 < 0 { // miss
		goto CharLoopBacktrack4
	}
	charloop_ending_pos5 += charloop_starting_pos5
	pos = charloop_ending_pos5
	slice = r.Runtext[pos:]

CharLoopEnd5:

	// Node: Atomic
	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	startingStackpos1 = r.Runstackpos
	loop_iteration1 = 0

LoopBody1:
	;
	r.StackPush(pos)

	loop_iteration1++

	// Node: Concatenate
	// Node: One(Ch = ')
	// Match '\”.
	if len(slice) == 0 || slice[0] != '\'' {
		goto LoopIterationNoMatch1
	}

	// Node: Atomic
	// Node: Alternate
	// Match with 6 alternative expressions, atomically.
	if len(slice) < 2 {
		goto LoopIterationNoMatch1
	}

	switch slice[1] {
	case 'S', 'T', 's', 't', 'ſ':
		pos += 2
		slice = r.Runtext[pos:]

	case 'R', 'r':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ee])
		// Match [Ee].
		if len(slice) < 3 || (slice[2]|0x20 != 'e') {
			goto LoopIterationNoMatch1
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'V', 'v':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ee])
		// Match [Ee].
		if len(slice) < 3 || (slice[2]|0x20 != 'e') {
			goto LoopIterationNoMatch1
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'M', 'm':
		pos += 2
		slice = r.Runtext[pos:]

	case 'L', 'l':
		// Node: SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
		// Match [Ll] exactly 2 times.
		if len(slice) < 3 ||
			(slice[1]|0x20 != 'l') ||
			(slice[2]|0x20 != 'l') {
			goto LoopIterationNoMatch1
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'D', 'd':
		pos += 2
		slice = r.Runtext[pos:]

	default:
		goto LoopIterationNoMatch1
	}

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration1 == 0 {
		goto LoopBody1
	}
	goto LoopEnd1

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch1:
	loop_iteration1--
	if loop_iteration1 < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto CharLoopBacktrack5
	}
	pos = r.StackPop()
	slice = r.Runtext[pos:]
LoopEnd1:
	r.Runstackpos = startingStackpos1 // Ensure any remaining backtracking state is removed.

	goto AlternationMatch

AlternationBranch1:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 2
	// Node: SetloopAtomic(Set = [\p{N}])(Min = 1, Max = 3)
	// Match [\p{N}] atomically at least 1 and at most 3 times.
	iteration4 = 0
	for iteration4 < 3 && iteration4 < len(slice) && unicode.In(slice[iteration4], unicode.N) {
		iteration4++
	}

	if iteration4 == 0 {
		goto AlternationBranch2
	}

	slice = slice[iteration4:]
	pos += iteration4

	goto AlternationMatch

AlternationBranch2:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 3
	// Node: Concatenate
	// Node: OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
	// Match ' ' atomically, optionally.
	if len(slice) > 0 && slice[0] == ' ' {
		slice = slice[1:]
		pos++
	}

	// Node: SetloopAtomic(Set = [^\s\p{L}\p{N}])(Min = 1, Max = inf)
	// Match [^\s\p{L}\p{N}] atomically at least once.
	iteration5 = 0
	for iteration5 < len(slice) && set_93b362dc942d41f83c0905ca6229c31c41f863d4ab4d2d9e17dae21b5b922113.CharIn(slice[iteration5]) {
		iteration5++
	}

	if iteration5 == 0 {
		goto AlternationBranch3
	}

	slice = slice[iteration5:]
	pos += iteration5

	// Node: SetloopAtomic(Set = [\n\r/])(Min = 0, Max = inf)
	// Match [\n\r/] atomically any number of times.
	iteration6 = helpers.IndexOfAnyExcept3(slice, '\n', '\r', '/')
	if iteration6 < 0 {
		iteration6 = len(slice)
	}

	slice = slice[iteration6:]
	pos += iteration6

	goto AlternationMatch

AlternationBranch3:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 4
	// Node: Concatenate
	// Node: Setloop(Set = [\s])(Min = 0, Max = inf)
	// Match [\s] greedily any number of times.
	charloop_starting_pos6 = pos

	iteration7 = 0
	for iteration7 < len(slice) && unicode.IsSpace(slice[iteration7]) {
		iteration7++
	}

	slice = slice[iteration7:]
	pos += iteration7

	charloop_ending_pos6 = pos
	goto CharLoopEnd6

CharLoopBacktrack6:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos6 >= charloop_ending_pos6 {
		goto AlternationBranch4
	}
	charloop_ending_pos6 = helpers.IndexOfAny2(r.Runtext[charloop_starting_pos6:charloop_ending_pos6], '\n', '\r')
	if charloop_ending_pos6 < 0 { // miss
		goto AlternationBranch4
	}
	charloop_ending_pos6 += charloop_starting_pos6
	pos = charloop_ending_pos6
	slice = r.Runtext[pos:]

CharLoopEnd6:

	// Node: SetloopAtomic(Set = [\n\r])(Min = 1, Max = inf)
	// Match [\n\r] atomically at least once.
	iteration8 = helpers.IndexOfAnyExcept2(slice, '\n', '\r')
	if iteration8 < 0 {
		iteration8 = len(slice)
	}

	if iteration8 == 0 {
		goto CharLoopBacktrack6
	}

	slice = slice[iteration8:]
	pos += iteration8

	goto AlternationMatch

AlternationBranch4:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 5
	// Node: Concatenate
	// Node: Setloop(Set = [\s])(Min = 1, Max = inf)
	// Match [\s] greedily at least once.
	charloop_starting_pos7 = pos

	iteration9 = 0
	for iteration9 < len(slice) && unicode.IsSpace(slice[iteration9]) {
		iteration9++
	}

	if iteration9 == 0 {
		goto AlternationBranch5
	}

	slice = slice[iteration9:]
	pos += iteration9

	charloop_ending_pos7 = pos
	charloop_starting_pos7++
	goto CharLoopEnd7

CharLoopBacktrack7:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos7 >= charloop_ending_pos7 {
		goto AlternationBranch5
	}
	charloop_ending_pos7--
	pos = charloop_ending_pos7
	slice = r.Runtext[pos:]

CharLoopEnd7:

	// Node: NegLook
	// Zero-width negative lookahead
	negativelookahead_starting_pos = pos

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	// Node: Set(Set = [^\s])
	// Match [^\s].
	if len(slice) == 0 || unicode.IsSpace(slice[0]) {
		goto NegativeLookaroundMatch
	}

	goto CharLoopBacktrack7

NegativeLookaroundMatch:
	pos = negativelookahead_starting_pos
	slice = r.Runtext[pos:]

	goto AlternationMatch

AlternationBranch5:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 6
	// Node: SetloopAtomic(Set = [\s])(Min = 1, Max = inf)
	// Match [\s] atomically at least once.
	iteration10 = 0
	for iteration10 < len(slice) && unicode.IsSpace(slice[iteration10]) {
		iteration10++
	}

	if iteration10 == 0 {
		return nil // The input didn't match.
	}

	slice = slice[iteration10:]
	pos += iteration10

AlternationMatch:
	;

	r.Runstackpos = atomic_stackpos

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture(index = 0, unindex = -1)
 Atomic
  Alternate
   Concatenate
    One(Ch = ')
    Atomic
     Alternate
      Set(Set = [R-Tr-tſ])
      Concatenate
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Vv])
       Set(Set = [Ee])
      Set(Set = [Mm])
      SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
      Set(Set = [Dd])
   Concatenate
    Setloop(Set = [^\n\r\p{L}\p{N}])(Min = 0, Max = 1)
    SetloopAtomic(Set = [\p{L}])(Min = 1, Max = inf)
   SetloopAtomic(Set = [\p{N}])(Min = 1, Max = 3)
   Concatenate
    OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
    SetloopAtomic(Set = [^\s\p{L}\p{N}])(Min = 1, Max = inf)
    SetloopAtomic(Set = [\n\r])(Min = 0, Max = inf)
   Concatenate
    Setloop(Set = [\s])(Min = 0, Max = inf)
    SetloopAtomic(Set = [\n\r])(Min = 1, Max = inf)
   Concatenate
    Setloop(Set = [\s])(Min = 1, Max = inf)
    NegLook
     Set(Set = [^\s])
   SetloopAtomic(Set = [\s])(Min = 1, Max = inf)
*/
// From encoding.go:153:29
// Pattern: "(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\\r\\n\\p{L}\\p{N}]?\\p{L}+|\\p{N}{1,3}| ?[^\\s\\p{L}\\p{N}]+[\\r\\n]*|\\s*[\\r\\n]+|\\s+(?!\\S)|\\s+"
// Options: regexp2.None
type reg_2_Engine struct{}

func (reg_2_Engine) Caps() map[int]int        { return nil }
func (reg_2_Engine) CapNames() map[string]int { return nil }
func (reg_2_Engine) CapsList() []string       { return nil }
func (reg_2_Engine) CapSize() int             { return 1 }

func (reg_2_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Empty matches aren't possible
	if pos < len(r.Runtext) {
		return true
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (reg_2_Engine) Execute(r *regexp2.Runner) error {
	atomic_stackpos := 0
	alternation_starting_pos := 0
	alternation_starting_pos1 := 0
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	iteration := 0
	iteration1 := 0
	iteration2 := 0
	iteration3 := 0
	var charloop_starting_pos1, charloop_ending_pos1 = 0, 0
	iteration4 := 0
	iteration5 := 0
	var charloop_starting_pos2, charloop_ending_pos2 = 0, 0
	iteration6 := 0
	negativelookahead_starting_pos := 0
	iteration7 := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Atomic
	// Atomic group.
	atomic_stackpos = r.Runstackpos

	// Node: Alternate
	// Match with 7 alternative expressions, atomically.
	alternation_starting_pos = pos

	// Branch 0
	// Node: Concatenate
	// Node: One(Ch = ')
	// Match '\”.
	if len(slice) == 0 || slice[0] != '\'' {
		goto AlternationBranch
	}

	// Node: Atomic
	// Node: Alternate
	// Match with 6 alternative expressions, atomically.
	alternation_starting_pos1 = pos

	// Branch 0
	// Node: Set(Set = [R-Tr-tſ])
	// Match [R-Tr-tſ].
	if len(slice) < 2 || !set_4a1357005dba18ced5af0bde8202a98fab7a3500c675e3cb1c60597d7a36436a.CharIn(slice[1]) {
		goto AlternationBranch1
	}

	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch1:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("re")) /* Match the string "re" (case-insensitive) */ {
		goto AlternationBranch2
	}

	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch2:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 2
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ve")) /* Match the string "ve" (case-insensitive) */ {
		goto AlternationBranch3
	}

	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch3:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 3
	// Node: Set(Set = [Mm])
	// Match [Mm].
	if len(slice) < 2 || (slice[1]|0x20 != 'm') {
		goto AlternationBranch4
	}

	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch4:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 4
	// Node: SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
	// Match [Ll] exactly 2 times.
	if len(slice) < 3 ||
		(slice[1]|0x20 != 'l') ||
		(slice[2]|0x20 != 'l') {
		goto AlternationBranch5
	}

	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch5:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 5
	// Node: Set(Set = [Dd])
	// Match [Dd].
	if len(slice) < 2 || (slice[1]|0x20 != 'd') {
		goto AlternationBranch
	}

	pos += 2
	slice = r.Runtext[pos:]

AlternationMatch1:
	;

	goto AlternationMatch

AlternationBranch:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Concatenate
	// Node: Setloop(Set = [^\n\r\p{L}\p{N}])(Min = 0, Max = 1)
	// Match [^\n\r\p{L}\p{N}] greedily, optionally.
	charloop_starting_pos = pos

	if len(slice) > 0 && set_4a7765fc40e8e8c7561121585d1545f9a51b44bf010a50f8b1b086a02ec1f07f.CharIn(slice[0]) {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos = pos
	goto CharLoopEnd

CharLoopBacktrack:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		goto AlternationBranch6
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:

	// Node: SetloopAtomic(Set = [\p{L}])(Min = 1, Max = inf)
	// Match [\p{L}] atomically at least once.
	iteration = 0
	for iteration < len(slice) && unicode.In(slice[iteration], unicode.L) {
		iteration++
	}

	if iteration == 0 {
		goto CharLoopBacktrack
	}

	slice = slice[iteration:]
	pos += iteration

	goto AlternationMatch

AlternationBranch6:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 2
	// Node: SetloopAtomic(Set = [\p{N}])(Min = 1, Max = 3)
	// Match [\p{N}] atomically at least 1 and at most 3 times.
	iteration1 = 0
	for iteration1 < 3 && iteration1 < len(slice) && unicode.In(slice[iteration1], unicode.N) {
		iteration1++
	}

	if iteration1 == 0 {
		goto AlternationBranch7
	}

	slice = slice[iteration1:]
	pos += iteration1

	goto AlternationMatch

AlternationBranch7:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 3
	// Node: Concatenate
	// Node: OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
	// Match ' ' atomically, optionally.
	if len(slice) > 0 && slice[0] == ' ' {
		slice = slice[1:]
		pos++
	}

	// Node: SetloopAtomic(Set = [^\s\p{L}\p{N}])(Min = 1, Max = inf)
	// Match [^\s\p{L}\p{N}] atomically at least once.
	iteration2 = 0
	for iteration2 < len(slice) && set_93b362dc942d41f83c0905ca6229c31c41f863d4ab4d2d9e17dae21b5b922113.CharIn(slice[iteration2]) {
		iteration2++
	}

	if iteration2 == 0 {
		goto AlternationBranch8
	}

	slice = slice[iteration2:]
	pos += iteration2

	// Node: SetloopAtomic(Set = [\n\r])(Min = 0, Max = inf)
	// Match [\n\r] atomically any number of times.
	iteration3 = helpers.IndexOfAnyExcept2(slice, '\n', '\r')
	if iteration3 < 0 {
		iteration3 = len(slice)
	}

	slice = slice[iteration3:]
	pos += iteration3

	goto AlternationMatch

AlternationBranch8:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 4
	// Node: Concatenate
	// Node: Setloop(Set = [\s])(Min = 0, Max = inf)
	// Match [\s] greedily any number of times.
	charloop_starting_pos1 = pos

	iteration4 = 0
	for iteration4 < len(slice) && unicode.IsSpace(slice[iteration4]) {
		iteration4++
	}

	slice = slice[iteration4:]
	pos += iteration4

	charloop_ending_pos1 = pos
	goto CharLoopEnd1

CharLoopBacktrack1:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos1 >= charloop_ending_pos1 {
		goto AlternationBranch9
	}
	charloop_ending_pos1 = helpers.IndexOfAny2(r.Runtext[charloop_starting_pos1:charloop_ending_pos1], '\n', '\r')
	if charloop_ending_pos1 < 0 { // miss
		goto AlternationBranch9
	}
	charloop_ending_pos1 += charloop_starting_pos1
	pos = charloop_ending_pos1
	slice = r.Runtext[pos:]

CharLoopEnd1:

	// Node: SetloopAtomic(Set = [\n\r])(Min = 1, Max = inf)
	// Match [\n\r] atomically at least once.
	iteration5 = helpers.IndexOfAnyExcept2(slice, '\n', '\r')
	if iteration5 < 0 {
		iteration5 = len(slice)
	}

	if iteration5 == 0 {
		goto CharLoopBacktrack1
	}

	slice = slice[iteration5:]
	pos += iteration5

	goto AlternationMatch

AlternationBranch9:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 5
	// Node: Concatenate
	// Node: Setloop(Set = [\s])(Min = 1, Max = inf)
	// Match [\s] greedily at least once.
	charloop_starting_pos2 = pos

	iteration6 = 0
	for iteration6 < len(slice) && unicode.IsSpace(slice[iteration6]) {
		iteration6++
	}

	if iteration6 == 0 {
		goto AlternationBranch10
	}

	slice = slice[iteration6:]
	pos += iteration6

	charloop_ending_pos2 = pos
	charloop_starting_pos2++
	goto CharLoopEnd2

CharLoopBacktrack2:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos2 >= charloop_ending_pos2 {
		goto AlternationBranch10
	}
	charloop_ending_pos2--
	pos = charloop_ending_pos2
	slice = r.Runtext[pos:]

CharLoopEnd2:

	// Node: NegLook
	// Zero-width negative lookahead
	negativelookahead_starting_pos = pos

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	// Node: Set(Set = [^\s])
	// Match [^\s].
	if len(slice) == 0 || unicode.IsSpace(slice[0]) {
		goto NegativeLookaroundMatch
	}

	goto CharLoopBacktrack2

NegativeLookaroundMatch:
	pos = negativelookahead_starting_pos
	slice = r.Runtext[pos:]

	goto AlternationMatch

AlternationBranch10:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 6
	// Node: SetloopAtomic(Set = [\s])(Min = 1, Max = inf)
	// Match [\s] atomically at least once.
	iteration7 = 0
	for iteration7 < len(slice) && unicode.IsSpace(slice[iteration7]) {
		iteration7++
	}

	if iteration7 == 0 {
		return nil // The input didn't match.
	}

	slice = slice[iteration7:]
	pos += iteration7

AlternationMatch:
	;

	r.Runstackpos = atomic_stackpos

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture(index = 0, unindex = -1)
 Atomic
  Alternate
   Concatenate
    One(Ch = ')
    Atomic
     Alternate
      Set(Set = [r-t])
      Multi(String = "re")
      Multi(String = "ve")
      One(Ch = m)
      Multi(String = "ll")
      One(Ch = d)
   Concatenate
    OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
    SetloopAtomic(Set = [\p{L}])(Min = 1, Max = inf)
   Concatenate
    OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
    SetloopAtomic(Set = [\p{N}])(Min = 1, Max = inf)
   Concatenate
    OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
    SetloopAtomic(Set = [^\s\p{L}\p{N}])(Min = 1, Max = inf)
   Concatenate
    Setloop(Set = [\s])(Min = 1, Max = inf)
    NegLook
     Set(Set = [^\s])
   SetloopAtomic(Set = [\s])(Min = 1, Max = inf)
*/
// From encoding.go:168:29
// Pattern: "'s|'t|'re|'ve|'m|'ll|'d| ?\\p{L}+| ?\\p{N}+| ?[^\\s\\p{L}\\p{N}]+|\\s+(?!\\S)|\\s+"
// Options: regexp2.None
type reg_3_Engine struct{}

func (reg_3_Engine) Caps() map[int]int        { return nil }
func (reg_3_Engine) CapNames() map[string]int { return nil }
func (reg_3_Engine) CapsList() []string       { return nil }
func (reg_3_Engine) CapSize() int             { return 1 }

func (reg_3_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Empty matches aren't possible
	if pos < len(r.Runtext) {
		return true
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (reg_3_Engine) Execute(r *regexp2.Runner) error {
	atomic_stackpos := 0
	alternation_starting_pos := 0
	alternation_starting_pos1 := 0
	iteration := 0
	iteration1 := 0
	iteration2 := 0
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	iteration3 := 0
	negativelookahead_starting_pos := 0
	iteration4 := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Atomic
	// Atomic group.
	atomic_stackpos = r.Runstackpos

	// Node: Alternate
	// Match with 6 alternative expressions, atomically.
	alternation_starting_pos = pos

	// Branch 0
	// Node: Concatenate
	// Node: One(Ch = ')
	// Match '\”.
	if len(slice) == 0 || slice[0] != '\'' {
		goto AlternationBranch
	}

	// Node: Atomic
	// Node: Alternate
	// Match with 6 alternative expressions, atomically.
	alternation_starting_pos1 = pos

	// Branch 0
	// Node: Set(Set = [r-t])
	// Match [r-t].
	if len(slice) < 2 || !helpers.IsBetween(slice[1], 'r', 't') {
		goto AlternationBranch1
	}

	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch1:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Multi(String = "re")
	// Match the string "re".
	if !helpers.StartsWith(slice[1:], []rune("re")) {
		goto AlternationBranch2
	}

	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch2:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 2
	// Node: Multi(String = "ve")
	// Match the string "ve".
	if !helpers.StartsWith(slice[1:], []rune("ve")) {
		goto AlternationBranch3
	}

	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch3:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 3
	// Node: One(Ch = m)
	// Match 'm'.
	if len(slice) < 2 || slice[1] != 'm' {
		goto AlternationBranch4
	}

	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch4:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 4
	// Node: Multi(String = "ll")
	// Match the string "ll".
	if !helpers.StartsWith(slice[1:], []rune("ll")) {
		goto AlternationBranch5
	}

	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch5:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 5
	// Node: One(Ch = d)
	// Match 'd'.
	if len(slice) < 2 || slice[1] != 'd' {
		goto AlternationBranch
	}

	pos += 2
	slice = r.Runtext[pos:]

AlternationMatch1:
	;

	goto AlternationMatch

AlternationBranch:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Concatenate
	// Node: OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
	// Match ' ' atomically, optionally.
	if len(slice) > 0 && slice[0] == ' ' {
		slice = slice[1:]
		pos++
	}

	// Node: SetloopAtomic(Set = [\p{L}])(Min = 1, Max = inf)
	// Match [\p{L}] atomically at least once.
	iteration = 0
	for iteration < len(slice) && unicode.In(slice[iteration], unicode.L) {
		iteration++
	}

	if iteration == 0 {
		goto AlternationBranch6
	}

	slice = slice[iteration:]
	pos += iteration

	goto AlternationMatch

AlternationBranch6:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 2
	// Node: Concatenate
	// Node: OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
	// Match ' ' atomically, optionally.
	if len(slice) > 0 && slice[0] == ' ' {
		slice = slice[1:]
		pos++
	}

	// Node: SetloopAtomic(Set = [\p{N}])(Min = 1, Max = inf)
	// Match [\p{N}] atomically at least once.
	iteration1 = 0
	for iteration1 < len(slice) && unicode.In(slice[iteration1], unicode.N) {
		iteration1++
	}

	if iteration1 == 0 {
		goto AlternationBranch7
	}

	slice = slice[iteration1:]
	pos += iteration1

	goto AlternationMatch

AlternationBranch7:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 3
	// Node: Concatenate
	// Node: OneloopAtomic(Ch = \ )(Min = 0, Max = 1)
	// Match ' ' atomically, optionally.
	if len(slice) > 0 && slice[0] == ' ' {
		slice = slice[1:]
		pos++
	}

	// Node: SetloopAtomic(Set = [^\s\p{L}\p{N}])(Min = 1, Max = inf)
	// Match [^\s\p{L}\p{N}] atomically at least once.
	iteration2 = 0
	for iteration2 < len(slice) && set_93b362dc942d41f83c0905ca6229c31c41f863d4ab4d2d9e17dae21b5b922113.CharIn(slice[iteration2]) {
		iteration2++
	}

	if iteration2 == 0 {
		goto AlternationBranch8
	}

	slice = slice[iteration2:]
	pos += iteration2

	goto AlternationMatch

AlternationBranch8:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 4
	// Node: Concatenate
	// Node: Setloop(Set = [\s])(Min = 1, Max = inf)
	// Match [\s] greedily at least once.
	charloop_starting_pos = pos

	iteration3 = 0
	for iteration3 < len(slice) && unicode.IsSpace(slice[iteration3]) {
		iteration3++
	}

	if iteration3 == 0 {
		goto AlternationBranch9
	}

	slice = slice[iteration3:]
	pos += iteration3

	charloop_ending_pos = pos
	charloop_starting_pos++
	goto CharLoopEnd

CharLoopBacktrack:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		goto AlternationBranch9
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:

	// Node: NegLook
	// Zero-width negative lookahead
	negativelookahead_starting_pos = pos

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	// Node: Set(Set = [^\s])
	// Match [^\s].
	if len(slice) == 0 || unicode.IsSpace(slice[0]) {
		goto NegativeLookaroundMatch
	}

	goto CharLoopBacktrack

NegativeLookaroundMatch:
	pos = negativelookahead_starting_pos
	slice = r.Runtext[pos:]

	goto AlternationMatch

AlternationBranch9:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 5
	// Node: SetloopAtomic(Set = [\s])(Min = 1, Max = inf)
	// Match [\s] atomically at least once.
	iteration4 = 0
	for iteration4 < len(slice) && unicode.IsSpace(slice[iteration4]) {
		iteration4++
	}

	if iteration4 == 0 {
		return nil // The input didn't match.
	}

	slice = slice[iteration4:]
	pos += iteration4

AlternationMatch:
	;

	r.Runstackpos = atomic_stackpos

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

// The set [^\n\r\p{L}\p{N}]
var set_4a7765fc40e8e8c7561121585d1545f9a51b44bf010a50f8b1b086a02ec1f07f = syntax.NewCharSetRuntime("\x01\x02\x00\x00\x00\x02\x00\x00\x00\n\n\r\r\x01L\x01N")

// The set [^\s\p{L}\p{N}]
var set_93b362dc942d41f83c0905ca6229c31c41f863d4ab4d2d9e17dae21b5b922113 = syntax.NewCharSetRuntime("\x01\x00\x00\x00\x00\x03\x00\x00\x00\x01 \x01L\x01N")

// The set [R-Tr-tſ]
var set_4a1357005dba18ced5af0bde8202a98fab7a3500c675e3cb1c60597d7a36436a = syntax.NewCharSetRuntime("\x00\x03\x00\x00\x00\x00\x00\x00\x00RTrtſſ")

func init() {
	regexp2.RegisterEngine("[^\\r\\n\\p{L}\\p{N}]?[\\p{Lu}\\p{Lt}\\p{Lm}\\p{Lo}\\p{M}]*[\\p{Ll}\\p{Lm}\\p{Lo}\\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?|[^\\r\\n\\p{L}\\p{N}]?[\\p{Lu}\\p{Lt}\\p{Lm}\\p{Lo}\\p{M}]+[\\p{Ll}\\p{Lm}\\p{Lo}\\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?|\\p{N}{1,3}| ?[^\\s\\p{L}\\p{N}]+[\\r\\n/]*|\\s*[\\r\\n]+|\\s+(?!\\S)|\\s+", regexp2.None, &reg_Engine{})
	regexp2.RegisterEngine("(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\\r\\n\\p{L}\\p{N}]?\\p{L}+|\\p{N}{1,3}| ?[^\\s\\p{L}\\p{N}]+[\\r\\n]*|\\s*[\\r\\n]+|\\s+(?!\\S)|\\s+", regexp2.None, &reg_2_Engine{})
	regexp2.RegisterEngine("'s|'t|'re|'ve|'m|'ll|'d| ?\\p{L}+| ?\\p{N}+| ?[^\\s\\p{L}\\p{N}]+|\\s+(?!\\S)|\\s+", regexp2.None, &reg_3_Engine{})
	var _ = helpers.Min
	var _ = syntax.NewCharSetRuntime
	var _ = unicode.IsDigit
}

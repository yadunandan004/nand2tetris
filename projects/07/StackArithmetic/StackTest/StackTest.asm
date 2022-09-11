// ---- Push Start ---- ## push constant 17
@17
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 17
@17
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Eq Start ---- ## eq
@0
A=M-1
D=M
A=A-1
M=D-M
D=A
@0
M=D+1
A=D
D=M
@negs0
D;JNE
@0
A=M-1
M=-1
D=0
(negs0)
@nege0
D;JEQ
@0
A=M-1
M=0
(nege0)
// ---- Eq End ----
// ---- Push Start ---- ## push constant 17
@17
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 16
@16
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Eq Start ---- ## eq
@0
A=M-1
D=M
A=A-1
M=D-M
D=A
@0
M=D+1
A=D
D=M
@negs1
D;JNE
@0
A=M-1
M=-1
D=0
(negs1)
@nege1
D;JEQ
@0
A=M-1
M=0
(nege1)
// ---- Eq End ----
// ---- Push Start ---- ## push constant 16
@16
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 17
@17
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Eq Start ---- ## eq
@0
A=M-1
D=M
A=A-1
M=D-M
D=A
@0
M=D+1
A=D
D=M
@negs2
D;JNE
@0
A=M-1
M=-1
D=0
(negs2)
@nege2
D;JEQ
@0
A=M-1
M=0
(nege2)
// ---- Eq End ----
// ---- Push Start ---- ## push constant 892
@892
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 891
@891
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Lt Start ---- ## lt
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
A=D
D=M
@jlts0
D;JGE
@0
A=M-1
M=-1
D=0
(jlts0)
@jlte0
D;JEQ
@0
A=M-1
M=0
(jlte0)
// ---- Lt End ----
// ---- Push Start ---- ## push constant 891
@891
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 892
@892
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Lt Start ---- ## lt
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
A=D
D=M
@jlts1
D;JGE
@0
A=M-1
M=-1
D=0
(jlts1)
@jlte1
D;JEQ
@0
A=M-1
M=0
(jlte1)
// ---- Lt End ----
// ---- Push Start ---- ## push constant 891
@891
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 891
@891
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Lt Start ---- ## lt
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
A=D
D=M
@jlts2
D;JGE
@0
A=M-1
M=-1
D=0
(jlts2)
@jlte2
D;JEQ
@0
A=M-1
M=0
(jlte2)
// ---- Lt End ----
// ---- Push Start ---- ## push constant 32767
@32767
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 32766
@32766
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Gt Start ---- ## gt
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
A=D
D=M
@jgts0
D;JLE
@0
A=M-1
M=-1
D=0
(jgts0)
@jgte0
D;JEQ
@0
A=M-1
M=0
(jgte0)
// ---- Gt End ----
// ---- Push Start ---- ## push constant 32766
@32766
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 32767
@32767
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Gt Start ---- ## gt
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
A=D
D=M
@jgts1
D;JLE
@0
A=M-1
M=-1
D=0
(jgts1)
@jgte1
D;JEQ
@0
A=M-1
M=0
(jgte1)
// ---- Gt End ----
// ---- Push Start ---- ## push constant 32766
@32766
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 32766
@32766
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Gt Start ---- ## gt
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
A=D
D=M
@jgts2
D;JLE
@0
A=M-1
M=-1
D=0
(jgts2)
@jgte2
D;JEQ
@0
A=M-1
M=0
(jgte2)
// ---- Gt End ----
// ---- Push Start ---- ## push constant 57
@57
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 31
@31
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 53
@53
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Add Start ---- ## add
@0
A=M-1
D=M
A=A-1
M=M+D
D=A
@0
M=D+1
// ---- Add End ----
// ---- Push Start ---- ## push constant 112
@112
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Sub Start ---- ## sub
@0
A=M-1
D=M
A=A-1
M=M-D
D=A
@0
M=D+1
// ---- Sub End ----
// ---- Neg Start ---- ## neg
@0
A=M-1
M=-M
// ---- Neg End ----
// ---- And Start ---- ## and
@0
A=M-1
D=M
A=A-1
M=D&M
D=A
@0
M=D+1
// ---- And End ----
// ---- Push Start ---- ## push constant 82
@82
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Or Start ---- ## or
@0
A=M-1
D=M
A=A-1
M=D|M
D=A
@0
M=D+1
// ---- Or End ----
// ---- Not Start ---- ## not
@0
A=M-1
M=!M
// ---- Not End ----

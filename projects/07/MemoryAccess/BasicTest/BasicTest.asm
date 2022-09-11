// ---- Push Start ---- ## push constant 10
@10
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop local 0
@0
D=A
@1
A=M
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push constant 21
@21
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 22
@22
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop argument 2
@2
D=A
@2
A=M
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Pop Start ---- ## pop argument 1
@1
D=A
@2
A=M
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push constant 36
@36
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop this 6
@6
D=A
@3
A=M
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push constant 42
@42
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 45
@45
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop that 5
@5
D=A
@4
A=M
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Pop Start ---- ## pop that 2
@2
D=A
@4
A=M
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push constant 510
@510
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop temp 6
@6
D=A
@5
D=A+D
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push local 0
@0
D=A
@1
A=M
A=A+D
D=M
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push that 5
@5
D=A
@4
A=M
A=A+D
D=M
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
// ---- Push Start ---- ## push argument 1
@1
D=A
@2
A=M
A=A+D
D=M
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
// ---- Push Start ---- ## push this 6
@6
D=A
@3
A=M
A=A+D
D=M
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push this 6
@6
D=A
@3
A=M
A=A+D
D=M
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
// ---- Push Start ---- ## push temp 6
@6
D=A
@5
A=A+D
D=M
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

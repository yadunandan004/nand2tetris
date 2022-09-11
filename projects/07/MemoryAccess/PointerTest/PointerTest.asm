// ---- Push Start ---- ## push constant 3030
@3030
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop pointer 0
@0
D=A
@50
D=A
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@3
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push constant 3040
@3040
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop pointer 1
@1
D=A
@51
D=A
@13
M=D
@0
A=M-1
D=M
@13
A=M
M=D
@4
M=D
@0
M=M-1
// ---- Pop End ----
// ---- Push Start ---- ## push constant 32
@32
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop this 2
@2
D=A
@3
A=M+D
D=A
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
// ---- Push Start ---- ## push constant 46
@46
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop that 6
@6
D=A
@4
A=M+D
D=A
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
// ---- Push Start ---- ## push pointer 0
@0
D=A
@50
D=M
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push pointer 1
@1
D=A
@51
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
// ---- Push Start ---- ## push this 2
@2
D=A
@3
A=M+D
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
// ---- Push Start ---- ## push that 6
@6
D=A
@4
A=M+D
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

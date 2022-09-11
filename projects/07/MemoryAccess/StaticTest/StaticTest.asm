// ---- Push Start ---- ## push constant 111
@111
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 333
@333
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push constant 888
@888
D=A
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Pop Start ---- ## pop static 8
@8
D=A
@30
A=M
A=A+D
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
// ---- Pop Start ---- ## pop static 3
@3
D=A
@30
A=M
A=A+D
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
// ---- Pop Start ---- ## pop static 1
@1
D=A
@30
A=M
A=A+D
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
// ---- Push Start ---- ## push static 3
@3
D=A
@30
A=M
A=A+D
D=M
@0
A=M
M=D
@0
M=M+1
// ---- Push End ----
// ---- Push Start ---- ## push static 1
@1
D=A
@30
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
// ---- Push Start ---- ## push static 8
@8
D=A
@30
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

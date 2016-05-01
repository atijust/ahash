// +build amd64

#include "textflag.h"

TEXT ·Distance(SB), NOSPLIT, $0-24
	MOVQ	hash1+0(FP), BX
	MOVQ	hash2+8(FP), BP
	XORQ    BP, BX
	BYTE $0xf3; BYTE $0x48; BYTE $0x0f; BYTE $0xb8; BYTE $0xdb // POPCNTQ BX, BX
	MOVQ	BX, ret+16(FP)
	RET

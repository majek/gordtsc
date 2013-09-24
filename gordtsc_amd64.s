// Don't insert stack check preamble.
#define NOSPLIT 4

#define RDTSCP BYTE $0x0f; \
                BYTE $0x01; \
	        BYTE $0xF9


TEXT ·Start(SB),NOSPLIT,$0
        CPUID
        RDTSC
        MOVL DX, ret+4(FP)
        MOVL AX, ret+0(FP)
        RET

TEXT ·Stop(SB),NOSPLIT,$0
        RDTSCP
        MOVL DX, ret+4(FP)
        MOVL AX, ret+0(FP)
        CPUID
        RET

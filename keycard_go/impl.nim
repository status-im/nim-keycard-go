proc Start*(params: cstring): cstring {.importc: "Start".}
proc Stop*(params: cstring): cstring {.importc: "Stop".}
proc Pair*(params: cstring): cstring {.importc: "Pair".}
proc OpenSecureChannel*(params: cstring): cstring {.importc: "OpenSecureChannel".}
proc VerifyPin*(params: cstring): cstring {.importc: "VerifyPin".}
proc GenerateKey*(params: cstring): cstring {.importc: "GenerateKey".}

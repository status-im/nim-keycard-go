proc Start*(): cstring {.importc: "Start".}
proc Stop*(): cstring {.importc: "Stop".}
proc Select*(): cstring {.importc: "Select".}
proc Pair*(params: cstring): cstring {.importc: "Pair".}
proc OpenSecureChannel*(params: cstring): cstring {.importc: "OpenSecureChannel".}
proc VerifyPin*(params: cstring): cstring {.importc: "VerifyPin".}
proc GenerateKey*(params: cstring): cstring {.importc: "GenerateKey".}

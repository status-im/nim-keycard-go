proc Start*(): cstring {.importc: "Start".}
proc Stop*(): cstring {.importc: "Stop".}
proc Select*(): cstring {.importc: "Select".}
proc Pair*(params: cstring): cstring {.importc: "Pair".}
proc OpenSecureChannel*(params: cstring): cstring {.importc: "OpenSecureChannel".}
proc VerifyPin*(params: cstring): cstring {.importc: "VerifyPin".}
proc GenerateKey*(params: cstring): cstring {.importc: "GenerateKey".}
proc DeriveKey*(params: cstring): cstring {.importc: "DeriveKey".}
proc SignWithPath*(params: cstring): cstring {.importc: "SignWithPath".}
proc ExportKey*(params: cstring): cstring {.importc: "ExportKey".}
proc LoadSeed*(params: cstring): cstring {.importc: "LoadSeed".}
proc Init*(params: cstring): cstring {.importc: "Init".}
proc Unpair*(params: cstring): cstring {.importc: "Unpair".}
proc GetStatusApplication*(): cstring {.importc: "GetStatusApplication".}
proc ChangePin*(): cstring {.importc: "ChangePin".}
proc ChangePuk*(): cstring {.importc: "ChangePuk".}
proc ChangePairingPassword*(): cstring {.importc: "ChangePairingPassword".}

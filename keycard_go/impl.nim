type KeycardSignalCallback* = proc(signal: cstring): void {.cdecl.}

proc start*(): cstring {.importc: "Start".}
proc stop*(): cstring {.importc: "Stop".}
proc select*(): cstring {.importc: "Select".}
proc pair*(params: cstring): cstring {.importc: "Pair".}
proc openSecureChannel*(params: cstring): cstring {.importc: "OpenSecureChannel".}
proc verifyPin*(params: cstring): cstring {.importc: "VerifyPin".}
proc generateKey*(): cstring {.importc: "GenerateKey".}
proc deriveKey*(params: cstring): cstring {.importc: "DeriveKey".}
proc signWithPath*(params: cstring): cstring {.importc: "SignWithPath".}
proc exportKey*(params: cstring): cstring {.importc: "ExportKey".}
proc loadSeed*(params: cstring): cstring {.importc: "LoadSeed".}
proc init*(params: cstring): cstring {.importc: "Init".}
proc unpair*(params: cstring): cstring {.importc: "Unpair".}
proc getStatusApplication*(): cstring {.importc: "GetStatusApplication".}
proc changePin*(): cstring {.importc: "ChangePin".}
proc changePuk*(): cstring {.importc: "ChangePuk".}
proc changePairingPassword*(): cstring {.importc: "ChangePairingPassword".}
proc free*(param: pointer) {.importc: "Free".}
proc setSignalEventCallback*(callback: KeycardSignalCallback) {.importc: "SetSignalEventCallback".}


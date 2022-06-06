type KeycardSignalCallback* = proc(signal: cstring): void {.cdecl.}

proc free*(param: pointer) {.importc: "Free".}
proc setSignalEventCallback*(callback: KeycardSignalCallback) {.importc: "KeycardSetSignalEventCallback".}

proc keycardInitFlow*(storageDir: cstring): cstring {.importc: "KeycardInitFlow".}
proc keycardStartFlow*(flowType: cint, jsonParams: cstring): cstring {.importc: "KeycardStartFlow".}
proc keycardResumeFlow*(jsonParams: cstring): cstring {.importc: "KeycardResumeFlow".}
proc keycardCancelFlow*(): cstring {.importc: "KeycardCancelFlow".}
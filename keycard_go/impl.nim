type KeycardSignalCallback* = proc(signal: cstring): void {.cdecl.}

proc free*(param: pointer) {.importc: "Free".}
proc setSignalEventCallback*(callback: KeycardSignalCallback) {.importc: "KeycardSetSignalEventCallback".}

proc keycardInitFlow*(storageDir: cstring): cstring {.importc: "KeycardInitFlow".}
proc keycardStartFlow*(flowType: cint, jsonParams: cstring): cstring {.importc: "KeycardStartFlow".}
proc keycardResumeFlow*(jsonParams: cstring): cstring {.importc: "KeycardResumeFlow".}
proc keycardCancelFlow*(): cstring {.importc: "KeycardCancelFlow".}

# availale in test mode only
proc mockedLibRegisterKeycard*(cardIndex: cint, readerState: cint, keycardState: cint, mockedKeycard: cstring, mockedKeycardHelper: cstring): cstring {.importc: "MockedLibRegisterKeycard".}
proc mockedLibReaderPluggedIn*(): cstring {.importc: "MockedLibReaderPluggedIn".}
proc mockedLibReaderUnplugged*(): cstring {.importc: "MockedLibReaderUnplugged".}
proc mockedLibKeycardInserted*(cardIndex: cint): cstring {.importc: "MockedLibKeycardInserted".}
proc mockedLibKeycardRemoved*(): cstring {.importc: "MockedLibKeycardRemoved".}
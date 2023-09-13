import ./keycard_go/impl as go_shim

export KeycardSignalCallback

proc keycardInitFlow*(storageDir: string): string =
  var funcOut = go_shim.keycardInitFlow(storageDir.cstring)
  defer: go_shim.free(funcOut)
  return $funcOut

proc keycardStartFlow*(flowType: int, jsonParams: string): string =
  var funcOut = go_shim.keycardStartFlow(flowType.cint, jsonParams.cstring)
  defer: go_shim.free(funcOut)
  return $funcOut

proc keycardResumeFlow*(jsonParams: string): string =
  var funcOut = go_shim.keycardResumeFlow(jsonParams.cstring)
  defer: go_shim.free(funcOut)
  return $funcOut

proc keycardCancelFlow*(): string =
  var funcOut = go_shim.keycardCancelFlow()
  defer: go_shim.free(funcOut)
  return $funcOut

proc setSignalEventCallback*(callback: KeycardSignalCallback) =
  go_shim.setSignalEventCallback(callback)

# availale in test mode only
proc mockedLibRegisterKeycard*(cardIndex: int, readerState: int, keycardState: int, mockedKeycard: string, mockedKeycardHelper: string): string =
  var funcOut = go_shim.mockedLibRegisterKeycard(cardIndex.cint, readerState.cint, keycardState.cint, mockedKeycard.cstring, mockedKeycardHelper.cstring)
  defer: go_shim.free(funcOut)
  return $funcOut

proc mockedLibReaderPluggedIn*(): string =
  var funcOut = go_shim.mockedLibReaderPluggedIn()
  defer: go_shim.free(funcOut)
  return $funcOut

proc mockedLibReaderUnplugged*(): string =
  var funcOut = go_shim.mockedLibReaderUnplugged()
  defer: go_shim.free(funcOut)
  return $funcOut

proc mockedLibKeycardInserted*(cardIndex: int): string =
  var funcOut = go_shim.mockedLibKeycardInserted(cardIndex.cint)
  defer: go_shim.free(funcOut)
  return $funcOut

proc mockedLibKeycardRemoved*(): string =
  var funcOut = go_shim.mockedLibKeycardRemoved()
  defer: go_shim.free(funcOut)
  return $funcOut
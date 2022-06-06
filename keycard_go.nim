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
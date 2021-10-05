import ./keycard_go/impl as go_shim

export KeycardSignalCallback

proc start*(): string =
  var funcOut = go_shim.start()
  defer: go_shim.free(funcOut)
  return $funcOut

proc stop*(): string =
  var funcOut = go_shim.stop()
  defer: go_shim.free(funcOut)
  return $funcOut

proc select*(): string =
  var funcOut = go_shim.select()
  defer: go_shim.free(funcOut)
  return $funcOut

proc pair*(params: cstring): string =
  var funcOut = go_shim.pair(params)
  defer: go_shim.free(funcOut)
  return $funcOut

proc openSecureChannel*(params: cstring): string =
  var funcOut = go_shim.openSecureChannel(params)
  defer: go_shim.free(funcOut)
  return $funcOut

proc verifyPin*(params: cstring): string =
  var funcOut = go_shim.verifyPin(params)
  defer: go_shim.free(funcOut)
  return $funcOut

proc exportKey*(params: cstring): string =
  var funcOut = go_shim.exportKey(params)
  defer: go_shim.free(funcOut)
  return $funcOut

proc getStatusApplication*(): string =
  var funcOut = go_shim.getStatusApplication()
  defer: go_shim.free(funcOut)
  return $funcOut

proc setSignalEventCallback*(callback: KeycardSignalCallback) =
  go_shim.setSignalEventCallback(callback)

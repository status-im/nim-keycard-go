import ./keycard_go/impl as go_shim

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


package main

// #cgo LDFLAGS: -shared
import "C"
import (
	"encoding/json"
	"fmt"
	"time"
	"unsafe"

	"github.com/status-im/nim-keycard-go/go/keycard/signal"
)

var kctx *keycardContext

func main() {
	// example()
}

func example() {
	signal.SetDefaultNodeNotificationHandler(func(jsonEvent string) {
		fmt.Printf("SIGNAL %+v\n", jsonEvent)
	})

	fmt.Printf("RUNNING EXAMPLE \n")
	res := Start()
	fmt.Printf("*** start %+v\n", C.GoString(res))
	time.Sleep(2)
	res = Select()
	fmt.Printf("*** select %+v\n", C.GoString(res))

	// res = Init(C.CString(`{"pin": "123456", "puk": "123456789012", "pairingPassword": "KeycardTest"}`))
	// fmt.Printf("*** OpenSecureChannel %+v\n", C.GoString(res))

	// res = Pair(C.CString(`{"pairingPassword": "KeycardTest"}`))
	// fmt.Printf("*** Pair %+v\n", C.GoString(res))

	res = OpenSecureChannel(C.CString(`{"index":1, "key": "33b0b458d19df44b009ea8142b64e041837667355250d13f3b84f389f6350cc8"}`))
	fmt.Printf("*** OpenSecureChannel %+v\n", C.GoString(res))

	res = VerifyPin(C.CString(`{"pin": "123456"}`))
	fmt.Printf("*** VerifyPin %+v\n", C.GoString(res))

	res = ChangePin(C.CString(`{"pin": "123456"}`))
	fmt.Printf("*** ChangePin %+v\n", C.GoString(res))

	res = ChangePuk(C.CString(`{"puk": "123456789012"}`))
	fmt.Printf("*** ChangePuk %+v\n", C.GoString(res))

	res = ChangePairingPassword(C.CString(`{"pairingPassword": "KeycardTest"}`))
	fmt.Printf("*** ChangePairingPassword %+v\n", C.GoString(res))

	res = GetStatusApplication()
	fmt.Printf("*** GetStatusApplication %+v\n", C.GoString(res))

	// res = Unpair(C.CString(`{"index": 1}`))
	// fmt.Printf("*** Unpair %+v\n", C.GoString(res))

	// res = GenerateKey()
	// fmt.Printf("*** GenerateKey %+v\n", C.GoString(res))

	res = DeriveKey(C.CString(`{"path":"m/1/2/3/4/5"}`))
	fmt.Printf("*** DeriveKey %+v\n", C.GoString(res))

	res = SignWithPath(C.CString(`{"data": "0000000000000000000000000000000000000000000000000000000000000000", "path":"m/1/2/3/4/5"}`))
	fmt.Printf("*** SignWithPath %+v\n", C.GoString(res))

	res = ExportKey(C.CString(`{"derive": true, "makeCurrent": false, "onlyPublic": false, "path": "m/43'/60'/1581'/0'/0"}`))
	fmt.Printf("*** ExportKey %+v\n", C.GoString(res))

	res = LoadSeed(C.CString(`{"seed": "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}`))
	fmt.Printf("*** LoadSeed %+v\n", C.GoString(res))

	res = Stop()
	fmt.Printf("*** stop %+v\n", C.GoString(res))
	time.Sleep(10 * time.Second)
}

//export Start
func Start() *C.char {
	var err error
	kctx, err = startKeycardContext()
	if err != nil {
		return retValue("err", err.Error())
	}
	return retValue("ok", true)
}

//export Select
func Select() *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	info, err := kctx.selectApplet()
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "applicationInfo", ApplicationInfo{
		Installed:              info.Installed,
		Initialized:            info.Initialized,
		InstanceUID:            info.InstanceUID,
		SecureChannelPublicKey: info.SecureChannelPublicKey,
		Version:                info.Version,
		AvailableSlots:         info.AvailableSlots,
		KeyUID:                 info.KeyUID,
		Capabilities:           Capability(info.Capabilities),
	})
}

//export Stop
func Stop() *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	if err := kctx.stop(); err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export Pair
func Pair(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params pairParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	pairingInfo, err := kctx.pair(params.PairingPassword)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "pairingInfo", PairingInfo{
		Key:   pairingInfo.Key,
		Index: pairingInfo.Index,
	})
}

//export OpenSecureChannel
func OpenSecureChannel(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params openSecureChannelParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.openSecureChannel(params.Index, params.Key)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export VerifyPin
func VerifyPin(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params verifyPinParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.verifyPin(params.Pin)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export GenerateKey
func GenerateKey() *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	keyUID, err := kctx.generateKey()
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "keyUID", keyUID)
}

//export DeriveKey
func DeriveKey(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params deriveKeyParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.deriveKey(params.Path)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export SignWithPath
func SignWithPath(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params signWithPathParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	sig, err := kctx.signWithPath(params.Data, params.Path)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "signature", Signature{
		PublicKey: hexString(sig.PubKey()),
		R:         hexString(sig.R()),
		S:         hexString(sig.S()),
		V:         sig.V(),
	})
}

//export ExportKey
func ExportKey(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params exportKeyParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	key, err := kctx.exportKey(params.Derive, params.MakeCurrent, params.OnlyPublic, params.Path)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "key", hexString(key))
}

//export LoadSeed
func LoadSeed(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params loadSeedParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	pubKey, err := kctx.loadSeed(params.Seed)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "publicKey", hexString(pubKey))
}

//export Init
func Init(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params initSeedParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.init(params.Pin, params.Puk, params.PairingPassword)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export Unpair
func Unpair(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params unpairParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.unpair(params.Index)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export GetStatusApplication
func GetStatusApplication() *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	status, err := kctx.getStatusApplication()
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true, "status", ApplicationStatus{
		PinRetryCount:  status.PinRetryCount,
		PUKRetryCount:  status.PUKRetryCount,
		KeyInitialized: status.KeyInitialized,
		Path:           status.Path,
	})
}

//export ChangePin
func ChangePin(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params changeSecretsParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.changePin(params.Pin)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export ChangePuk
func ChangePuk(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params changeSecretsParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.changePuk(params.Puk)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export ChangePairingPassword
func ChangePairingPassword(jsonParams *C.char) *C.char {
	if kctx == nil {
		l("select: not started")
		return retValue("error", "not started")
	}

	var params changeSecretsParams
	if err := json.Unmarshal([]byte(C.GoString(jsonParams)), &params); err != nil {
		return retValue("error", err.Error())
	}

	err := kctx.changePairingPassword(params.PairingPassword)
	if err != nil {
		return retValue("error", err.Error())
	}

	return retValue("ok", true)
}

//export SetSignalEventCallback
func SetSignalEventCallback(cb unsafe.Pointer) {
	signal.SetSignalEventCallback(cb)
}

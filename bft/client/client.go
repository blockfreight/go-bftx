package bftcli

import (
	"fmt"
	"sync"	//https://golang.org/pkg/sync/#Mutex

	. "github.com/tendermint/go-common"
	"github.com/tendermint/abci/types"
)

type Client interface {
	//Separates the definition of objects of the implementation
	//Declares equal than Structures
	//Helps to call different Structure's functions that share each other the same Inferface
	Service	//?	go-common/service.go
		//Start() (bool, error)
		//OnStart() error
		//Stop() bool
		//OnStop()
		//Reset() (bool, error)
		//OnReset() error
		//IsRunning() bool
		//String() string

	/*Temp
	SetResponseCallback(Callback)
	Error() error

	FlushAsync() *ReqRes
	EchoAsync(msg string) *ReqRes
	InfoAsync() *ReqRes
	SetOptionAsync(key string, value string) *ReqRes
	DeliverTxAsync(tx []byte) *ReqRes
	CheckTxAsync(tx []byte) *ReqRes
	QueryAsync(tx []byte) *ReqRes
	CommitAsync() *ReqRes

	FlushSync() error
	EchoSync(msg string) (res types.Result)
	InfoSync() (resInfo types.ResponseInfo, err error)
	SetOptionSync(key string, value string) (res types.Result)
	DeliverTxSync(tx []byte) (res types.Result)
	CheckTxSync(tx []byte) (res types.Result)
	QuerySync(tx []byte) (res types.Result)
	CommitSync() (res types.Result)

	InitChainAsync(validators []*types.Validator) *ReqRes
	BeginBlockAsync(hash []byte, header *types.Header) *ReqRes
	EndBlockAsync(height uint64) *ReqRes

	InitChainSync(validators []*types.Validator) (err error)
	BeginBlockSync(hash []byte, header *types.Header) (err error)
	EndBlockSync(height uint64) (resEndBlock types.ResponseEndBlock, err error)*/
}

//It returns a new Client or an determined error 
func NewClient(addr, transport string, mustConnect bool) (client Client, err error) {
	fmt.Println("\nCreating client...\n") //JCNM
	switch transport {
	case "socket":
		client, err = NewSocketClient(addr, mustConnect)
		/*var cli1 Client
		client = cli1*/	//JCNM
	case "grpc":
		//client, err = NewGRPCClient(addr, mustConnect)
		var cli2 Client
		client = cli2
	default:
		err = fmt.Errorf("Unknown bft transport %s", transport)

	}
	//var err error
	return
}

//----------------------------------------

type Callback func(*types.Request, *types.Response)

//----------------------------------------

type ReqRes struct {
	*types.Request
	*sync.WaitGroup
	*types.Response // Not set atomically, so be sure to use WaitGroup.

	mtx  sync.Mutex
	done bool                  // Gets set to true once *after* WaitGroup.Done().
	cb   func(*types.Response) // A single callback that may be set.
}

func NewReqRes(req *types.Request) *ReqRes {
	return &ReqRes{
		Request:   req,
		WaitGroup: waitGroup1(),
		Response:  nil,

		done: false,
		cb:   nil,
	}
}

// Sets the callback for this ReqRes atomically.
// If reqRes is already done, calls cb immediately.
// NOTE: reqRes.cb should not change if reqRes.done.
// NOTE: only one callback is supported.
func (reqRes *ReqRes) SetCallback(cb func(res *types.Response)) {
	reqRes.mtx.Lock()

	if reqRes.done {
		reqRes.mtx.Unlock()
		cb(reqRes.Response)
		return
	}

	defer reqRes.mtx.Unlock()
	reqRes.cb = cb
}

func (reqRes *ReqRes) GetCallback() func(*types.Response) {
	reqRes.mtx.Lock()
	defer reqRes.mtx.Unlock()
	return reqRes.cb
}

// NOTE: it should be safe to read reqRes.cb without locks after this.
func (reqRes *ReqRes) SetDone() {
	reqRes.mtx.Lock()
	reqRes.done = true
	reqRes.mtx.Unlock()
}

func waitGroup1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}

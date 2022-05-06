package gethrpc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// RPCRequest represents a jsonrpc request object.
//
// See: http://www.jsonrpc.org/specification#request_object
type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      uint        `json:"id"`
}

// RPCResponse represents a jsonrpc response object.
// If no rpc specific error occurred Error field is nil.
//
// See: http://www.jsonrpc.org/specification#response_object
type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      uint        `json:"id"`
}

// BatchResponse a list of jsonrpc response objects as a result of a batch request
//
// if you are interested in the response of a specific request use: GetResponseOf(request)
type BatchResponse struct {
	rpcResponses []RPCResponse
}

// RPCError represents a jsonrpc error object if an rpc error occurred.
//
// See: http://www.jsonrpc.org/specification#error_object
type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *RPCError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}

// ParityRPCClient sends jsonrpc requests over http to the provided rpc backend.
// ParityRPCClient is created using the factory function NewRPCClient().
type Instance struct {
	endpoint        string
	httpClient      *http.Client
	customHeaders   map[string]string
	autoIncrementID bool
	nextID          uint
	idMutex         sync.Mutex
}

// NewRPCClient returns a new ParityRPCClient instance with default configuration (no custom headers, default http.Instance, autoincrement ids).
// Endpoint is the rpc-service url to which the rpc requests are sent.
func NewRPCClient(endpoint string) *Instance {
	client := &Instance{
		endpoint:        endpoint,
		httpClient:      http.DefaultClient,
		autoIncrementID: true,
		nextID:          0,
		customHeaders:   make(map[string]string),
	}

	return client
}

// NewRPCRequestObject creates and returns a raw RPCRequest structure.
// It is mainly used when building batch requests. For single requests use ParityRPCClient.Call().
// RPCRequest struct can also be created directly, but this function sets the ID and the jsonrpc field to the correct values.
func (client *Instance) NewRPCRequestObject(method string, params ...interface{}) *RPCRequest {
	client.idMutex.Lock()
	rpcRequest := RPCRequest{
		ID:      client.nextID,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}
	if client.autoIncrementID == true {
		client.nextID++
	}
	client.idMutex.Unlock()

	if len(params) == 0 {
		rpcRequest.Params = nil
	}

	return &rpcRequest
}

// Call sends an jsonrpc request over http to the rpc-service url that was provided on Instance creation.
//
// If something went wrong on the network / http level or if json parsing failed it returns an error.
//
// If something went wrong on the rpc-service / protocol level the Error field of the returned RPCResponse is set
// and contains information about the error.
//
// If the request was successful the Error field is nil and the Result field of the RPCRespnse struct contains the rpc result.
func (client *Instance) Call(method string, params ...interface{}) (*RPCResponse, error) {
	var p interface{}

	if len(params) != 0 {
		p = params
	}

	httpRequest, err := client.newRequest(method, p)

	if err != nil {
		return nil, err
	}

	return client.doCall(httpRequest)
}

// CallNamed sends an jsonrpc request over http to the rpc-service url that was provided on Instance creation.
// This differs from Call() by sending named, rather than positional, arguments.
//
// If something went wrong on the network / http level or if json parsing failed it returns an error.
//
// If something went wrong on the rpc-service / protocol level the Error field of the returned RPCResponse is set
// and contains information about the error.
//
// If the request was successful the Error field is nil and the Result field of the RPCRespnse struct contains the rpc result.
func (client *Instance) CallNamed(method string, params map[string]interface{}) (*RPCResponse, error) {
	httpRequest, err := client.newRequest(method, params)
	if err != nil {
		return nil, err
	}
	return client.doCall(httpRequest)
}

func (client *Instance) doCall(req *http.Request) (*RPCResponse, error) {
	httpResponse, err := client.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()
	decoder := json.NewDecoder(httpResponse.Body)
	decoder.UseNumber()
	rpcResponse := RPCResponse{}
	err = decoder.Decode(&rpcResponse)

	if err != nil {
		return nil, err
	}

	return &rpcResponse, nil
}

// Batch sends a jsonrpc batch request to the rpc-service.
// The parameter is a list of requests the could be one of:
//	RPCRequest
//	RPCNotification.
//
// The batch requests returns a list of RPCResponse structs.
func (client *Instance) Batch(requests ...interface{}) (*BatchResponse, error) {
	for _, r := range requests {
		switch r := r.(type) {
		default:
			return nil, fmt.Errorf("Invalid parameter: %s", r)
		case *RPCRequest:
		}
	}

	httpRequest, err := client.newBatchRequest(requests...)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	rpcResponses := []RPCResponse{}
	decoder := json.NewDecoder(httpResponse.Body)
	decoder.UseNumber()
	err = decoder.Decode(&rpcResponses)
	if err != nil {
		return nil, err
	}

	return &BatchResponse{rpcResponses: rpcResponses}, nil
}

// SetAutoIncrementID if set to true, the id field of an rpcjson request will be incremented automatically
func (client *Instance) SetAutoIncrementID(flag bool) {
	client.autoIncrementID = flag
}

// SetNextID can be used to manually set the next id / reset the id.
func (client *Instance) SetNextID(id uint) {
	client.idMutex.Lock()
	client.nextID = id
	client.idMutex.Unlock()
}

// SetBasicAuth is a helper function that sets the header for the given basic authentication credentials.
// To reset / disable authentication just set username or password to an empty string value.
func (client *Instance) SetBasicAuth(username string, password string) {
	if username == "" || password == "" {
		delete(client.customHeaders, "Authorization")
		return
	}
	auth := username + ":" + password
	client.customHeaders["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func (client *Instance) newRequest(method string, params interface{}) (*http.Request, error) {
	client.idMutex.Lock()
	rpcRequest := RPCRequest{
		ID:      client.nextID,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}
	if client.autoIncrementID == true {
		client.nextID++
	}
	client.idMutex.Unlock()

	body, err := json.Marshal(rpcRequest)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", client.endpoint, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	for k, v := range client.customHeaders {
		request.Header.Add(k, v)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return request, nil
}

func (client *Instance) newBatchRequest(requests ...interface{}) (*http.Request, error) {

	body, err := json.Marshal(requests)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", client.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for k, v := range client.customHeaders {
		request.Header.Add(k, v)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return request, nil
}

// UpdateRequestID updates the ID of an RPCRequest structure.
//
// This is used if a request is sent another time and the request should get an updated id.
//
// This does only make sense when used on with Batch() since Call() and Notififcation() do update the id automatically.
func (client *Instance) UpdateRequestID(rpcRequest *RPCRequest) {
	if rpcRequest == nil {
		return
	}
	client.idMutex.Lock()
	defer client.idMutex.Unlock()
	rpcRequest.ID = client.nextID
	if client.autoIncrementID == true {
		client.nextID++
	}
}

// GetStringList converts the rpc response to []string and returns it.
//
// If result was not a string an error is returned.
func (rpcResponse *RPCResponse) GetStringList() ([]string, error) {
	val, ok := rpcResponse.Result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("could not parse []interface{} list from %s", rpcResponse.Result)
	}

	s := make([]string, len(val))

	for k, v := range val {
		s[k] = v.(string)
	}

	return s, nil
}

// GetResponseOf returns the rpc response of the corresponding request by matching the id.
//
// For this method to work, autoincrementID should be set to true (default).
func (batchResponse *BatchResponse) GetResponseOf(request *RPCRequest) (*RPCResponse, error) {
	if request == nil {
		return nil, errors.New("parameter cannot be nil")
	}

	for _, elem := range batchResponse.rpcResponses {
		if elem.ID == request.ID {
			return &elem, nil
		}
	}

	return nil, fmt.Errorf("element with id %d not found", request.ID)
}

func CheckRPCError(response *RPCResponse, err error) (*RPCResponse, error) {
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response, nil
}

func (it *RPCResponse) ToJSONBytes() ([]byte, error) {
	if it.Error != nil {
		return nil, it.Error
	}

	if it.Result == nil {
		return nil, fmt.Errorf("response returned without error but result is empty")
	}

	return json.Marshal(it.Result)
}

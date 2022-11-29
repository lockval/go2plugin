package go2plugin

type Input struct {
	UID      string         // Which user called (call,login,watch)
	KSUID    string         // system call ID (system call)
	Info     string         // login Info(login)
	WATCHUID string         // Which ID to observe (watch)
	Requ     map[string]any // User Request Parameters (call)

	GetResp *GetAndLockResp // Data returned from GetAndLock

	Throw            func(Code int, Error string)                    // Throw throw an error to the client
	Log              func(v any)                                     // Log Output an arbitrary data to the console
	Sleep            func(ms int64, ksuid, cmd string, obj any)      // Sleep Call cmd("xxx/xxx", obj) after timing ms, ksuid:must have value
	GetSubValAll     func(id, key string) *GetOpt                    // GetSubValAll Get all the data of the key
	GetAndLock       func()                                          // GetAndLock Acquire and lock (can only be called once)
	DiscardAndUnlock func() (resp *PutAndUnlockResp)                 // DiscardAndUnlock discard all edits (can only be called once)
	PutAndUnlock     func() (resp *PutAndUnlockResp)                 // PutAndUnlock change and unlock (can only be called once)
	GetSubVal        func(id, key string, subkeys ...string) *GetOpt // GetSubVal Get some data, when the subkeys are empty, it is to get all
	PutSubVal        func(id, key string, kvs ...string) *PutOpt     // PutSubVal set key value

	get *GetAndLockRequ
	put *PutAndUnlockRequ
}

// GetOpt GetOpt
type GetOpt struct {
	input   *Input
	id, key string
}

// Link Link
func (gsv *GetOpt) Link(slaveKeys ...string) *GetOpt {
	LinkMaster := NewBase62UUID()
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].LinkMaster = LinkMaster
	for _, key := range slaveKeys {
		if key == gsv.key {
			continue
		}
		_, ok := gsv.input.get.IDKey[gsv.id].KeySub[key]
		if !ok {
			gsv.input.get.IDKey[gsv.id].KeySub[key] = &GetRequSub{}
		}
		gsv.input.get.IDKey[gsv.id].KeySub[key].LinkSlave = LinkMaster

		_, ok = gsv.input.put.IDKey[gsv.id].KeySub[key]
		if !ok {
			gsv.input.put.IDKey[gsv.id].KeySub[key] = &PutRequSub{}
		}

	}
	return gsv
}

// Max Max
func (gsv *GetOpt) Max(num uint32, v int64) *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].MaxNum = num
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].MaxVal = v
	return gsv
}

// Min Min
func (gsv *GetOpt) Min(num uint32, v int64) *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].MinNum = num
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].MinVal = v
	return gsv
}

// Range Range
func (gsv *GetOpt) Range(v int32) *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Range = v
	return gsv
}

// Search Search
func (gsv *GetOpt) Search(v string) *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Search = v
	return gsv
}

// Random Random
func (gsv *GetOpt) Random(v uint32) *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Random = v
	return gsv
}

// Sum Sum
func (gsv *GetOpt) Sum() *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Sum = true
	return gsv
}

// Len Len
func (gsv *GetOpt) Len() *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Len = true
	return gsv
}

// Unique Unique
func (gsv *GetOpt) Unique() *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Unique = true
	return gsv
}

// Group Group
func (gsv *GetOpt) Group() *GetOpt {
	gsv.input.get.IDKey[gsv.id].KeySub[gsv.key].Group = true
	return gsv
}

// PutOpt PutOpt
type PutOpt struct {
	input   *Input
	id, key string
}

// Clear Clear
func (psv *PutOpt) Clear() *PutOpt {
	psv.input.put.IDKey[psv.id].KeySub[psv.key].Clear = true
	return psv
}

// List List
func (psv *PutOpt) List(val int32) *PutOpt {
	psv.input.put.IDKey[psv.id].KeySub[psv.key].ListVal = val
	psv.input.put.IDKey[psv.id].KeySub[psv.key].List = true
	return psv
}

// Link Link
func (psv *PutOpt) Link(slaveKeys ...string) *PutOpt {
	LinkMaster := NewBase62UUID()
	psv.input.put.IDKey[psv.id].KeySub[psv.key].LinkMaster = LinkMaster
	for _, key := range slaveKeys {
		if key == psv.key {
			continue
		}

		_, ok := psv.input.put.IDKey[psv.id].KeySub[key]
		if !ok {
			psv.input.put.IDKey[psv.id].KeySub[key] = &PutRequSub{}
		}
		psv.input.put.IDKey[psv.id].KeySub[key].LinkSlave = LinkMaster

	}
	return psv
}

func InputGet(input *Input) *GetAndLockRequ {
	return input.get
}

func InputPut(input *Input) *PutAndUnlockRequ {
	return input.put
}

func SetInputGet(input *Input, get *GetAndLockRequ) {
	input.get = get
}

func SetInputPut(input *Input, put *PutAndUnlockRequ) {
	input.put = put
}

func NewGetOpt(input *Input, id, key string) *GetOpt {
	return &GetOpt{input: input, id: id, key: key}
}

func NewPutOpt(input *Input, id, key string) *PutOpt {
	return &PutOpt{input: input, id: id, key: key}
}

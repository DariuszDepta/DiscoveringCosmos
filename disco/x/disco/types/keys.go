package types

const (
	// ModuleName defines the module name
	ModuleName = "disco"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_disco"
)

var (
	ParamsKey = []byte("p_disco")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

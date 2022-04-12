package dutils

type TCMD string

const (
	closeCmd TCMD = "CLOSE"
	setCmd   TCMD = "SET"
	getCmd   TCMD = "GET"
	delCmd   TCMD = "DEL"
	keyCmd   TCMD = "KEY"
)

type command[K comparable, V interface{}] struct {
	kind  TCMD
	key   K
	value V
}

func NewCommand[K comparable, V any](kind TCMD, key K, value V) command[K, V] {
	return command[K, V]{
		kind,
		key,
		value,
	}
}

type dmap[K comparable, V interface{}] struct {
	poom     chan *command[K, V]
	internal map[K]V
	chGet    chan V
	chKeys   chan []K
	chSet    chan V
	chDel    chan K
}

func (m *dmap[K, V]) run() {
	for cmd := range m.poom {
		switch cmd.kind {
		case closeCmd:
			close(m.poom)
			close(m.chGet)
			close(m.chSet)
			close(m.chKeys)
			close(m.chDel)
			m.internal = nil
			return
		case getCmd:
			m.chGet <- m.internal[cmd.key]
		case setCmd:
			m.internal[cmd.key] = cmd.value
			m.chSet <- cmd.value
		case delCmd:
			delete(m.internal, cmd.key)
			m.chDel <- cmd.key
		case keyCmd:
			keys := make([]K, 0, len(m.internal))
			for key := range m.internal {
				keys = append(keys, key)
			}
			m.chKeys <- keys
		}
	}
}

func (m *dmap[K, V]) pushCmd(cmd *command[K, V]) {
	m.poom <- cmd
}

// Get a value from Map
func (m *dmap[K, V]) Get(key K) interface{} {
	get := command[K, V]{
		kind: getCmd,
		key:  key,
	}
	go m.pushCmd(&get)
	return <-m.chGet
}

func (m *dmap[K, V]) Del(key K) {
	del := command[K, V]{
		kind: delCmd,
		key:  key,
	}
	go m.pushCmd(&del)
	<-m.chDel

}

// Set a value for given key
func (m *dmap[K, V]) Set(key K, value V) V {
	set := command[K, V]{
		kind:  setCmd,
		key:   key,
		value: value,
	}
	go m.pushCmd(&set)
	return <-m.chSet
}

// Return a slice of map key
func (m *dmap[K, V]) Keys() []K {
	key := command[K, V]{
		kind: keyCmd,
	}
	go m.pushCmd(&key)
	return <-m.chKeys
}

// Close the map
func (m *dmap[K, V]) Close() {
	m.pushCmd(&command[K, V]{
		kind: closeCmd,
	})
}

// Create a safe map, using one channel to process get/set
func NewMap[K comparable, V interface{}]() *dmap[K, V] {
	m := &dmap[K, V]{
		poom:     make(chan *command[K, V]),
		internal: make(map[K]V),
		chGet:    make(chan V),
		chSet:    make(chan V),
		chDel:    make(chan K),
		chKeys:   make(chan []K),
	}
	go m.run()
	return m
}

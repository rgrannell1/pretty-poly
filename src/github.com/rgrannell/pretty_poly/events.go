
package pretty_poly





type Emitter struct {
	listeners map[ interface{ } ][ ]Listener
}

type Listener struct {
	callback func(...interface{ })
}





func NewEmitter( ) *Emitter {

	emitter := new(Emitter)

	emitter.listeners = make(map[interface{ }][ ]Listener)

	return emitter

}



func (self *Emitter) On(event string, callback func(...interface{ })) *Emitter {

	if _, ok := self.listeners[event]; !ok {
		self.listeners[event] = [ ]Listener{ }
	}

	self.listeners[event] = append(self.listeners[event], Listener{callback})

	return self

}

func (self *Emitter) Emit(event string, args ...interface{ }) *Emitter {

	for _, listener := range(self.listeners[event]) {
		go listener.callback(args...)
	}

	return self

}






const (
	EVENT_DRAW_IMAGE = "DRAW_IMAGE"
)

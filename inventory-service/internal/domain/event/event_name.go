package event

type EventName string

func NewEventName(name string) (EventName, error) {
	if name == "" {
		return "", ErrEventNameEmpty
	}
	return EventName(name), nil
}

func RehydrateEventName(name string) EventName {
	return EventName(name)
}

func (e EventName) String() string {
	return string(e)
}

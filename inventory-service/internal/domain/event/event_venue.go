package event

type EventVenue string

func NewEventVenue(name string) (EventVenue, error) {
	if name == "" {
		return "", ErrEventVenueNameEmpty
	}
	return EventVenue(name), nil
}

func RehydrateEventVenue(venueName string) EventVenue {
	return EventVenue(venueName)
}

func (e EventVenue) String() string {
	return string(e)
}

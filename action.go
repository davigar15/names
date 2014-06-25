// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ActionTagKind is used to identify the Tag type
	ActionTagKind = "action"

	// ActionMarker is the identifier used to join filterable
	// prefixes for Action Id's with unique suffixes
	ActionMarker = "_a_"
)

// IsAction returns whether actionId is a valid actionId
// Valid action ids include the names.ActionMarker token that delimits
// a prefix that can be used for filtering, and a suffix that should be
// unique.  The prefix should match the name rules for units
func IsAction(actionId string) bool {
	_, ok := parseActionId(actionId)
	return ok
}

// ActionTag is a Tag type for representing Action entities, which
// are records of queued actions for a given unit
type ActionTag struct {
	unit     UnitTag
	sequence int
}

// String returns a string that shows the type and id of an ActionTag
func (t ActionTag) String() string {
	return t.Kind() + "-" + t.Id()
}

// Kind exposes the ActionTagKind value to identify what kind of Tag this is
func (t ActionTag) Kind() string { return ActionTagKind }

// Id returns the id of the Action this Tag represents
func (t ActionTag) Id() string { return fmt.Sprintf("%s%s%d", t.unit.Id(), ActionMarker, t.sequence) }

// NewActionTag returns the tag for the action with the given id.
func NewActionTag(tag UnitTag, sequence int) ActionTag {
	return ActionTag{unit: tag, sequence: sequence}
}

// UnitTag returns the UnitTag that the ActionTag is queued for
func (t ActionTag) UnitTag() UnitTag {
	return t.unit
}

// Sequence returns the unique suffix of the ActionTag
func (t ActionTag) Sequence() int {
    return t.sequence
}

// parseActionId creates an ActionTag from an actionId
// Id.  It returns false if the actionId cannot be parsed otherwise true
func parseActionId(actionId string) (ActionTag, bool) {
	bad := ActionTag{}
	parts := strings.Split(actionId, ActionMarker)
	// must have exactly one ActionMarker token
	if len(parts) != 2 {
		return bad, false
	}
	// first part must be a unit name
	tag, ok := tagFromUnitName(parts[0])
	if !ok {
		return bad, false
	}

	sl := len(parts[1])
	// sequence has to be at least one digit long
	if sl == 0 {
		return bad, false
	}
	// sequence cannot have leading zero if more than
	// one digit long
	if sl > 1 && strings.HasPrefix(parts[1], "0") {
		return bad, false
	}
	// sequence must be a number (it's generated as int)
	sequence, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return bad, false
	}
	return ActionTag{unit: tag, sequence: int(sequence)}, true
}

// ParseActionTag parses a action tag string.
func ParseActionTag(actionTag string) (ActionTag, error) {
	tag, err := ParseTag(actionTag)
	if err != nil {
		return ActionTag{}, err
	}
	st, ok := tag.(ActionTag)
	if !ok {
		return ActionTag{}, invalidTagError(actionTag, ActionTagKind)
	}
	return st, nil
}
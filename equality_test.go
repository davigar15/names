package names

import (
	gc "launchpad.net/gocheck"
)

var tagEqualityTests = []struct {
	expected Tag
	want     Tag
}{
	{NewMachineTag("0"), MachineTag{id: "0"}},
	{NewMachineTag("10/lxc/1"), MachineTag{id: "10-lxc-1"}},
	{NewUnitTag("mysql/1"), UnitTag{name: "mysql-1"}},
	{NewServiceTag("ceph"), ServiceTag{name: "ceph"}},
	{NewRelationTag("wordpress:haproxy"), RelationTag{key: "wordpress.haproxy"}},
	{NewEnvironTag("local"), EnvironTag{uuid: "local"}},
	{NewUserTag("admin"), UserTag{name: "admin"}},
	{NewNetworkTag("eth0"), NetworkTag{name: "eth0"}},
	{makeActionTag("foo" + ActionMarker + "321"), ActionTag{}},
	{makeActionTag("foo/0" + ActionMarker + "321"), ActionTag{unit: NewUnitTag("foo/0"), sequence: 321}},
}

type equalitySuite struct{}

var _ = gc.Suite(&equalitySuite{})

func (s *equalitySuite) TestTagEquality(c *gc.C) {
	for _, tt := range tagEqualityTests {
		c.Check(tt.want, gc.Equals, tt.expected)
	}
}

func makeActionTag(actionId string) ActionTag {
	if tag, ok := ParseActionId(actionId); ok {
		return tag
	}
	return ActionTag{}
}
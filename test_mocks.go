package main

import (
    "evohome"
)

type MockSystem struct {
    Id string
    Type string
    Zones []evohome.Zone
}

var testZone = evohome.Zone {
    Id: "98765",
    Name: "TestZone",
    ModelType: "MockZone",
    ZoneType: "MockZone",
    // TemperatureStatus: nil,
    // HeatSetPointStatus: nil,
    // Schedules: nil,
}

func (t *MockSystem) Zone(name string) (*evohome.Zone) {
    return &testZone
}

func (t *MockSystem) ZoneNames() ([]string) {
    return []string{"TestZone"}
}

func (t *MockSystem) ZoneNamesWithOverride() ([]string) {
    return []string{"TestZone"}
}

func (t *MockSystem) ZonesMap() (map[string]*evohome.Zone) {
    return map[string]*evohome.Zone {
        "TestZone": &testZone,
    }
}

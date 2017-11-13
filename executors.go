package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
    "github.com/araddon/dateparse"
    "github.com/ryanuber/columnize"
)

func executor(s string) {
    s = strings.TrimSpace(s)
    if s == "" {
        return
    } else if s == "quit" || s == "exit" {
        fmt.Println("Bye!")
        os.Exit(0)
        return
    }

    if !clientInitialized() {
        error("Evohome client not initialized")
    }

    t := client.TemperatureControlSystem()

    splitCmd := strings.Split(s, " ")
    first := splitCmd[0]
    switch first {
        case "cancel":
            if len(splitCmd) != 5 {
                error("Invalid command")
                return
            }

            re := regexp.MustCompile(`(?i)^cancel zone (?P<zone>[a-z\d]+) (?P<action>[a-z]+) (?P<attribute>[a-z]+)$`)
            matches := RegexSubMatchMap(re, s)
            if len(matches) == 0 {
                error("Invalid command")
                return
            }

            zone := t.Zone(matches["zone"])
            if zone == nil {
                error("Invalid zone '" + matches["zone"] + "'")
                return
            }

            if matches["action"] != "temperature" {
                error("Invalid action '" + matches["action"] + "'")
                return
            }

            if matches["attribute"] != "override" {
                error("Invalid attribute type '" + matches["attribute"] + "'")
                return
            }

            err := zone.CancelTemperatureOverride()
            if err != nil {
                error("An error occurred while cancelling temperature override for zone '" + zone.Name + "'")
            }
        case "set":
            if !(len(splitCmd) == 5 || len(splitCmd) >= 7) {
                error("Invalid command")
                return
            }

            re := regexp.MustCompile(`(?i)^set zone (?P<zone>[a-z\d]+) (?P<action>[a-z]+) (?P<value>\d+(\.[\d]+)?)(\suntil (?P<time>.+))?$`)
            matches := RegexSubMatchMap(re, s)
            if len(matches) == 0 {
                error("Invalid command")
                return
            }

            zone := t.Zone(matches["zone"])
            if zone == nil {
                error("Invalid zone '" + matches["zone"] + "'")
                return
            }

            if matches["action"] != "temperature" {
                error("Invalid action '" + matches["action"] + "'")
                return
            }

            temperature, err := strconv.ParseFloat(matches["value"], 32)
            if err != nil {
                error("Invalid temperature value")
                return
            }

            if matches["time"] == "" {
                err := zone.SetTemperature(float32(temperature), time.Time{})
                if err != nil {
                    error("An error occurred while configuring temperature for zone '" + zone.Name + "'")
                }
                return
            }

            until, parseErr := dateparse.ParseLocal(matches["time"])
            if parseErr != nil {
                error("Invalid date/time '" + matches["time"] + "'; use format 'yyyy/mm/dd hh:mm'")
                return
            }

            err = zone.SetTemperature(float32(temperature), until)
            if err != nil {
                error("An error occurred while configuring temperature for zone '" + zone.Name + "'")
            }
        case "show":
            if len(splitCmd) <= 1 {
                error("Invalid command")
                return
            }

            second := splitCmd[1]
            if second == "zone" {
                if len(splitCmd) < 3 {
                    error("Invalid command")
                    return
                }

                re := regexp.MustCompile(`(?i)^show zone (?P<zone>[a-z\d]+)(\sschedule)?$`)
                matches := RegexSubMatchMap(re, s)
                if len(matches) == 0 {
                    error("Invalid command")
                    return
                }

                zone := t.Zone(matches["zone"])
                if zone == nil {
                    error("Invalid zone '" + matches["zone"] + "'")
                    return
                }

                if strings.Contains(s, "schedule") {
                    output := []string {
                        "Day of week | Time | Temperature",
                        "---- | ---- | ----",
                    }
                    for _, schedule := range zone.Schedules.DailySchedules {
                        for i, switchPoint := range schedule.SwitchPoints {
                            var day string
                            if i == 0 {
                                day = schedule.DayOfWeek
                            } else {
                                day = ""
                            }
                            output = append(output, fmt.Sprintf("%s|%s|%.1f",
                                    day, switchPoint.Time, switchPoint.Temperature))
                        }
                    }
                    fmt.Println(columnize.SimpleFormat(output))
                } else {
                    fmt.Printf("%-20s %s\n", "Name:", zone.Name)
                    fmt.Printf("%-20s %s\n", "Zone type:", zone.ZoneType)
                    fmt.Printf("%-20s %s\n", "Model type:", zone.ModelType)
                    if zone.TemperatureStatus.IsAvailable {
                        fmt.Printf("%-20s %.1f\n", "Temperature:", zone.TemperatureStatus.Temperature)
                    }

                    var override string
                    if zone.HeatSetPointStatus.SetPointMode == "PermanentOverride" {
                        override = "(permanent override)"
                    } else if zone.HeatSetPointStatus.SetPointMode == "TemporaryOverride" {
                        override = "(temporary override)"
                    }

                    fmt.Printf("%-20s %.1f %s\n", "Target temperature:",
                            zone.HeatSetPointStatus.TargetTemperature,
                            override)
                }
            } else if second == "zones" {
                output := []string {
                    "Name | Type | Temperature",
                    "---- | ---- | ----",
                }
                for _, zone := range t.Zones {
                    temperature := "-"
                    if zone.TemperatureStatus.IsAvailable {
                        temperature = fmt.Sprintf("%.1f", zone.TemperatureStatus.Temperature)
                    }

                    targetTemperature := fmt.Sprintf("%.1f", zone.HeatSetPointStatus.TargetTemperature)
                    if temperature != targetTemperature {
                        output = append(output, fmt.Sprintf("%s|%s|%s* (%s)",
                                zone.Name, zone.ZoneType, temperature, targetTemperature))
                    } else {
                        output = append(output, fmt.Sprintf("%s|%s|%s",
                                zone.Name, zone.ZoneType, temperature))
                    }
                }
                fmt.Println(columnize.SimpleFormat(output))
            }
        default:
            error("Invalid command")
            return
    }
}

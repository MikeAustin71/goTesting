package main


func main() {

  // typeBaseName -> fullName -> typeDeclarationName -> typeDirectoryPath ->[]filepath -> tzName
  //
  m := make(map[string]map[string]map[string]map[string]map[string]string)

  m["America"] = make(map[string]map[string]map[string]map[string]string)
  m["America"]["AmericaTimeZones"] =  make(map[string]map[string]map[string]string)
  m["America"]["AmericaTimeZones"]["americaTimeZones"] = make(map[string]map[string]string)
  m["America"]["AmericaTimeZones"]["americaTimeZones"]["D:\\zoneinfo\\America"] =
    make(map[string]string)

  m["America"]["AmericaTimeZones"]["americaTimeZones"]["D:\\zoneinfo\\America"]["D:\\zoneinfo\\America\\Chicago"] =
    "America/Chicago"
  m["America"]["AmericaTimeZones"]["americaTimeZones"]["D:\\zoneinfo\\America"]["D:\\zoneinfo\\America\\Denver"] =
    "America/Denver"
  m["America"]["AmericaTimeZones"]["americaTimeZones"]["D:\\zoneinfo\\America"]["D:\\zoneinfo\\America\\America\\New_York"] =
    "America/New_York"
  m["America"]["AmericaTimeZones"]["americaTimeZones"]["D:\\zoneinfo\\America"]["D:\\zoneinfo\\America\\America\\Los_Angeles"] =
    "America/Los_Angeles"



}

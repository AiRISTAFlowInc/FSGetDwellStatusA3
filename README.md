# GetDwellStatus

This activity checks dwell status for a given ItemId. If no device logs exist, then DwellStatus is false indicating absent.

## NOTE

GracePeriod >= 2 minutes. Zone data takes about one minute to be logged in the database.

## Configuration

### Input:

| Name        | Type   | Description                                                                                                                                             |
| :---------- | :----- | :------------------------------------------------------------------------------------------------------------------------------------------------------ |
| IP          | string | IP address followed by port. e.g. "32.41.13.112:322"                                                                                                    |
| CustomerId  | string | Customer unique identifier. e.g. "43"                                                                                                                   |
| Username    | string | Username used to log into Sofia. e.g. "adminUser"                                                                                                       |
| Password    | string | Password used to log into Sofia. e.g. "adminPassword"                                                                                                   |
| MAC         | string | Device MAC to monitor dwell status. e.g. "C4:CB:6B:11:22:00"                                                                                            |
| GracePeriod | string | Minutes allowed away from zone before Exit Alert. e.g. "10"                                                                                             |
| ZoneItem    | string | Monitored zone. Accepts either Zone name, Zone Id, Or Zone Obj. e.g. "Hallway 3", "1234", OR {"ZoneID":5215,"ZoneName":"Audiotorium","ZoneType":"Open"} |

### Output:

| Name        | Type    | Description                                     |
| :---------- | :------ | :---------------------------------------------- |
| DwellStatus | boolean | Returns True if still in zone, False if absent. |

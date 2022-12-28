# 🔌 The `evse` file(s) - EVSE configuration

EVSE settings files can be found in the `evse` folder. To add and configure the EVSE, simply add a new file with the
file structure as in the example. The client will scan the folder at boot, validate the configuration files and add the
EVSE with the provided configuration.

## Attributes

| Attribute  |                Description                |            Possible values             | 
|:----------:|:-----------------------------------------:|:--------------------------------------:|
|   evseId   |              ID of the EVSE               |                   >1                   |
| connectors | List of available connectors for the EVSE ||
|    evcc    |            Charging Controller            ||
|   status   |            Status of the EVSE             | "Available", "Charging", "Faulted",... |
|  session   |          Persistent session data          ||
| powerMeter |    Power Meter configuration for EVSE     || 
|  maxPower  |       Maximum power output for EVSE       |                   >0                   |

### evcc

|   Attribute   |              Description              |  Possible values   | 
|:-------------:|:-------------------------------------:|:------------------:|
|     type      |         Type of the EVCC used         | "Relay", "Phoenix" |
|   relayPin    |   Attribute specific for the Relay.   |    Any GPIO pin    |
| inverseLogic  |   Attribute specific for the Relay.   |     true,false     |
| deviceAddress | Attribute specific for smarter EVCCs. |  Any string value  |

### powerMeter

|      Attribute       |            Description            | Possible values | 
|:--------------------:|:---------------------------------:|:---------------:|
|      isEnabled       | Enable or disable the power meter |   true, false   |
|         type         |      Type of the power meter      |    "CS5460A"    |
|    powerMeterPin     |  Attribute specific for CS5460A   |  Any GPIO pin   |
|        spiBus        |  Attribute specific for CS5460A   |       0,1       |
|     consumption      |  Attribute specific for CS5460A   |                 |
|     shuntOffset      |  Attribute specific for CS5460A   |  Default: 1337  |
| voltageDividerOffset |  Attribute specific for CS5460A   |  Default: 0.01  |

### connectors

|   Attribute   |       Description       |            Possible values            | 
|:-------------:|:-----------------------:|:-------------------------------------:|
|  connectorId  |      connector ID       |          Default:"ChargePi"           |
|     type      |  type of the connector  | "Schuko", "Type1","Type2", "CCS", ... |
|    status     | Status of the connector |     "Available", "Charging", ...      |

Example EVSE configuration:

```json
{
  "evseId": 1,
  "maxPower": 6.2,
  "connectors": [
    {
      "connectorId": 1,
      "type": "Schuko",
      "status": "Available"
    }
  ],
  "evcc": {
    "type": "Relay",
    "relayPin": 26,
    "inverseLogic": false
  },
  "status": "Available",
  "session": {
    "isActive": false,
    "transactionId": "",
    "tagId": "",
    "started": "",
    "consumption": []
  },
  "powerMeter": {
    "enabled": false,
    "type": "CS5460A",
    "powerMeterPin": 25,
    "spiBus": 0,
    "consumption": 0.0,
    "shuntOffset": 0.055,
    "voltageDividerOffset": 1333
  }
}
```
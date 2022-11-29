---
# !!! DO NOT MODIFY !!!
# This file is auto-generated by the gendoc tool based on the source code of the app.
description: This section describes the configuration parameters and their types for WASP.
keywords:
- IOTA Node 
- Hornet Node
- WASP Node
- Smart Contracts
- L2
- Configuration
- JSON
- Customize
- Config
- reference
---


# Core Configuration

WASP uses a JSON standard format as a config file. If you are unsure about JSON syntax, you can find more information in the [official JSON specs](https://www.json.org).

You can change the path of the config file by using the `-c` or `--config` argument while executing `wasp` executable.

For example:
```bash
wasp -c config_defaults.json
```

You can always get the most up-to-date description of the config parameters by running:

```bash
wasp -h --full
```

## <a id="app"></a> 1. Application

| Name                      | Description                                            | Type    | Default value |
| ------------------------- | ------------------------------------------------------ | ------- | ------------- |
| checkForUpdates           | Whether to check for updates of the application or not | boolean | true          |
| [shutdown](#app_shutdown) | Configuration for shutdown                             | object  |               |

### <a id="app_shutdown"></a> Shutdown

| Name                     | Description                                                                                            | Type   | Default value |
| ------------------------ | ------------------------------------------------------------------------------------------------------ | ------ | ------------- |
| stopGracePeriod          | The maximum time to wait for background processes to finish during shutdown before terminating the app | string | "5m"          |
| [log](#app_shutdown_log) | Configuration for Shutdown Log                                                                         | object |               |

### <a id="app_shutdown_log"></a> Shutdown Log

| Name     | Description                                         | Type    | Default value  |
| -------- | --------------------------------------------------- | ------- | -------------- |
| enabled  | Whether to store self-shutdown events to a log file | boolean | true           |
| filePath | The file path to the self-shutdown log              | string  | "shutdown.log" |

Example:

```json
  {
    "app": {
      "checkForUpdates": true,
      "shutdown": {
        "stopGracePeriod": "5m",
        "log": {
          "enabled": true,
          "filePath": "shutdown.log"
        }
      }
    }
  }
```

## <a id="logger"></a> 2. Logger

| Name              | Description                                                                 | Type    | Default value |
| ----------------- | --------------------------------------------------------------------------- | ------- | ------------- |
| level             | The minimum enabled logging level                                           | string  | "info"        |
| disableCaller     | Stops annotating logs with the calling function's file name and line number | boolean | true          |
| disableStacktrace | Disables automatic stacktrace capturing                                     | boolean | false         |
| stacktraceLevel   | The level stacktraces are captured and above                                | string  | "panic"       |
| encoding          | The logger's encoding (options: "json", "console")                          | string  | "console"     |
| outputPaths       | A list of URLs, file paths or stdout/stderr to write logging output to      | array   | stdout        |
| disableEvents     | Prevents log messages from being raced as events                            | boolean | true          |

Example:

```json
  {
    "logger": {
      "level": "info",
      "disableCaller": true,
      "disableStacktrace": false,
      "stacktraceLevel": "panic",
      "encoding": "console",
      "outputPaths": [
        "stdout"
      ],
      "disableEvents": true
    }
  }
```

## <a id="inx"></a> 3. INX

| Name                  | Description                                                                                        | Type   | Default value    |
| --------------------- | -------------------------------------------------------------------------------------------------- | ------ | ---------------- |
| address               | The INX address to which to connect to                                                             | string | "localhost:9029" |
| maxConnectionAttempts | The amount of times the connection to INX will be attempted before it fails (1 attempt per second) | uint   | 30               |
| targetNetworkName     | The network name on which the node should operate on (optional)                                    | string | ""               |

Example:

```json
  {
    "inx": {
      "address": "localhost:9029",
      "maxConnectionAttempts": 30,
      "targetNetworkName": ""
    }
  }
```

## <a id="db"></a> 4. Database

| Name                                     | Description                                                                      | Type    | Default value |
| ---------------------------------------- | -------------------------------------------------------------------------------- | ------- | ------------- |
| engine                                   | The used database engine (rocksdb/mapdb)                                         | string  | "rocksdb"     |
| [consensusJournal](#db_consensusjournal) | Configuration for consensusJournal                                               | object  |               |
| [chainState](#db_chainstate)             | Configuration for chainState                                                     | object  |               |
| debugSkipHealthCheck                     | Ignore the check for corrupted databases (should only be used for debug reasons) | boolean | false         |

### <a id="db_consensusjournal"></a> ConsensusJournal

| Name | Description                                       | Type   | Default value             |
| ---- | ------------------------------------------------- | ------ | ------------------------- |
| path | The path to the consensus journal database folder | string | "waspdb/chains/consensus" |

### <a id="db_chainstate"></a> ChainState

| Name | Description                                  | Type   | Default value        |
| ---- | -------------------------------------------- | ------ | -------------------- |
| path | The path to the chain state databases folder | string | "waspdb/chains/data" |

Example:

```json
  {
    "db": {
      "engine": "rocksdb",
      "consensusJournal": {
        "path": "waspdb/chains/consensus"
      },
      "chainState": {
        "path": "waspdb/chains/data"
      },
      "debugSkipHealthCheck": false
    }
  }
```

## <a id="p2p"></a> 5. P2p

| Name               | Description                                             | Type   | Default value |
| ------------------ | ------------------------------------------------------- | ------ | ------------- |
| identityPrivateKey | Private key used to derive the node identity (optional) | string | ""            |
| [db](#p2p_db)      | Configuration for Database                              | object |               |

### <a id="p2p_db"></a> Database

| Name | Description                  | Type   | Default value     |
| ---- | ---------------------------- | ------ | ----------------- |
| path | The path to the p2p database | string | "waspdb/p2pstore" |

Example:

```json
  {
    "p2p": {
      "identityPrivateKey": "",
      "db": {
        "path": "waspdb/p2pstore"
      }
    }
  }
```

## <a id="registry"></a> 6. Registry

| Name                                   | Description                    | Type   | Default value |
| -------------------------------------- | ------------------------------ | ------ | ------------- |
| [chains](#registry_chains)             | Configuration for chains       | object |               |
| [dkShares](#registry_dkshares)         | Configuration for dkShares     | object |               |
| [trustedPeers](#registry_trustedpeers) | Configuration for trustedPeers | object |               |

### <a id="registry_chains"></a> Chains

| Name     | Description                         | Type   | Default value                |
| -------- | ----------------------------------- | ------ | ---------------------------- |
| filePath | The path to the chain registry file | string | "waspdb/chain_registry.json" |

### <a id="registry_dkshares"></a> DkShares

| Name     | Description                                          | Type   | Default value          |
| -------- | ---------------------------------------------------- | ------ | ---------------------- |
| filePath | The path to the distributed key shares registry file | string | "waspdb/dkshares.json" |

### <a id="registry_trustedpeers"></a> TrustedPeers

| Name     | Description                                 | Type   | Default value               |
| -------- | ------------------------------------------- | ------ | --------------------------- |
| filePath | The path to the trusted peers registry file | string | "waspdb/trusted_peers.json" |

Example:

```json
  {
    "registry": {
      "chains": {
        "filePath": "waspdb/chain_registry.json"
      },
      "dkShares": {
        "filePath": "waspdb/dkshares.json"
      },
      "trustedPeers": {
        "filePath": "waspdb/trusted_peers.json"
      }
    }
  }
```

## <a id="peering"></a> 7. Peering

| Name  | Description                                          | Type   | Default value  |
| ----- | ---------------------------------------------------- | ------ | -------------- |
| netID | Node host address as it is recognized by other peers | string | "0.0.0.0:4000" |
| port  | Port for Wasp committee connection/peering           | int    | 4000           |

Example:

```json
  {
    "peering": {
      "netID": "0.0.0.0:4000",
      "port": 4000
    }
  }
```

## <a id="chains"></a> 8. Chains

| Name                             | Description                                                          | Type    | Default value |
| -------------------------------- | -------------------------------------------------------------------- | ------- | ------------- |
| broadcastUpToNPeers              | Number of peers an offledger request is broadcasted to               | int     | 2             |
| broadcastInterval                | Time between re-broadcast of offledger requests                      | string  | "5s"          |
| apiCacheTTL                      | Time to keep processed offledger requests in api cache               | string  | "5m"          |
| pullMissingRequestsFromCommittee | Whether or not to pull missing requests from other committee members | boolean | true          |

Example:

```json
  {
    "chains": {
      "broadcastUpToNPeers": 2,
      "broadcastInterval": "5s",
      "apiCacheTTL": "5m",
      "pullMissingRequestsFromCommittee": true
    }
  }
```

## <a id="rawblocks"></a> 9. Raw Blocks

| Name      | Description                              | Type    | Default value |
| --------- | ---------------------------------------- | ------- | ------------- |
| enabled   | Whether the raw blocks plugin is enabled | boolean | false         |
| directory | The raw blocks path                      | string  | "blocks"      |

Example:

```json
  {
    "rawBlocks": {
      "enabled": false,
      "directory": "blocks"
    }
  }
```

## <a id="profiling"></a> 10. Profiling

| Name        | Description                                       | Type    | Default value    |
| ----------- | ------------------------------------------------- | ------- | ---------------- |
| enabled     | Whether the profiling plugin is enabled           | boolean | false            |
| bindAddress | The bind address on which the profiler listens on | string  | "localhost:6060" |

Example:

```json
  {
    "profiling": {
      "enabled": false,
      "bindAddress": "localhost:6060"
    }
  }
```

## <a id="wal"></a> 11. Write-Ahead Logging

| Name      | Description                       | Type    | Default value |
| --------- | --------------------------------- | ------- | ------------- |
| enabled   | Whether the WAL plugin is enabled | boolean | true          |
| directory | Path to logs folder               | string  | "wal"         |

Example:

```json
  {
    "wal": {
      "enabled": true,
      "directory": "wal"
    }
  }
```

## <a id="metrics"></a> 12. Metrics

| Name        | Description                             | Type    | Default value    |
| ----------- | --------------------------------------- | ------- | ---------------- |
| enabled     | Whether the metrics plugin is enabled   | boolean | true             |
| bindAddress | The bind address for the node dashboard | string  | "127.0.0.1:2112" |

Example:

```json
  {
    "metrics": {
      "enabled": true,
      "bindAddress": "127.0.0.1:2112"
    }
  }
```

## <a id="webapi"></a> 13. Web API

| Name                      | Description                                              | Type    | Default value    |
| ------------------------- | -------------------------------------------------------- | ------- | ---------------- |
| enabled                   | Whether the web api plugin is enabled                    | boolean | true             |
| nodeOwnerAddresses        | Defines a list of node owner addresses (bech32)          | array   |                  |
| bindAddress               | The bind address for the node web api                    | string  | "127.0.0.1:9090" |
| debugRequestLoggerEnabled | Whether the debug logging for requests should be enabled | boolean | false            |
| [auth](#webapi_auth)      | Configuration for auth                                   | object  |                  |

### <a id="webapi_auth"></a> Auth

| Name                        | Description                            | Type   | Default value |
| --------------------------- | -------------------------------------- | ------ | ------------- |
| scheme                      | Selects which authentication to choose | string | "jwt"         |
| [jwt](#webapi_auth_jwt)     | Configuration for JWT Auth             | object |               |
| [basic](#webapi_auth_basic) | Configuration for Basic Auth           | object |               |
| [ip](#webapi_auth_ip)       | Configuration for IP-based Auth        | object |               |

### <a id="webapi_auth_jwt"></a> JWT Auth

| Name     | Description        | Type   | Default value |
| -------- | ------------------ | ------ | ------------- |
| duration | Jwt token lifetime | string | "24h"         |

### <a id="webapi_auth_basic"></a> Basic Auth

| Name     | Description                                     | Type   | Default value |
| -------- | ----------------------------------------------- | ------ | ------------- |
| username | The username which grants access to the service | string | "wasp"        |

### <a id="webapi_auth_ip"></a> IP-based Auth

| Name      | Description                                          | Type  | Default value |
| --------- | ---------------------------------------------------- | ----- | ------------- |
| whitelist | A list of ips that are allowed to access the service | array | 127.0.0.1     |

Example:

```json
  {
    "webapi": {
      "enabled": true,
      "nodeOwnerAddresses": [],
      "bindAddress": "127.0.0.1:9090",
      "debugRequestLoggerEnabled": false,
      "auth": {
        "scheme": "jwt",
        "jwt": {
          "duration": "24h"
        },
        "basic": {
          "username": "wasp"
        },
        "ip": {
          "whitelist": [
            "127.0.0.1"
          ]
        }
      }
    }
  }
```

## <a id="nanomsg"></a> 14. nanomsg

| Name    | Description                              | Type    | Default value |
| ------- | ---------------------------------------- | ------- | ------------- |
| enabled | Whether the publisher plugin is enabled  | boolean | true          |
| port    | The port for the nanomsg event publisher | int     | 5550          |

Example:

```json
  {
    "nanomsg": {
      "enabled": true,
      "port": 5550
    }
  }
```

## <a id="dashboard"></a> 15. Dashboard

| Name                      | Description                                              | Type    | Default value    |
| ------------------------- | -------------------------------------------------------- | ------- | ---------------- |
| enabled                   | Whether the dashboard plugin is enabled                  | boolean | true             |
| bindAddress               | The bind address for the node dashboard                  | string  | "127.0.0.1:7000" |
| exploreAddressURL         | URL to add as href to addresses in the dashboard         | string  | ""               |
| debugRequestLoggerEnabled | Whether the debug logging for requests should be enabled | boolean | false            |
| [auth](#dashboard_auth)   | Configuration for auth                                   | object  |                  |

### <a id="dashboard_auth"></a> Auth

| Name                           | Description                            | Type   | Default value |
| ------------------------------ | -------------------------------------- | ------ | ------------- |
| scheme                         | Selects which authentication to choose | string | "basic"       |
| [jwt](#dashboard_auth_jwt)     | Configuration for JWT Auth             | object |               |
| [basic](#dashboard_auth_basic) | Configuration for Basic Auth           | object |               |
| [ip](#dashboard_auth_ip)       | Configuration for IP-based Auth        | object |               |

### <a id="dashboard_auth_jwt"></a> JWT Auth

| Name     | Description        | Type   | Default value |
| -------- | ------------------ | ------ | ------------- |
| duration | Jwt token lifetime | string | "24h"         |

### <a id="dashboard_auth_basic"></a> Basic Auth

| Name     | Description                                     | Type   | Default value |
| -------- | ----------------------------------------------- | ------ | ------------- |
| username | The username which grants access to the service | string | "wasp"        |

### <a id="dashboard_auth_ip"></a> IP-based Auth

| Name      | Description                                          | Type  | Default value |
| --------- | ---------------------------------------------------- | ----- | ------------- |
| whitelist | A list of ips that are allowed to access the service | array | 127.0.0.1     |

Example:

```json
  {
    "dashboard": {
      "enabled": true,
      "bindAddress": "127.0.0.1:7000",
      "exploreAddressURL": "",
      "debugRequestLoggerEnabled": false,
      "auth": {
        "scheme": "basic",
        "jwt": {
          "duration": "24h"
        },
        "basic": {
          "username": "wasp"
        },
        "ip": {
          "whitelist": [
            "127.0.0.1"
          ]
        }
      }
    }
  }
```

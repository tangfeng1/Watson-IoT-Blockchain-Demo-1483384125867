package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {
                                    "authorityRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "insurerRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "lastTrip": {
                                        "properties": {
                                            "ctx_sub_trips": {
                                                "items": {
                                                    "avg_speed": {
                                                        "type": "number"
                                                    },
                                                    "ctx_features": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "context_category": {
                                                                "type": "string"
                                                            },
                                                            "context_category_id": {
                                                                "type": "number"
                                                            },
                                                            "context_id": {
                                                                "type": "number"
                                                            },
                                                            "context_name": {
                                                                "type": "string"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "driver_id": {
                                                        "type": "string"
                                                    },
                                                    "driving_behavior_details": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "behavior_name": {
                                                                "type": "string"
                                                            },
                                                            "driver_id": {
                                                                "type": "string"
                                                            },
                                                            "end_latitude": {
                                                                "type": "number"
                                                            },
                                                            "end_longitude": {
                                                                "type": "number"
                                                            },
                                                            "end_time": {
                                                                "type": "number"
                                                            },
                                                            "generated_time": {
                                                                "type": "number"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "mo_id": {
                                                                "type": "string"
                                                            },
                                                            "start_latitude": {
                                                                "type": "number"
                                                            },
                                                            "start_longitude": {
                                                                "type": "number"
                                                            },
                                                            "start_time": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "end_latitude": {
                                                        "type": "number"
                                                    },
                                                    "end_longitude": {
                                                        "type": "number"
                                                    },
                                                    "end_time": {
                                                        "type": "number"
                                                    },
                                                    "id": {
                                                        "properties": {
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "length": {
                                                        "type": "number"
                                                    },
                                                    "mo_id": {
                                                        "type": "string"
                                                    },
                                                    "start_latitude": {
                                                        "type": "number"
                                                    },
                                                    "start_longitude": {
                                                        "type": "number"
                                                    },
                                                    "start_time": {
                                                        "type": "number"
                                                    },
                                                    "trip_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "driver_id": {
                                                "type": "string"
                                            },
                                            "end_altitude": {
                                                "type": "number"
                                            },
                                            "end_latitude": {
                                                "type": "number"
                                            },
                                            "end_longitude": {
                                                "type": "number"
                                            },
                                            "end_time": {
                                                "type": "number"
                                            },
                                            "generated_time": {
                                                "type": "number"
                                            },
                                            "id": {
                                                "properties": {
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "mo_id": {
                                                "type": "string"
                                            },
                                            "start_altitude": {
                                                "type": "number"
                                            },
                                            "start_latitude": {
                                                "type": "number"
                                            },
                                            "start_longitude": {
                                                "type": "number"
                                            },
                                            "start_time": {
                                                "type": "number"
                                            },
                                            "trip_features": {
                                                "items": {
                                                    "feature_name": {
                                                        "type": "string"
                                                    },
                                                    "feature_value": {
                                                        "type": "string"
                                                    },
                                                    "id": {
                                                        "type": "number"
                                                    },
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "trip_id": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "lastKnownLocation": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "timestamp": {
                                "description": "Event timestamp.",
                                "type": "string"
                            },
                            "vehicleID": {
                                "properties": {
                                    "driverID": {
                                        "items": {
                                            "licenceNumber": {
                                                "type": "string"
                                            },
                                            "licenceValidityFrom": {
                                                "type": "string"
                                            },
                                            "licenceValidityTo": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "vin": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "assetID"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAsset function",
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssets": {
            "description": "Delete the state of all assets. No arguments are accepted. For each managed asset, the state and history are erased, and the asset is removed if necessary from recent states.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAllAssets function",
                    "enum": [
                        "deleteAllAssets"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'assetID'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an 'assetID' for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAsset": {
            "description": "Delete one or more properties from an asset's state. Argument is a JSON encoded string containing an 'assetID' and an array of qualified property names. For example, in an event object containing common and custom properties objects, the argument might look like {'assetID':'A1',['common.location', 'custom.carrier', 'custom.temperature']} and the result of that invoke would be the removal of the location, carrier and temperature properties. The missing temperature would clear a 'OVERTEMP' alert when the rules engine runs.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' with a list of qualified property names.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "qualPropsToDelete": {
                                "items": {
                                    "description": "The qualified name of a property. E.g. 'event.common.carrier', 'event.custom.temperature', etc.",
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "required": [
                            "assetID",
                            "qualPropsToDelete"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deletePropertiesFromAsset function",
                    "enum": [
                        "deletePropertiesFromAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "nickname": {
                                "default": "IOT_BLOCKCHAIN_DEMO",
                                "description": "The nickname of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "version"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssets": {
            "description": "Returns the state of all assets as an array of JSON encoded strings. Accepts no arguments. For each managed asset, the state is read from the ledger and added to the returned array. Array is sorted by 'assetID'.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAllAssets function",
                    "enum": [
                        "readAllAssets"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states, often for different assets",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {
                                    "authorityRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "insurerRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "lastTrip": {
                                        "properties": {
                                            "ctx_sub_trips": {
                                                "items": {
                                                    "avg_speed": {
                                                        "type": "number"
                                                    },
                                                    "ctx_features": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "context_category": {
                                                                "type": "string"
                                                            },
                                                            "context_category_id": {
                                                                "type": "number"
                                                            },
                                                            "context_id": {
                                                                "type": "number"
                                                            },
                                                            "context_name": {
                                                                "type": "string"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "driver_id": {
                                                        "type": "string"
                                                    },
                                                    "driving_behavior_details": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "behavior_name": {
                                                                "type": "string"
                                                            },
                                                            "driver_id": {
                                                                "type": "string"
                                                            },
                                                            "end_latitude": {
                                                                "type": "number"
                                                            },
                                                            "end_longitude": {
                                                                "type": "number"
                                                            },
                                                            "end_time": {
                                                                "type": "number"
                                                            },
                                                            "generated_time": {
                                                                "type": "number"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "mo_id": {
                                                                "type": "string"
                                                            },
                                                            "start_latitude": {
                                                                "type": "number"
                                                            },
                                                            "start_longitude": {
                                                                "type": "number"
                                                            },
                                                            "start_time": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "end_latitude": {
                                                        "type": "number"
                                                    },
                                                    "end_longitude": {
                                                        "type": "number"
                                                    },
                                                    "end_time": {
                                                        "type": "number"
                                                    },
                                                    "id": {
                                                        "properties": {
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "length": {
                                                        "type": "number"
                                                    },
                                                    "mo_id": {
                                                        "type": "string"
                                                    },
                                                    "start_latitude": {
                                                        "type": "number"
                                                    },
                                                    "start_longitude": {
                                                        "type": "number"
                                                    },
                                                    "start_time": {
                                                        "type": "number"
                                                    },
                                                    "trip_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "driver_id": {
                                                "type": "string"
                                            },
                                            "end_altitude": {
                                                "type": "number"
                                            },
                                            "end_latitude": {
                                                "type": "number"
                                            },
                                            "end_longitude": {
                                                "type": "number"
                                            },
                                            "end_time": {
                                                "type": "number"
                                            },
                                            "generated_time": {
                                                "type": "number"
                                            },
                                            "id": {
                                                "properties": {
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "mo_id": {
                                                "type": "string"
                                            },
                                            "start_altitude": {
                                                "type": "number"
                                            },
                                            "start_latitude": {
                                                "type": "number"
                                            },
                                            "start_longitude": {
                                                "type": "number"
                                            },
                                            "start_time": {
                                                "type": "number"
                                            },
                                            "trip_features": {
                                                "items": {
                                                    "feature_name": {
                                                        "type": "string"
                                                    },
                                                    "feature_value": {
                                                        "type": "string"
                                                    },
                                                    "id": {
                                                        "type": "number"
                                                    },
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "trip_id": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "args": {
                                        "items": {
                                            "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "function": {
                                        "description": "function that created this state object",
                                        "type": "string"
                                    },
                                    "redirectedFromFunction": {
                                        "description": "function that originally received the event",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "lastKnownLocation": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "timestamp": {
                                "description": "Event timestamp.",
                                "type": "string"
                            },
                            "txntimestamp": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnuuid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "vehicleID": {
                                "properties": {
                                    "driverID": {
                                        "items": {
                                            "licenceNumber": {
                                                "type": "string"
                                            },
                                            "licenceValidityFrom": {
                                                "type": "string"
                                            },
                                            "licenceValidityTo": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "vin": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. The arg is an 'assetID' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an 'assetID' for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                    "properties": {
                        "alerts": {
                            "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                            "properties": {
                                "active": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "TRIP_GOOD",
                                            "TRIP_DANGEROUS",
                                            "RATING_DRIVER_INCREASE",
                                            "RATING_DRIVER_DECREASE",
                                            "RATING_INSURER_INCREASE",
                                            "RATING_INSURER_DECREASE"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "cleared": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "TRIP_GOOD",
                                            "TRIP_DANGEROUS",
                                            "RATING_DRIVER_INCREASE",
                                            "RATING_DRIVER_DECREASE",
                                            "RATING_INSURER_INCREASE",
                                            "RATING_INSURER_DECREASE"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "raised": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "TRIP_GOOD",
                                            "TRIP_DANGEROUS",
                                            "RATING_DRIVER_INCREASE",
                                            "RATING_DRIVER_DECREASE",
                                            "RATING_INSURER_INCREASE",
                                            "RATING_INSURER_DECREASE"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                }
                            },
                            "type": "object"
                        },
                        "assetID": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "A contract-specific indication that this asset is compliant.",
                            "type": "boolean"
                        },
                        "extension": {
                            "description": "Application-managed state. Opaque to contract.",
                            "properties": {
                                "authorityRating": {
                                    "properties": {
                                        "currentRating": {
                                            "type": "number"
                                        },
                                        "lastIncrement": {
                                            "type": "number"
                                        },
                                        "lastIncrementTime": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "insurerRating": {
                                    "properties": {
                                        "currentRating": {
                                            "type": "number"
                                        },
                                        "lastIncrement": {
                                            "type": "number"
                                        },
                                        "lastIncrementTime": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "lastTrip": {
                                    "properties": {
                                        "ctx_sub_trips": {
                                            "items": {
                                                "avg_speed": {
                                                    "type": "number"
                                                },
                                                "ctx_features": {
                                                    "items": {
                                                        "behavior_id": {
                                                            "type": "string"
                                                        },
                                                        "context_category": {
                                                            "type": "string"
                                                        },
                                                        "context_category_id": {
                                                            "type": "number"
                                                        },
                                                        "context_id": {
                                                            "type": "number"
                                                        },
                                                        "context_name": {
                                                            "type": "string"
                                                        },
                                                        "id": {
                                                            "type": "number"
                                                        },
                                                        "sub_trip_id": {
                                                            "type": "string"
                                                        },
                                                        "tenant_id": {
                                                            "type": "string"
                                                        },
                                                        "trip_id": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "driver_id": {
                                                    "type": "string"
                                                },
                                                "driving_behavior_details": {
                                                    "items": {
                                                        "behavior_id": {
                                                            "type": "string"
                                                        },
                                                        "behavior_name": {
                                                            "type": "string"
                                                        },
                                                        "driver_id": {
                                                            "type": "string"
                                                        },
                                                        "end_latitude": {
                                                            "type": "number"
                                                        },
                                                        "end_longitude": {
                                                            "type": "number"
                                                        },
                                                        "end_time": {
                                                            "type": "number"
                                                        },
                                                        "generated_time": {
                                                            "type": "number"
                                                        },
                                                        "id": {
                                                            "type": "number"
                                                        },
                                                        "mo_id": {
                                                            "type": "string"
                                                        },
                                                        "start_latitude": {
                                                            "type": "number"
                                                        },
                                                        "start_longitude": {
                                                            "type": "number"
                                                        },
                                                        "start_time": {
                                                            "type": "number"
                                                        },
                                                        "sub_trip_id": {
                                                            "type": "string"
                                                        },
                                                        "tenant_id": {
                                                            "type": "string"
                                                        },
                                                        "trip_id": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "end_latitude": {
                                                    "type": "number"
                                                },
                                                "end_longitude": {
                                                    "type": "number"
                                                },
                                                "end_time": {
                                                    "type": "number"
                                                },
                                                "id": {
                                                    "properties": {
                                                        "sub_trip_id": {
                                                            "type": "string"
                                                        },
                                                        "tenant_id": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "length": {
                                                    "type": "number"
                                                },
                                                "mo_id": {
                                                    "type": "string"
                                                },
                                                "start_latitude": {
                                                    "type": "number"
                                                },
                                                "start_longitude": {
                                                    "type": "number"
                                                },
                                                "start_time": {
                                                    "type": "number"
                                                },
                                                "trip_id": {
                                                    "type": "string"
                                                },
                                                "trip_uuid": {
                                                    "type": "string"
                                                }
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "driver_id": {
                                            "type": "string"
                                        },
                                        "end_altitude": {
                                            "type": "number"
                                        },
                                        "end_latitude": {
                                            "type": "number"
                                        },
                                        "end_longitude": {
                                            "type": "number"
                                        },
                                        "end_time": {
                                            "type": "number"
                                        },
                                        "generated_time": {
                                            "type": "number"
                                        },
                                        "id": {
                                            "properties": {
                                                "tenant_id": {
                                                    "type": "string"
                                                },
                                                "trip_uuid": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "mo_id": {
                                            "type": "string"
                                        },
                                        "start_altitude": {
                                            "type": "number"
                                        },
                                        "start_latitude": {
                                            "type": "number"
                                        },
                                        "start_longitude": {
                                            "type": "number"
                                        },
                                        "start_time": {
                                            "type": "number"
                                        },
                                        "trip_features": {
                                            "items": {
                                                "feature_name": {
                                                    "type": "string"
                                                },
                                                "feature_value": {
                                                    "type": "string"
                                                },
                                                "id": {
                                                    "type": "number"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                },
                                                "trip_uuid": {
                                                    "type": "string"
                                                }
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "trip_id": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "lastEvent": {
                            "description": "function and string parameter that created this state object",
                            "properties": {
                                "args": {
                                    "items": {
                                        "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                        "type": "string"
                                    },
                                    "type": "array"
                                },
                                "function": {
                                    "description": "function that created this state object",
                                    "type": "string"
                                },
                                "redirectedFromFunction": {
                                    "description": "function that originally received the event",
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "lastKnownLocation": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        },
                        "timestamp": {
                            "description": "Event timestamp.",
                            "type": "string"
                        },
                        "txntimestamp": {
                            "description": "Transaction timestamp matching that in the blockchain.",
                            "type": "string"
                        },
                        "txnuuid": {
                            "description": "Transaction UUID matching that in the blockchain.",
                            "type": "string"
                        },
                        "vehicleID": {
                            "properties": {
                                "driverID": {
                                    "items": {
                                        "licenceNumber": {
                                            "type": "string"
                                        },
                                        "licenceValidityFrom": {
                                            "type": "string"
                                        },
                                        "licenceValidityTo": {
                                            "type": "string"
                                        },
                                        "name": {
                                            "type": "string"
                                        }
                                    },
                                    "minItems": 1,
                                    "type": "array"
                                },
                                "vin": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetHistory": {
            "description": "Requests a specified number of history states for an assets. Returns an array of states sorted with the most recent first. The 'assetID' property is required and the count property is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' with item 'count'.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "count": {
                                "type": "integer"
                            }
                        },
                        "required": [
                            "assetID"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetHistory function",
                    "enum": [
                        "readAssetHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {
                                    "authorityRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "insurerRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "lastTrip": {
                                        "properties": {
                                            "ctx_sub_trips": {
                                                "items": {
                                                    "avg_speed": {
                                                        "type": "number"
                                                    },
                                                    "ctx_features": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "context_category": {
                                                                "type": "string"
                                                            },
                                                            "context_category_id": {
                                                                "type": "number"
                                                            },
                                                            "context_id": {
                                                                "type": "number"
                                                            },
                                                            "context_name": {
                                                                "type": "string"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "driver_id": {
                                                        "type": "string"
                                                    },
                                                    "driving_behavior_details": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "behavior_name": {
                                                                "type": "string"
                                                            },
                                                            "driver_id": {
                                                                "type": "string"
                                                            },
                                                            "end_latitude": {
                                                                "type": "number"
                                                            },
                                                            "end_longitude": {
                                                                "type": "number"
                                                            },
                                                            "end_time": {
                                                                "type": "number"
                                                            },
                                                            "generated_time": {
                                                                "type": "number"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "mo_id": {
                                                                "type": "string"
                                                            },
                                                            "start_latitude": {
                                                                "type": "number"
                                                            },
                                                            "start_longitude": {
                                                                "type": "number"
                                                            },
                                                            "start_time": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "end_latitude": {
                                                        "type": "number"
                                                    },
                                                    "end_longitude": {
                                                        "type": "number"
                                                    },
                                                    "end_time": {
                                                        "type": "number"
                                                    },
                                                    "id": {
                                                        "properties": {
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "length": {
                                                        "type": "number"
                                                    },
                                                    "mo_id": {
                                                        "type": "string"
                                                    },
                                                    "start_latitude": {
                                                        "type": "number"
                                                    },
                                                    "start_longitude": {
                                                        "type": "number"
                                                    },
                                                    "start_time": {
                                                        "type": "number"
                                                    },
                                                    "trip_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "driver_id": {
                                                "type": "string"
                                            },
                                            "end_altitude": {
                                                "type": "number"
                                            },
                                            "end_latitude": {
                                                "type": "number"
                                            },
                                            "end_longitude": {
                                                "type": "number"
                                            },
                                            "end_time": {
                                                "type": "number"
                                            },
                                            "generated_time": {
                                                "type": "number"
                                            },
                                            "id": {
                                                "properties": {
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "mo_id": {
                                                "type": "string"
                                            },
                                            "start_altitude": {
                                                "type": "number"
                                            },
                                            "start_latitude": {
                                                "type": "number"
                                            },
                                            "start_longitude": {
                                                "type": "number"
                                            },
                                            "start_time": {
                                                "type": "number"
                                            },
                                            "trip_features": {
                                                "items": {
                                                    "feature_name": {
                                                        "type": "string"
                                                    },
                                                    "feature_value": {
                                                        "type": "string"
                                                    },
                                                    "id": {
                                                        "type": "number"
                                                    },
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "trip_id": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "args": {
                                        "items": {
                                            "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "function": {
                                        "description": "function that created this state object",
                                        "type": "string"
                                    },
                                    "redirectedFromFunction": {
                                        "description": "function that originally received the event",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "lastKnownLocation": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "timestamp": {
                                "description": "Event timestamp.",
                                "type": "string"
                            },
                            "txntimestamp": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnuuid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "vehicleID": {
                                "properties": {
                                    "driverID": {
                                        "items": {
                                            "licenceNumber": {
                                                "type": "string"
                                            },
                                            "licenceValidityFrom": {
                                                "type": "string"
                                            },
                                            "licenceValidityTo": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "vin": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readRecentStates": {
            "description": "Returns the state of recently updated assets as an array of objects sorted with the most recently updated asset first. Each asset appears exactly once up to a maxmum of 20 in this version of the contract.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readRecentStates function",
                    "enum": [
                        "readRecentStates"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "TRIP_GOOD",
                                                "TRIP_DANGEROUS",
                                                "RATING_DRIVER_INCREASE",
                                                "RATING_DRIVER_DECREASE",
                                                "RATING_INSURER_INCREASE",
                                                "RATING_INSURER_DECREASE"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {
                                    "authorityRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "insurerRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "lastTrip": {
                                        "properties": {
                                            "ctx_sub_trips": {
                                                "items": {
                                                    "avg_speed": {
                                                        "type": "number"
                                                    },
                                                    "ctx_features": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "context_category": {
                                                                "type": "string"
                                                            },
                                                            "context_category_id": {
                                                                "type": "number"
                                                            },
                                                            "context_id": {
                                                                "type": "number"
                                                            },
                                                            "context_name": {
                                                                "type": "string"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "driver_id": {
                                                        "type": "string"
                                                    },
                                                    "driving_behavior_details": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "behavior_name": {
                                                                "type": "string"
                                                            },
                                                            "driver_id": {
                                                                "type": "string"
                                                            },
                                                            "end_latitude": {
                                                                "type": "number"
                                                            },
                                                            "end_longitude": {
                                                                "type": "number"
                                                            },
                                                            "end_time": {
                                                                "type": "number"
                                                            },
                                                            "generated_time": {
                                                                "type": "number"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "mo_id": {
                                                                "type": "string"
                                                            },
                                                            "start_latitude": {
                                                                "type": "number"
                                                            },
                                                            "start_longitude": {
                                                                "type": "number"
                                                            },
                                                            "start_time": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "end_latitude": {
                                                        "type": "number"
                                                    },
                                                    "end_longitude": {
                                                        "type": "number"
                                                    },
                                                    "end_time": {
                                                        "type": "number"
                                                    },
                                                    "id": {
                                                        "properties": {
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "length": {
                                                        "type": "number"
                                                    },
                                                    "mo_id": {
                                                        "type": "string"
                                                    },
                                                    "start_latitude": {
                                                        "type": "number"
                                                    },
                                                    "start_longitude": {
                                                        "type": "number"
                                                    },
                                                    "start_time": {
                                                        "type": "number"
                                                    },
                                                    "trip_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "driver_id": {
                                                "type": "string"
                                            },
                                            "end_altitude": {
                                                "type": "number"
                                            },
                                            "end_latitude": {
                                                "type": "number"
                                            },
                                            "end_longitude": {
                                                "type": "number"
                                            },
                                            "end_time": {
                                                "type": "number"
                                            },
                                            "generated_time": {
                                                "type": "number"
                                            },
                                            "id": {
                                                "properties": {
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "mo_id": {
                                                "type": "string"
                                            },
                                            "start_altitude": {
                                                "type": "number"
                                            },
                                            "start_latitude": {
                                                "type": "number"
                                            },
                                            "start_longitude": {
                                                "type": "number"
                                            },
                                            "start_time": {
                                                "type": "number"
                                            },
                                            "trip_features": {
                                                "items": {
                                                    "feature_name": {
                                                        "type": "string"
                                                    },
                                                    "feature_value": {
                                                        "type": "string"
                                                    },
                                                    "id": {
                                                        "type": "number"
                                                    },
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "trip_id": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "args": {
                                        "items": {
                                            "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "function": {
                                        "description": "function that created this state object",
                                        "type": "string"
                                    },
                                    "redirectedFromFunction": {
                                        "description": "function that originally received the event",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "lastKnownLocation": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "timestamp": {
                                "description": "Event timestamp.",
                                "type": "string"
                            },
                            "txntimestamp": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnuuid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "vehicleID": {
                                "properties": {
                                    "driverID": {
                                        "items": {
                                            "licenceNumber": {
                                                "type": "string"
                                            },
                                            "licenceValidityFrom": {
                                                "type": "string"
                                            },
                                            "licenceValidityTo": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "vin": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "setCreateOnUpdate": {
            "description": "Allow updateAsset to redirect to createAsset when 'assetID' does not exist.",
            "properties": {
                "args": {
                    "description": "True for redirect allowed, false for error on asset does not exist.",
                    "items": {
                        "setCreateOnUpdate": {
                            "type": "boolean"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "setCreateOnUpdate function",
                    "enum": [
                        "setCreateOnUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level in the contract.",
            "properties": {
                "args": {
                    "description": "logging levels indicate what you see",
                    "items": {
                        "logLevel": {
                            "enum": [
                                "CRITICAL",
                                "ERROR",
                                "WARNING",
                                "NOTICE",
                                "INFO",
                                "DEBUG"
                            ],
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "setLoggingLevel function",
                    "enum": [
                        "setLoggingLevel"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. The 'assetID' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {
                                    "authorityRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "insurerRating": {
                                        "properties": {
                                            "currentRating": {
                                                "type": "number"
                                            },
                                            "lastIncrement": {
                                                "type": "number"
                                            },
                                            "lastIncrementTime": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "lastTrip": {
                                        "properties": {
                                            "ctx_sub_trips": {
                                                "items": {
                                                    "avg_speed": {
                                                        "type": "number"
                                                    },
                                                    "ctx_features": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "context_category": {
                                                                "type": "string"
                                                            },
                                                            "context_category_id": {
                                                                "type": "number"
                                                            },
                                                            "context_id": {
                                                                "type": "number"
                                                            },
                                                            "context_name": {
                                                                "type": "string"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "driver_id": {
                                                        "type": "string"
                                                    },
                                                    "driving_behavior_details": {
                                                        "items": {
                                                            "behavior_id": {
                                                                "type": "string"
                                                            },
                                                            "behavior_name": {
                                                                "type": "string"
                                                            },
                                                            "driver_id": {
                                                                "type": "string"
                                                            },
                                                            "end_latitude": {
                                                                "type": "number"
                                                            },
                                                            "end_longitude": {
                                                                "type": "number"
                                                            },
                                                            "end_time": {
                                                                "type": "number"
                                                            },
                                                            "generated_time": {
                                                                "type": "number"
                                                            },
                                                            "id": {
                                                                "type": "number"
                                                            },
                                                            "mo_id": {
                                                                "type": "string"
                                                            },
                                                            "start_latitude": {
                                                                "type": "number"
                                                            },
                                                            "start_longitude": {
                                                                "type": "number"
                                                            },
                                                            "start_time": {
                                                                "type": "number"
                                                            },
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            },
                                                            "trip_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "end_latitude": {
                                                        "type": "number"
                                                    },
                                                    "end_longitude": {
                                                        "type": "number"
                                                    },
                                                    "end_time": {
                                                        "type": "number"
                                                    },
                                                    "id": {
                                                        "properties": {
                                                            "sub_trip_id": {
                                                                "type": "string"
                                                            },
                                                            "tenant_id": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "length": {
                                                        "type": "number"
                                                    },
                                                    "mo_id": {
                                                        "type": "string"
                                                    },
                                                    "start_latitude": {
                                                        "type": "number"
                                                    },
                                                    "start_longitude": {
                                                        "type": "number"
                                                    },
                                                    "start_time": {
                                                        "type": "number"
                                                    },
                                                    "trip_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "driver_id": {
                                                "type": "string"
                                            },
                                            "end_altitude": {
                                                "type": "number"
                                            },
                                            "end_latitude": {
                                                "type": "number"
                                            },
                                            "end_longitude": {
                                                "type": "number"
                                            },
                                            "end_time": {
                                                "type": "number"
                                            },
                                            "generated_time": {
                                                "type": "number"
                                            },
                                            "id": {
                                                "properties": {
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "mo_id": {
                                                "type": "string"
                                            },
                                            "start_altitude": {
                                                "type": "number"
                                            },
                                            "start_latitude": {
                                                "type": "number"
                                            },
                                            "start_longitude": {
                                                "type": "number"
                                            },
                                            "start_time": {
                                                "type": "number"
                                            },
                                            "trip_features": {
                                                "items": {
                                                    "feature_name": {
                                                        "type": "string"
                                                    },
                                                    "feature_value": {
                                                        "type": "string"
                                                    },
                                                    "id": {
                                                        "type": "number"
                                                    },
                                                    "tenant_id": {
                                                        "type": "string"
                                                    },
                                                    "trip_uuid": {
                                                        "type": "string"
                                                    }
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "trip_id": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "lastKnownLocation": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "timestamp": {
                                "description": "Event timestamp.",
                                "type": "string"
                            },
                            "vehicleID": {
                                "properties": {
                                    "driverID": {
                                        "items": {
                                            "licenceNumber": {
                                                "type": "string"
                                            },
                                            "licenceValidityFrom": {
                                                "type": "string"
                                            },
                                            "licenceValidityTo": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "vin": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "assetID"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "assetIDKey": {
            "description": "An object containing only an 'assetID' for use as an argument to read or delete.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "assetIDandCount": {
            "description": "Requested 'assetID' with item 'count'.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                }
            },
            "required": [
                "assetID"
            ],
            "type": "object"
        },
        "event": {
            "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "extension": {
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {
                        "authorityRating": {
                            "properties": {
                                "currentRating": {
                                    "type": "number"
                                },
                                "lastIncrement": {
                                    "type": "number"
                                },
                                "lastIncrementTime": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "insurerRating": {
                            "properties": {
                                "currentRating": {
                                    "type": "number"
                                },
                                "lastIncrement": {
                                    "type": "number"
                                },
                                "lastIncrementTime": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "lastTrip": {
                            "properties": {
                                "ctx_sub_trips": {
                                    "items": {
                                        "avg_speed": {
                                            "type": "number"
                                        },
                                        "ctx_features": {
                                            "items": {
                                                "behavior_id": {
                                                    "type": "string"
                                                },
                                                "context_category": {
                                                    "type": "string"
                                                },
                                                "context_category_id": {
                                                    "type": "number"
                                                },
                                                "context_id": {
                                                    "type": "number"
                                                },
                                                "context_name": {
                                                    "type": "string"
                                                },
                                                "id": {
                                                    "type": "number"
                                                },
                                                "sub_trip_id": {
                                                    "type": "string"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                },
                                                "trip_id": {
                                                    "type": "string"
                                                }
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "driver_id": {
                                            "type": "string"
                                        },
                                        "driving_behavior_details": {
                                            "items": {
                                                "behavior_id": {
                                                    "type": "string"
                                                },
                                                "behavior_name": {
                                                    "type": "string"
                                                },
                                                "driver_id": {
                                                    "type": "string"
                                                },
                                                "end_latitude": {
                                                    "type": "number"
                                                },
                                                "end_longitude": {
                                                    "type": "number"
                                                },
                                                "end_time": {
                                                    "type": "number"
                                                },
                                                "generated_time": {
                                                    "type": "number"
                                                },
                                                "id": {
                                                    "type": "number"
                                                },
                                                "mo_id": {
                                                    "type": "string"
                                                },
                                                "start_latitude": {
                                                    "type": "number"
                                                },
                                                "start_longitude": {
                                                    "type": "number"
                                                },
                                                "start_time": {
                                                    "type": "number"
                                                },
                                                "sub_trip_id": {
                                                    "type": "string"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                },
                                                "trip_id": {
                                                    "type": "string"
                                                }
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "end_latitude": {
                                            "type": "number"
                                        },
                                        "end_longitude": {
                                            "type": "number"
                                        },
                                        "end_time": {
                                            "type": "number"
                                        },
                                        "id": {
                                            "properties": {
                                                "sub_trip_id": {
                                                    "type": "string"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "length": {
                                            "type": "number"
                                        },
                                        "mo_id": {
                                            "type": "string"
                                        },
                                        "start_latitude": {
                                            "type": "number"
                                        },
                                        "start_longitude": {
                                            "type": "number"
                                        },
                                        "start_time": {
                                            "type": "number"
                                        },
                                        "trip_id": {
                                            "type": "string"
                                        },
                                        "trip_uuid": {
                                            "type": "string"
                                        }
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "driver_id": {
                                    "type": "string"
                                },
                                "end_altitude": {
                                    "type": "number"
                                },
                                "end_latitude": {
                                    "type": "number"
                                },
                                "end_longitude": {
                                    "type": "number"
                                },
                                "end_time": {
                                    "type": "number"
                                },
                                "generated_time": {
                                    "type": "number"
                                },
                                "id": {
                                    "properties": {
                                        "tenant_id": {
                                            "type": "string"
                                        },
                                        "trip_uuid": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "mo_id": {
                                    "type": "string"
                                },
                                "start_altitude": {
                                    "type": "number"
                                },
                                "start_latitude": {
                                    "type": "number"
                                },
                                "start_longitude": {
                                    "type": "number"
                                },
                                "start_time": {
                                    "type": "number"
                                },
                                "trip_features": {
                                    "items": {
                                        "feature_name": {
                                            "type": "string"
                                        },
                                        "feature_value": {
                                            "type": "string"
                                        },
                                        "id": {
                                            "type": "number"
                                        },
                                        "tenant_id": {
                                            "type": "string"
                                        },
                                        "trip_uuid": {
                                            "type": "string"
                                        }
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "trip_id": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "lastKnownLocation": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "timestamp": {
                    "description": "Event timestamp.",
                    "type": "string"
                },
                "vehicleID": {
                    "properties": {
                        "driverID": {
                            "items": {
                                "licenceNumber": {
                                    "type": "string"
                                },
                                "licenceValidityFrom": {
                                    "type": "string"
                                },
                                "licenceValidityTo": {
                                    "type": "string"
                                },
                                "name": {
                                    "type": "string"
                                }
                            },
                            "minItems": 1,
                            "type": "array"
                        },
                        "vin": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "required": [
                "assetID"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "IOT_BLOCKCHAIN_DEMO",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
            "properties": {
                "alerts": {
                    "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                    "properties": {
                        "active": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "TRIP_GOOD",
                                    "TRIP_DANGEROUS",
                                    "RATING_DRIVER_INCREASE",
                                    "RATING_DRIVER_DECREASE",
                                    "RATING_INSURER_INCREASE",
                                    "RATING_INSURER_DECREASE"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "cleared": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "TRIP_GOOD",
                                    "TRIP_DANGEROUS",
                                    "RATING_DRIVER_INCREASE",
                                    "RATING_DRIVER_DECREASE",
                                    "RATING_INSURER_INCREASE",
                                    "RATING_INSURER_DECREASE"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "raised": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "TRIP_GOOD",
                                    "TRIP_DANGEROUS",
                                    "RATING_DRIVER_INCREASE",
                                    "RATING_DRIVER_DECREASE",
                                    "RATING_INSURER_INCREASE",
                                    "RATING_INSURER_DECREASE"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "compliant": {
                    "description": "A contract-specific indication that this asset is compliant.",
                    "type": "boolean"
                },
                "extension": {
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {
                        "authorityRating": {
                            "properties": {
                                "currentRating": {
                                    "type": "number"
                                },
                                "lastIncrement": {
                                    "type": "number"
                                },
                                "lastIncrementTime": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "insurerRating": {
                            "properties": {
                                "currentRating": {
                                    "type": "number"
                                },
                                "lastIncrement": {
                                    "type": "number"
                                },
                                "lastIncrementTime": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "lastTrip": {
                            "properties": {
                                "ctx_sub_trips": {
                                    "items": {
                                        "avg_speed": {
                                            "type": "number"
                                        },
                                        "ctx_features": {
                                            "items": {
                                                "behavior_id": {
                                                    "type": "string"
                                                },
                                                "context_category": {
                                                    "type": "string"
                                                },
                                                "context_category_id": {
                                                    "type": "number"
                                                },
                                                "context_id": {
                                                    "type": "number"
                                                },
                                                "context_name": {
                                                    "type": "string"
                                                },
                                                "id": {
                                                    "type": "number"
                                                },
                                                "sub_trip_id": {
                                                    "type": "string"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                },
                                                "trip_id": {
                                                    "type": "string"
                                                }
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "driver_id": {
                                            "type": "string"
                                        },
                                        "driving_behavior_details": {
                                            "items": {
                                                "behavior_id": {
                                                    "type": "string"
                                                },
                                                "behavior_name": {
                                                    "type": "string"
                                                },
                                                "driver_id": {
                                                    "type": "string"
                                                },
                                                "end_latitude": {
                                                    "type": "number"
                                                },
                                                "end_longitude": {
                                                    "type": "number"
                                                },
                                                "end_time": {
                                                    "type": "number"
                                                },
                                                "generated_time": {
                                                    "type": "number"
                                                },
                                                "id": {
                                                    "type": "number"
                                                },
                                                "mo_id": {
                                                    "type": "string"
                                                },
                                                "start_latitude": {
                                                    "type": "number"
                                                },
                                                "start_longitude": {
                                                    "type": "number"
                                                },
                                                "start_time": {
                                                    "type": "number"
                                                },
                                                "sub_trip_id": {
                                                    "type": "string"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                },
                                                "trip_id": {
                                                    "type": "string"
                                                }
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "end_latitude": {
                                            "type": "number"
                                        },
                                        "end_longitude": {
                                            "type": "number"
                                        },
                                        "end_time": {
                                            "type": "number"
                                        },
                                        "id": {
                                            "properties": {
                                                "sub_trip_id": {
                                                    "type": "string"
                                                },
                                                "tenant_id": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "length": {
                                            "type": "number"
                                        },
                                        "mo_id": {
                                            "type": "string"
                                        },
                                        "start_latitude": {
                                            "type": "number"
                                        },
                                        "start_longitude": {
                                            "type": "number"
                                        },
                                        "start_time": {
                                            "type": "number"
                                        },
                                        "trip_id": {
                                            "type": "string"
                                        },
                                        "trip_uuid": {
                                            "type": "string"
                                        }
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "driver_id": {
                                    "type": "string"
                                },
                                "end_altitude": {
                                    "type": "number"
                                },
                                "end_latitude": {
                                    "type": "number"
                                },
                                "end_longitude": {
                                    "type": "number"
                                },
                                "end_time": {
                                    "type": "number"
                                },
                                "generated_time": {
                                    "type": "number"
                                },
                                "id": {
                                    "properties": {
                                        "tenant_id": {
                                            "type": "string"
                                        },
                                        "trip_uuid": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "mo_id": {
                                    "type": "string"
                                },
                                "start_altitude": {
                                    "type": "number"
                                },
                                "start_latitude": {
                                    "type": "number"
                                },
                                "start_longitude": {
                                    "type": "number"
                                },
                                "start_time": {
                                    "type": "number"
                                },
                                "trip_features": {
                                    "items": {
                                        "feature_name": {
                                            "type": "string"
                                        },
                                        "feature_value": {
                                            "type": "string"
                                        },
                                        "id": {
                                            "type": "number"
                                        },
                                        "tenant_id": {
                                            "type": "string"
                                        },
                                        "trip_uuid": {
                                            "type": "string"
                                        }
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "trip_id": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "lastEvent": {
                    "description": "function and string parameter that created this state object",
                    "properties": {
                        "args": {
                            "items": {
                                "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "function": {
                            "description": "function that created this state object",
                            "type": "string"
                        },
                        "redirectedFromFunction": {
                            "description": "function that originally received the event",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "lastKnownLocation": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "timestamp": {
                    "description": "Event timestamp.",
                    "type": "string"
                },
                "txntimestamp": {
                    "description": "Transaction timestamp matching that in the blockchain.",
                    "type": "string"
                },
                "txnuuid": {
                    "description": "Transaction UUID matching that in the blockchain.",
                    "type": "string"
                },
                "vehicleID": {
                    "properties": {
                        "driverID": {
                            "items": {
                                "licenceNumber": {
                                    "type": "string"
                                },
                                "licenceValidityFrom": {
                                    "type": "string"
                                },
                                "licenceValidityTo": {
                                    "type": "string"
                                },
                                "name": {
                                    "type": "string"
                                }
                            },
                            "minItems": 1,
                            "type": "array"
                        },
                        "vin": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        }
    }
}`
package version

const dataTable = `[
  {
    "vid": 1,
    "name": "LIFX",
    "defaults": {
      "hev": false,
      "color": false,
      "chain": false,
      "matrix": false,
      "relays": false,
      "buttons": false,
      "infrared": false,
      "multizone": false,
      "temperature_range": null,
      "extended_multizone": false
    },
    "products": [
      {
        "pid": 1,
        "name": "LIFX Original 1000",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 3,
        "name": "LIFX Color 650",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 10,
        "name": "LIFX White 800 (Low Voltage)",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            6500
          ]
        },
        "upgrades": []
      },
      {
        "pid": 11,
        "name": "LIFX White 800 (High Voltage)",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            6500
          ]
        },
        "upgrades": []
      },
      {
        "pid": 15,
        "name": "LIFX Color 1000",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 18,
        "name": "LIFX White 900 BR30 (Low Voltage)",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 19,
        "name": "LIFX White 900 BR30 (High Voltage)",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 20,
        "name": "LIFX Color 1000 BR30",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 22,
        "name": "LIFX Color 1000",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 27,
        "name": "LIFX A19",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 28,
        "name": "LIFX BR30",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 29,
        "name": "LIFX A19 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 30,
        "name": "LIFX BR30 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 31,
        "name": "LIFX Z",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": true,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 32,
        "name": "LIFX Z",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": true,
          "temperature_range": [
            2500,
            9000
          ],
          "min_ext_mz_firmware": 1532997580,
          "min_ext_mz_firmware_components": [
            2,
            77
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 77,
            "features": {
              "extended_multizone": true
            }
          },
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 36,
        "name": "LIFX Downlight",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 37,
        "name": "LIFX Downlight",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 38,
        "name": "LIFX Beam",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": true,
          "temperature_range": [
            2500,
            9000
          ],
          "min_ext_mz_firmware": 1532997580,
          "min_ext_mz_firmware_components": [
            2,
            77
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 77,
            "features": {
              "extended_multizone": true
            }
          },
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 39,
        "name": "LIFX Downlight White to Warm",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 40,
        "name": "LIFX Downlight",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 43,
        "name": "LIFX A19",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 44,
        "name": "LIFX BR30",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 45,
        "name": "LIFX A19 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 46,
        "name": "LIFX BR30 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": [
          {
            "major": 2,
            "minor": 80,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 49,
        "name": "LIFX Mini Color",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 50,
        "name": "LIFX Mini White to Warm",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            6500
          ]
        },
        "upgrades": [
          {
            "major": 3,
            "minor": 70,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 51,
        "name": "LIFX Mini White",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            2700
          ]
        },
        "upgrades": []
      },
      {
        "pid": 52,
        "name": "LIFX GU10",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 53,
        "name": "LIFX GU10",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 55,
        "name": "LIFX Tile",
        "features": {
          "color": true,
          "chain": true,
          "matrix": true,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 57,
        "name": "LIFX Candle",
        "features": {
          "color": true,
          "chain": false,
          "matrix": true,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 59,
        "name": "LIFX Mini Color",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 60,
        "name": "LIFX Mini White to Warm",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            6500
          ]
        },
        "upgrades": [
          {
            "major": 3,
            "minor": 70,
            "features": {
              "temperature_range": [
                1500,
                9000
              ]
            }
          }
        ]
      },
      {
        "pid": 61,
        "name": "LIFX Mini White",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            2700
          ]
        },
        "upgrades": []
      },
      {
        "pid": 62,
        "name": "LIFX A19",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 63,
        "name": "LIFX BR30",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 64,
        "name": "LIFX A19 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 65,
        "name": "LIFX BR30 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 66,
        "name": "LIFX Mini White",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            2700
          ]
        },
        "upgrades": []
      },
      {
        "pid": 68,
        "name": "LIFX Candle",
        "features": {
          "color": true,
          "chain": false,
          "matrix": true,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 70,
        "name": "LIFX Switch",
        "features": {
          "color": false,
          "relays": true,
          "chain": false,
          "matrix": false,
          "buttons": true,
          "infrared": false,
          "multizone": false
        },
        "upgrades": []
      },
      {
        "pid": 71,
        "name": "LIFX Switch",
        "features": {
          "color": false,
          "relays": true,
          "chain": false,
          "matrix": false,
          "buttons": true,
          "infrared": false,
          "multizone": false
        },
        "upgrades": []
      },
      {
        "pid": 81,
        "name": "LIFX Candle White to Warm",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2200,
            6500
          ]
        },
        "upgrades": []
      },
      {
        "pid": 82,
        "name": "LIFX Filament Clear",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2100,
            2100
          ]
        },
        "upgrades": []
      },
      {
        "pid": 85,
        "name": "LIFX Filament Amber",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2000,
            2000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 87,
        "name": "LIFX Mini White",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            2700
          ]
        },
        "upgrades": []
      },
      {
        "pid": 88,
        "name": "LIFX Mini White",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2700,
            2700
          ]
        },
        "upgrades": []
      },
      {
        "pid": 89,
        "name": "LIFX Switch",
        "features": {
          "color": false,
          "relays": true,
          "chain": false,
          "matrix": false,
          "buttons": true,
          "infrared": false,
          "multizone": false
        },
        "upgrades": []
      },
      {
        "pid": 90,
        "name": "LIFX Clean",
        "features": {
          "hev": true,
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 91,
        "name": "LIFX Color",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 92,
        "name": "LIFX Color",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 93,
        "name": "LIFX A19 US",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 94,
        "name": "LIFX BR30",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 96,
        "name": "LIFX Candle White to Warm",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2200,
            6500
          ]
        },
        "upgrades": []
      },
      {
        "pid": 97,
        "name": "LIFX A19",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 98,
        "name": "LIFX BR30",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 99,
        "name": "LIFX Clean",
        "features": {
          "hev": true,
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 100,
        "name": "LIFX Filament Clear",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2100,
            2100
          ]
        },
        "upgrades": []
      },
      {
        "pid": 101,
        "name": "LIFX Filament Amber",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            2000,
            2000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 109,
        "name": "LIFX A19 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 110,
        "name": "LIFX BR30 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 111,
        "name": "LIFX A19 Night Vision",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 112,
        "name": "LIFX BR30 Night Vision Intl",
        "features": {
          "color": true,
          "chain": false,
          "matrix": false,
          "infrared": true,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 113,
        "name": "LIFX Mini WW US",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      },
      {
        "pid": 114,
        "name": "LIFX Mini WW Intl",
        "features": {
          "color": false,
          "chain": false,
          "matrix": false,
          "infrared": false,
          "multizone": false,
          "temperature_range": [
            1500,
            9000
          ]
        },
        "upgrades": []
      }
    ]
  }
]
`

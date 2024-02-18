# go-tile-server
[![MIT License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/philipp-meier/go-tile-server/blob/main/LICENSE)  
Simple tile server for `.mbtiles`-files written in Go.  
Based on the [C# tile-server](https://github.com/philipp-meier/tile-server) repository.

## Quickstart
Run `go run .` to start the application.  
Open `index.html` in the browser to get an interactive [leaflet](https://leafletjs.com/) map for your tiles.

## Stress testing
Install [k6](https://k6.io/docs/get-started/installation/) by Grafana Labs and run the following command: `k6 run test/k6.js`

## Statistics
Statistics from a local [bombardier](https://github.com/codesenberg/bombardier) benchmark:
```
./bombardier http://localhost:5296/tiles/5/15/19 -d 120s           
Bombarding http://localhost:5296/tiles/5/15/19 for 2m0s using 125 connection(s)
[=============================================================================================================================================================================================================================] 2m0s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     20621.63    3287.19   30625.50
  Latency        6.06ms     2.85ms   110.69ms
  HTTP codes:
    1xx - 0, 2xx - 2474205, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   399.18MB/s
```

## Used mbtiles file
https://ftp.gwdg.de/pub/misc/openstreetmap/openandromaps/world/OAM-World-1-8-min-J80.mbtiles

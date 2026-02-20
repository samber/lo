# Benchmark

## Summary

Benchmarks show that running SIMD operators on small datasets is slower:

```txt
BenchmarkSumInt8/small/Fallback-lo-4             203616572        5.875 ns/op
BenchmarkSumInt8/small/SSE-x16-4                 100000000        12.04 ns/op
BenchmarkSumInt8/small/AVX2-x32-4                 64041816        17.93 ns/op
BenchmarkSumInt8/small/AVX512-x64-4               26947528        44.75 ns/op
```

But much much faster on big datasets:

```txt
BenchmarkSumInt8/xlarge/Fallback-lo-4               247677       4860 ns/op
BenchmarkSumInt8/xlarge/SSE-x16-4                  3851040      311.4 ns/op
BenchmarkSumInt8/xlarge/AVX2-x32-4                 7100002      169.2 ns/op
BenchmarkSumInt8/xlarge/AVX512-x64-4              10107534      118.1 ns/op
```

## Run

```bash
export GOEXPERIMENT=simd
cd exp/simd/
go test -bench ./... -run=^Benchmark -benchmem -bench
```

## Result

| Benchmark                                      | Iterations | Time/op     | Bytes/op | Allocs/op   |
| ---------------------------------------------- | ---------- | ----------- | -------- | ----------- |
| BenchmarkContainsInt8/tiny/SSE-x16-4           | 217810574  | 5.519 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/tiny/AVX2-x32-4          | 232276010  | 5.133 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/tiny/AVX512-x64-4        | 256556498  | 4.689 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/small/SSE-x16-4          | 369299803  | 3.287 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/small/AVX2-x32-4         | 127514481  | 9.329 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/small/AVX512-x64-4       | 122257533  | 9.789 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/medium/SSE-x16-4         | 202462611  | 6.115 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/medium/AVX2-x32-4        | 287002275  | 4.223 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/medium/AVX512-x64-4      | 382635258  | 3.125 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/large/SSE-x16-4          | 270810375  | 4.425 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/large/AVX2-x32-4         | 316008705  | 3.796 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/large/AVX512-x64-4       | 447781525  | 2.689 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/xlarge/SSE-x16-4         | 120643213  | 10.02 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/xlarge/AVX2-x32-4        | 172479368  | 6.773 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/xlarge/AVX512-x64-4      | 268865846  | 4.469 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/massive/SSE-x16-4        | 189045388  | 6.332 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/massive/AVX2-x32-4       | 258672679  | 4.686 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/massive/AVX512-x64-4     | 311598692  | 3.860 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/tiny/SSE-x8-4           | 273258456  | 4.349 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/tiny/AVX2-x16-4         | 273500064  | 4.365 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/tiny/AVX512-x32-4       | 263483358  | 4.543 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/small/SSE-x8-4          | 273197698  | 4.322 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/small/AVX2-x16-4        | 355109746  | 3.365 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/small/AVX512-x32-4      | 121548804  | 9.874 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/medium/SSE-x8-4         | 126075469  | 9.524 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/medium/AVX2-x16-4       | 183189763  | 6.075 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/medium/AVX512-x32-4     | 273689677  | 4.350 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/large/SSE-x8-4          | 38848131   | 30.83 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/large/AVX2-x16-4        | 71321032   | 16.56 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/large/AVX512-x32-4      | 132036739  | 9.144 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/xlarge/SSE-x8-4         | 10515756   | 113.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/xlarge/AVX2-x16-4       | 20209566   | 58.50 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/xlarge/AVX512-x32-4     | 43699071   | 28.06 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/massive/SSE-x8-4        | 1336891    | 896.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/massive/AVX2-x16-4      | 2646244    | 453.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/massive/AVX512-x32-4    | 5799852    | 207.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/tiny/SSE-x4-4           | 352252977  | 3.332 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/tiny/AVX2-x8-4          | 233631776  | 5.113 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/tiny/AVX512-x16-4       | 262243407  | 4.596 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/small/SSE-x4-4          | 203636325  | 5.878 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/small/AVX2-x8-4         | 259371962  | 4.758 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/small/AVX512-x16-4      | 356743093  | 3.356 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/medium/SSE-x4-4         | 69843790   | 16.86 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/medium/AVX2-x8-4        | 126707199  | 9.452 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/medium/AVX512-x16-4     | 210475352  | 5.671 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/large/SSE-x4-4          | 20589945   | 58.38 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/large/AVX2-x8-4         | 39030094   | 30.56 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/large/AVX512-x16-4      | 86611118   | 14.30 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/xlarge/SSE-x4-4         | 13936316   | 86.03 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/xlarge/AVX2-x8-4        | 26901062   | 44.37 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/xlarge/AVX512-x16-4     | 60009277   | 20.75 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/massive/SSE-x4-4        | 25489658   | 46.95 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/massive/AVX2-x8-4       | 48848792   | 24.43 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/massive/AVX512-x16-4    | 100000000  | 11.31 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/tiny/SSE-x2-4           | 293674094  | 4.107 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/tiny/AVX2-x4-4          | 373632538  | 3.222 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/tiny/AVX512-x8-4        | 244335542  | 4.950 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/small/SSE-x2-4          | 129728997  | 9.246 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/small/AVX2-x4-4         | 206096462  | 5.800 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/small/AVX512-x8-4       | 300555229  | 4.008 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/medium/SSE-x2-4         | 38974218   | 30.69 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/medium/AVX2-x4-4        | 72598192   | 16.18 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/medium/AVX512-x8-4      | 147323280  | 8.131 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/large/SSE-x2-4          | 10538320   | 113.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/large/AVX2-x4-4         | 20545034   | 58.38 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/large/AVX512-x8-4       | 46294234   | 25.73 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/xlarge/SSE-x2-4         | 2656102    | 451.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/xlarge/AVX2-x4-4        | 5203542    | 230.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/xlarge/AVX512-x8-4      | 12310575   | 97.33 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/massive/SSE-x2-4        | 335900     | 3556 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/massive/AVX2-x4-4       | 657386     | 1783 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/massive/AVX512-x8-4     | 1493280    | 796.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/tiny/SSE-x16-4          | 210596146  | 5.700 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/tiny/AVX2-x32-4         | 230931388  | 5.197 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/tiny/AVX512-x64-4       | 242121343  | 4.886 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/small/SSE-x16-4         | 377043338  | 3.187 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/small/AVX2-x32-4        | 129512683  | 9.253 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/small/AVX512-x64-4      | 126131306  | 9.521 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/medium/SSE-x16-4        | 189831553  | 6.381 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/medium/AVX2-x32-4       | 292308152  | 4.071 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/medium/AVX512-x64-4     | 366734703  | 3.270 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/large/SSE-x16-4         | 60894787   | 19.18 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/large/AVX2-x32-4        | 122211781  | 9.905 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/large/AVX512-x64-4      | 201761496  | 6.009 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/xlarge/SSE-x16-4        | 38853038   | 35.79 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/xlarge/AVX2-x32-4       | 78938445   | 15.22 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/xlarge/AVX512-x64-4     | 148406178  | 8.044 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/massive/SSE-x16-4       | 68406747   | 16.96 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/massive/AVX2-x32-4      | 124138452  | 10.03 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/massive/AVX512-x64-4    | 224304250  | 5.350 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/tiny/SSE-x8-4          | 218804454  | 5.491 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/tiny/AVX2-x16-4        | 218739122  | 5.498 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/tiny/AVX512-x32-4      | 213586394  | 5.610 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/small/SSE-x8-4         | 285711435  | 4.074 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/small/AVX2-x16-4       | 376051632  | 3.225 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/small/AVX512-x32-4     | 122171212  | 9.828 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/medium/SSE-x8-4        | 129853016  | 9.255 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/medium/AVX2-x16-4      | 207015410  | 5.799 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/medium/AVX512-x32-4    | 283134154  | 4.244 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/large/SSE-x8-4         | 39025154   | 30.73 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/large/AVX2-x16-4       | 73705192   | 16.19 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/large/AVX512-x32-4     | 127408450  | 9.394 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/xlarge/SSE-x8-4        | 10563104   | 113.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/xlarge/AVX2-x16-4      | 20532660   | 58.34 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/xlarge/AVX512-x32-4    | 39390879   | 30.43 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/massive/SSE-x8-4       | 1338374    | 896.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/massive/AVX2-x16-4     | 2642508    | 453.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/massive/AVX512-x32-4   | 5214861    | 232.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/tiny/SSE-x4-4          | 338625207  | 3.526 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/tiny/AVX2-x8-4         | 234937918  | 5.069 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/tiny/AVX512-x16-4      | 260805802  | 4.607 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/small/SSE-x4-4         | 196105390  | 6.078 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/small/AVX2-x8-4        | 271456726  | 4.388 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/small/AVX512-x16-4     | 356633224  | 3.328 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/medium/SSE-x4-4        | 70967472   | 16.66 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/medium/AVX2-x8-4       | 126340971  | 9.402 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/medium/AVX512-x16-4    | 208789647  | 5.764 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/large/SSE-x4-4         | 26834684   | 44.41 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/large/AVX2-x8-4        | 50715457   | 23.52 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/large/AVX512-x16-4     | 100000000  | 11.20 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/xlarge/SSE-x4-4        | 5199847    | 230.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/xlarge/AVX2-x8-4       | 10525443   | 113.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/xlarge/AVX512-x16-4    | 23942347   | 49.88 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/massive/SSE-x4-4       | 670052     | 1782 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/massive/AVX2-x8-4      | 1338936    | 902.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/massive/AVX512-x16-4   | 3023478    | 395.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/tiny/SSE-x2-4          | 281334750  | 4.100 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/tiny/AVX2-x4-4         | 367084185  | 3.218 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/tiny/AVX512-x8-4       | 237374930  | 4.888 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/small/SSE-x2-4         | 129761154  | 9.240 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/small/AVX2-x4-4        | 207136089  | 5.813 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/small/AVX512-x8-4      | 300139946  | 4.198 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/medium/SSE-x2-4        | 39123123   | 30.62 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/medium/AVX2-x4-4       | 72446277   | 16.55 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/medium/AVX512-x8-4     | 146686130  | 8.231 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/large/SSE-x2-4         | 10551722   | 113.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/large/AVX2-x4-4        | 20594492   | 58.39 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/large/AVX512-x8-4      | 46978982   | 25.74 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/xlarge/SSE-x2-4        | 2654188    | 452.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/xlarge/AVX2-x4-4       | 5206354    | 230.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/xlarge/AVX512-x8-4     | 12328929   | 97.16 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/massive/SSE-x2-4       | 336739     | 3619 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/massive/AVX2-x4-4      | 664974     | 1782 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/massive/AVX512-x8-4    | 1508221    | 795.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/tiny/SSE-x4-4         | 343127668  | 3.443 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/tiny/AVX2-x8-4        | 268992757  | 4.437 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/tiny/AVX512-x16-4     | 217254800  | 5.542 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/small/SSE-x4-4        | 212643310  | 5.718 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/small/AVX2-x8-4       | 289509478  | 4.171 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/small/AVX512-x16-4    | 312890053  | 3.868 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/medium/SSE-x4-4       | 80126814   | 15.40 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/medium/AVX2-x8-4      | 131377432  | 9.139 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/medium/AVX512-x16-4   | 192624958  | 6.225 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/large/SSE-x4-4        | 24038194   | 49.75 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/large/AVX2-x8-4       | 42452012   | 28.41 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/large/AVX512-x16-4    | 74948697   | 16.26 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/xlarge/SSE-x4-4       | 5967651    | 199.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/xlarge/AVX2-x8-4      | 11176159   | 107.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/xlarge/AVX512-x16-4   | 21256507   | 53.30 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/massive/SSE-x4-4      | 794624     | 1511 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/massive/AVX2-x8-4     | 1368274    | 888.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/massive/AVX512-x16-4  | 2972845    | 401.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/tiny/SSE-x2-4         | 281273576  | 4.349 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/tiny/AVX2-x4-4        | 344326042  | 3.407 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/tiny/AVX512-x8-4      | 212695608  | 5.642 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/small/SSE-x2-4        | 131755514  | 9.074 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/small/AVX2-x4-4       | 203556692  | 5.893 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/small/AVX512-x8-4     | 257131224  | 4.619 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/medium/SSE-x2-4       | 42138819   | 27.86 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/medium/AVX2-x4-4      | 78823770   | 15.68 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/medium/AVX512-x8-4    | 126535322  | 9.443 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/large/SSE-x2-4        | 12221720   | 97.53 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/large/AVX2-x4-4       | 22674674   | 52.34 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/large/AVX512-x8-4     | 42609235   | 28.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/xlarge/SSE-x2-4       | 2720515    | 440.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/xlarge/AVX2-x4-4      | 5717328    | 202.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/xlarge/AVX512-x8-4    | 11988493   | 106.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/massive/SSE-x2-4      | 350553     | 3543 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/massive/AVX2-x4-4     | 693710     | 1753 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/massive/AVX512-x8-4   | 1282303    | 936.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsWorstCase/SSE-x4-4            | 5161525    | 232.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsWorstCase/AVX2-x8-4           | 10623906   | 112.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsWorstCase/AVX512-x16-4        | 24401241   | 49.09 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsBestCase/SSE-x4-4             | 402412988  | 2.970 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsBestCase/AVX2-x8-4            | 402852975  | 2.917 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsBestCase/AVX512-x16-4         | 404357092  | 2.971 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/tiny/SSE-x4-4        | 373490160  | 3.199 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/tiny/AVX2-x8-4       | 217269805  | 5.487 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/tiny/AVX512-x16-4    | 243944257  | 4.972 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/small/SSE-x4-4       | 206713904  | 5.778 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/small/AVX2-x8-4      | 286690113  | 4.179 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/small/AVX512-x16-4   | 362335026  | 3.315 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/medium/SSE-x4-4      | 72029425   | 16.53 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/medium/AVX2-x8-4     | 127715524  | 9.247 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/medium/AVX512-x16-4  | 201068770  | 6.009 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/large/SSE-x4-4       | 20579641   | 58.85 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/large/AVX2-x8-4      | 38948685   | 30.72 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/large/AVX512-x16-4   | 83767486   | 14.73 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/xlarge/SSE-x4-4      | 5193009    | 230.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/xlarge/AVX2-x8-4     | 10552609   | 113.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/xlarge/AVX512-x16-4  | 24272108   | 49.37 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/massive/SSE-x4-4     | 670375     | 1781 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/massive/AVX2-x8-4    | 1339380    | 896.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/massive/AVX512-x16-4 | 3035526    | 395.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8ByWidth/SSE-x16-4         | 91492479   | 13.20 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8ByWidth/AVX2-x32-4        | 140600484  | 8.452 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8ByWidth/AVX512-x64-4      | 221681901  | 5.441 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64SteadyState/SSE-x2-4     | 619605     | 1945 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64SteadyState/AVX2-x4-4    | 1238368    | 968.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64SteadyState/AVX512-x8-4  | 2807049    | 427.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/Fallback-lo-4           | 203616572  | 5.875 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/SSE-x16-4               | 100000000  | 12.04 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/AVX2-x32-4              | 64041816   | 17.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/AVX512-x64-4            | 26947528   | 44.75 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/Fallback-lo-4          | 15706714   | 76.44 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/SSE-x16-4              | 97807059   | 12.27 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/AVX2-x32-4             | 59851663   | 16.78 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/AVX512-x64-4           | 27894214   | 43.09 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/Fallback-lo-4           | 1972309    | 608.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/SSE-x16-4               | 26922098   | 52.48 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/AVX2-x32-4              | 37505374   | 31.81 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/AVX512-x64-4            | 23313915   | 52.22 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/Fallback-lo-4          | 247677     | 4860 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/SSE-x16-4              | 3851040    | 311.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/AVX2-x32-4             | 7100002    | 169.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/AVX512-x64-4           | 10107534   | 118.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/Fallback-lo-4          | 202912227  | 5.889 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/SSE-x8-4               | 145487781  | 8.287 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/AVX2-x16-4             | 75542098   | 15.60 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/AVX512-x32-4           | 71501834   | 16.70 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/Fallback-lo-4         | 15750567   | 76.40 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/SSE-x8-4              | 67642022   | 17.71 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/AVX2-x16-4            | 64991774   | 18.42 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/AVX512-x32-4          | 69041605   | 17.01 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/Fallback-lo-4          | 1969231    | 610.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/SSE-x8-4               | 13688154   | 86.79 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/AVX2-x16-4             | 15430106   | 78.01 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/AVX512-x32-4           | 37842558   | 31.85 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/Fallback-lo-4         | 250257     | 4815 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/SSE-x8-4              | 1948285    | 635.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/AVX2-x16-4            | 2571985    | 467.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/AVX512-x32-4          | 6853812    | 174.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/Fallback-lo-4          | 203750773  | 5.886 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/SSE-x4-4               | 171269833  | 6.978 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/AVX2-x8-4              | 136651821  | 8.805 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/AVX512-x16-4           | 74063526   | 16.13 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/Fallback-lo-4         | 15769312   | 76.37 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/SSE-x4-4              | 37296025   | 37.70 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/AVX2-x8-4             | 59583854   | 20.17 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/AVX512-x16-4          | 71383305   | 16.83 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/Fallback-lo-4          | 1971070    | 608.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/SSE-x4-4               | 5084606    | 237.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/AVX2-x8-4              | 10197510   | 117.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/AVX512-x16-4           | 16171777   | 73.91 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/Fallback-lo-4         | 249922     | 4823 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/SSE-x4-4              | 656119     | 1787 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/AVX2-x8-4             | 1325719    | 906.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/AVX512-x16-4          | 2637379    | 454.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/Fallback-lo-4          | 340475902  | 3.538 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/SSE-x2-4               | 159884966  | 7.505 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/AVX2-x4-4              | 160129171  | 7.488 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/AVX512-x8-4            | 133788973  | 8.975 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/Fallback-lo-4         | 30947872   | 38.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/SSE-x2-4              | 27422982   | 44.66 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/AVX2-x4-4             | 37786063   | 38.13 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/AVX512-x8-4           | 70712422   | 17.16 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/Fallback-lo-4          | 3736202    | 321.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/SSE-x2-4               | 3839794    | 314.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/AVX2-x4-4              | 5031933    | 238.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/AVX512-x8-4            | 9745372    | 111.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/Fallback-lo-4         | 497773     | 2426 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/SSE-x2-4              | 411196     | 2718 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/AVX2-x4-4             | 651866     | 1794 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/AVX512-x8-4           | 1311506    | 915.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/Fallback-lo-4        | 338983024  | 3.554 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/SSE-x4-4             | 171496161  | 6.985 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/AVX2-x8-4            | 123336990  | 9.747 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/AVX512-x16-4         | 39568072   | 30.27 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/Fallback-lo-4       | 30962536   | 38.98 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/SSE-x4-4            | 30815191   | 37.90 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/AVX2-x8-4           | 59834246   | 19.99 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/AVX512-x16-4        | 58653121   | 20.13 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/Fallback-lo-4        | 3731205    | 321.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/SSE-x4-4             | 4324430    | 277.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/AVX2-x8-4            | 9017343    | 133.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/AVX512-x16-4         | 14832765   | 80.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/Fallback-lo-4       | 496615     | 2429 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/SSE-x4-4            | 510874     | 2349 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/AVX2-x8-4           | 1000000    | 1167 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/AVX512-x16-4        | 1989156    | 603.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/Fallback-lo-4        | 339932667  | 3.528 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/SSE-x2-4             | 137921846  | 8.816 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/AVX2-x4-4            | 163056916  | 7.298 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/AVX512-x8-4          | 100000000  | 11.19 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/Fallback-lo-4       | 31006650   | 38.94 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/SSE-x2-4            | 20184583   | 59.12 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/AVX2-x4-4           | 39255891   | 35.84 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/AVX512-x8-4         | 59392477   | 20.05 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/Fallback-lo-4        | 3734490    | 322.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/SSE-x2-4             | 2101627    | 570.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/AVX2-x4-4            | 4326589    | 277.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/AVX512-x8-4          | 8676978    | 138.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/Fallback-lo-4       | 496094     | 2436 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/SSE-x2-4            | 253662     | 4712 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/AVX2-x4-4           | 508648     | 2403 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/AVX512-x8-4         | 987357     | 1206 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/Fallback-lo-4         | 313575626  | 3.821 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/SSE-x4-4              | 142947225  | 8.401 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/AVX2-x8-4             | 100000000  | 10.23 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/AVX512-x16-4          | 67735402   | 17.57 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/Fallback-lo-4        | 30579666   | 39.33 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/SSE-x4-4             | 35956716   | 33.27 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/AVX2-x8-4            | 55783533   | 21.60 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/AVX512-x16-4         | 67622037   | 17.81 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/Fallback-lo-4         | 3729078    | 322.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/SSE-x4-4              | 5028604    | 239.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/AVX2-x8-4             | 10095586   | 119.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/AVX512-x16-4          | 19380102   | 61.84 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/Fallback-lo-4        | 499718     | 2433 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/SSE-x4-4             | 665932     | 1791 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/AVX2-x8-4            | 1324820    | 905.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/AVX512-x16-4         | 2617774    | 458.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/Fallback-lo-4       | 268642384  | 4.463 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/SSE-x2-4            | 121571077  | 9.921 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/AVX2-x4-4           | 128876829  | 9.288 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/AVX512-x8-4         | 72002227   | 16.78 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/Fallback-lo-4      | 29774340   | 40.35 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/SSE-x2-4           | 17057521   | 72.58 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/AVX2-x4-4          | 37395951   | 32.44 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/AVX512-x8-4        | 48551912   | 24.65 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/Fallback-lo-4       | 3679522    | 326.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/SSE-x2-4            | 2097200    | 572.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/AVX2-x4-4           | 4292325    | 279.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/AVX512-x8-4         | 8150467    | 146.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/Fallback-lo-4      | 495211     | 2460 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/SSE-x2-4           | 250432     | 4859 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/AVX2-x4-4          | 507294     | 2350 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/AVX512-x8-4        | 984428     | 1212 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/small/SSE-x4-4               | 244127626  | 4.923 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/small/AVX2-x8-4              | 239254977  | 4.876 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/small/AVX512-x16-4           | 100000000  | 11.72 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/medium/SSE-x4-4              | 36288748   | 32.21 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/medium/AVX2-x8-4             | 58453642   | 20.07 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/medium/AVX512-x16-4          | 71356076   | 16.79 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/large/SSE-x4-4               | 4519351    | 265.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/large/AVX2-x8-4              | 8000348    | 150.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/large/AVX512-x16-4           | 18772939   | 63.76 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/xlarge/SSE-x4-4              | 598405     | 2015 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/xlarge/AVX2-x8-4             | 1000000    | 1187 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/xlarge/AVX512-x16-4          | 2546474    | 471.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/small/SSE-x2-4             | 178538296  | 6.742 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/small/AVX2-x4-4            | 202859059  | 5.914 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/small/AVX512-x8-4          | 58871919   | 20.32 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/medium/SSE-x2-4            | 18199695   | 65.97 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/medium/AVX2-x4-4           | 32141115   | 37.84 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/medium/AVX512-x8-4         | 40255204   | 29.90 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/large/SSE-x2-4             | 2050605    | 584.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/large/AVX2-x4-4            | 4049390    | 294.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/large/AVX512-x8-4          | 4924977    | 243.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/xlarge/SSE-x2-4            | 253034     | 8193 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/xlarge/AVX2-x4-4           | 506792     | 2393 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/xlarge/AVX512-x8-4         | 528156     | 2244 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/small/SSE-x4-4               | 230192032  | 5.196 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/small/AVX2-x8-4              | 237162885  | 5.025 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/small/AVX512-x16-4           | 90529538   | 13.11 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/medium/SSE-x4-4              | 32832202   | 35.26 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/medium/AVX2-x8-4             | 66369580   | 18.03 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/medium/AVX512-x16-4          | 70957309   | 16.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/large/SSE-x4-4               | 4232040    | 291.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/large/AVX2-x8-4              | 8597635    | 138.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/large/AVX512-x16-4           | 15311949   | 78.88 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/xlarge/SSE-x4-4              | 508044     | 2184 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/xlarge/AVX2-x8-4             | 1000000    | 1040 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/xlarge/AVX512-x16-4          | 1962882    | 613.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/small/SSE-x2-4             | 173400889  | 6.905 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/small/AVX2-x4-4            | 122223142  | 9.815 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/small/AVX512-x8-4          | 43791807   | 27.39 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/medium/SSE-x2-4            | 15801436   | 75.77 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/medium/AVX2-x4-4           | 28346755   | 40.54 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/medium/AVX512-x8-4         | 31498221   | 39.45 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/large/SSE-x2-4             | 2002983    | 599.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/large/AVX2-x4-4            | 3938358    | 304.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/large/AVX512-x8-4          | 4483813    | 267.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/xlarge/SSE-x2-4            | 252200     | 4745 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/xlarge/AVX2-x4-4           | 503626     | 2373 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/xlarge/AVX512-x8-4         | 609382     | 2056 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/Fallback-lo-4          | 685263     | 1756 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/SSE-x16-4              | 7270306    | 165.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/AVX2-x32-4             | 12081855   | 98.97 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/AVX512-x64-4           | 14816949   | 80.89 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/Fallback-lo-4     | 496068     | 2430 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/SSE-x2-4          | 460347     | 2634 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/AVX2-x4-4         | 664824     | 1792 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/AVX512-x8-4       | 1311258    | 914.6 ns/op | 0 B/op   | 0 allocs/op |

```

goos: linux
goarch: amd64
pkg: github.com/samber/lo/exp/simd
cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
PASS
ok      github.com/samber/lo/exp/simd        596.328s
```

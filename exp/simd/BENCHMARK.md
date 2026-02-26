# Benchmark

## Summary

Benchmarks show that running SIMD operations on small datasets is slower:

```txt
BenchmarkSumInt8/small/Fallback-lo-2            	248740710	        5.218 ns/op
BenchmarkSumInt8/small/AVX-x16-2                	126181464	        9.485 ns/op
BenchmarkSumInt8/small/AVX2-x32-2               	 73059427	        14.44 ns/op
BenchmarkSumInt8/small/AVX512-x64-2             	 49913169	        24.41 ns/op
```

But SIMD is much faster on large datasets:

```txt
BenchmarkSumInt8/xlarge/Fallback-lo-2           	  273898	         4383 ns/op
BenchmarkSumInt8/xlarge/AVX-x16-2               	 6928408	        173.1 ns/op
BenchmarkSumInt8/xlarge/AVX2-x32-2              	12639586	        94.09 ns/op
BenchmarkSumInt8/xlarge/AVX512-x64-2            	13509693	        89.67 ns/op
```

## Run

```bash
export GOEXPERIMENT=simd
cd exp/simd/
go test -bench ./... -run=^Benchmark -benchmem -bench
```

```bash
# get instruction set
cat /proc/cpuinfo
```

## Result

```
archsimd.X86: AVX=true AVX2=true AVX512=true
goos: linux
goarch: amd64
pkg: github.com/samber/lo/exp/simd
cpu: AMD EPYC 9454P 48-Core Processor    

...

PASS
ok  	github.com/samber/lo/exp/simd	596.213s
```

| Benchmark                                      | Iterations | Time/op     | Bytes/op | Allocs/op   |
| ---------------------------------------------- | ---------- | ----------- | -------- | ----------- |
| BenchmarkContainsInt8/tiny/AVX512-x16-2        | 312359204  | 3.625 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/tiny/AVX512-x32-2        | 277194441  | 4.531 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/tiny/AVX512-x64-2        | 336853209  | 3.401 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/small/AVX512-x16-2       | 449132103  | 2.670 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/small/AVX512-x32-2       | 148648339  | 8.332 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/small/AVX512-x64-2       | 143124861  | 7.982 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/medium/AVX512-x16-2      | 276816714  | 4.302 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/medium/AVX512-x32-2      | 345774957  | 3.529 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/medium/AVX512-x64-2      | 449868722  | 2.669 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/large/AVX512-x16-2          | 100000000  | 10.68 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/large/AVX512-x32-2         | 172934200  | 6.941 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/large/AVX512-x64-2       | 280992625  | 4.384 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/xlarge/AVX512-x16-2         | 187189599  | 6.203 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/xlarge/AVX512-x32-2        | 274289563  | 4.042 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/xlarge/AVX512-x64-2      | 375048555  | 2.953 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/massive/AVX512-x16-2        | 86434948   | 14.02 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/massive/AVX512-x32-2       | 153742346  | 8.012 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8/massive/AVX512-x64-2     | 259404483  | 5.214 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/tiny/AVX512-x8-2           | 270309470  | 4.315 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/tiny/AVX512-x16-2         | 264874646  | 4.281 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/tiny/AVX512-x32-2       | 328810479  | 3.593 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/small/AVX512-x8-2          | 374742561  | 3.206 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/small/AVX512-x16-2        | 449838870  | 2.678 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/small/AVX512-x32-2      | 143845734  | 8.484 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/medium/AVX512-x8-2         | 185415590  | 6.448 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/medium/AVX512-x16-2       | 273780868  | 4.268 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/medium/AVX512-x32-2     | 350067484  | 3.431 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/large/AVX512-x8-2       | 61109778   | 19.66 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/large/AVX512-x16-2      | 100000000  | 10.74 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/large/AVX512-x32-2      | 182886646  | 6.575 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/xlarge/AVX512-x8-2      | 15220682   | 71.53 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/xlarge/AVX512-x16-2     | 31876572   | 37.57 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/xlarge/AVX512-x32-2     | 61992217   | 19.55 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/massive/AVX512-x8-2     | 4372000    | 262.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/massive/AVX512-x16-2    | 9019658    | 131.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt16/massive/AVX512-x32-2    | 16568430   | 74.25 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/tiny/AVX512-x4-2        | 499209442  | 2.406 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/tiny/AVX512-x8-2        | 350479609  | 3.433 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/tiny/AVX512-x16-2       | 280918554  | 4.309 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/small/AVX512-x4-2       | 299561596  | 4.028 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/small/AVX512-x8-2       | 374064310  | 3.205 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/small/AVX512-x16-2      | 499219765  | 2.418 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/medium/AVX512-x4-2      | 100000000  | 10.42 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/medium/AVX512-x8-2      | 187391635  | 6.403 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/medium/AVX512-x16-2     | 307955800  | 3.875 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/large/AVX512-x4-2       | 33256420   | 36.05 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/large/AVX512-x8-2       | 62421526   | 19.23 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/large/AVX512-x16-2      | 100000000  | 10.36 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/xlarge/AVX512-x4-2      | 8328856    | 144.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/xlarge/AVX512-x8-2      | 17039037   | 71.14 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/xlarge/AVX512-x16-2     | 28740241   | 41.77 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/massive/AVX512-x4-2     | 3525885    | 332.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/massive/AVX512-x8-2     | 7318027    | 164.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt32/massive/AVX512-x16-2    | 12181366   | 99.08 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/tiny/AVX512-x2-2        | 409014308  | 2.934 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/tiny/AVX512-x4-2        | 449210791  | 2.667 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/tiny/AVX512-x8-2        | 280998146  | 4.293 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/small/AVX512-x2-2       | 195631429  | 6.172 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/small/AVX512-x4-2       | 281272394  | 4.308 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/small/AVX512-x8-2       | 408933924  | 3.044 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/medium/AVX512-x2-2      | 63006909   | 18.94 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/medium/AVX512-x4-2      | 100000000  | 10.67 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/medium/AVX512-x8-2      | 197411126  | 6.016 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/large/AVX512-x2-2       | 17098578   | 70.57 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/large/AVX512-x4-2       | 32558013   | 37.07 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/large/AVX512-x8-2       | 57629485   | 20.94 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/xlarge/AVX512-x2-2      | 4286155    | 281.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/xlarge/AVX512-x4-2      | 8344772    | 143.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/xlarge/AVX512-x8-2      | 14428276   | 83.14 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/massive/AVX512-x2-2     | 1000000    | 1012 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/massive/AVX512-x4-2     | 2350525    | 510.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64/massive/AVX512-x8-2     | 3773523    | 318.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/tiny/AVX512-x16-2       | 338880315  | 3.332 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/tiny/AVX512-x32-2       | 320784217  | 3.559 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/tiny/AVX512-x64-2       | 341599854  | 3.331 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/small/AVX512-x16-2      | 449579424  | 2.670 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/small/AVX512-x32-2      | 140368142  | 8.648 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/small/AVX512-x64-2      | 146828888  | 8.182 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/medium/AVX512-x16-2     | 374443974  | 3.472 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/medium/AVX512-x32-2     | 449271607  | 2.672 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/medium/AVX512-x64-2     | 598525731  | 2.018 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/large/AVX512-x16-2      | 254828565  | 4.956 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/large/AVX512-x32-2      | 407777484  | 2.938 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/large/AVX512-x64-2      | 443472316  | 2.666 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/xlarge/AVX512-x16-2     | 162196827  | 7.867 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/xlarge/AVX512-x32-2     | 268324950  | 4.518 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/xlarge/AVX512-x64-2     | 400437789  | 2.952 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/massive/AVX512-x16-2    | 214548872  | 5.640 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/massive/AVX512-x32-2    | 348431553  | 3.391 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint8/massive/AVX512-x64-2    | 459781908  | 2.455 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/tiny/AVX512-x8-2       | 276271912  | 4.297 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/tiny/AVX512-x16-2      | 281145528  | 4.270 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/tiny/AVX512-x32-2      | 315343911  | 3.667 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/small/AVX512-x8-2      | 374632351  | 3.204 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/small/AVX512-x16-2     | 449355727  | 2.670 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/small/AVX512-x32-2     | 138088146  | 8.395 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/medium/AVX512-x8-2     | 187276191  | 6.582 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/medium/AVX512-x16-2    | 281107980  | 4.306 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/medium/AVX512-x32-2    | 358850328  | 3.516 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/large/AVX512-x8-2      | 59025931   | 19.98 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/large/AVX512-x16-2     | 100000000  | 10.68 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/large/AVX512-x32-2     | 179631354  | 6.569 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/xlarge/AVX512-x8-2     | 16576267   | 71.63 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/xlarge/AVX512-x16-2    | 32578981   | 36.96 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/xlarge/AVX512-x32-2    | 61464870   | 19.44 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/massive/AVX512-x8-2    | 2153736    | 557.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/massive/AVX512-x16-2   | 4225728    | 281.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint16/massive/AVX512-x32-2   | 7829936    | 145.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/tiny/AVX512-x4-2       | 499390296  | 2.403 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/tiny/AVX512-x8-2       | 362964080  | 3.342 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/tiny/AVX512-x16-2      | 281063364  | 4.268 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/small/AVX512-x4-2      | 293867554  | 4.004 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/small/AVX512-x8-2      | 374510434  | 3.203 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/small/AVX512-x16-2     | 499714206  | 2.402 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/medium/AVX512-x4-2     | 100000000  | 10.42 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/medium/AVX512-x8-2     | 187258657  | 6.405 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/medium/AVX512-x16-2    | 312999210  | 3.881 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/large/AVX512-x4-2      | 33298366   | 36.02 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/large/AVX512-x8-2      | 62409421   | 19.23 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/large/AVX512-x16-2     | 100000000  | 10.10 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/xlarge/AVX512-x4-2     | 7948898    | 143.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/xlarge/AVX512-x8-2     | 17021738   | 70.49 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/xlarge/AVX512-x16-2    | 28742320   | 41.77 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/massive/AVX512-x4-2    | 1595774    | 751.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/massive/AVX512-x8-2    | 3094242    | 381.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint32/massive/AVX512-x16-2   | 5080051    | 238.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/tiny/AVX512-x2-2       | 374760351  | 3.203 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/tiny/AVX512-x4-2       | 498763054  | 2.419 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/tiny/AVX512-x8-2       | 319635274  | 3.582 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/small/AVX512-x2-2      | 187032452  | 6.447 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/small/AVX512-x4-2      | 299546244  | 4.009 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/small/AVX512-x8-2      | 373937659  | 3.207 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/medium/AVX512-x2-2     | 62413118   | 19.23 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/medium/AVX512-x4-2     | 113978791  | 10.42 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/medium/AVX512-x8-2     | 186965330  | 6.484 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/large/AVX512-x2-2      | 17005768   | 70.57 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/large/AVX512-x4-2      | 33286495   | 36.69 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/large/AVX512-x8-2      | 61486065   | 19.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/xlarge/AVX512-x2-2     | 4154370    | 280.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/xlarge/AVX512-x4-2     | 8371358    | 148.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/xlarge/AVX512-x8-2     | 14193795   | 72.36 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/massive/AVX512-x2-2    | 1773937    | 676.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/massive/AVX512-x4-2    | 3500168    | 343.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsUint64/massive/AVX512-x8-2    | 7097266    | 249.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/tiny/AVX512-x4-2      | 410522160  | 2.675 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/tiny/AVX512-x8-2      | 308565882  | 3.814 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/tiny/AVX512-x16-2     | 315331897  | 3.755 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/small/AVX512-x4-2     | 278219434  | 4.642 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/small/AVX512-x8-2     | 362945481  | 3.287 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/small/AVX512-x16-2    | 408523153  | 2.941 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/medium/AVX512-x4-2    | 100000000  | 10.77 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/medium/AVX512-x8-2    | 186186376  | 6.409 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/medium/AVX512-x16-2   | 264255108  | 4.619 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/large/AVX512-x4-2      | 33028701   | 36.27 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/large/AVX512-x8-2      | 62465360   | 19.53 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/large/AVX512-x16-2    | 108213310  | 10.95 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/xlarge/AVX512-x4-2    | 8359381    | 143.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/xlarge/AVX512-x8-2    | 17042701   | 70.46 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/xlarge/AVX512-x16-2   | 31806921   | 37.13 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/massive/AVX512-x4-2   | 1000000    | 1100 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/massive/AVX512-x8-2   | 2164672    | 554.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat32/massive/AVX512-x16-2  | 4201453    | 293.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/tiny/AVX512-x2-2      | 362183925  | 3.223 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/tiny/AVX512-x4-2      | 449021466  | 2.687 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/tiny/AVX512-x8-2      | 320176149  | 3.820 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/small/AVX512-x2-2     | 187139116  | 6.415 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/small/AVX512-x4-2     | 280722585  | 4.300 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/small/AVX512-x8-2     | 335670502  | 3.472 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/medium/AVX512-x2-2    | 62343927   | 19.23 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/medium/AVX512-x4-2    | 112332902  | 10.69 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/medium/AVX512-x8-2    | 179610780  | 6.741 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/large/AVX512-x2-2     | 16996959   | 70.51 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/large/AVX512-x4-2     | 33017950   | 36.29 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/large/AVX512-x8-2     | 60322328   | 19.73 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/xlarge/AVX512-x2-2    | 4141281    | 282.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/xlarge/AVX512-x4-2    | 7856590    | 145.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/xlarge/AVX512-x8-2    | 16623739   | 72.06 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/massive/AVX512-x2-2   | 541202     | 2195 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/massive/AVX512-x4-2   | 1000000    | 1158 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsFloat64/massive/AVX512-x8-2   | 2115301    | 560.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsWorstCase/AVX512-x4-2         | 7651734    | 145.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsWorstCase/AVX512-x8-2         | 14921599   | 70.49 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsWorstCase/AVX512-x16-2        | 28708478   | 41.38 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsBestCase/AVX512-x4-2          | 534237578  | 2.136 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsBestCase/AVX512-x8-2          | 561252645  | 2.159 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsBestCase/AVX512-x16-2         | 560396454  | 2.137 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/tiny/AVX512-x4-2     | 499649139  | 2.401 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/tiny/AVX512-x8-2     | 329743240  | 3.421 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/tiny/AVX512-x16-2    | 280516392  | 4.276 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/small/AVX512-x4-2    | 299373171  | 4.006 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/small/AVX512-x8-2    | 374407988  | 3.267 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/small/AVX512-x16-2   | 486948346  | 2.424 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/medium/AVX512-x4-2   | 100000000  | 10.41 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/medium/AVX512-x8-2   | 182899621  | 6.412 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/medium/AVX512-x16-2  | 311969776  | 3.829 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/large/AVX512-x4-2    | 33309816   | 36.04 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/large/AVX512-x8-2    | 59912676   | 19.74 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/large/AVX512-x16-2   | 100000000  | 10.65 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/xlarge/AVX512-x4-2   | 8346818    | 143.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/xlarge/AVX512-x8-2   | 16980399   | 70.54 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/xlarge/AVX512-x16-2  | 28676455   | 42.94 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/massive/AVX512-x4-2  | 1000000    | 1151 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/massive/AVX512-x8-2  | 2161594    | 555.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsNegative/massive/AVX512-x16-2 | 3549094    | 350.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8ByWidth/AVX512-x16-2      | 331533141  | 3.222 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8ByWidth/AVX512-x32-2      | 408741681  | 3.193 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt8ByWidth/AVX512-x64-2      | 365382873  | 3.241 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64SteadyState/AVX512-x2-2  | 5722603    | 211.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64SteadyState/AVX512-x4-2  | 11711869   | 103.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkContainsInt64SteadyState/AVX512-x8-2  | 19671033   | 61.36 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/Fallback-lo-2           | 248740710  | 5.218 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/AVX-x16-2               | 126181464  | 9.485 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/AVX2-x32-2              | 73059427   | 14.44 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/small/AVX512-x64-2            | 49913169   | 24.41 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/Fallback-lo-2          | 17278075   | 69.96 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/AVX-x16-2              | 100000000  | 10.58 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/AVX2-x32-2             | 91620999   | 13.10 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/medium/AVX512-x64-2           | 54082130   | 22.20 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/Fallback-lo-2           | 2006178    | 576.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/AVX-x16-2               | 41836690   | 27.82 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/AVX2-x32-2              | 51735399   | 23.04 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/large/AVX512-x64-2            | 40861586   | 29.40 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/Fallback-lo-2          | 273898     | 4383 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/AVX-x16-2              | 6928408    | 173.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/AVX2-x32-2             | 12639586   | 94.09 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8/xlarge/AVX512-x64-2           | 13509693   | 89.67 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/Fallback-lo-2          | 249444103  | 5.012 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/AVX-x8-2               | 244927230  | 5.052 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/AVX2-x16-2             | 122088517  | 9.715 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/small/AVX512-x32-2           | 54098370   | 22.00 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/Fallback-lo-2         | 15782683   | 72.54 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/AVX-x8-2              | 100000000  | 10.51 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/AVX2-x16-2            | 100000000  | 10.75 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/medium/AVX512-x32-2          | 56147455   | 21.38 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/Fallback-lo-2          | 2173214    | 598.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/AVX-x8-2               | 26319481   | 44.73 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/AVX2-x16-2             | 40459519   | 27.91 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/large/AVX512-x32-2           | 39359752   | 31.28 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/Fallback-lo-2         | 273932     | 4382 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/AVX-x8-2              | 3557265    | 331.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/AVX2-x16-2            | 6930166    | 173.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt16/xlarge/AVX512-x32-2          | 12100244   | 97.01 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/Fallback-lo-2          | 249566539  | 4.808 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/AVX-x4-2               | 259250019  | 4.581 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/AVX2-x8-2              | 232858933  | 5.404 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/small/AVX512-x16-2           | 100000000  | 11.18 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/Fallback-lo-2         | 17274441   | 72.28 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/AVX-x4-2              | 58400258   | 20.56 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/AVX2-x8-2             | 110851756  | 10.67 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/medium/AVX512-x16-2          | 106593603  | 11.25 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/Fallback-lo-2          | 2171817    | 551.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/AVX-x4-2               | 8270253    | 146.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/AVX2-x8-2              | 22234518   | 46.06 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/large/AVX512-x16-2           | 37448763   | 32.31 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/Fallback-lo-2         | 273699     | 4559 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/AVX-x4-2              | 1000000    | 1102 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/AVX2-x8-2             | 3586887    | 332.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt32/xlarge/AVX512-x16-2          | 7214437    | 170.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/Fallback-lo-2          | 417473124  | 2.886 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/AVX-x2-2               | 287521756  | 4.169 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/AVX2-x4-2              | 277783513  | 4.311 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/small/AVX512-x8-2            | 172823103  | 6.993 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/Fallback-lo-2         | 34022653   | 35.27 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/AVX-x2-2              | 49241248   | 24.05 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/AVX2-x4-2             | 78897342   | 14.58 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/medium/AVX512-x8-2           | 84361297   | 14.03 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/Fallback-lo-2          | 3680988    | 282.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/AVX-x2-2               | 6293607    | 170.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/AVX2-x4-2              | 12739849   | 91.28 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/large/AVX512-x8-2            | 25508130   | 46.30 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/Fallback-lo-2         | 546321     | 2283 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/AVX-x2-2              | 877434     | 1289 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/AVX2-x4-2             | 1845892    | 650.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64/xlarge/AVX512-x8-2           | 2148355    | 550.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/Fallback-lo-2        | 411100770  | 2.951 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/AVX-x4-2             | 264013596  | 4.572 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/AVX2-x8-2            | 174478266  | 6.911 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/small/AVX512-x16-2         | 61182673   | 19.78 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/Fallback-lo-2       | 33815070   | 35.68 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/AVX-x4-2            | 58238188   | 20.66 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/AVX2-x8-2           | 91316544   | 13.26 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/medium/AVX512-x16-2        | 80046624   | 15.08 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/Fallback-lo-2        | 4304168    | 278.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/AVX-x4-2             | 6198957    | 184.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/AVX2-x8-2            | 12260169   | 86.60 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/large/AVX512-x16-2         | 22147112   | 45.34 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/Fallback-lo-2       | 546901     | 2193 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/AVX-x4-2            | 736503     | 1622 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/AVX2-x8-2           | 1493887    | 810.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat32/xlarge/AVX512-x16-2        | 2959298    | 393.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/Fallback-lo-2        | 410778070  | 3.043 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/AVX-x2-2             | 254156008  | 4.714 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/AVX2-x4-2            | 227604434  | 5.323 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/small/AVX512-x8-2          | 170099748  | 7.115 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/Fallback-lo-2       | 33646345   | 35.78 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/AVX-x2-2            | 32931152   | 34.92 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/AVX2-x4-2           | 75389446   | 16.79 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/medium/AVX512-x8-2         | 89826181   | 13.33 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/Fallback-lo-2        | 4293837    | 302.8 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/AVX-x2-2             | 3146601    | 381.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/AVX2-x4-2            | 6373876    | 184.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/large/AVX512-x8-2          | 13464712   | 88.96 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/Fallback-lo-2       | 545764     | 2193 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/AVX-x2-2            | 368846     | 3390 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/AVX2-x4-2           | 709940     | 1613 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumFloat64/xlarge/AVX512-x8-2         | 1480214    | 808.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/Fallback-lo-2         | 411529147  | 3.043 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/AVX-x4-2              | 204428401  | 5.872 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/AVX2-x8-2             | 187573928  | 6.214 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/small/AVX512-x16-2          | 98346700   | 12.12 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/Fallback-lo-2        | 33481442   | 35.72 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/AVX-x4-2             | 52042394   | 22.12 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/AVX2-x8-2            | 96288541   | 13.44 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/medium/AVX512-x16-2         | 100995780  | 11.90 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/Fallback-lo-2         | 4296570    | 289.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/AVX-x4-2              | 7743022    | 146.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/AVX2-x8-2             | 24355988   | 46.26 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/large/AVX512-x16-2          | 37322655   | 32.89 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/Fallback-lo-2        | 547008     | 2193 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/AVX-x4-2             | 1087246    | 1112 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/AVX2-x8-2            | 1386868    | 761.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanInt32/xlarge/AVX512-x16-2         | 7166142    | 170.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/Fallback-lo-2       | 349760005  | 3.449 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/AVX-x2-2            | 189674538  | 6.293 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/AVX2-x4-2           | 159228600  | 7.531 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/small/AVX512-x8-2         | 110196433  | 10.89 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/Fallback-lo-2      | 32968618   | 36.17 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/AVX-x2-2           | 30863817   | 37.69 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/AVX2-x4-2          | 62428772   | 19.66 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/medium/AVX512-x8-2        | 77140984   | 15.54 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/Fallback-lo-2       | 4281057    | 280.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/AVX-x2-2            | 3057349    | 389.4 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/AVX2-x4-2           | 6509438    | 185.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/large/AVX512-x8-2         | 12668032   | 93.50 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/Fallback-lo-2      | 545898     | 2288 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/AVX-x2-2           | 367671     | 4048 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/AVX2-x4-2          | 739941     | 1621 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMeanFloat64/xlarge/AVX512-x8-2        | 1434867    | 811.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/small/AVX-x4-2               | 312338268  | 3.860 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/small/AVX2-x8-2              | 238034872  | 5.042 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/small/AVX512-x16-2           | 152600943  | 6.661 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/medium/AVX-x4-2              | 61051266   | 19.73 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/medium/AVX2-x8-2             | 91792144   | 13.11 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/medium/AVX512-x16-2          | 99994540   | 12.18 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/large/AVX-x4-2               | 8604774    | 140.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/large/AVX2-x8-2              | 15581037   | 77.56 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/large/AVX512-x16-2           | 30512421   | 40.24 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/xlarge/AVX-x4-2              | 1000000    | 1110 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/xlarge/AVX2-x8-2             | 2158272    | 557.2 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinInt32/xlarge/AVX512-x16-2          | 4253668    | 282.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/small/AVX-x2-2             | 264129410  | 4.544 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/small/AVX2-x4-2            | 299587609  | 4.008 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/small/AVX512-x8-2          | 100000000  | 10.05 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/medium/AVX-x2-2            | 32778514   | 36.93 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/medium/AVX2-x4-2           | 53356347   | 20.30 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/medium/AVX512-x8-2         | 74832976   | 16.21 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/large/AVX-x2-2             | 3863326    | 300.0 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/large/AVX2-x4-2            | 7670576    | 146.5 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/large/AVX512-x8-2          | 14017984   | 78.21 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/xlarge/AVX-x2-2            | 492739     | 2195 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/xlarge/AVX2-x4-2           | 1000000    | 1103 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMinFloat64/xlarge/AVX512-x8-2         | 2145290    | 560.3 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/small/AVX-x4-2               | 306585705  | 3.860 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/small/AVX2-x8-2              | 237347997  | 5.086 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/small/AVX512-x16-2           | 201433966  | 6.130 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/medium/AVX-x4-2              | 60759631   | 19.92 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/medium/AVX2-x8-2             | 90934662   | 13.13 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/medium/AVX512-x16-2          | 98517944   | 12.18 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/large/AVX-x4-2               | 8590542    | 139.6 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/large/AVX2-x8-2              | 15770372   | 77.69 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/large/AVX512-x16-2           | 30197324   | 39.32 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/xlarge/AVX-x4-2              | 1000000    | 1104 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/xlarge/AVX2-x8-2             | 2152038    | 562.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxInt32/xlarge/AVX512-x16-2          | 3917990    | 296.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/small/AVX-x2-2             | 249617162  | 4.816 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/small/AVX2-x4-2            | 207017514  | 5.855 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/small/AVX512-x8-2          | 66520290   | 17.74 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/medium/AVX-x2-2            | 32307492   | 36.92 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/medium/AVX2-x4-2           | 57306838   | 20.77 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/medium/AVX512-x8-2         | 56911946   | 21.12 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/large/AVX-x2-2             | 4259366    | 287.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/large/AVX2-x4-2            | 7905420    | 148.9 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/large/AVX512-x8-2          | 14100686   | 83.43 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/xlarge/AVX-x2-2            | 545378     | 2243 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/xlarge/AVX2-x4-2           | 1000000    | 1113 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkMaxFloat64/xlarge/AVX512-x8-2         | 2119741    | 565.7 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/Fallback-lo-2          | 896775     | 1335 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/AVX-x16-2              | 12557700   | 94.52 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/AVX2-x32-2             | 18702537   | 55.03 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt8ByWidth/AVX512-x64-2           | 21342572   | 56.10 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/Fallback-lo-2     | 513738     | 2195 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/AVX-x2-2          | 928376     | 1296 ns/op  | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/AVX2-x4-2         | 1836968    | 888.1 ns/op | 0 B/op   | 0 allocs/op |
| BenchmarkSumInt64SteadyState/AVX512-x8-2       | 2141715    | 551.3 ns/op | 0 B/op   | 0 allocs/op |

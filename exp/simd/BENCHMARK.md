# Benchmark

## Summary

Compares two vector widths: an AMD EPYC 9454P (AVX-512, 512-bit vectors, up to 16 int32
lanes) against arm64 NEON (128-bit vectors, 4 int32 lanes, `ubicloud-standard-2-arm`). Both
ran with `GOMAXPROCS=2`, on inputs of 8 / 128 / 1024 / 8192 elements.

### Speedup vs `lo.*` (xlarge, 8192 elements)

| Operation               | amd64 (AVX-512) | arm64 (NEON)   |
| ----------------------- | --------------- | -------------- |
| Sum Int8                | 27.9x           | 8.7x           |
| Sum Int16               | 17.9x           | 5.5x           |
| Sum Int32               | 6.9x            | 1.7x           |
| Sum Int64               | 2.8x            | 0.98x (no win) |
| Sum Float32             | 13.5x           | 2.0x           |
| Sum Float64             | 7.1x            | 1.3x           |
| Min Int32               | 14.2x           | 2.7x           |
| Min Float64             | 10.2x           | 1.7x           |
| Clamp Int32             | 6.2x            | 6.1x           |
| Clamp Float64           | 2.5x            | 3.8x           |
| Contains Int64 (miss)   | 1.9x            | 0.78x (slower) |
| Contains Float64 (miss) | 1.9x            | 0.79x (slower) |

### Findings

- Speedup scales with lane count. amd64's vectors are 4x wider (512 vs 128 bit), and
  `Sum`/`Min`/`Max` on 8-32 bit element types show a roughly proportional 3-4x larger
  speedup on amd64 than on arm64.
- 64-bit elements barely benefit on arm64 (only 2 lanes/vector), and can regress: `Sum
Int64` at xlarge is a wash (0.98x), and `Contains` on Int64/Float64 (miss, xlarge) is
  about 25% slower under SIMD on arm64. The mask-fold-every-8-vectors batching amortizes
  over only 16 elements per cycle at 2 lanes/vector, versus 64 elements per cycle on amd64.
- `Clamp` is nearly lane-count-independent: its win comes from skipping the per-element
  function-call and branch overhead of `lo.Clamp`, not from vector width, so arm64 keeps
  pace with (and sometimes beats) amd64.
- `Contains` with an early hit is slower under SIMD on both platforms (e.g. amd64
  ContainsInt8/small/hit-first: lo=1.49ns vs simd=6.18ns). Batching the mask check every
  `containsBlock=8` vectors trades early-exit speed for miss/late-hit throughput — a fixed
  cost on both architectures, not a regression specific to one.
- Small inputs (8 elements) favor the scalar path on both platforms for `Sum`/`Min`/`Max`:
  fixed vector setup cost dominates below one full vector's worth of data.

## Run

```bash
export GOEXPERIMENT=simd
cd exp/simd/
go test -run=^$ -bench=. -benchmem ./...
```

```bash
# force a specific vector width, or force the scalar fallback with 0
GODEBUG=simd=128 go test -run=^$ -bench=. -benchmem ./...
GODEBUG=simd=0   go test -run=^$ -bench=. -benchmem ./...
```

## Results

### amd64 (AVX-512)

```txt
simd: useSIMD=true lanes8=64 lanes16=32 lanes32=16 lanes64=8
goos: linux
goarch: amd64
pkg: github.com/samber/lo/exp/simd
cpu: AMD EPYC 9454P 48-Core Processor
BenchmarkContainsInt8/small/miss/lo-2         	258590469	         4.873 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/miss/simd-2       	100000000	        12.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-first/lo-2    	746566902	         1.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-first/simd-2  	166063821	         6.179 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-mid/lo-2      	417460252	         2.984 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-mid/simd-2    	120501679	         9.060 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-last/lo-2     	297413970	         3.812 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-last/simd-2   	122279126	         9.853 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/miss/lo-2        	21178443	        47.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/miss/simd-2      	77388469	        14.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-first/lo-2   	847617650	         1.646 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-first/simd-2 	59468092	        17.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-mid/lo-2     	34128739	        29.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-mid/simd-2   	100000000	        12.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-last/lo-2    	25065469	        42.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-last/simd-2  	100000000	        13.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/miss/lo-2         	31231299	        39.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/miss/simd-2       	74019148	        18.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-first/lo-2    	741736572	         1.597 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-first/simd-2  	74457026	        21.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-mid/lo-2      	27794562	        41.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-mid/simd-2    	74253519	        21.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-last/lo-2     	16731894	        75.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-last/simd-2   	64175052	        19.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/miss/lo-2        	31031833	        46.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/miss/simd-2      	74761477	        16.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-first/lo-2   	886308862	         1.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-first/simd-2 	74385810	        16.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-mid/lo-2     	45080029	        27.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-mid/simd-2   	71762157	        16.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-last/lo-2    	23040543	        52.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-last/simd-2  	75295464	        16.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/miss/lo-2        	272166064	         4.306 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/miss/simd-2      	143611341	         8.675 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-first/lo-2   	712440501	         1.561 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-first/simd-2 	193880863	         6.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-mid/lo-2     	375078192	         3.022 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-mid/simd-2   	171406761	         6.996 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-last/lo-2    	322390484	         3.792 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-last/simd-2  	153217083	         7.745 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/miss/lo-2       	28433482	        42.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/miss/simd-2     	90882747	        15.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-first/lo-2  	696932468	         1.553 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-first/simd-2	95348328	        12.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-mid/lo-2    	49993876	        25.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-mid/simd-2  	79206555	        12.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-last/lo-2   	32296876	        37.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-last/simd-2 	94944873	        12.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/miss/lo-2        	 4123124	       291.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/miss/simd-2      	17753541	        75.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-first/lo-2   	570499130	         1.813 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-first/simd-2 	93741966	        14.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-mid/lo-2     	 6735867	       169.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-mid/simd-2   	23077220	        51.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-last/lo-2    	 4019640	       318.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-last/simd-2  	16570404	        95.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/miss/lo-2       	  482202	      2403 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/miss/simd-2     	 1982820	       639.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-first/lo-2  	735247777	         1.706 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-first/simd-2	93812905	        13.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-mid/lo-2    	44889517	        28.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-mid/simd-2  	95267710	        13.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-last/lo-2   	23570908	        51.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-last/simd-2 	58265901	        21.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/miss/lo-2        	260468460	         4.884 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/miss/simd-2      	100000000	        10.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-first/lo-2   	523228150	         2.038 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-first/simd-2 	91179624	        12.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-mid/lo-2     	284786263	         4.039 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-mid/simd-2   	80576648	        13.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-last/lo-2    	286026655	         4.011 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-last/simd-2  	107635684	         9.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/miss/lo-2       	26121252	        43.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/miss/simd-2     	56498504	        25.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-first/lo-2  	603157316	         1.873 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-first/simd-2	85209267	        13.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-mid/lo-2    	51278792	        25.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-mid/simd-2  	54627612	        22.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-last/lo-2   	29775832	        48.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-last/simd-2 	51090681	        27.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/miss/lo-2        	 3480914	       299.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/miss/simd-2      	 8725729	       150.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-first/lo-2   	660332406	         1.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-first/simd-2 	92847571	        14.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-mid/lo-2     	 8148780	       155.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-mid/simd-2   	15726166	        79.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-last/lo-2    	 3880444	       311.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-last/simd-2  	 8311950	       138.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/miss/lo-2       	  543722	      2347 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/miss/simd-2     	 1077140	      1216 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-first/lo-2  	582159559	         1.943 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-first/simd-2	93735734	        13.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-mid/lo-2    	44993198	        26.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-mid/simd-2  	57991215	        21.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-last/lo-2   	22240648	        52.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-last/simd-2 	39257781	        30.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/miss/lo-2      	263744562	         4.910 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/miss/simd-2    	110686003	        10.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-first/lo-2 	800239761	         1.675 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-first/simd-2	100000000	        11.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-mid/lo-2   	434927352	         2.971 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-mid/simd-2 	124090588	        10.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-last/lo-2  	336679656	         3.681 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-last/simd-2	130138957	        10.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/miss/lo-2     	26741660	        49.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/miss/simd-2   	50410518	        26.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-first/lo-2	738986659	         1.361 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-first/simd-2	84278013	        15.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-mid/lo-2  	37861258	        32.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-mid/simd-2	46584414	        27.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-last/lo-2 	22960437	        46.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-last/simd-2	53695180	        23.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/miss/lo-2      	 3524479	       358.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/miss/simd-2    	 8473680	       144.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-first/lo-2 	760730401	         1.550 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-first/simd-2	85564095	        15.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-mid/lo-2   	 6367954	       189.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-mid/simd-2 	13155859	        88.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-last/lo-2  	 3163206	       383.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-last/simd-2	 7382221	       165.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/miss/lo-2     	  498136	      2775 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/miss/simd-2   	  882256	      1443 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-first/lo-2	730575022	         1.679 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-first/simd-2	82845156	        15.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-mid/lo-2  	43674962	        32.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-mid/simd-2	52087334	        24.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-last/lo-2 	20047483	        60.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-last/simd-2	39308262	        30.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/small/lo-2                   	344053036	         3.576 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/small/simd-2                 	48484386	        24.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/medium/lo-2                  	30230888	        40.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/medium/simd-2                	46501662	        26.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/large/lo-2                   	 3257307	       342.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/large/simd-2                 	38451621	        31.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/xlarge/lo-2                  	  466400	      2576 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/xlarge/simd-2                	13143464	        92.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/small/lo-2                  	317628589	         3.875 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/small/simd-2                	76036317	        16.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/medium/lo-2                 	30946586	        38.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/medium/simd-2               	71961078	        16.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/large/lo-2                  	 3441673	       346.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/large/simd-2                	39078126	        30.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/xlarge/lo-2                 	  468649	      2697 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/xlarge/simd-2               	 7565733	       150.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/small/lo-2                  	304213428	         3.625 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/small/simd-2                	100000000	        14.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/medium/lo-2                 	25264540	        43.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/medium/simd-2               	64266990	        19.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/large/lo-2                  	 2906858	       371.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/large/simd-2                	26669353	        48.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/xlarge/lo-2                 	  477314	      2438 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/xlarge/simd-2               	 3902476	       356.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/small/lo-2                  	336338990	         3.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/small/simd-2                	116219266	        10.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/medium/lo-2                 	24315469	        41.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/medium/simd-2               	49670037	        21.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/large/lo-2                  	 3341492	       349.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/large/simd-2                	15373803	        80.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/xlarge/lo-2                 	  540304	      2507 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/xlarge/simd-2               	 1646770	       881.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/small/lo-2                	174620245	         6.944 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/small/simd-2              	100000000	        12.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/medium/lo-2               	12684345	       103.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/medium/simd-2             	72784719	        17.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/large/lo-2                	 1447813	       826.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/large/simd-2              	22004838	        61.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/xlarge/lo-2               	  180746	      6611 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/xlarge/simd-2             	 2388896	       489.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/small/lo-2                	167798367	         7.192 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/small/simd-2              	131432814	         9.717 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/medium/lo-2               	12276523	       103.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/medium/simd-2             	67104901	        17.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/large/lo-2                	 1447917	       829.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/large/simd-2              	11921593	       113.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/xlarge/lo-2               	  181279	      6788 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/xlarge/simd-2             	 1266904	       953.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/min/lo-2           	300567751	         3.828 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/min/simd-2         	139868398	         8.673 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/max/lo-2           	308610484	         3.860 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/max/simd-2         	131786589	         9.070 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/min/lo-2          	18918738	        66.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/min/simd-2        	70622557	        19.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/max/lo-2          	19207419	        65.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/max/simd-2        	66366148	        20.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/min/lo-2           	 2194827	       555.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/min/simd-2         	27969184	        41.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/max/lo-2           	 2154427	       587.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/max/simd-2         	24299485	        48.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/min/lo-2          	  266289	      4462 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/min/simd-2        	 3702694	       314.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/max/lo-2          	  274068	      4421 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/max/simd-2        	 3619138	       333.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/min/lo-2         	171644604	         7.243 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/min/simd-2       	65596477	        17.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/max/lo-2         	169380402	         6.882 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/max/simd-2       	94213588	        12.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/min/lo-2        	11517909	       106.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/min/simd-2      	46245289	        27.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/max/lo-2        	10827914	       123.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/max/simd-2      	48010387	        28.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/min/lo-2         	 1218367	      1058 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/min/simd-2       	 8835969	       128.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/max/lo-2         	 1244212	      1019 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/max/simd-2       	10061354	       107.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/min/lo-2        	  166238	      8233 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/min/simd-2      	 1424132	       805.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/max/lo-2        	  142473	      8942 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/max/simd-2      	 1000000	      1066 ns/op	       0 B/op	       0 allocs/op
BenchmarkClampInt32/small/lo-2                	66321675	        19.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkClampInt32/small/simd-2              	43225483	        25.51 ns/op	      32 B/op	       1 allocs/op
BenchmarkClampInt32/medium/lo-2               	 3376581	       348.7 ns/op	     512 B/op	       1 allocs/op
BenchmarkClampInt32/medium/simd-2             	11056546	       114.8 ns/op	     512 B/op	       1 allocs/op
BenchmarkClampInt32/large/lo-2                	  450595	      2776 ns/op	    4096 B/op	       1 allocs/op
BenchmarkClampInt32/large/simd-2              	 1686592	       707.3 ns/op	    4096 B/op	       1 allocs/op
BenchmarkClampInt32/xlarge/lo-2               	   56334	     22259 ns/op	   32768 B/op	       1 allocs/op
BenchmarkClampInt32/xlarge/simd-2             	  383905	      3574 ns/op	   32768 B/op	       1 allocs/op
BenchmarkClampFloat64/small/lo-2              	32228316	        35.51 ns/op	      64 B/op	       1 allocs/op
BenchmarkClampFloat64/small/simd-2            	41585523	        27.92 ns/op	      64 B/op	       1 allocs/op
BenchmarkClampFloat64/medium/lo-2             	 2785234	       423.3 ns/op	    1024 B/op	       1 allocs/op
BenchmarkClampFloat64/medium/simd-2           	 7016434	       172.9 ns/op	    1024 B/op	       1 allocs/op
BenchmarkClampFloat64/large/lo-2              	  349032	      3193 ns/op	    8192 B/op	       1 allocs/op
BenchmarkClampFloat64/large/simd-2            	  888154	      1331 ns/op	    8192 B/op	       1 allocs/op
BenchmarkClampFloat64/xlarge/lo-2             	   46261	     25138 ns/op	   65536 B/op	       1 allocs/op
BenchmarkClampFloat64/xlarge/simd-2           	  186994	     10151 ns/op	   65536 B/op	       1 allocs/op
PASS
ok  	github.com/samber/lo/exp/simd	338.404s
```

### arm64 (NEON)

```txt
simd: useSIMD=true lanes8=16 lanes16=8 lanes32=4 lanes64=2
goos: linux
goarch: arm64
pkg: github.com/samber/lo/exp/simd
BenchmarkContainsInt8/small/miss/lo-2         	137006245	         8.760 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/miss/simd-2       	73249340	        14.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-first/lo-2    	441079110	         2.699 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-first/simd-2  	135510302	         8.962 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-mid/lo-2      	186966316	         6.449 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-mid/simd-2    	97672782	        12.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-last/lo-2     	142161350	         8.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/small/hit-last/simd-2   	84128817	        14.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/miss/lo-2        	12797007	        89.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/miss/simd-2      	47489017	        25.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-first/lo-2   	445481127	         2.697 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-first/simd-2 	45007200	        25.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-mid/lo-2     	25167320	        46.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-mid/simd-2   	48042814	        25.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-last/lo-2    	13216838	        90.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/medium/hit-last/simd-2  	44360389	        25.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/miss/lo-2         	12512444	        94.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/miss/simd-2       	28459406	        43.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-first/lo-2    	436439814	         2.710 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-first/simd-2  	47582415	        25.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-mid/lo-2      	13911404	        86.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-mid/simd-2    	48105369	        25.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-last/lo-2     	 6102433	       167.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/large/hit-last/simd-2   	28124726	        41.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/miss/lo-2        	12499442	        95.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/miss/simd-2      	29132605	        42.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-first/lo-2   	444513262	         2.693 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-first/simd-2 	48135163	        25.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-mid/lo-2     	18556674	        64.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-mid/simd-2   	48140337	        25.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-last/lo-2    	 9575859	       125.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt8/xlarge/hit-last/simd-2  	28025019	        41.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/miss/lo-2        	162047974	         7.404 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/miss/simd-2      	75506649	        16.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-first/lo-2   	504571672	         2.392 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-first/simd-2 	78218086	        15.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-mid/lo-2     	237777654	         5.050 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-mid/simd-2   	78970851	        15.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-last/lo-2    	169049674	         7.085 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/small/hit-last/simd-2  	77261839	        15.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/miss/lo-2       	12688476	        88.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/miss/simd-2     	18429182	        65.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-first/lo-2  	495413870	         2.383 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-first/simd-2	52785583	        22.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-mid/lo-2    	26387823	        45.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-mid/simd-2  	24030200	        50.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-last/lo-2   	13613529	        88.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/medium/hit-last/simd-2 	17889247	        64.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/miss/lo-2        	 1711374	       703.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/miss/simd-2      	 2612956	       463.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-first/lo-2   	500786902	         2.377 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-first/simd-2 	53372235	        22.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-mid/lo-2     	 3333408	       358.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-mid/simd-2   	 4809026	       252.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-last/lo-2    	 1712554	       700.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/large/hit-last/simd-2  	 2576718	       461.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/miss/lo-2       	  215704	      5529 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/miss/simd-2     	  292855	      3565 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-first/lo-2  	503753553	         2.390 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-first/simd-2	53524976	        22.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-mid/lo-2    	19008756	        63.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-mid/simd-2  	24042795	        50.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-last/lo-2   	 9720858	       124.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt32/xlarge/hit-last/simd-2 	12760754	        93.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/miss/lo-2        	168856111	         7.106 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/miss/simd-2      	63486894	        18.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-first/lo-2   	506103865	         2.365 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-first/simd-2 	64977679	        18.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-mid/lo-2     	236541609	         5.057 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-mid/simd-2   	61865874	        18.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-last/lo-2    	177210231	         6.758 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/small/hit-last/simd-2  	57167575	        18.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/miss/lo-2       	13559020	        88.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/miss/simd-2     	 9366336	       125.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-first/lo-2  	501215614	         2.370 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-first/simd-2	53465071	        22.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-mid/lo-2    	26389981	        45.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-mid/simd-2  	15703723	        76.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-last/lo-2   	13501939	        88.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/medium/hit-last/simd-2 	 9378918	       126.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/miss/lo-2        	 1712763	       700.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/miss/simd-2      	 1345393	       901.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-first/lo-2   	507027400	         2.366 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-first/simd-2 	47157206	        22.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-mid/lo-2     	 3340345	       356.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-mid/simd-2   	 2461178	       470.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-last/lo-2    	 1713714	       703.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/large/hit-last/simd-2  	 1341931	       894.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/miss/lo-2       	  207632	      5570 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/miss/simd-2     	  170432	      7124 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-first/lo-2  	505347420	         2.371 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-first/simd-2	53230467	        22.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-mid/lo-2    	18824450	        63.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-mid/simd-2  	12773811	        91.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-last/lo-2   	 9693945	       123.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsInt64/xlarge/hit-last/simd-2 	 6668452	       181.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/miss/lo-2      	136832542	         8.789 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/miss/simd-2    	63899839	        18.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-first/lo-2 	503597872	         2.378 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-first/simd-2	62177582	        18.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-mid/lo-2   	187028780	         6.458 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-mid/simd-2 	66439002	        18.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-last/lo-2  	141672436	         8.428 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/small/hit-last/simd-2	66701499	        18.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/miss/lo-2     	13376398	        90.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/miss/simd-2   	 9432597	       125.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-first/lo-2	502758975	         2.377 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-first/simd-2	53679554	        22.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-mid/lo-2  	25565072	        47.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-mid/simd-2	15721094	        77.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-last/lo-2 	13391666	        89.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/medium/hit-last/simd-2	 9311408	       128.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/miss/lo-2      	 1707831	       704.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/miss/simd-2    	 1338388	       900.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-first/lo-2 	496458430	         2.381 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-first/simd-2	52696478	        22.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-mid/lo-2   	 3352152	       357.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-mid/simd-2 	 2579428	       464.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-last/lo-2  	 1710384	       701.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/large/hit-last/simd-2	 1346439	       897.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/miss/lo-2     	  212464	      5564 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/miss/simd-2   	  171996	      7051 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-first/lo-2	489202886	         2.387 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-first/simd-2	49176135	        22.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-mid/lo-2  	18309127	        64.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-mid/simd-2	13398700	        89.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-last/lo-2 	 9648476	       124.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsFloat64/xlarge/hit-last/simd-2	 6662628	       182.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/small/lo-2                   	197486548	         6.083 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/small/simd-2                 	42468735	        24.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/medium/lo-2                  	15576266	        73.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/medium/simd-2                	37733427	        29.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/large/lo-2                   	 1974452	       605.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/large/simd-2                 	13962170	        86.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/xlarge/lo-2                  	  236778	      4783 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt8/xlarge/simd-2                	 2162647	       548.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/small/lo-2                  	142403330	         8.436 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/small/simd-2                	71547817	        16.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/medium/lo-2                 	13088924	        90.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/medium/simd-2               	39407416	        30.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/large/lo-2                  	 1605288	       747.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/large/simd-2                	 8126799	       143.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/xlarge/lo-2                 	  203371	      5798 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt16/xlarge/simd-2               	 1000000	      1063 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/small/lo-2                  	197902393	         6.063 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/small/simd-2                	66953673	        15.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/medium/lo-2                 	16271486	        73.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/medium/simd-2               	21470592	        54.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/large/lo-2                  	 1985832	       610.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/large/simd-2                	 3331377	       357.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/xlarge/lo-2                 	  236628	      4774 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt32/xlarge/simd-2               	  427998	      2826 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/small/lo-2                  	141805894	         8.424 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/small/simd-2                	77401177	        15.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/medium/lo-2                 	13424770	        89.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/medium/simd-2               	12129513	        96.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/large/lo-2                  	 1714032	       699.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/large/simd-2                	 1669855	       719.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/xlarge/lo-2                 	  212902	      5546 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumInt64/xlarge/simd-2               	  216127	      5655 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/small/lo-2                	142123837	         8.439 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/small/simd-2              	72145560	        15.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/medium/lo-2               	13430660	        89.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/medium/simd-2             	22270309	        54.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/large/lo-2                	 1713061	       700.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/large/simd-2              	 3297398	       365.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/xlarge/lo-2               	  215550	      5531 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat32/xlarge/simd-2             	  432170	      2807 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/small/lo-2                	142442215	         8.475 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/small/simd-2              	75010501	        15.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/medium/lo-2               	13289500	        89.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/medium/simd-2             	15224515	        76.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/large/lo-2                	 1719090	       698.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/large/simd-2              	 2201510	       540.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/xlarge/lo-2               	  213517	      5543 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumFloat64/xlarge/simd-2             	  286266	      4200 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/min/lo-2           	142470356	         8.447 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/min/simd-2         	75993615	        16.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/max/lo-2           	142484416	         8.454 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/small/max/simd-2         	75328999	        16.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/min/lo-2          	13353237	        89.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/min/simd-2        	26530100	        46.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/max/lo-2          	13131878	        89.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/medium/max/simd-2        	23302201	        45.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/min/lo-2           	 1698438	       707.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/min/simd-2         	 4385018	       273.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/max/lo-2           	 1683387	       707.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/large/max/simd-2         	 4350010	       274.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/min/lo-2          	  211755	      5645 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/min/simd-2        	  573325	      2098 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/max/lo-2          	  211996	      5682 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxInt32/xlarge/max/simd-2        	  516504	      2102 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/min/lo-2         	122796840	         9.768 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/min/simd-2       	62294169	        19.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/max/lo-2         	131742105	         9.132 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/small/max/simd-2       	66451216	        18.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/min/lo-2        	 8769506	       137.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/min/simd-2      	10309614	       116.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/max/lo-2        	 7546413	       159.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/medium/max/simd-2      	10432866	       116.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/min/lo-2         	  861070	      1384 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/min/simd-2       	 1408084	       865.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/max/lo-2         	  756825	      1376 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/large/max/simd-2       	 1392196	       862.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/min/lo-2        	  106204	     11252 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/min/simd-2      	  177849	      6774 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/max/lo-2        	  106603	     11233 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinMaxFloat64/xlarge/max/simd-2      	  180156	      6718 ns/op	       0 B/op	       0 allocs/op
BenchmarkClampInt32/small/lo-2                	45785798	        25.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkClampInt32/small/simd-2              	23673775	        51.21 ns/op	      32 B/op	       1 allocs/op
BenchmarkClampInt32/medium/lo-2               	 1689514	       724.5 ns/op	     512 B/op	       1 allocs/op
BenchmarkClampInt32/medium/simd-2             	 5762042	       206.8 ns/op	     512 B/op	       1 allocs/op
BenchmarkClampInt32/large/lo-2                	  220647	      5846 ns/op	    4096 B/op	       1 allocs/op
BenchmarkClampInt32/large/simd-2              	  975466	      1208 ns/op	    4096 B/op	       1 allocs/op
BenchmarkClampInt32/xlarge/lo-2               	   26809	     43919 ns/op	   32768 B/op	       1 allocs/op
BenchmarkClampInt32/xlarge/simd-2             	  176863	      7260 ns/op	   32768 B/op	       1 allocs/op
BenchmarkClampFloat64/small/lo-2              	15235107	        81.72 ns/op	      64 B/op	       1 allocs/op
BenchmarkClampFloat64/small/simd-2            	16819980	        70.96 ns/op	      64 B/op	       1 allocs/op
BenchmarkClampFloat64/medium/lo-2             	 1274198	       936.6 ns/op	    1024 B/op	       1 allocs/op
BenchmarkClampFloat64/medium/simd-2           	 3112430	       378.1 ns/op	    1024 B/op	       1 allocs/op
BenchmarkClampFloat64/large/lo-2              	  158911	      7288 ns/op	    8192 B/op	       1 allocs/op
BenchmarkClampFloat64/large/simd-2            	  452385	      2745 ns/op	    8192 B/op	       1 allocs/op
BenchmarkClampFloat64/xlarge/lo-2             	   21015	     57075 ns/op	   65536 B/op	       1 allocs/op
BenchmarkClampFloat64/xlarge/simd-2           	   80240	     14947 ns/op	   65536 B/op	       1 allocs/op
PASS
ok  	github.com/samber/lo/exp/simd	326.896s
```

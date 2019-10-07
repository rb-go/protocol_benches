# protocol benches

[![Build Status](https://travis-ci.org/riftbit/protocol_benches.svg?branch=master)](https://travis-ci.org/riftbit/protocol_benches)

I decided to make some benchmarks to compare few libraries. And results are shocked me.
Решил провести небольшие сравнительные тесты по скорости работы разных библиотек для Golang. Результаты меня, мягко говоря, шокировали. 

Easyjson в результате тестов оказался самый медленный при Marshal, причем даже стандартная библиотека encode/json (в тесте Benchmark_std_*) работает на маршал значительно быстрее (на 50%)
Так же он оказался медленнее старенького ffjson (в случае прегенерации кода и при не использовании вложенных структур, а типов).

Что я делаю не так??? Может быть в коде что-то не то?
Пулл-реквесты приветствуются.

Тесты с пометкой Min - это minified json, у него убраны переносы строк и лишние пробелы. Full - это отформатированный json с отступами и переносами.

### Tested Packages

 - `msgpack_vmihailenco` - https://github.com/vmihailenco/msgpack (v4.0.4+incompatible)
 - `easyjson` - https://github.com/mailru/easyjson (v0.7.0)
 - `ffjson` - https://github.com/pquerna/ffjson (v0.0.0-20190930134022-aa0246cd15f7)
 - `jsoniterator` - https://github.com/json-iterator (v1.1.7)
 - `std` - https://godoc.org/encoding/json (same as golang)
 - `protobuf` - todo
 - `gogoproto` - todo
 
### todo

 - check why Benchmark_ffjson_Marshal_generated generates so much allocations

### golang 1.13 on windows 10 x64
```
Benchmark_msgpack_vmihailenco_Marshal-4        	  166664	     31693 ns/op	  164080 B/op	       6 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   13209	    484745 ns/op	  175224 B/op	       3 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   12784	    466364 ns/op	  176421 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   12740	    474725 ns/op	  177033 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   12520	    482108 ns/op	  181627 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   12261	    515945 ns/op	  180195 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_generated-4           	    9523	    606111 ns/op	  433784 B/op	    1502 allocs/op
Benchmark_easyjson_Marshal-4                   	    9230	    645829 ns/op	  174755 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	    9375	    648427 ns/op	  174632 B/op	      16 allocs/op

Benchmark_msgpack_vmihailenco_Unmarshal-4      	   12834	    451457 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	    8332	    724316 ns/op	  214862 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	    7791	    769478 ns/op	  214862 B/op	    1010 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	    7406	    809209 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	    7057	    857163 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	    6249	    963354 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    6184	    988357 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    5713	   1079993 ns/op	  603479 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    5607	   1097914 ns/op	  613786 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    5356	   1126027 ns/op	  603479 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    5127	   1157987 ns/op	  613787 B/op	    2816 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2751	   2165396 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2690	   2149069 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2316	   2623919 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2316	   2577286 ns/op	  349664 B/op	    1659 allocs/op
PASS
ok  	github.com/riftbit/protocol_benches	175.739s
```


### golang 1.11.13 linux (travis ci)
```

```


### golang 1.12.10 linux (travis ci)
```

```

### golang 1.13.1 linux (travis ci)
```

```

### golang tip linux (travis ci) - go version devel +a0894ea5b5 Mon Oct 7 18:50:14 2019 +0000 linux/amd64
```

```
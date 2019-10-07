# protocol benches

I decided to make some benchmarks to compare few libraries. And results are shocked me.
Решил провести небольшие сравнительные тесты по скорости работы разных библиотек для Golang. Результаты меня, мягко говоря, шокировали. 

Easyjson в результате тестов оказался самый медленный при Marshal, причем даже стандартная библиотека encode/json (в тесте Benchmark_std_*) работает на маршал значительно быстрее (на 50%)
Так же он оказался медленнее старенького ffjson (в случае прегенерации кода и при не использовании вложенных структур, а типов).

Что я делаю не так??? Может быть в коде что-то не то?
Пулл-реквесты приветствуются.

Тесты с пометкой Min - это minified json, у него убраны переносы строк и лишние пробелы. Full - это отформатированный json с отступами и переносами.

```
goos: windows
goarch: amd64
pkg: github.com/riftbit/jsonBenches

Benchmark_msgpack_vmihailenco_Marshal-8        	  191079	     31730 ns/op	  164080 B/op	       6 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-8      	   13440	    446057 ns/op	  175394 B/op	       3 allocs/op
Benchmark_jsoniterator_Marshal_fast-8          	   12955	    470166 ns/op	  182647 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-8              	   12889	    464815 ns/op	  182091 B/op	       2 allocs/op
Benchmark_std_Marshal-8                        	   12679	    466598 ns/op	  183459 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-8               	   12549	    476452 ns/op	  183039 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_generated-8           	    9835	    596645 ns/op	  434371 B/op	    1502 allocs/op
Benchmark_easyjson_Marshal-8                   	    8823	    675847 ns/op	  174991 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-8             	    8570	    679581 ns/op	  175119 B/op	      16 allocs/op

Benchmark_msgpack_vmihailenco_Unmarshal-8      	   14190	    428048 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-8     	    8218	    737041 ns/op	  214990 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-8    	    7594	    784303 ns/op	  214991 B/op	    1010 allocs/op
Benchmark_easyjson_Unmarshal_Min-8             	    7406	    816904 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-8       	    7406	    820686 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-8            	    6184	    958279 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-8      	    6315	    966588 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-8         	    5606	   1061006 ns/op	  613877 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-8        	    5308	   1143557 ns/op	  613877 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-8   	    5503	   1128289 ns/op	  603566 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-8    	    5824	   1071429 ns/op	  603566 B/op	    2108 allocs/op
Benchmark_std_Unmarshal_Min-8                  	    2702	   2174685 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-8        	    2690	   2210781 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-8       	    2271	   2639807 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-8                 	    2238	   2572388 ns/op	  349665 B/op	    1659 allocs/op

PASS
ok  	github.com/riftbit/jsonBenches	174.960s
?   	github.com/riftbit/jsonBenches/structs	[no test files]
?   	github.com/riftbit/jsonBenches/vars	[no test files]
```
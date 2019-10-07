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

## Benchmark results

### golang 1.13 on windows 10 x64

```bash
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
```


### golang 1.11.13 linux (travis ci)

```bash
Benchmark_msgpack_vmihailenco_Marshal-4        	  200000	     36088 ns/op	  164080 B/op	       5 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   20000	    483571 ns/op	  178474 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   20000	    497696 ns/op	  179118 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   10000	    551037 ns/op	  177714 B/op	       3 allocs/op
Benchmark_std_Marshal-4                        	   10000	    551902 ns/op	  176970 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   10000	    560600 ns/op	  177264 B/op	       2 allocs/op
Benchmark_easyjson_Marshal-4                   	   10000	    694860 ns/op	  173646 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	   10000	    704054 ns/op	  173638 B/op	      16 allocs/op
Benchmark_ffjson_Marshal_generated-4           	   10000	    732218 ns/op	  433329 B/op	    1502 allocs/op

Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    534900 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	   10000	    832883 ns/op	  214932 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	   10000	    910375 ns/op	  214933 B/op	    1011 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	   10000	    928399 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	   10000	    938205 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	   10000	   1113408 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	   10000	   1167865 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    5000	   1234806 ns/op	  605105 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    5000	   1254914 ns/op	  615412 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    5000	   1298931 ns/op	  605105 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    5000	   1320642 ns/op	  615412 B/op	    2877 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    3000	   2993079 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2000	   3008480 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2000	   3429728 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2000	   3470857 ns/op	  349408 B/op	    1655 allocs/op
```


### golang 1.12.10 linux (travis ci)

```bash
Benchmark_msgpack_vmihailenco_Marshal-4        	  200000	     33898 ns/op	  164080 B/op	       6 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   20000	    443825 ns/op	  176237 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   20000	    445067 ns/op	  176185 B/op	       3 allocs/op
Benchmark_std_Marshal-4                        	   20000	    448964 ns/op	  176529 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   20000	    463498 ns/op	  177788 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   20000	    473627 ns/op	  179246 B/op	       2 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	   10000	    597866 ns/op	  173750 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	   10000	    608611 ns/op	  173778 B/op	      16 allocs/op
Benchmark_ffjson_Marshal_generated-4           	   10000	    651874 ns/op	  434165 B/op	    1502 allocs/op

Benchmark_msgpack_vmihailenco_Unmarshal-4      	   20000	    493865 ns/op	  217745 B/op	    2914 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	   10000	    811723 ns/op	  214991 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	   10000	    861470 ns/op	  214966 B/op	    1011 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	   10000	    893878 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	   10000	    904700 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	   10000	   1051424 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	   10000	   1052504 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	   10000	   1188355 ns/op	  603478 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    5000	   1221668 ns/op	  613788 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    5000	   1252551 ns/op	  603478 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    5000	   1271024 ns/op	  613786 B/op	    2816 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    3000	   2884785 ns/op	  349412 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    3000	   2896155 ns/op	  349412 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2000	   3305551 ns/op	  349413 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2000	   3321658 ns/op	  349414 B/op	    1655 allocs/op
```

### golang 1.13.1 linux (travis ci)

```bash
Benchmark_msgpack_vmihailenco_Marshal-4        	  153388	     39698 ns/op	  164080 B/op	       6 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   14192	    418450 ns/op	  174859 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   13009	    462408 ns/op	  176679 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   12814	    465422 ns/op	  177760 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   12415	    486271 ns/op	  177159 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   12141	    490734 ns/op	  176437 B/op	       2 allocs/op
Benchmark_easyjson_Marshal-4                   	    9525	    627824 ns/op	  173848 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	    9368	    619305 ns/op	  173815 B/op	      16 allocs/op
Benchmark_ffjson_Marshal_generated-4           	    8662	    701555 ns/op	  432370 B/op	    1501 allocs/op

Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    524200 ns/op	  217745 B/op	    2914 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	    6592	    914319 ns/op	  214831 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	    6544	    844495 ns/op	  214801 B/op	    1010 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	    6540	    910280 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	    6380	    913498 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	    5630	   1067339 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    5630	   1079019 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    4677	   1284890 ns/op	  603466 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    4548	   1316086 ns/op	  613774 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    4438	   1340666 ns/op	  603467 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    4018	   1379697 ns/op	  613773 B/op	    2816 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2724	   2211588 ns/op	  349667 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2674	   2251020 ns/op	  349667 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2287	   2609672 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2284	   2618486 ns/op	  349667 B/op	    1659 allocs/op
```

### golang tip linux (travis ci) - go version devel +a0894ea5b5 Mon Oct 7 18:50:14 2019 +0000 linux/amd64

```bash
Benchmark_msgpack_vmihailenco_Marshal-4        	  146298	     39127 ns/op	  164080 B/op	       6 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   14431	    410855 ns/op	  174937 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   13166	    455004 ns/op	  177902 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   13023	    459394 ns/op	  177531 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   12457	    484857 ns/op	  178183 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   12064	    493877 ns/op	  177727 B/op	       2 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	    9774	    627446 ns/op	  173909 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	    9507	    629377 ns/op	  173814 B/op	      16 allocs/op
Benchmark_ffjson_Marshal_generated-4           	    8812	    694523 ns/op	  432278 B/op	    1501 allocs/op

Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    527506 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	    6900	    865734 ns/op	  214820 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	    6679	    909629 ns/op	  214829 B/op	    1010 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	    6673	    909896 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	    6519	    911578 ns/op	  320291 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    5665	   1077787 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	    5439	   1074901 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    4651	   1302185 ns/op	  613772 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    4557	   1275244 ns/op	  603462 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    4440	   1338005 ns/op	  603464 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    4266	   1360309 ns/op	  613773 B/op	    2816 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2776	   2176288 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2772	   2201790 ns/op	  349667 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2314	   2586071 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2302	   2554858 ns/op	  349668 B/op	    1659 allocs/op
```

## Compare across go versions (only changed values)

### 1.11 to 1.12

```bash
benchmark                                        old ns/op     new ns/op     delta
Benchmark_easyjson_Marshal-4                     694860        608611        -12.41%
Benchmark_easyjson_Marshal_lexer-4               704054        597866        -15.08%
Benchmark_easyjson_Unmarshal_Full-4              1167865       1052504       -9.88%
Benchmark_easyjson_Unmarshal_Min-4               938205        904700        -3.57%
Benchmark_easyjson_Unmarshal_Full_lexer-4        1113408       1051424       -5.57%
Benchmark_easyjson_Unmarshal_Min_lexer-4         928399        893878        -3.72%
Benchmark_ffjson_Marshal_simple-4                560600        443825        -20.83%
Benchmark_ffjson_Marshal_simple_pooling-4        551037        445067        -19.23%
Benchmark_ffjson_Marshal_generated-4             732218        651874        -10.97%
Benchmark_ffjson_Unmarshal_Full_simple-4         3470857       3321658       -4.30%
Benchmark_ffjson_Unmarshal_Min_simple-4          3008480       2884785       -4.11%
Benchmark_ffjson_Unmarshal_Full_generated-4      910375        861470        -5.37%
Benchmark_ffjson_Unmarshal_Min_generated-4       832883        811723        -2.54%
Benchmark_jsoniterator_Marshal-4                 497696        473627        -4.84%
Benchmark_jsoniterator_Marshal_fast-4            483571        463498        -4.15%
Benchmark_jsoniterator_Unmarshal_Full-4          1320642       1271024       -3.76%
Benchmark_jsoniterator_Unmarshal_Min-4           1254914       1221668       -2.65%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     1298931       1252551       -3.57%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      1234806       1188355       -3.76%
Benchmark_msgpack_vmihailenco_Marshal-4          36088         33898         -6.07%
Benchmark_msgpack_vmihailenco_Unmarshal-4        534900        493865        -7.67%
Benchmark_std_Marshal-4                          551902        448964        -18.65%
Benchmark_std_Unmarshal_Full-4                   3429728       3305551       -3.62%
Benchmark_std_Unmarshal_Min-4                    2993079       2896155       -3.24%

benchmark                                        old allocs     new allocs     delta
Benchmark_jsoniterator_Unmarshal_Full-4          2877           2816           -2.12%
Benchmark_jsoniterator_Unmarshal_Min-4           2877           2816           -2.12%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     2169           2108           -2.81%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      2169           2108           -2.81%
Benchmark_msgpack_vmihailenco_Marshal-4          5              6              +20.00%

benchmark                                        old bytes     new bytes     delta
Benchmark_easyjson_Marshal-4                     173646        173778        +0.08%
Benchmark_easyjson_Marshal_lexer-4               173638        173750        +0.06%
Benchmark_easyjson_Unmarshal_Min-4               320288        320289        +0.00%
Benchmark_easyjson_Unmarshal_Full_lexer-4        320289        320290        +0.00%
Benchmark_easyjson_Unmarshal_Min_lexer-4         320288        320290        +0.00%
Benchmark_ffjson_Marshal_simple-4                177264        176237        -0.58%
Benchmark_ffjson_Marshal_simple_pooling-4        177714        176185        -0.86%
Benchmark_ffjson_Marshal_generated-4             433329        434165        +0.19%
Benchmark_ffjson_Unmarshal_Full_simple-4         349408        349414        +0.00%
Benchmark_ffjson_Unmarshal_Min_simple-4          349408        349412        +0.00%
Benchmark_ffjson_Unmarshal_Full_generated-4      214933        214966        +0.02%
Benchmark_ffjson_Unmarshal_Min_generated-4       214932        214991        +0.03%
Benchmark_jsoniterator_Marshal-4                 179118        179246        +0.07%
Benchmark_jsoniterator_Marshal_fast-4            178474        177788        -0.38%
Benchmark_jsoniterator_Unmarshal_Full-4          615412        613786        -0.26%
Benchmark_jsoniterator_Unmarshal_Min-4           615412        613788        -0.26%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     605105        603478        -0.27%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      605105        603478        -0.27%
Benchmark_msgpack_vmihailenco_Unmarshal-4        217744        217745        +0.00%
Benchmark_std_Marshal-4                          176970        176529        -0.25%
Benchmark_std_Unmarshal_Full-4                   349408        349413        +0.00%
Benchmark_std_Unmarshal_Min-4                    349408        349412        +0.00%
```

### 1.12 to 1.13

```bash
benchmark                                        old ns/op     new ns/op     delta
Benchmark_easyjson_Marshal-4                     608611        627824        +3.16%
Benchmark_easyjson_Marshal_lexer-4               597866        619305        +3.59%
Benchmark_easyjson_Unmarshal_Full-4              1052504       1067339       +1.41%
Benchmark_easyjson_Unmarshal_Min-4               904700        910280        +0.62%
Benchmark_easyjson_Unmarshal_Full_lexer-4        1051424       1079019       +2.62%
Benchmark_easyjson_Unmarshal_Min_lexer-4         893878        913498        +2.19%
Benchmark_ffjson_Marshal_simple-4                443825        462408        +4.19%
Benchmark_ffjson_Marshal_simple_pooling-4        445067        418450        -5.98%
Benchmark_ffjson_Marshal_generated-4             651874        701555        +7.62%
Benchmark_ffjson_Unmarshal_Full_simple-4         3321658       2609672       -21.43%
Benchmark_ffjson_Unmarshal_Min_simple-4          2884785       2251020       -21.97%
Benchmark_ffjson_Unmarshal_Full_generated-4      861470        914319        +6.13%
Benchmark_ffjson_Unmarshal_Min_generated-4       811723        844495        +4.04%
Benchmark_jsoniterator_Marshal-4                 473627        490734        +3.61%
Benchmark_jsoniterator_Marshal_fast-4            463498        486271        +4.91%
Benchmark_jsoniterator_Unmarshal_Full-4          1271024       1379697       +8.55%
Benchmark_jsoniterator_Unmarshal_Min-4           1221668       1316086       +7.73%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     1252551       1340666       +7.03%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      1188355       1284890       +8.12%
Benchmark_msgpack_vmihailenco_Marshal-4          33898         39698         +17.11%
Benchmark_msgpack_vmihailenco_Unmarshal-4        493865        524200        +6.14%
Benchmark_std_Marshal-4                          448964        465422        +3.67%
Benchmark_std_Unmarshal_Full-4                   3305551       2618486       -20.79%
Benchmark_std_Unmarshal_Min-4                    2896155       2211588       -23.64%

benchmark                                       old allocs     new allocs     delta
Benchmark_ffjson_Marshal_generated-4            1502           1501           -0.07%
Benchmark_ffjson_Unmarshal_Full_simple-4        1655           1659           +0.24%
Benchmark_ffjson_Unmarshal_Min_simple-4         1655           1659           +0.24%
Benchmark_ffjson_Unmarshal_Full_generated-4     1011           1010           -0.10%
Benchmark_ffjson_Unmarshal_Min_generated-4      1011           1010           -0.10%
Benchmark_std_Unmarshal_Full-4                  1655           1659           +0.24%
Benchmark_std_Unmarshal_Min-4                   1655           1659           +0.24%

benchmark                                        old bytes     new bytes     delta
Benchmark_easyjson_Marshal-4                     173778        173848        +0.04%
Benchmark_easyjson_Marshal_lexer-4               173750        173815        +0.04%
Benchmark_easyjson_Unmarshal_Min-4               320289        320290        +0.00%
Benchmark_ffjson_Marshal_simple-4                176237        176679        +0.25%
Benchmark_ffjson_Marshal_simple_pooling-4        176185        174859        -0.75%
Benchmark_ffjson_Marshal_generated-4             434165        432370        -0.41%
Benchmark_ffjson_Unmarshal_Full_simple-4         349414        349668        +0.07%
Benchmark_ffjson_Unmarshal_Min_simple-4          349412        349667        +0.07%
Benchmark_ffjson_Unmarshal_Full_generated-4      214966        214831        -0.06%
Benchmark_ffjson_Unmarshal_Min_generated-4       214991        214801        -0.09%
Benchmark_jsoniterator_Marshal-4                 179246        176437        -1.57%
Benchmark_jsoniterator_Marshal_fast-4            177788        177159        -0.35%
Benchmark_jsoniterator_Unmarshal_Full-4          613786        613773        -0.00%
Benchmark_jsoniterator_Unmarshal_Min-4           613788        613774        -0.00%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     603478        603467        -0.00%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      603478        603466        -0.00%
Benchmark_std_Marshal-4                          176529        177760        +0.70%
Benchmark_std_Unmarshal_Full-4                   349413        349667        +0.07%
Benchmark_std_Unmarshal_Min-4                    349412        349667        +0.07%
```

### 1.13 to tip

```bash
benchmark                                        old ns/op     new ns/op     delta
Benchmark_easyjson_Marshal-4                     627824        629377        +0.25%
Benchmark_easyjson_Marshal_lexer-4               619305        627446        +1.31%
Benchmark_easyjson_Unmarshal_Full-4              1067339       1074901       +0.71%
Benchmark_easyjson_Unmarshal_Min-4               910280        909896        -0.04%
Benchmark_easyjson_Unmarshal_Full_lexer-4        1079019       1077787       -0.11%
Benchmark_easyjson_Unmarshal_Min_lexer-4         913498        911578        -0.21%
Benchmark_ffjson_Marshal_simple-4                462408        455004        -1.60%
Benchmark_ffjson_Marshal_simple_pooling-4        418450        410855        -1.82%
Benchmark_ffjson_Marshal_generated-4             701555        694523        -1.00%
Benchmark_ffjson_Unmarshal_Full_simple-4         2609672       2554858       -2.10%
Benchmark_ffjson_Unmarshal_Min_simple-4          2251020       2176288       -3.32%
Benchmark_ffjson_Unmarshal_Full_generated-4      914319        909629        -0.51%
Benchmark_ffjson_Unmarshal_Min_generated-4       844495        865734        +2.51%
Benchmark_jsoniterator_Marshal-4                 490734        493877        +0.64%
Benchmark_jsoniterator_Marshal_fast-4            486271        484857        -0.29%
Benchmark_jsoniterator_Unmarshal_Full-4          1379697       1360309       -1.41%
Benchmark_jsoniterator_Unmarshal_Min-4           1316086       1302185       -1.06%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     1340666       1338005       -0.20%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      1284890       1275244       -0.75%
Benchmark_msgpack_vmihailenco_Marshal-4          39698         39127         -1.44%
Benchmark_msgpack_vmihailenco_Unmarshal-4        524200        527506        +0.63%
Benchmark_std_Marshal-4                          465422        459394        -1.30%
Benchmark_std_Unmarshal_Full-4                   2618486       2586071       -1.24%
Benchmark_std_Unmarshal_Min-4                    2211588       2201790       -0.44%

benchmark                                        old bytes     new bytes     delta
Benchmark_easyjson_Marshal-4                     173848        173814        -0.02%
Benchmark_easyjson_Marshal_lexer-4               173815        173909        +0.05%
Benchmark_easyjson_Unmarshal_Min-4               320290        320289        -0.00%
Benchmark_easyjson_Unmarshal_Min_lexer-4         320290        320291        +0.00%
Benchmark_ffjson_Marshal_simple-4                176679        177902        +0.69%
Benchmark_ffjson_Marshal_simple_pooling-4        174859        174937        +0.04%
Benchmark_ffjson_Marshal_generated-4             432370        432278        -0.02%
Benchmark_ffjson_Unmarshal_Min_simple-4          349667        349668        +0.00%
Benchmark_ffjson_Unmarshal_Full_generated-4      214831        214829        -0.00%
Benchmark_ffjson_Unmarshal_Min_generated-4       214801        214820        +0.01%
Benchmark_jsoniterator_Marshal-4                 176437        177727        +0.73%
Benchmark_jsoniterator_Marshal_fast-4            177159        178183        +0.58%
Benchmark_jsoniterator_Unmarshal_Min-4           613774        613772        -0.00%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     603467        603464        -0.00%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      603466        603462        -0.00%
Benchmark_msgpack_vmihailenco_Unmarshal-4        217745        217744        -0.00%
Benchmark_std_Marshal-4                          177760        177531        -0.13%
Benchmark_std_Unmarshal_Full-4                   349667        349666        -0.00%

```

### 1.11 to 1.13

```bash
benchmark                                        old ns/op     new ns/op     delta
Benchmark_easyjson_Marshal-4                     694860        627824        -9.65%
Benchmark_easyjson_Marshal_lexer-4               704054        619305        -12.04%
Benchmark_easyjson_Unmarshal_Full-4              1167865       1067339       -8.61%
Benchmark_easyjson_Unmarshal_Min-4               938205        910280        -2.98%
Benchmark_easyjson_Unmarshal_Full_lexer-4        1113408       1079019       -3.09%
Benchmark_easyjson_Unmarshal_Min_lexer-4         928399        913498        -1.61%
Benchmark_ffjson_Marshal_simple-4                560600        462408        -17.52%
Benchmark_ffjson_Marshal_simple_pooling-4        551037        418450        -24.06%
Benchmark_ffjson_Marshal_generated-4             732218        701555        -4.19%
Benchmark_ffjson_Unmarshal_Full_simple-4         3470857       2609672       -24.81%
Benchmark_ffjson_Unmarshal_Min_simple-4          3008480       2251020       -25.18%
Benchmark_ffjson_Unmarshal_Full_generated-4      910375        914319        +0.43%
Benchmark_ffjson_Unmarshal_Min_generated-4       832883        844495        +1.39%
Benchmark_jsoniterator_Marshal-4                 497696        490734        -1.40%
Benchmark_jsoniterator_Marshal_fast-4            483571        486271        +0.56%
Benchmark_jsoniterator_Unmarshal_Full-4          1320642       1379697       +4.47%
Benchmark_jsoniterator_Unmarshal_Min-4           1254914       1316086       +4.87%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     1298931       1340666       +3.21%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      1234806       1284890       +4.06%
Benchmark_msgpack_vmihailenco_Marshal-4          36088         39698         +10.00%
Benchmark_msgpack_vmihailenco_Unmarshal-4        534900        524200        -2.00%
Benchmark_std_Marshal-4                          551902        465422        -15.67%
Benchmark_std_Unmarshal_Full-4                   3429728       2618486       -23.65%
Benchmark_std_Unmarshal_Min-4                    2993079       2211588       -26.11%

benchmark                                        old allocs     new allocs     delta
Benchmark_ffjson_Marshal_generated-4             1502           1501           -0.07%
Benchmark_ffjson_Unmarshal_Full_simple-4         1655           1659           +0.24%
Benchmark_ffjson_Unmarshal_Min_simple-4          1655           1659           +0.24%
Benchmark_ffjson_Unmarshal_Full_generated-4      1011           1010           -0.10%
Benchmark_ffjson_Unmarshal_Min_generated-4       1011           1010           -0.10%
Benchmark_jsoniterator_Unmarshal_Full-4          2877           2816           -2.12%
Benchmark_jsoniterator_Unmarshal_Min-4           2877           2816           -2.12%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     2169           2108           -2.81%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      2169           2108           -2.81%
Benchmark_msgpack_vmihailenco_Marshal-4          5              6              +20.00%
Benchmark_std_Unmarshal_Full-4                   1655           1659           +0.24%
Benchmark_std_Unmarshal_Min-4                    1655           1659           +0.24%

benchmark                                        old bytes     new bytes     delta
Benchmark_easyjson_Marshal-4                     173646        173848        +0.12%
Benchmark_easyjson_Marshal_lexer-4               173638        173815        +0.10%
Benchmark_easyjson_Unmarshal_Min-4               320288        320290        +0.00%
Benchmark_easyjson_Unmarshal_Full_lexer-4        320289        320290        +0.00%
Benchmark_easyjson_Unmarshal_Min_lexer-4         320288        320290        +0.00%
Benchmark_ffjson_Marshal_simple-4                177264        176679        -0.33%
Benchmark_ffjson_Marshal_simple_pooling-4        177714        174859        -1.61%
Benchmark_ffjson_Marshal_generated-4             433329        432370        -0.22%
Benchmark_ffjson_Unmarshal_Full_simple-4         349408        349668        +0.07%
Benchmark_ffjson_Unmarshal_Min_simple-4          349408        349667        +0.07%
Benchmark_ffjson_Unmarshal_Full_generated-4      214933        214831        -0.05%
Benchmark_ffjson_Unmarshal_Min_generated-4       214932        214801        -0.06%
Benchmark_jsoniterator_Marshal-4                 179118        176437        -1.50%
Benchmark_jsoniterator_Marshal_fast-4            178474        177159        -0.74%
Benchmark_jsoniterator_Unmarshal_Full-4          615412        613773        -0.27%
Benchmark_jsoniterator_Unmarshal_Min-4           615412        613774        -0.27%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     605105        603467        -0.27%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      605105        603466        -0.27%
Benchmark_msgpack_vmihailenco_Unmarshal-4        217744        217745        +0.00%
Benchmark_std_Marshal-4                          176970        177760        +0.45%
Benchmark_std_Unmarshal_Full-4                   349408        349667        +0.07%
Benchmark_std_Unmarshal_Min-4                    349408        349667        +0.07%
```

### 1.11 to tip

```bash
benchmark                                        old ns/op     new ns/op     delta
Benchmark_easyjson_Marshal-4                     694860        629377        -9.42%
Benchmark_easyjson_Marshal_lexer-4               704054        627446        -10.88%
Benchmark_easyjson_Unmarshal_Full-4              1167865       1074901       -7.96%
Benchmark_easyjson_Unmarshal_Min-4               938205        909896        -3.02%
Benchmark_easyjson_Unmarshal_Full_lexer-4        1113408       1077787       -3.20%
Benchmark_easyjson_Unmarshal_Min_lexer-4         928399        911578        -1.81%
Benchmark_ffjson_Marshal_simple-4                560600        455004        -18.84%
Benchmark_ffjson_Marshal_simple_pooling-4        551037        410855        -25.44%
Benchmark_ffjson_Marshal_generated-4             732218        694523        -5.15%
Benchmark_ffjson_Unmarshal_Full_simple-4         3470857       2554858       -26.39%
Benchmark_ffjson_Unmarshal_Min_simple-4          3008480       2176288       -27.66%
Benchmark_ffjson_Unmarshal_Full_generated-4      910375        909629        -0.08%
Benchmark_ffjson_Unmarshal_Min_generated-4       832883        865734        +3.94%
Benchmark_jsoniterator_Marshal-4                 497696        493877        -0.77%
Benchmark_jsoniterator_Marshal_fast-4            483571        484857        +0.27%
Benchmark_jsoniterator_Unmarshal_Full-4          1320642       1360309       +3.00%
Benchmark_jsoniterator_Unmarshal_Min-4           1254914       1302185       +3.77%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     1298931       1338005       +3.01%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      1234806       1275244       +3.27%
Benchmark_msgpack_vmihailenco_Marshal-4          36088         39127         +8.42%
Benchmark_msgpack_vmihailenco_Unmarshal-4        534900        527506        -1.38%
Benchmark_std_Marshal-4                          551902        459394        -16.76%
Benchmark_std_Unmarshal_Full-4                   3429728       2586071       -24.60%
Benchmark_std_Unmarshal_Min-4                    2993079       2201790       -26.44%

benchmark                                        old allocs     new allocs     delta
Benchmark_ffjson_Marshal_generated-4             1502           1501           -0.07%
Benchmark_ffjson_Unmarshal_Full_simple-4         1655           1659           +0.24%
Benchmark_ffjson_Unmarshal_Min_simple-4          1655           1659           +0.24%
Benchmark_ffjson_Unmarshal_Full_generated-4      1011           1010           -0.10%
Benchmark_ffjson_Unmarshal_Min_generated-4       1011           1010           -0.10%
Benchmark_jsoniterator_Unmarshal_Full-4          2877           2816           -2.12%
Benchmark_jsoniterator_Unmarshal_Min-4           2877           2816           -2.12%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     2169           2108           -2.81%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      2169           2108           -2.81%
Benchmark_msgpack_vmihailenco_Marshal-4          5              6              +20.00%
Benchmark_std_Unmarshal_Full-4                   1655           1659           +0.24%
Benchmark_std_Unmarshal_Min-4                    1655           1659           +0.24%

benchmark                                        old bytes     new bytes     delta
Benchmark_easyjson_Marshal-4                     173646        173814        +0.10%
Benchmark_easyjson_Marshal_lexer-4               173638        173909        +0.16%
Benchmark_easyjson_Unmarshal_Min-4               320288        320289        +0.00%
Benchmark_easyjson_Unmarshal_Full_lexer-4        320289        320290        +0.00%
Benchmark_easyjson_Unmarshal_Min_lexer-4         320288        320291        +0.00%
Benchmark_ffjson_Marshal_simple-4                177264        177902        +0.36%
Benchmark_ffjson_Marshal_simple_pooling-4        177714        174937        -1.56%
Benchmark_ffjson_Marshal_generated-4             433329        432278        -0.24%
Benchmark_ffjson_Unmarshal_Full_simple-4         349408        349668        +0.07%
Benchmark_ffjson_Unmarshal_Min_simple-4          349408        349668        +0.07%
Benchmark_ffjson_Unmarshal_Full_generated-4      214933        214829        -0.05%
Benchmark_ffjson_Unmarshal_Min_generated-4       214932        214820        -0.05%
Benchmark_jsoniterator_Marshal-4                 179118        177727        -0.78%
Benchmark_jsoniterator_Marshal_fast-4            178474        178183        -0.16%
Benchmark_jsoniterator_Unmarshal_Full-4          615412        613773        -0.27%
Benchmark_jsoniterator_Unmarshal_Min-4           615412        613772        -0.27%
Benchmark_jsoniterator_Unmarshal_Full_fast-4     605105        603464        -0.27%
Benchmark_jsoniterator_Unmarshal_Min_fast-4      605105        603462        -0.27%
Benchmark_std_Marshal-4                          176970        177531        +0.32%
Benchmark_std_Unmarshal_Full-4                   349408        349666        +0.07%
Benchmark_std_Unmarshal_Min-4                    349408        349667        +0.07%
```
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

### golang 1.13 on windows 10 x64 (ssd hdd)
```
goos: windows
goarch: amd64
pkg: github.com/riftbit/protocol_benches
Benchmark_easyjson_Marshal                     	    8332	    707274 ns/op	  173384 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	    9090	    651376 ns/op	  174765 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-8                   	    9229	    658034 ns/op	  175027 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer               	    8107	    723942 ns/op	  173384 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	    9091	    677593 ns/op	  174517 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-8             	    9230	    660022 ns/op	  175126 B/op	      16 allocs/op
Benchmark_easyjson_Unmarshal_Full              	    5940	   1133502 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	    6060	   1011882 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-8            	    5998	   1002835 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min               	    5940	    953199 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	    7228	    846015 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-8             	    7058	    836072 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer        	    5606	   1087763 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    6184	    975744 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-8      	    6121	    980721 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer         	    6592	    923240 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	    7498	    818752 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-8       	    7406	    816230 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_ffjson_Marshal_simple                	   12894	    456336 ns/op	  174722 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   12380	    478837 ns/op	  183222 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-8              	   12994	    467678 ns/op	  182845 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling        	   13177	    447977 ns/op	  174778 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   13434	    457273 ns/op	  175304 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-8      	   13428	    448392 ns/op	  175395 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_generated             	    9229	    641999 ns/op	  428072 B/op	    1500 allocs/op
Benchmark_ffjson_Marshal_generated-4           	   10000	    576700 ns/op	  434338 B/op	    1502 allocs/op
Benchmark_ffjson_Marshal_generated-8           	   10000	    580500 ns/op	  434617 B/op	    1502 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple         	    2221	   2763619 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2289	   2585847 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-8       	    2343	   2618003 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple          	    2563	   2413190 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2739	   2178899 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-8        	    2764	   2280028 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated      	    7500	    827600 ns/op	  214519 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	    7402	    808700 ns/op	  214854 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-8    	    7692	    812402 ns/op	  215005 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated       	    7692	    784321 ns/op	  214518 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	    8107	    770075 ns/op	  214855 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-8     	    8218	    751035 ns/op	  215000 B/op	    1010 allocs/op
Benchmark_jsoniterator_Marshal                 	   10000	    512300 ns/op	  174723 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   12429	    482822 ns/op	  177987 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-8               	   12523	    480156 ns/op	  184083 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast            	   10000	    504400 ns/op	  174723 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   12561	    478385 ns/op	  176655 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-8          	   12396	    478622 ns/op	  182933 B/op	       2 allocs/op
Benchmark_jsoniterator_Unmarshal_Full          	    4687	   1358439 ns/op	  613706 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    5134	   1217959 ns/op	  613787 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-8        	    5356	   1209298 ns/op	  613877 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min           	    4316	   1182113 ns/op	  613706 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    5404	   1083827 ns/op	  613786 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-8         	    5262	   1162106 ns/op	  613877 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast     	    4411	   1254364 ns/op	  603402 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    5172	   1171307 ns/op	  603478 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-8   	    4761	   1164251 ns/op	  603565 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast      	    4544	   1270467 ns/op	  603402 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    5308	   1120949 ns/op	  603477 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-8    	    5606	   1118266 ns/op	  603564 B/op	    2108 allocs/op
Benchmark_msgpack_vmihailenco_Marshal          	  151513	     35337 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-4        	  146697	     37063 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-8        	  169970	     36495 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal        	   10000	    513300 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-4      	   13742	    431669 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-8      	   13904	    427215 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_std_Marshal                          	   12986	    464577 ns/op	  174723 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   12289	    490032 ns/op	  182090 B/op	       2 allocs/op
Benchmark_std_Marshal-8                        	   12692	    483691 ns/op	  182619 B/op	       2 allocs/op
Benchmark_std_Unmarshal_Full                   	    2127	   2883404 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2197	   2611746 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-8                 	    2316	   2626081 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min                    	    2316	   2490934 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2596	   2317796 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-8                  	    2438	   2278917 ns/op	  349666 B/op	    1659 allocs/op
PASS
ok  	github.com/riftbit/protocol_benches	516.756s
```


### golang 1.11.13 linux (travis ci)
```
goos: linux
goarch: amd64
pkg: github.com/riftbit/protocol_benches
Benchmark_easyjson_Marshal                     	   10000	    671863 ns/op	  173492 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	   10000	    639857 ns/op	  173595 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-8                   	   10000	    655085 ns/op	  174560 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer               	   10000	    669908 ns/op	  173420 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	   10000	    639954 ns/op	  173622 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-8             	   10000	    646417 ns/op	  174421 B/op	      16 allocs/op
Benchmark_easyjson_Unmarshal_Full              	   10000	   1097880 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	   10000	   1041487 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-8            	   10000	   1062925 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min               	   10000	    950158 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	   10000	    883135 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-8             	   10000	    892970 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer        	   10000	   1100583 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	   10000	   1052554 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-8      	   10000	   1062175 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer         	   10000	    936831 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	   10000	    867748 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-8       	   10000	    886021 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_ffjson_Marshal_simple                	   10000	    537988 ns/op	  175544 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   10000	    548174 ns/op	  178144 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-8              	   10000	    557820 ns/op	  186551 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling        	   10000	    541074 ns/op	  175165 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   10000	    549680 ns/op	  178476 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-8      	   10000	    558651 ns/op	  186497 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_generated             	   10000	    671007 ns/op	  429614 B/op	    1500 allocs/op
Benchmark_ffjson_Marshal_generated-4           	   10000	    684101 ns/op	  434610 B/op	    1502 allocs/op
Benchmark_ffjson_Marshal_generated-8           	   10000	    696994 ns/op	  435103 B/op	    1502 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple         	    2000	   3575307 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2000	   3438493 ns/op	  349409 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-8       	    2000	   3506322 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple          	    2000	   3142175 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2000	   2997570 ns/op	  349409 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-8        	    2000	   3019313 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated      	   10000	    922248 ns/op	  214764 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	   10000	    903394 ns/op	  214929 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-8    	   10000	    915755 ns/op	  215089 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated       	   10000	    861211 ns/op	  214760 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	   10000	    837388 ns/op	  214928 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-8     	   10000	    850143 ns/op	  215096 B/op	    1011 allocs/op
Benchmark_jsoniterator_Marshal                 	   10000	    508526 ns/op	  174809 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   20000	    496680 ns/op	  178646 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-8               	   10000	    510416 ns/op	  196470 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast            	   20000	    496029 ns/op	  175366 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   20000	    482175 ns/op	  177658 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-8          	   20000	    497078 ns/op	  186689 B/op	       2 allocs/op
Benchmark_jsoniterator_Unmarshal_Full          	    5000	   1351878 ns/op	  615337 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    5000	   1367216 ns/op	  615411 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-8        	    5000	   1369324 ns/op	  615497 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Min           	    5000	   1285233 ns/op	  615337 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    5000	   1266591 ns/op	  615412 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-8         	    5000	   1294884 ns/op	  615497 B/op	    2877 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast     	    5000	   1325579 ns/op	  605033 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    5000	   1327240 ns/op	  605105 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-8   	    5000	   1345249 ns/op	  605188 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast      	    5000	   1272656 ns/op	  605033 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    5000	   1251008 ns/op	  605105 B/op	    2169 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-8    	    5000	   1276379 ns/op	  605189 B/op	    2169 allocs/op
Benchmark_msgpack_vmihailenco_Marshal          	  200000	     36879 ns/op	  164080 B/op	       5 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-4        	  200000	     37300 ns/op	  164080 B/op	       5 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-8        	  200000	     41548 ns/op	  164080 B/op	       5 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal        	   10000	    596998 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    538852 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-8      	   10000	    549708 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_std_Marshal                          	   10000	    546099 ns/op	  174899 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   10000	    551888 ns/op	  176441 B/op	       2 allocs/op
Benchmark_std_Marshal-8                        	   10000	    563170 ns/op	  183031 B/op	       2 allocs/op
Benchmark_std_Unmarshal_Full                   	    2000	   3559540 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2000	   3438774 ns/op	  349409 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Full-8                 	    2000	   3476588 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min                    	    2000	   3191906 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2000	   3009291 ns/op	  349409 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min-8                  	    2000	   3043760 ns/op	  349408 B/op	    1655 allocs/op
PASS
ok  	github.com/riftbit/protocol_benches	557.856s
```


### golang 1.12.10 linux (travis ci)
```
goos: linux
goarch: amd64
pkg: github.com/riftbit/protocol_benches
Benchmark_easyjson_Marshal                     	   10000	    660039 ns/op	  173420 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	   10000	    623010 ns/op	  174022 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-8                   	   10000	    626703 ns/op	  174092 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer               	   10000	    666962 ns/op	  173400 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	   10000	    650877 ns/op	  173886 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-8             	   10000	    621871 ns/op	  174032 B/op	      16 allocs/op
Benchmark_easyjson_Unmarshal_Full              	   10000	   1167213 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	   10000	   1120061 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-8            	   10000	   1151448 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min               	   10000	   1040834 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	   10000	    968843 ns/op	  320291 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-8             	   10000	    964545 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer        	   10000	   1209451 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    5000	   1169616 ns/op	  320292 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-8      	   10000	   1144476 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer         	   10000	   1046530 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	   10000	    967989 ns/op	  320291 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-8       	   10000	    964801 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_ffjson_Marshal_simple                	   20000	    457917 ns/op	  174869 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   20000	    474142 ns/op	  176531 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-8              	   20000	    471574 ns/op	  183415 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling        	   20000	    455135 ns/op	  174921 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   20000	    469345 ns/op	  176917 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-8      	   10000	    503221 ns/op	  183390 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_generated             	   10000	    786108 ns/op	  429763 B/op	    1500 allocs/op
Benchmark_ffjson_Marshal_generated-4           	   10000	    703522 ns/op	  438041 B/op	    1503 allocs/op
Benchmark_ffjson_Marshal_generated-8           	   10000	    748950 ns/op	  436301 B/op	    1503 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple         	    2000	   3576568 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2000	   3485584 ns/op	  349412 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-8       	    2000	   3496325 ns/op	  349410 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple          	    2000	   3280902 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2000	   3056371 ns/op	  349413 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-8        	    3000	   2925945 ns/op	  349409 B/op	    1655 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated      	   10000	    918070 ns/op	  214767 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	   10000	    921316 ns/op	  215025 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-8    	   10000	    919648 ns/op	  215128 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated       	   10000	    896870 ns/op	  214760 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	   10000	    857144 ns/op	  214991 B/op	    1011 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-8     	   10000	    834355 ns/op	  215103 B/op	    1011 allocs/op
Benchmark_jsoniterator_Marshal                 	   10000	    539617 ns/op	  175581 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   10000	    513644 ns/op	  181092 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-8               	   10000	    513635 ns/op	  184545 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast            	   10000	    512659 ns/op	  174809 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   20000	    494451 ns/op	  179718 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-8          	   20000	    482160 ns/op	  189735 B/op	       2 allocs/op
Benchmark_jsoniterator_Unmarshal_Full          	    5000	   1398610 ns/op	  613706 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    5000	   1374351 ns/op	  613781 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-8        	    5000	   1386204 ns/op	  613870 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min           	    5000	   1419255 ns/op	  613706 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    5000	   1313172 ns/op	  613781 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-8         	    5000	   1335261 ns/op	  613870 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast     	    5000	   1380553 ns/op	  603401 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    5000	   1325752 ns/op	  603473 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-8   	    5000	   1412850 ns/op	  603561 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast      	    5000	   1371775 ns/op	  603401 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    5000	   1324933 ns/op	  603473 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-8    	    5000	   1342838 ns/op	  603559 B/op	    2108 allocs/op
Benchmark_msgpack_vmihailenco_Marshal          	  200000	     40584 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-4        	  200000	     38998 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-8        	   10000	    633021 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    544395 ns/op	  217746 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-8      	   10000	    567901 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_std_Marshal                          	   20000	    466803 ns/op	  175016 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   10000	    504000 ns/op	  180838 B/op	       2 allocs/op
Benchmark_std_Marshal-8                        	   10000	    500202 ns/op	  180275 B/op	       2 allocs/op
Benchmark_std_Unmarshal_Full                   	    2000	   3505632 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2000	   3472047 ns/op	  349411 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Full-8                 	    2000	   3488302 ns/op	  349410 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min                    	    2000	   3207523 ns/op	  349408 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2000	   3035851 ns/op	  349412 B/op	    1655 allocs/op
Benchmark_std_Unmarshal_Min-8                  	    2000	   2948831 ns/op	  349409 B/op	    1655 allocs/op
PASS
ok  	github.com/riftbit/protocol_benches	604.426s
```

### golang 1.13.1 linux (travis ci)
```
goos: linux
goarch: amd64
pkg: github.com/riftbit/protocol_benches
Benchmark_easyjson_Marshal                     	    8890	    651375 ns/op	  173382 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	    9439	    613754 ns/op	  173882 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-8                   	    9570	    615241 ns/op	  173928 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer               	    9218	    654341 ns/op	  173382 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	    9525	    618758 ns/op	  173817 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-8             	    9679	    608630 ns/op	  173880 B/op	      16 allocs/op
Benchmark_easyjson_Unmarshal_Full              	    5636	   1089631 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	    5619	   1060478 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-8            	    5665	   1061604 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min               	    6627	    946368 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	    6508	    901911 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-8             	    6342	    893759 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer        	    5587	   1094529 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    5539	   1066460 ns/op	  320291 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-8      	    5562	   1069073 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer         	    6645	    943780 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	    6704	    909832 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-8       	    6430	    903438 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_ffjson_Marshal_simple                	   13734	    435489 ns/op	  174722 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   13058	    458638 ns/op	  176672 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-8              	   12580	    473172 ns/op	  181614 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling        	   14065	    418184 ns/op	  174777 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   14330	    417505 ns/op	  174899 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-8      	   14390	    414899 ns/op	  174987 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_generated             	    8961	    669564 ns/op	  428092 B/op	    1500 allocs/op
Benchmark_ffjson_Marshal_generated-4           	    8653	    684078 ns/op	  432527 B/op	    1501 allocs/op
Benchmark_ffjson_Marshal_generated-8           	    8823	    680591 ns/op	  433040 B/op	    1502 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple         	    2239	   2651017 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2306	   2594276 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-8       	    2311	   2599736 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple          	    2676	   2246941 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2733	   2194152 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-8        	    2775	   2173742 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated      	    6894	    887573 ns/op	  214522 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	    6721	    890108 ns/op	  214798 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-8    	    6268	    890750 ns/op	  215010 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated       	    7407	    830191 ns/op	  214521 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	    7128	    832680 ns/op	  214813 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-8     	    7282	    830902 ns/op	  215023 B/op	    1010 allocs/op
Benchmark_jsoniterator_Marshal                 	   12337	    486709 ns/op	  174723 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   12351	    494514 ns/op	  178214 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-8               	   10000	    506876 ns/op	  180252 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast            	   12446	    481618 ns/op	  174723 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   12408	    479663 ns/op	  177436 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-8          	   12314	    488677 ns/op	  180962 B/op	       2 allocs/op
Benchmark_jsoniterator_Unmarshal_Full          	    4419	   1342109 ns/op	  613705 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    4431	   1350518 ns/op	  613774 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-8        	    4252	   1369936 ns/op	  613858 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min           	    4671	   1274179 ns/op	  613705 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    4725	   1295117 ns/op	  613771 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-8         	    4594	   1315302 ns/op	  613858 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast     	    4510	   1303554 ns/op	  603400 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    4524	   1307715 ns/op	  603467 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-8   	    4602	   1345298 ns/op	  603545 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast      	    4728	   1259832 ns/op	  603400 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    4600	   1272532 ns/op	  603466 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-8    	    4598	   1279453 ns/op	  603544 B/op	    2108 allocs/op
Benchmark_msgpack_vmihailenco_Marshal          	  172640	     32503 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-4        	  171348	     36004 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-8        	  149588	     40587 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal        	    9921	    574285 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    525644 ns/op	  217746 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-8      	   10000	    528410 ns/op	  217745 B/op	    2914 allocs/op
Benchmark_std_Marshal                          	   13467	    446608 ns/op	  174723 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   12564	    474894 ns/op	  176700 B/op	       2 allocs/op
Benchmark_std_Marshal-8                        	   12688	    474153 ns/op	  179936 B/op	       2 allocs/op
Benchmark_std_Unmarshal_Full                   	    2217	   2716848 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2264	   2619880 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-8                 	    2300	   2631120 ns/op	  349667 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min                    	    2616	   2262099 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2721	   2204920 ns/op	  349667 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-8                  	    2690	   2225012 ns/op	  349666 B/op	    1659 allocs/op
PASS
ok  	github.com/riftbit/protocol_benches	501.357s
```

### golang tip linux (travis ci) - go version devel +a0894ea5b5 Mon Oct 7 18:50:14 2019 +0000 linux/amd64
```
goos: linux
goarch: amd64
pkg: github.com/riftbit/protocol_benches
Benchmark_easyjson_Marshal                     	    8608	    665497 ns/op	  173383 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-4                   	    9626	    638992 ns/op	  173758 B/op	      16 allocs/op
Benchmark_easyjson_Marshal-8                   	    9354	    634134 ns/op	  174021 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer               	    9328	    666628 ns/op	  173383 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-4             	    9675	    637596 ns/op	  173909 B/op	      16 allocs/op
Benchmark_easyjson_Marshal_lexer-8             	    9424	    633786 ns/op	  174032 B/op	      16 allocs/op
Benchmark_easyjson_Unmarshal_Full              	    5522	   1129473 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-4            	    5511	   1089187 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full-8            	    5427	   1090213 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min               	    6439	    973150 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-4             	    6289	    928195 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min-8             	    6570	    937497 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer        	    5442	   1136706 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-4      	    5521	   1100443 ns/op	  320291 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Full_lexer-8      	    5482	   1111180 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer         	    6304	    979115 ns/op	  320288 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-4       	    6391	    937977 ns/op	  320290 B/op	    1146 allocs/op
Benchmark_easyjson_Unmarshal_Min_lexer-8       	    6570	    930832 ns/op	  320289 B/op	    1146 allocs/op
Benchmark_ffjson_Marshal_simple                	   13519	    450867 ns/op	  174723 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-4              	   12862	    464258 ns/op	  177521 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple-8              	   12772	    473420 ns/op	  180042 B/op	       2 allocs/op
Benchmark_ffjson_Marshal_simple_pooling        	   14227	    414040 ns/op	  174778 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-4      	   14326	    412641 ns/op	  174982 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_simple_pooling-8      	   14505	    413741 ns/op	  175268 B/op	       3 allocs/op
Benchmark_ffjson_Marshal_generated             	    8475	    685443 ns/op	  428057 B/op	    1500 allocs/op
Benchmark_ffjson_Marshal_generated-4           	    8078	    722435 ns/op	  432560 B/op	    1501 allocs/op
Benchmark_ffjson_Marshal_generated-8           	    8256	    741354 ns/op	  433663 B/op	    1502 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple         	    2236	   2694240 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-4       	    2246	   2727344 ns/op	  349667 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_simple-8       	    2194	   2613250 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple          	    2684	   2263082 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-4        	    2686	   2248729 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Min_simple-8        	    2678	   2232470 ns/op	  349665 B/op	    1659 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated      	    6715	    929395 ns/op	  214520 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-4    	    6621	    921652 ns/op	  214808 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Full_generated-8    	    6514	    940288 ns/op	  215028 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated       	    7086	    875134 ns/op	  214520 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-4     	    6732	    869943 ns/op	  214827 B/op	    1010 allocs/op
Benchmark_ffjson_Unmarshal_Min_generated-8     	    6771	    879518 ns/op	  215020 B/op	    1010 allocs/op
Benchmark_jsoniterator_Marshal                 	   10000	    502477 ns/op	  174723 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-4               	   10000	    502024 ns/op	  179115 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal-8               	   10000	    511471 ns/op	  183855 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast            	   12013	    500936 ns/op	  174723 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-4          	   12114	    499178 ns/op	  177714 B/op	       2 allocs/op
Benchmark_jsoniterator_Marshal_fast-8          	   10000	    514709 ns/op	  184798 B/op	       2 allocs/op
Benchmark_jsoniterator_Unmarshal_Full          	    4318	   1394501 ns/op	  613705 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-4        	    4202	   1411116 ns/op	  613773 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full-8        	    4210	   1444718 ns/op	  613858 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min           	    4431	   1361386 ns/op	  613705 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-4         	    4418	   1351432 ns/op	  613773 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Min-8         	    4352	   1395998 ns/op	  613857 B/op	    2816 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast     	    4432	   1374864 ns/op	  603402 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Full_fast-4   	    4207	   1429342 ns/op	  603546 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast      	    4448	   1317903 ns/op	  603401 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-4    	    4365	   1321470 ns/op	  603463 B/op	    2108 allocs/op
Benchmark_jsoniterator_Unmarshal_Min_fast-8    	    4363	   1365278 ns/op	  603546 B/op	    2108 allocs/op
Benchmark_msgpack_vmihailenco_Marshal          	  149264	     34789 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-4        	  139072	     42924 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Marshal-8        	  123862	     49073 ns/op	  164080 B/op	       6 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal        	    9382	    593332 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-4      	   10000	    542608 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_msgpack_vmihailenco_Unmarshal-8      	   10000	    554475 ns/op	  217744 B/op	    2914 allocs/op
Benchmark_std_Marshal                          	   13200	    452864 ns/op	  174723 B/op	       2 allocs/op
Benchmark_std_Marshal-4                        	   12774	    471077 ns/op	  177034 B/op	       2 allocs/op
Benchmark_std_Marshal-8                        	   12492	    477549 ns/op	  180252 B/op	       2 allocs/op
Benchmark_std_Unmarshal_Full                   	    2210	   2695364 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-4                 	    2310	   2631863 ns/op	  349668 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Full-8                 	    2323	   2629961 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min                    	    2622	   2270267 ns/op	  349664 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-4                  	    2684	   2236143 ns/op	  349666 B/op	    1659 allocs/op
Benchmark_std_Unmarshal_Min-8                  	    2689	   2285286 ns/op	  349665 B/op	    1659 allocs/op
PASS
ok  	github.com/riftbit/protocol_benches	487.359s
```
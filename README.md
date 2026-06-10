# 100 Go задач без решений

Содержимое:
- `slice/` — 50 задач на слайсы;
- `mapslice/` — 50 задач на map + slice;
- к каждой задаче есть отдельный файл `_test.go`;
- реализации отсутствуют: внутри функций только `TODO` и zero value return.

Запуск тестов:

```bash
go test ./...
```

Индекс задач:
1. `slice.CountEven` — возвращает количество чётных чисел в src.
2. `slice.CountOdd` — возвращает количество нечётных чисел в src.
3. `slice.SumNegative` — возвращает сумму всех отрицательных чисел из src.
4. `slice.SumNonNegative` — возвращает сумму чисел, которые больше либо равны нулю.
5. `slice.ProductNonZero` — возвращает произведение всех ненулевых элементов.
6. `slice.MaxValue` — возвращает максимальное значение в src.
7. `slice.MinValue` — возвращает минимальное значение в src.
8. `slice.ContainsValue` — сообщает, встречается ли target в src.
9. `slice.AllPositive` — возвращает true, если все элементы src строго положительные.
10. `slice.AnyNegative` — возвращает true, если в src есть хотя бы одно отрицательное число.
11. `slice.FirstPositiveIndex` — возвращает индекс первого положительного числа.
12. `slice.LastIndex` — возвращает индекс последнего вхождения target в src.
13. `slice.RemoveAll` — возвращает новый слайс без всех элементов, равных target.
14. `slice.FilterEven` — возвращает новый слайс только с чётными числами в исходном порядке.
15. `slice.FilterOdd` — возвращает новый слайс только с нечётными числами в исходном порядке.
16. `slice.FilterPositive` — возвращает новый слайс только со строго положительными числами.
17. `slice.DoubleSlice` — возвращает новый слайс, где каждый элемент src умножен на 2.
18. `slice.AbsSlice` — возвращает новый слайс из модулей элементов src.
19. `slice.PrefixSums` — возвращает слайс префиксных сумм: элемент i равен сумме src[0:i+1].
20. `slice.RunningMax` — возвращает слайс текущих максимумов при проходе слева направо.
21. `slice.RotateLeft` — возвращает новый слайс, циклически сдвинутый влево на k позиций.
22. `slice.RotateRight` — возвращает новый слайс, циклически сдвинутый вправо на k позиций.
23. `slice.IsSortedAsc` — сообщает, отсортирован ли src по неубыванию.
24. `slice.IsSortedDesc` — сообщает, отсортирован ли src по невозрастанию.
25. `slice.UniqueInts` — возвращает уникальные значения в порядке первого появления.
26. `slice.IntersectInts` — возвращает уникальные элементы, которые есть и в a, и в b, в порядке первого появления в a.
27. `slice.DifferenceInts` — возвращает элементы из a, которых нет в b.
28. `slice.MoveZerosEnd` — возвращает новый слайс, где все нули перенесены в конец, а порядок ненулевых элементов сохранён.
29. `slice.RemoveDuplicatesSorted` — для отсортированного src удаляет соседние дубликаты и возвращает уникальные элементы.
30. `slice.SecondMax` — возвращает второй по величине уникальный элемент.
31. `slice.CountGreaterThan` — возвращает количество элементов src, которые строго больше limit.
32. `slice.CountLessThan` — возвращает количество элементов src, которые строго меньше limit.
33. `slice.CountInRange` — возвращает количество элементов в диапазоне [left, right] включительно.
34. `slice.ClampSlice` — возвращает новый слайс, где значения меньше min заменены на min, а больше max — на max.
35. `slice.RepeatElements` — повторяет каждый элемент src times раз подряд.
36. `slice.TakeEverySecond` — возвращает элементы с индексами 0, 2, 4 и так далее.
37. `slice.FlattenMatrix` — разворачивает матрицу в один слайс построчно.
38. `slice.MatrixRowSums` — возвращает суммы строк матрицы.
39. `slice.DiagonalSum` — возвращает сумму главной диагонали квадратной или прямоугольной матрицы, пока существуют и строка, и столбец.
40. `slice.TransposeRect` — транспонирует прямоугольную матрицу.
41. `slice.MergeAlternating` — возвращает новый слайс, чередуя элементы a и b.
42. `slice.PadRight` — возвращает новый слайс длиной не меньше size, дополняя справа value.
43. `slice.TrimZerosEdges` — удаляет нули только с начала и конца слайса.
44. `slice.DropEveryNth` — возвращает слайс без каждого n-го элемента при счёте с 1.
45. `slice.IndexesOf` — возвращает все индексы, на которых встречается target.
46. `slice.ReverseCopy` — возвращает новый слайс с элементами src в обратном порядке.
47. `slice.InsertAt` — возвращает новый слайс, где value вставлен по index.
48. `slice.RemoveAt` — возвращает новый слайс без элемента по index.
49. `slice.ReplaceRange` — возвращает копию src, где элементы с индексами [from, to) заменены на value.
50. `slice.SliceWithoutNegatives` — возвращает новый слайс без отрицательных чисел.
51. `mapslice.SumMapValues` — возвращает сумму всех значений map.
52. `mapslice.HasKey` — сообщает, есть ли key в map.
53. `mapslice.MergeCounts` — объединяет две map со счётчиками, складывая значения одинаковых ключей.
54. `mapslice.FilterPositiveMap` — возвращает новую map только с парами, где значение строго положительное.
55. `mapslice.WordLengthsMap` — возвращает map, где ключ — слово, значение — длина слова в байтах.
56. `mapslice.CountByLength` — считает, сколько строк каждой длины встречается в words.
57. `mapslice.FirstIndexMap` — возвращает map значение -> индекс первого появления этого значения.
58. `mapslice.LastIndexMap` — возвращает map значение -> индекс последнего появления этого значения.
59. `mapslice.PositionsMap` — возвращает map значение -> все индексы этого значения в порядке возрастания.
60. `mapslice.AreDisjoint` — возвращает true, если у двух слайсов нет общих значений.
61. `mapslice.IsSubset` — возвращает true, если каждое уникальное значение subset встречается в set.
62. `mapslice.UniqueStrings` — возвращает уникальные строки в порядке первого появления.
63. `mapslice.DuplicateStrings` — возвращает строки, которые встречаются больше одного раза, в порядке их первого повторного появления.
64. `mapslice.CountUniqueInts` — возвращает количество уникальных значений в src.
65. `mapslice.CountUniqueStrings` — возвращает количество уникальных строк в words.
66. `mapslice.RemoveBannedInts` — возвращает элементы src, которых нет в banned или для которых banned[value] == false.
67. `mapslice.KeepAllowedInts` — возвращает элементы src, для которых allowed[value] == true.
68. `mapslice.StringFirstIndexes` — возвращает map строка -> индекс первого появления.
69. `mapslice.StringLastIndexes` — возвращает map строка -> индекс последнего появления.
70. `mapslice.SameStringFrequencies` — возвращает true, если в a и b одинаковые частоты строк.
71. `mapslice.IsPermutationInts` — возвращает true, если b является перестановкой a с учётом количества повторов.
72. `mapslice.MissingNumbersInRange` — возвращает числа из диапазона [left,right], которых нет в src, по возрастанию.
73. `mapslice.CommonStrings` — возвращает уникальные строки, которые есть и в a, и в b, в порядке первого появления в a.
74. `mapslice.CountDuplicatedInts` — возвращает количество различных значений, которые встречаются больше одного раза.
75. `mapslice.IntFrequency` — возвращает map значение -> количество вхождений.
76. `mapslice.StringFrequency` — возвращает map строка -> количество вхождений.
77. `mapslice.MostFrequentInt` — возвращает самое частое число.
78. `mapslice.IncrementMapValues` — возвращает новую map, где к каждому значению прибавлен delta.
79. `mapslice.KeysWithValue` — возвращает ключи, у которых значение равно value, в лексикографическом порядке.
80. `mapslice.EqualStringIntMaps` — сообщает, равны ли две map по набору ключей и значений.
81. `mapslice.MapFromPairs` — создаёт map из параллельных слайсов keys и values.
82. `mapslice.CountTrueValues` — возвращает количество ключей со значением true.
83. `mapslice.RemoveKeys` — возвращает копию m без ключей из keys.
84. `mapslice.OnlyKeys` — возвращает новую map только с указанными ключами, если они есть в m.
85. `mapslice.GroupByParity` — группирует числа по чётности в ключи "even" и "odd".
86. `mapslice.GroupByFirstLetter` — группирует непустые строки по первому байту.
87. `mapslice.IndexWordsByLength` — возвращает map длина слова -> индексы слов этой длины.
88. `mapslice.CountWordsWithPrefix` — возвращает количество строк, которые начинаются с prefix.
89. `mapslice.RemoveStopWords` — возвращает words без строк, для которых stop[word] == true.
90. `mapslice.SetUnionInts` — возвращает объединение уникальных значений: сначала новые элементы из a, затем новые элементы из b.
91. `mapslice.SetIntersectionStrings` — возвращает уникальные строки, которые есть в обоих слайсах, в порядке первого появления в a.
92. `mapslice.SymmetricDifferenceInts` — возвращает уникальные значения, которые встречаются только в одном из двух слайсов.
93. `mapslice.BuildIntSet` — строит set из src в виде map[int]bool, где каждое встреченное значение имеет true.
94. `mapslice.BuildStringSet` — строит set из words в виде map[string]bool.
95. `mapslice.ToggleIntSet` — возвращает копию set, последовательно переключая значения: если value был true — становится false/удаляется, иначе становится true.
96. `mapslice.RenameMapKey` — возвращает копию m, где значение oldKey перенесено в newKey.
97. `mapslice.MapValuesToKeys` — возвращает map значение -> слайс ключей с этим значением, ключи внутри каждого слайса должны быть отсортированы лексикографически.
98. `mapslice.CountPairSums` — возвращает количество пар индексов i < j, для которых src[i] + src[j] == target.
99. `mapslice.CanMakePairSum` — возвращает true, если существуют два разных индекса с суммой target.
100. `mapslice.CountWordsByLastLetter` — считает непустые строки по последнему байту.

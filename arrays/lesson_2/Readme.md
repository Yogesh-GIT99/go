# slices
1. selecting parts of arrays using slices
2. more ways of selecting slices

# Deep Dive look on slices: ( ref code ): 
- slice refers to a array, golang doesnot create new mem location for slices, it simply ref slice to the array, from which it is created.
- slices one sliced from left cannot be retained back during operation ( though it still exist in array), but the right side spaces are always accessible and len can be increased anytime to its full capacity ( ref code)
# how to use go routines in a actual project. 
1. go routines donot support return a value ( Note: simple return can be used ), there is different way of handling err in go routines. ( ref code)
2. Managing channels with select statemets ( to handle errorChan case )
3. deferring code execution with "defer" ( to stop closing file manaully)
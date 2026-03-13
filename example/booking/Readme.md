# creating a booking platform

preparing stmt vs direct exec/query stmt:

1. using prepare stmt is 100% optional. you can directly use exec/query.
2. But when executing multiple statements, preparing stmt optimise execution to a great extent, leads to better performance ( given stmt is closed at the end of multiple executions not at every executions ). for a single run it makes no difference. 
3. use exec when you want to manipulate data in database and use query when you want to fetch data. 
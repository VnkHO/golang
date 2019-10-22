# The Reader Interface

| Source of Input                              |     | Returns ?!? |     | To Print It...            |
| -------------------------------------------- | --- | ----------- | --- | ------------------------- |
| HTTP Request Body                            | ->  | []flargen   | ->  | func printHTTP([]flargen) |
| Text file on hard drive                      | ->  | []string    | ->  | func printFile([]string)  |
| Image file on hard drive                     | ->  | jpegne      | ->  | func printImage(jpegne)   |
| User entering text into command line         | ->  | []byte      | ->  | func printText([]byte)    |
| Data from analog sensor plugged into machine | ->  | []float     | ->  | func printData([]float)   |

Source of Input ------> Reader -------> []byte (Output data that anyone can work with)

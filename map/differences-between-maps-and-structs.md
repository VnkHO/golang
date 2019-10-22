# Differences Between Maps and Structs

| Map                                                 | Struct                                                        |
| --------------------------------------------------- | ------------------------------------------------------------- |
| All keys must be the same type                      | Values can be of different type                               |
| All value must be the same type                     | Keys don't support indexing                                   |
| Keys are indexed - we can iterate over them         | Value Type!                                                   |
| Use to represent a collection of related properties | You need to know all the different fields at compile time     |
| Don't need to know all the keys at compile time     | Use to represent a "thing" with a lot of different properties |
| Reference Type!                                     |                                                               |

General guidance is that you will want to use a **map** whenever you are **representing a collection of very closely related properties**.

# GoDatabase

## Main Parts
1. Start with a *B+tree*, the data structure for querying and manipulating the data.
2. Make it *durable*, that’s what makes a DB different from a file.
3. Relational DB with *concurrent* transactions on top of the copy-on-write B+tree KV.
4. A SQL-like *query language*, the finishing touch.

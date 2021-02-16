## S3 as a database

The project aims to replace a traditional database, although i would consider Amazon S3 as a database in that it is simply a large key-value store and effectivley operates as a no-frills no-SQL database where the Key is the filename, and the Value is the contents of the file.
Considering that the needs of this project are to simply store json files via a GUID and retreive them the same way, S3 is actually great choice for the POC as it is significantly cheaper than the alternatives.
In fact I have read that orders older than 12 months on amazon are actually stored in S3, but i would like to clarify that this IS a database in all but name , and as such i feel that it is a great choice and i would't suggest this strategy be changed, but we shouldn't be under the illusion that we have completley satisfied the principle premise of the POC.


You are "considering using AWS S3 bucket instead of a NoSQL database", but the fact is that Amazon S3 effectively is a NoSQL database.

It is a very large Key-Value store. The Key is the filename, the Value is the contents of the file.

If your needs are simply "Store a value with this key" and "Retrieve a value with this key", then it would work just fine!

In fact, old orders on Amazon.com (more than a year old) are apparently archived to Amazon S3 since they are read-only (no returns, no changes).

While slower than DynamoDB, Amazon S3 certainly costs significantly less for storage!
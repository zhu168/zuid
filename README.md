#zuid

zuid is a unique id generator, and the generated result is a sortable 16byte array or 32byte string.

The implementation feature is that the first 8byte is the Snowflake algorithm, and the last 8byte is the last 8byte of the uuid.
The advantage of this algorithm is that it has most of the advantages of Snowflake, especially it is relatively friendly to the database primary key index, and obtains a certain security of uuid, while limiting the length, and is easy to implement and transplant.

```go
func main() {
	//If you need to set the timestamp of the Snowflake algorithm in advance, 
	//the default is 2023-08-30 00:00:00 UTC
	Epoch = uint64(1682908800000)
	// Create a ZUID instance and pass in the worker node ID
	zuid, err := NewZUID(1)
	if err != nil {
		fmt.Println("Error creating ZUID instance:", err)
		return
	}
	for i := 0; i < 10; i++ {
		//Generate a zuid
		id := zuid.NextIDSimple()
		fmt.Println("ZUID:", id, len(id))
	}
}

```
output
```
ZUID: 009c4c6e67801000a1dd70157184a3be 32
ZUID: 009c4c6e67801001982a7155b2a3333d 32
ZUID: 009c4c6e67801002808694c9f5516a4c 32
ZUID: 009c4c6e67801003a66f9761247c534f 32
ZUID: 009c4c6e67801004a223b62fff7f8e89 32
ZUID: 009c4c6e67801005b8c6eeb091c2feb5 32
ZUID: 009c4c6e6780100694c17fd9ef31e17e 32
ZUID: 009c4c6e6780100780290d0b87a0bec0 32
ZUID: 009c4c6e67801008b96d2a77c3de3b61 32
ZUID: 009c4c6e678010099ba963cc9f041271 32
```
# NBT

Maybe This is a simple NBT parser.

## Gets

You can get the package with go get command.

```shell
go get -u github.com/beito123/nbt
```

*-u is a option updating the package*

## License

This is licensed by MIT License.
See LICENSE file.

## Examples

### Read

```go
func main() {
	// Data: Unnamed(compound){ world_name(string): jagajaga, hardcore(byte): 18 }

	// Raw data
	// 	0x0000: 0a 00 00 01 00 08 68 61 72 64 63 6f | ......hardco
	// 	0x000c: 72 65 12 08 00 0a 77 6f 72 6c 64 5f | re....world_
	// 	0x0018: 6e 61 6d 65 00 08 6a 61 67 61 6a 61 | name..jagaja
	// 	0x0024: 67 61 00                            | ga.

	// Compressed bytes // level.dat and more are saved as compressed with gzip
	data := []byte{0x1F, 0x8B, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x03, 0xE3, 0x62, 0x60, 0x60, 0x64, 0xE0, 0xC8, 0x48, 0x2C, 0x4A,
		0x49, 0xCE, 0x2F, 0x4A, 0x15, 0xE2, 0x60, 0xE0, 0x2A, 0xCF, 0x2F,
		0xCA, 0x49, 0x89, 0xCF, 0x4B, 0xCC, 0x4D, 0x65, 0xE0, 0xC8, 0x4A,
		0x4C, 0x4F, 0x04, 0x61, 0x06, 0x00, 0x7A, 0x8D, 0xB4, 0xF9, 0x27, 0x00, 0x00, 0x00}

	// Read after uncompressing
	// If you want to load compressed reader, you should use this
	stream, err := nbt.FromReader(bytes.NewBuffer(data), nbt.BigEndian)
	if err != nil {
		panic(err)
	}

	// If you read from file
	// stream, err := nbt.FromFile("./level.dat", nbt.BigEndian)

	// Read a tag from read bytes
	tag, err := stream.ReadTag()
	if err != nil {
		panic(err)
	}

	// Convert to String // For dump
	str, err := tag.ToString()
	if err != nil {
		panic(err)
	}

	// read data: { hardcore(Byte): 18, world_name(String): jagajaga }
	fmt.Printf("read data: %s", str)
}
```

### Write

```go
func main(){
	// tag = Unnamed(Compound){ hardcore(Byte): 18, world_name(String): jagajaga }
	tag := &nbt.Compound{
		Value: map[string]nbt.Tag{
			"hardcore": &nbt.Byte{
				Value: 18,
			},
			"world_name": &nbt.String{
				Value: "jagajaga",
			},
		},
	}

	// New stream with big endian
	// Big endian is used for level.dat and world data
	// Little endian is used for leveldb
	stream := nbt.NewStream(nbt.BigEndian)

	// Write a tag
	err := stream.WriteTag(tag)
	if err != nil {
		panic(err)
	}

	// Output file:
	// 0x0000: 0a 00 00 01 00 08 68 61 72 64 63 6f | ......hardco
	// 0x000c: 72 65 12 08 00 0a 77 6f 72 6c 64 5f | re....world_
	// 0x0018: 6e 61 6d 65 00 08 6a 61 67 61 6a 61 | name..jagaja
	// 0x0024: 67 61 00                            | ga.
	ioutil.WriteFile("./nbt.dat", stream.Bytes(), os.ModePerm)

	// If you want to compress for level.dat etc...
	data, err := nbt.Compress(stream)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./nbt_compressed.dat", data, os.ModePerm)
}
```

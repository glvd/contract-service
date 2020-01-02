output:
 ```
abigen --sol ./contract/sol/writeable.sol --pkg dwriteable --out ./contract/dwriteabl
e/dwriteable.go
```

```
java:
abigen --sol ./contract/sol/dmessage.sol --pkg com.dhash.dmessage --lang java --out ./
contract/dmessage/dmessage.java
go:
abigen --sol ./contract/sol/dmessage.sol --pkg dmessage --out ./contract/dmessage/dmes
sage.go
```

```
java:
abigen --sol ./contract/sol/dnode.sol --pkg com.dhash.dnode --lang java --out \
./contract/dnode/dnode.java
go:
abigen --sol ./contract/sol/dnode.sol --pkg dnode --out ./contract/dnode/dnode.go
```

```
java:
 abigen --sol ./contract/sol/dtag.sol --pkg com.dhash.dtag --lang java --out ./contract
 /dtag/dtag.java
go:
 abigen --pkg dtag --out ./contract/dtag/dtag.go --abi ./contract/sol/dtag.abi
```

```
java:
 abigen --sol ./contract/sol/dhashcoin.sol --pkg com.dhash.dhashcoin --lang java --out ./contract
 /dtag/dhashcoin.java
go:
 abigen --pkg dhashcoin --out ./contract/dhashcoin/dhashcoin.go  --sol ./contract/sol/dhashcoin.sol
```

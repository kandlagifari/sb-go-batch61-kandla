# Run the code using shell script

```shell
chmod +x run.sh
./run.sh
```

# API Testing

```shell
curl -sL -X GET http://localhost:30069/nilai-mahasiswa | jq
```

```shell
curl -sL -X POST http://localhost:30069/nilai-mahasiswa \
    -u admin:password123 \
    -H "Content-Type: application/json" \
    -d '{"Nama": "Luffy","MataKuliah": "One Piece","Nilai": 85}' | jq
```

![Alt text](gif/api-testing.gif)

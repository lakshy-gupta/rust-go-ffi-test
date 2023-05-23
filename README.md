# rust-go-ffi-test

```
cd mylib
cargo build --release
cd ..
go build -ldflags="-r mylib/target/release"
./rust-go-ffi-test
```


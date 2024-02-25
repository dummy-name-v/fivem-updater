## FiveM binaries with no extra dependencies updater

### Build
```
go build -o updater.exe
```

### Usage
> put the executable in your FiveM server directory
```
./updater.exe -windows -o path
```

- The platform flag is optional and will fallback to your OS ``(window/linux)``
- The out flag is optional and will fallback to the folder ``binaries`` in the current working directory

@echo off

if "%~3"=="" (
    echo Usage: %0 ^<proto_path^> ^<output_file^> ^<filename^>
    exit /b 1
)

set "proto_path=%~1"
set "output_file=%~2"
set "filename=%~3"

protoc --go_out="%output_file%" --go_opt=paths=source_relative --go-grpc_out="%output_file%" --go-grpc_opt=paths=source_relative "%filename%.proto"
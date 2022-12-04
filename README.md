# Scan-for-ipv4
## About project
This project is a learning proejct to develop my skills in golang. And learning about networking and ipv4 protocol.

It is basic and not that optimist but I am working on it

## Features
What can the project do
- Scan and ping public ip addresses
- Calculate all possible ips when connected to network
- Ping all ips on private network

## How it works
1. Download tar file that corresponds to your setup
2. Extract binary(executable) from tar file
3. Run the file from terminal 
4. It will display help section

## Example
### Linux
This will download and unpack the tar file
```bash
wget https://github.com/theredwiking/scan-for-ipv4/releases/download/v0.0.2/scan4_0.0.2_Linux_x86_64.tar.gz && tar -xvf scan4_0.0.2_Linux_x86_64.tar.gz
```

To run the program
```bash
./scan4
```

That will output following
```bash
Scan for all public or private IPv4 addresses
This tool is for scanning for all public IPv4 addresses
but it can also be used to scan for local IPv4 addresses

Usage:
  scan-for-ipv4 [command]

Available Commands:
  calculate   calculates all possible IPv4 addresses
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  private     For working with internal IPv4
  public      Scans for public IPv4 addresses

Flags:
  -h, --help     help for scan-for-ipv4
  -t, --toggle   Help message for toggle

Use "scan-for-ipv4 [command] --help" for more information about a command.
```
Here you can see all possible command and can get help how to use them
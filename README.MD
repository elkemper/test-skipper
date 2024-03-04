# test-skipper (tssk) - Manage skipped tests in the repository

Using commands 'skip' or 'unskip' you can easily modify the 'skipped.cfg' file that will contain your test id's in plain text comma separated format.

### Usage:
tssk skip <test_ids>  
tssk unskip <test_ids>

### Commands:
skip - Add tests to the `skipped.cfg` file  
unskip - Remove tests from the `skipped.cfg` file 

### Examples:
tssk skip 123 2453  
tssk skip 123, 234  
tssk unskip 123  

## Build 

```bash
go build -ldflags  -w -o tssk cmd/tssk/*.go 
```

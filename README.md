# Making a http call to an IPFS command using go
## Pre-requisite
- Vanilla IPFS must be installed in the system and its PATH must be set.
- The file to be read must be located in the MFS (here, loremipsum.txt in MFS containing the placeholder text).
- Run the "ipfs daemon".

## Procedure
- Write a program in go using net/http package.
- Use either Post or NewRequest function in which the kubo RPC API URL for the respecive IPFS command passed (here, ipfs files read --offset --count).
- Now, execute "go run" command for the program written.
- The IPFS command is called using http request.

*Note: Set offset and count value such that the file content is read in the powers of 2*


## Program

![Screenshot (93)](https://github.com/gayathrisuresh150501/go-http/assets/77912752/052b4daa-13c5-488b-b298-ce0f7ae48e80)



## Output

![Screenshot (94)](https://github.com/gayathrisuresh150501/go-http/assets/77912752/a40cb73e-18d5-448d-a3ca-ce20082d1777)

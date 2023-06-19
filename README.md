# Making a http call to an IPFS command using go
## Pre-requisite
- Vanilla IPFS must be installed in the system and its PATH must be set.
- The file to be read must be located in the MFS (here, loremipsum.txt in MFS containing the placeholder text).
- Run the "ipfs daemon".

## Procedure
- Write a program in go using net/http package.
- Use either "Post" or "NewRequest" function in which the kubo RPC API URL for the respective IPFS command passed (here, ipfs files read --offset --count).
- Now, execute "go run" command for the program written.
- The IPFS command is called using http request.

*Note: Set offset and count value such that the file content is read in the powers of 2*


## Program

![Screenshot (95)](https://github.com/gayathrisuresh150501/go-http/assets/77912752/44f152d5-b5b2-46cc-aa8c-f12bd539fb6c)



## Output

![Screenshot (96)](https://github.com/gayathrisuresh150501/go-http/assets/77912752/1330ee7b-f63c-41be-a9e4-a1f2133be54b)

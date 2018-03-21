/*
STEPS
1) socket establishment, 
2) case/switch for packet type/condition, 
3) buffer array check for one password, 
4) send file (handle little/bigEndian)
5) terminate connection/hash file


Protocol specification - according to example in C
---------------------
        case JOIN_REQ:
            break;
        case PASS_RESP:
            break;
        default:
            quitFlag = 0;
            break;
			
define the packet formats:
JOIN_REQ = 1
PASS_REQ = 2
PASS_ACCEPT = 3
DATA = 5
TERMINATE = 6 (has file digest/hash)
REJECT = 7
			
*/

package main






import (
// all support packages go here
	"crypto/sha1" //handle file hash
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv" //convert int to strings and vice versa
)

int main(int argc, char *argv[])
{


	/*
	NEED TO HANDLE PACKET CONNECTION IN GO, C CODE BELOW
	DEFINE AS FUNCTION
	*/
    // The packet counter
    unsigned int packetcou = 1;

	/*
	NEED TO HANDLE FILE HASHING
	DEFINE AS FUNCTION?
	*/
    // The result of the hash function
    char sha1_hash[20];

	/*
	NEED TO DEFINE UDP SOCKET CONNECTION
	DEFINE AS FUNCTION
	*/
	
    // Socket descriptor and additional variables for handling the connection
    int sockfd;
    struct addrinfo hints, *servinfo, *p;
    int rv, numbytes;

    // Instance of the union packet and while-flag
    Packets packet;
    int quitFlag = 1;

    // Needed for the sendto() and recvfrom()
    struct sockaddr_storage their_addr;
    socklen_t addr_len = sizeof their_addr;

    /*
    SETUP THE SOCKET CONNECTION,
    AND BE SURE THAT A CONNECTION IS ACTUALLY BEEN MADE.

    Hint: Use the methods getaddrinfo() socket() and bind()
    */

    /*
     COMMUNICATE WITH THE CLIENT,
     BE SURE TO HANDLE THE NETWORK BYTE ORDER

     Hint: Use the methods sendto(), recvfrom() and bzero()
    */

    while(quitFlag)
    {
        switch(ntohs(packet.ctrl_msg.header))
        {
        case JOIN_REQ:
            break;
        case PASS_RESP:
            break;
        default:
            quitFlag = 0;
            break;
        }
    }
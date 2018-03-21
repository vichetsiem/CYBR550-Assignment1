/*
STEPS
1) socket establishment, 
2) case/switch for packet type/condition, 
3) buffer array check for one password, 
4) send file (handle little/bigEndian)
5) terminate connection/hash file


Protocol specification - according to example in C
---------------------
        case PASS_REQ:
            break;
        case PASS_ACCEPT:
            break;
        case DATA:
            break;
        case TERMINATE:
            break;
        case REJECT:
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
    // Variable to contain the received data (easily handle on the heap)
    char *datamsg = (char *) malloc(sizeof datamsg);

    // The packet counter
    unsigned int packetcou = 1;

    // The result of the hash function
    char sha1_hash[20];

    // Socket descriptor and additional variables for handling the connection
    int sockfd;
    struct addrinfo hints, *servinfo, *p;
    int rv, numbytes;

    // Instance of the union packet and while-flag
    Packets packet;
    int quitFlag = 1;

    // Need for the sendto and recvfrom
    struct sockaddr_storage their_addr;
    socklen_t addr_len = sizeof their_addr;

    /*
    SETUP THE SOCKET CONNECTION,
    AND BE SURE THAT A CONNECTION IS ACTUALLY BEEN MADE.

    Hint: Use the methods getaddrinfo() and socket()
    */

    /*
     COMMUNICATE WITH THE SERVER,
     BE SURE TO HANDLE THE NETWORK BYTE ORDER

     Hint: Use the methods sendto(), recvfrom() and bzero()
    */
    while(quitFlag)
    {
        switch(ntohs(packet.ctrl_msg.header))
        {
        case PASS_REQ:
            break;
        case PASS_ACCEPT:
            break;
        case DATA:
            break;
        case TERMINATE:
            break;
        case REJECT:
            break;
        default:
            quitFlag = 0;
            break;
        }
    }

    // Do some cleaning
    freeaddrinfo(servinfo);
    free(datamsg);
    close(sockfd);
}
#ifndef _SOCKETS_H_
#define _SOCKETS_H_

#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h> 

int open_client( char *host_name, int port_no );
int open_server( int port_no );

#endif


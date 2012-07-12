#include "sockets.h"

int open_client( char *host_name, int port_no )
{

	int sock;
	struct sockaddr_in serv_addr;
	struct hostent * record;
	struct in_addr *address;

	record = gethostbyname( host_name );
	
	if( record == NULL )
	{
		fprintf( stderr, "Unable to find host\n" );
		exit(1);
	} 

	sock = socket( AF_INET, SOCK_STREAM, 0 );
		
	address = (struct in_addr *) record->h_addr;

	if( sock < 0 )
	{
		fprintf( stderr, "Unable to create client socket\nExitting program\n" );
		exit(1);
	}

	memset( &serv_addr, 0, sizeof( serv_addr ) );
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_addr = *address;
	serv_addr.sin_port = htons( port_no );

	if( connect( sock, (struct sockaddr *) &serv_addr, sizeof( serv_addr ) ) < 0 )
	{
		fprintf( stderr, "Error, was unable to connect to socket\n" );
		exit(1);
	}

	return sock;
}

int open_server( int port_no )
{

	int sock, newsock;
	int client_size;
	struct sockaddr_in serv_addr, cli_addr;

	sock = socket( AF_INET, SOCK_STREAM, 0 );

	client_size = sizeof( cli_addr );
	
	if( sock < 0 )
	{
		fprintf( stderr, "Unable to create server socket\n" );
		exit(1);
	}

	memset( &serv_addr, 0, sizeof( serv_addr ) );
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_addr.s_addr = htonl( INADDR_ANY );
	serv_addr.sin_port = htons( port_no );
	
	if( bind( sock, (struct sockadd *) &serv_addr, sizeof( serv_addr ) ) < 0 )
	{
		fprintf( stderr, "Failed to bind to port #: %d\n", port_no );
		exit(1);
	}
		
	if( listen( sock, 1 ) < 0 )
	{
		fprintf( stderr, "Failed to listen on port\n" );	
		exit(1);
	}

	newsock = accept( sock, (struct serv_addr *) &cli_addr, &client_size );
	if( newsock < 0 )
	{
		fprintf( stderr, "Unable to accept client connection\n" );
		exit(1);
	}

	return newsock;
}

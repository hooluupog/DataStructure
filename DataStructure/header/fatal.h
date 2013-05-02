/*
 * fatal.h
 *
 *  Created on: 2013-4-2
 *      Author: darkblue
 */

#ifndef FATAL_H_
#define FATAL_H_

#include "stdio.h"
#include "stdlib.h"

#define Error( Str )        FatalError( Str )
#define FatalError( Str )   fprintf( stderr, "%s\n", Str ), exit( 1 )



#endif /* FATAL_H_ */

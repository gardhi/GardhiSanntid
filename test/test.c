

#include <stdio.h>

#define BUFLEN 512

int main(){

	char buf[BUFLEN];
	char message[BUFLEN];

	while(1){
		printf("---------------------\n");
		printf("Input: ");
		fgets(message, BUFLEN, stdin);

		printf("---------------------\n");
		printf("Output: %s\n",message );
	}

}
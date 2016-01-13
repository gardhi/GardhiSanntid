

#include <stdio.h>
#include <pthread.h>

int count = 0;

void* dec(){

	//int *count_ptr = (int *)dec_ptr;

	for(int i = 0; i< 1000000; i++){
		count--;
	};
	return NULL;
};



void* inc(){

	//int *count_ptr = (int *)inc_ptr;

	for(int i = 0; i<1000000; i++){
		count++;
	};

	return NULL;
};



int main(){
	

	pthread_t inc_thr;
	pthread_t dec_thr;

	pthread_create(&inc_thr, NULL, inc, NULL);

	pthread_create(&dec_thr, NULL, dec, NULL);


	pthread_join(inc_thr, NULL);
	pthread_join(dec_thr, NULL);

	printf("updated %i \n", count);
	printf("Result = %i \n", count);


	return 0;
};




/*
if(pthread_create(&inc_thr, NULL, inc, &count)){

		fprintf(stderr, "error creating thread \n");

	};
*/
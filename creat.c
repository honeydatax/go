#include <stdio.h>
#include <stdlib.h>

int main(){
	FILE *f1;
	f1=fopen("currency.dat","w+");
	fprintf(f1,"CREATE TABLE CURRENCYS (\n");
	fprintf(f1,"ID INT NOT NULL ,\n");
	fprintf(f1,"CLIENT INT NOT NULL ,\n");
	fprintf(f1,"DATES CHAR(10) NOT NULL  ,\n");
	fprintf(f1,"REFS VARCHAR(25) NOT NULL, \n");
	fprintf(f1,"VALUE CURRENCY NOT NULL, \n");
	fprintf(f1,"PRIMARY KEY (ID)\n");
	fprintf(f1,");\n");
	fprintf(f1,".tables\n");
	fprintf(f1,".schema CURRENCYS\n");
	fprintf(f1,".quit\n");
	fclose(f1);
	printf("creating tables....\n");
	system("sqlite currency.db < currency.dat");
	return 0;
}




















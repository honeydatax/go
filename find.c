#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *find(FILE *f1,char *ccc);
void finds(FILE *f1,char *ccc);

int main(){
	FILE *f1;
	char cc[190];
	char *ccc=cc;
	int i=0;
	printf("Press control-c to exit\n");
	do{
		f1=fopen("currency.dat","w+");
		fprintf(f1,"select * from CURRENCYS where CLIENT = ");
		ccc=find (f1,ccc);
		fprintf(f1,";\n");
		fprintf(f1,"select SUM(VALUE) from CURRENCYS where CLIENT = ");
		finds (f1,ccc);
		fprintf(f1,";\n");
		fprintf(f1,".quit\n");
		fclose(f1);
		printf("................................\n");
		system("sqlite currency.db < currency.dat");
	}while(i==0);
	return 0;
}


char *find(FILE *f1,char *ccc){
	char *cc;
	cc=ccc;
	printf("find CLIENT number :? ");
	cc=fgets(cc,181,stdin);
	cc=strstr(cc,"\n");
	if(cc!=NULL) cc[0]=0;
	cc=ccc;
	printf("\n");
	finds(f1,ccc);
	return ccc;
}


void finds(FILE *f1,char *ccc){
	fprintf(f1," %s ",ccc);
	
}


















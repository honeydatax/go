# /usr/bin/sh
echo "press control-break to exit."
for b in {0..50}
do
printf "%s" "select * from MOVS where CLIENT = " > ./data/find.sql
echo 'enter client number yours refs  ,'
read a
echo $a >> ./data/find.sql
echo ';' >> ./data/find.sql
echo '.quit' >> ./data/find.sql
sqlite ./data/currency.dbf < ./data/find.sql
printf "%s" "select SUM(VALUE) from MOVS where CLIENT = " > ./data/find.sql
echo 'SUM:  ,'
echo $a >> ./data/find.sql
echo ';' >> ./data/find.sql
echo '.quit' >> ./data/find.sql
sqlite ./data/currency.dbf < ./data/find.sql
echo ------------------------------------------
done
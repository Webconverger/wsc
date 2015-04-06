export webc_id=just_testing
homepage=http://example.com
while true
do
	homepage=$(./client surf2 $homepage)
done

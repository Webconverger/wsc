export webc_id=just_testing
homepage=http://example.com
while true
do
	echo Launching $homepage
	homepage="$(./wsc surf2 $homepage)"
	echo cycle
done

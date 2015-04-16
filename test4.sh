export webc_id=just_testing
homepage=http://example.com
webck=firefox
while true
do
	homepage="$(echo ${rhomepage:-$homepage} | urldecode)"
	echo Launching with "${homepage}"
	rhomepage="$(./wsc ${webck:-/usr/bin/surf2} $homepage)"
done

-include secret/server.mk

update:
	ssh -l root ${server_ip} 'cd ${home_path}; . ./util/update.sh'


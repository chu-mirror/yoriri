-include secret/server.mk

reboot:
	ssh -l root ${server_ip} 'cd ${home_path}; . ./util/reboot.sh'


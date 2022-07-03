-include secret/server.mk

reboot:
	ssh -l root ${server_ip} '. util/reboot.sh'


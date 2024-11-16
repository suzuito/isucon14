/bin/alp: /tmp/alp_linux_amd64.tar.gz
	cd /tmp && tar xzvf $< && sudo mv alp $@

/tmp/alp_linux_amd64.tar.gz:
	cd /tmp && curl -LO https://github.com/tkuchiki/alp/releases/download/v1.0.21/alp_linux_amd64.tar.gz


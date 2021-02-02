centos:
	docker-compose up -d

ssh:
	ssh -F ssh.cfg localhost

server-setup:
	ansible-playbook server-setup.yml -i hosts 

deploy-dev:
	ansible-playbook hugo-deploy.yml -i hosts -l dev

deploy-staging:
	ansible-playbook hugo-deploy.yml -i hosts -l dev
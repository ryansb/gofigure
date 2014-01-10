# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.box = "opscode-f19-64"
    config.vm.box_url = "https://opscode-vm-bento.s3.amazonaws.com/vagrant/opscode-fedora-19_provisionerless.box"
    config.vm.provision :ansible do |ansible|
        ansible.playbook = "ansible/mongo.yml"
        ansible.inventory_path = "ansible/hosts"
        ansible.sudo = true
        ansible.host_key_checking = false
    end
    config.vm.network :private_network, ip: "172.16.0.101"
    config.vm.provider :virtualbox do |vb|
        vb.gui = false
        vb.customize ["modifyvm", :id, "--memory", "600"]
        vb.customize ["modifyvm", :id, "--cpus", "2"]
        # This allows symlinks to be created within the /vagrant root directory
        vb.customize ["setextradata", :id, "VBoxInternal2/SharedFoldersEnableSymlinksCreate/v-root", "1"]
    end
end

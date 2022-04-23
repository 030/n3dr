Vagrant.require_version ">= 2.2.18"

Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |v|
    v.gui = true
    v.customize ['setextradata', :id, 'GUI/ScaleFactor', '2.0']
    v.customize ["setextradata", "global", "GUI/SuppressMessages", "all" ]
    v.customize ["modifyvm", :id, "--clipboard", "bidirectional"]
    v.customize ["modifyvm", :id, "--memory", "8192"]
    v.customize ["modifyvm", :id, "--cpus", "4"]
  end

  config.vm.define "n3dr" do |n3dr|
    n3dr.vm.guest = :windows
    n3dr.vm.communicator = "winrm"
    n3dr.winrm.username = "vagrant"
    n3dr.winrm.password = "vagrant"
    n3dr.vm.box = "win2016/n3dr"
    n3dr.vm.hostname = "n3dr"
    n3dr.vm.network "public_network",
      ip: ENV['VAGRANT_N3DR_IP'],
      bridge: ENV['VAGRANT_N3DR_NETWORK_ADAPTER']
      n3dr.vm.provision "shell",
      path: "vagrant/scripts/n3dr.ps1"
    n3dr.vm.provision "windows-update", filters: [
      "exclude:$_.Title -like '*Preview*'",
      "include:$_.Title -like '*Cumulative Update for *'",
      "include:$_.AutoSelectOnWebSites"]
  end

  config.vm.define "nexus3" do |nexus3|
    nexus3.vm.box = "ubuntu/focal64"
    nexus3.vm.hostname = "nexus3"
    nexus3.vm.network "public_network",
      ip: ENV['VAGRANT_NEXUS3_IP'],
      bridge: ENV['VAGRANT_N3DR_NETWORK_ADAPTER']
    nexus3.vm.provision "shell" do |s|
      s.path = "vagrant/scripts/nexus3.sh"
      s.env =  { "N3DR_APT_GPG_SECRET" => ENV['N3DR_APT_GPG_SECRET'] }
    end
  end
end

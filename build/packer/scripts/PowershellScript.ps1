function InstallNuget {
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

    # https://answers.microsoft.com/en-us/windows/forum/all/trying-to-install-program-using-powershell-and/4c3ac2b2-ebd4-4b2a-a673-e283827da143
    Set-ItemProperty -Path 'HKLM:\SOFTWARE\Wow6432Node\Microsoft\.NetFramework\v4.0.30319' -Name 'SchUseStrongCrypto' -Value '1' -Type DWord
    Set-ItemProperty -Path 'HKLM:\SOFTWARE\Microsoft\.NetFramework\v4.0.30319' -Name 'SchUseStrongCrypto' -Value '1' -Type DWord

    Install-PackageProvider -Name Nuget -Force
}

function PrivateNetwork {
    Install-Module WindowsBox.Network -Force
    Set-NetworkToPrivate
}

function DisableAutoLogon {
    Install-Module WindowsBox.AutoLogon -Force
    Disable-AutoLogon
}

function InstallChocolatey {
    Write-Output "Installing Chocolatey..."
    Set-ExecutionPolicy Bypass -Scope Process -Force;
    [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; 
    Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
    Write-Output "Chocolatey installation has been completed"
}

function InstallChocolateyPackages {
    choco install -y golang
}

function SetVagrantAccount {
    Install-Module WindowsBox.VagrantAccount -Force
    Set-VagrantAccount
}

function InstallVMGuestTools {
    Install-Module WindowsBox.VMGuestTools -Force
    Install-VMGuestTools
}

function OptimizeDiskUsage {
    Install-Module WindowsBox.Compact -Force
    Optimize-DiskUsage
}

function DisableHibernate {
    Install-Module WindowsBox.Hibernation -Force
    Disable-Hibernation
}

function EnableRDP {
    Install-Module WindowsBox.RDP -Force
    Enable-RDP
}

function DisableUAC {
    Install-Module WindowsBox.UAC -Force
    Disable-UAC
}

function ExplorerConfig {
    Install-Module WindowsBox.Explorer -Force
    Set-ExplorerConfiguration
}

function Main {
    InstallNuget
    PrivateNetwork
    DisableAutoLogon
    DisableUAC
    EnableRDP
    ExplorerConfig
    DisableHibernate
    SetVagrantAccount
    InstallVMGuestTools
    InstallChocolatey
    InstallChocolateyPackages
    OptimizeDiskUsage
}

Main

printBanner() {
    echo " ____                            ____ ___ "
    echo "/ ___|  ___  _ __   __ _ _ __   / ___|_ _|"
    echo "\___ \ / _ \| '_ \ / _  | '__| | |    | | "
    echo " ___) | (_) | | | | (_| | |    | |___ | | "
    echo "|____/ \___/|_| |_|\__,_|_|     \____|___|"
    echo ""
    echo "has been installed with success"
}

install() {
    wget https://github.com/odair-pedro/sonarci/releases/latest/download/sonarci-linux-x64.tar.gz
    tar -xf sonarci-linux-x64.tar.gz sonarci
    rm sonarci-linux-x64.tar.gz

    mv ./sonarci /usr/local/bin/sonarci
    chmod +x /usr/local/bin/docker-compose

    printBanner
}

install
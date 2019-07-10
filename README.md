## Setup

### Install Nginx
* brew install nginx
* start nginx: `./nginx`

### Install Yarn
* brew install yarn

### Setup Project
* git clone `git@github.com:hmarcelino/m3u-proxy.git`
* cd m3u-proxy
* copy `config/config-dev.yml` and change necessary configuration
* run `yarn run run -f <path-to-yaml-file>`

### Setup crontab  
* Add to crontab: 
```$bash
$> crontab -e

0  *  *  *  *  cd ${PATH_TO_THIS_FOLDER} && date >> output.log && yarn run prod 2>&1 >> output.log
```

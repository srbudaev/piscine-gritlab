curl -s https://01.gritlab.ax/assets/superhero/all.json | jq ' .[] | select(.id =='$HERO_ID') | .connections | .relatives' | sed 's/"//g'

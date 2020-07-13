git add .
dt=$(date '+%A-%b-%d-%Y-%H-%M-%S');
git commit -m "hello world-$dt"
git push --set-upstream origin master
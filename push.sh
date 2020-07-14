echo -n "Enter url git project > "
read url
echo -n "Please select a name for commit > "
read namecommit
git init
git add .
git commit -m "$namecommit"
git remote add origin $url
git pull --rebase origin master
git push origin master 
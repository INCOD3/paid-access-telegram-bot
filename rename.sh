# parse command line argument 
# if the first argument is missing, print usage

if [ $# -eq 0 ]
then
  echo "Usage: $0 name_of_go_package"
  exit 1
fi

find . -name "*.go" -exec sed -i '' -e "s/telegram-bot-template/$1/g" {} \;
sed -i '' -e "s/telegram-bot-template/$1/g" go.mod

